package shared

import (
	"github.com/pkg/errors"
)

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
	memory         []int
	ip             int
	instructionSet map[OpCode]Instruction
	read           func() int
	write          func(out int)
	halt           bool
}

func NewComputer(memory []int, inputCB func() int, outputCB func(out int)) *Computer {
	c := &Computer{
		memory: memory,
		ip:     0,
		read:   inputCB,
		write:  outputCB,
		halt:   false,
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

func (c *Computer) Run(input ...int) error {
	var err error
	for !c.halt {
		ins := c.fetchNextInstruction()
		err = c.execute(ins)
		if err != nil {
			err = errors.Wrapf(err, "execution error. ip: %d", c.ip)
			c.halt = true
		}
	}
	return err
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

func (c *Computer) addExec(in Instruction, args []int) {
	result := c.readMem(args[0]) + c.readMem(args[1])
	c.writeInt(args[2], result)
	c.incrementIP(in.argCount + 1)
}

func (c *Computer) haltExec(Instruction, []int) {
	c.halt = true
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
	c.write(val)
	c.incrementIP(in.argCount + 1)
}

func (c *Computer) readExec(in Instruction, args []int) {
	input := c.read()
	c.writeInt(args[0], input)
	c.incrementIP(in.argCount + 1)
}

func (c *Computer) multExec(in Instruction, args []int) {
	result := c.readMem(args[0]) * c.readMem(args[1])
	c.writeInt(args[2], result)
	c.incrementIP(in.argCount + 1)
}
