package shared

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

type onErrorFunc = func(cid int, err error)

type OpCode int

type Instruction struct {
	opCode   int
	argCount int // including opCode e.g. len(ADD) = 4: opCode,arg0,arg1,arg3
	execute  func(in Instruction, args []int /* paramModes []ParameterMode */)
}

type ParameterMode int

const (
	ParamModePositional ParameterMode = 0
	ParamModeImmediate  ParameterMode = 1
)

type Status int

const (
	StatusCreated Status = 0
	StatusRunning Status = 1
	StatusWaiting Status = 2
	StatusHalted  Status = 3
)

func (s Status) String() string {
	switch s {
	case 0:
		return "CREATED"
	case 1:
		return "RUNNING"
	case 2:
		return "WAITING"
	case 3:
		return "HALTED"
	default:
		return "INVALID_STATUS"
	}
}

func (p *ParameterMode) Parse(in int) error {
	switch in {
	case 0:
		*p = ParamModePositional
	case 1:
		*p = ParamModeImmediate
	default:
		return errors.Errorf("invalid parameter mode: %d", in)
	}
	return nil
}

type Computer struct {
	id             int
	originalMem    []int
	memory         []int
	ip             int
	instructionSet map[OpCode]Instruction
	onError        onErrorFunc
	error          chan error
	status         Status
	input          chan int
	output         chan int
	debug          bool
}

func NewComputer(id int, memory string, onError onErrorFunc) *Computer {
	mem, err := StrToMem(memory)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	c := &Computer{
		id:          id,
		originalMem: CloneSlice(mem),
		memory:      CloneSlice(mem),
		ip:          0,
		onError:     onError,
		input:       make(chan int, 10),
		output:      make(chan int, 10),
		status:      StatusCreated,
	}
	instructions := []Instruction{
		{
			opCode:   1,
			argCount: 3,
			execute:  c.addExec,
		},
		{
			opCode:   2,
			argCount: 3,
			execute:  c.multExec,
		},
		{
			opCode:   3,
			argCount: 1,
			execute:  c.readExec,
		},
		{
			opCode:   4,
			argCount: 1,
			execute:  c.writeExec,
		},
		{
			opCode:   5,
			argCount: 2,
			execute:  c.jumpTrueExec,
		},
		{
			opCode:   6,
			argCount: 2,
			execute:  c.jumpFalseExec,
		},
		{
			opCode:   7,
			argCount: 3,
			execute:  c.lessThanExec,
		},
		{
			opCode:   8,
			argCount: 3,
			execute:  c.equalsExec,
		},
		{
			opCode:   99,
			argCount: 0,
			execute:  c.haltExec,
		},
	}

	iSet := make(map[OpCode]Instruction)
	for _, ins := range instructions {
		iSet[OpCode(ins.opCode)] = ins
	}
	c.instructionSet = iSet
	return c
}

func (c *Computer) SetDebug(val bool) {
	c.debug = val
}

func (c *Computer) Reset() {
	c.memory = CloneSlice(c.originalMem)
	c.setStatus(StatusCreated)
	close(c.input)
	close(c.output)
	c.input = make(chan int, 10)
	c.output = make(chan int, 10)
	c.ip = 0
}

func (c *Computer) QueueInput(val int) {
	if c.Status() == StatusHalted {
		return
	}
	c.input <- val
}

func (c *Computer) log(format string, args ...interface{}) {
	if !c.debug {
		return
	}
	c.log(format+"\n", args...)
}

func (c *Computer) Run() {
	if c.Status() != StatusCreated {
		return
	}
	c.log("[%d] started\n", c.id)
	var err error
	c.setStatus(StatusRunning)
	for c.status == StatusRunning {
		ins := c.fetchNextInstruction()
		err = c.execute(ins)
		if err != nil {
			err = errors.Wrapf(err, "execution error. ip: %d", c.ip)
			c.error <- err
			c.setStatus(StatusHalted)
		}
	}
}

func (c *Computer) WaitForOutput() int {
	return <-c.output
}

func (c *Computer) GetOutput() (int, error) {
	out, ok := <-c.output
	if !ok {
		return 0, errors.New("no output")
	}
	return out, nil
}

func (c *Computer) Error() error {
	if err, ok := <-c.error; ok {
		return err
	}
	return nil
}

func (c *Computer) writeError() {
	err := <-c.error
	c.onError(c.id, err)
}

func (c *Computer) Shutdown() {

	if c.Status() == StatusHalted {
		return
	}
	c.setStatus(StatusHalted)

	close(c.input)
	close(c.output)

	c.log("[%d] shutdown\n", c.id)
}

func (c *Computer) readMemChunk(start, length int) []int {
	return c.memory[start : start+length]
}

func (c *Computer) fetchNextInstruction() int {
	return c.memory[c.ip]
}

func (c *Computer) incrementIP(amount int) {
	c.ip += amount
}

func (c *Computer) execute(insCode int) error {
	op := getOpCode(insCode)

	ins, ok := c.instructionSet[op]
	if !ok {
		return errors.Errorf("opcode %d not supported by system", op)
	}

	modes, err := parameterModes(insCode, ins.argCount)
	if err != nil {
		return err
	}
	args := c.fetchArgs(ins.argCount, modes)
	ins.execute(ins, args)
	return nil
}

func (c *Computer) fetchArgs(length int, modes []ParameterMode) []int {
	args := c.readMemChunk(c.ip+1, length)
	params := make([]int, len(args))

	for i := 0; i < len(args); i++ {
		switch modes[i] {
		case ParamModePositional:
			params[i] = args[i]
		case ParamModeImmediate:
			params[i] = c.ip + 1 + i
		}
	}
	return params
}

func parameterModes(code int, length int) ([]ParameterMode, error) {
	// skip first 2 digits because opCode
	code /= 100

	var modes []ParameterMode
	for i := 0; i < length; i++ {
		m := code % 10
		var p ParameterMode
		if err := p.Parse(m); err != nil {
			return nil, err
		}
		modes = append(modes, p)
		code /= 10
	}
	return modes, nil
}

func getOpCode(in int) OpCode {
	op := in % 100
	return OpCode(op)
}

func (c *Computer) writeInt(pos int, val int) {
	c.memory[pos] = val
}

func (c *Computer) readMem(pos int) int {
	return c.readMemChunk(pos, 1)[0]
}

func (c *Computer) setIP(val int) {
	c.ip = val
}

func (c *Computer) setStatus(status Status) {
	c.log("[%d] status changed: %s\n", c.id, status)
	c.status = status
}

func (c *Computer) addExec(in Instruction, args []int) {
	result := c.readMem(args[0]) + c.readMem(args[1])
	c.writeInt(args[2], result)
	c.incrementIP(in.argCount + 1)
}

func (c *Computer) haltExec(Instruction, []int) {
	c.setStatus(StatusHalted)
}

func (c *Computer) equalsExec(in Instruction, args []int) {
	val := 0
	if c.readMem(args[0]) == c.readMem(args[1]) {
		val = 1
	}
	c.writeInt(args[2], val)
	c.incrementIP(in.argCount + 1)
}

func (c *Computer) lessThanExec(in Instruction, args []int) {
	val := 0
	if c.readMem(args[0]) < c.readMem(args[1]) {
		val = 1
	}
	c.writeInt(args[2], val)
	c.incrementIP(in.argCount + 1)
}

func (c *Computer) jumpTrueExec(in Instruction, args []int) {
	if c.readMem(args[0]) != 0 {
		c.setIP(c.readMem(args[1]))
		return
	}
	c.incrementIP(in.argCount + 1)
}

func (c *Computer) jumpFalseExec(in Instruction, args []int) {
	if c.readMem(args[0]) == 0 {
		c.setIP(c.readMem(args[1]))
		return
	}
	c.incrementIP(in.argCount + 1)
}

func (c *Computer) writeExec(in Instruction, args []int) {
	val := c.readMem(args[0])
	c.output <- val
	//c.onWrite(c.id, val)
	c.incrementIP(in.argCount + 1)
}

func (c *Computer) readExec(in Instruction, args []int) {
	c.setStatus(StatusWaiting)
	val := <-c.input
	c.setStatus(StatusRunning)
	//c.log("%d waiting for input\n", c.id)
	//val := c.onRead(c.id)
	c.log("[%d] input received: %d\n", c.id, val)

	c.writeInt(args[0], val)
	c.incrementIP(in.argCount + 1)
}

func (c *Computer) multExec(in Instruction, args []int) {
	result := c.readMem(args[0]) * c.readMem(args[1])
	c.writeInt(args[2], result)
	c.incrementIP(in.argCount + 1)
}

func (c *Computer) Status() Status {
	return c.status
}
