package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	// opcode: length
	opcodes = map[int]func(opcode, ic int, modes []int, program []int) int{
		99: haltF,
		1:  add,
		2:  mult,
		3:  read,
		4:  printVal,
		5:  jumpCond, // jump-if-true
		6:  jumpCond, // jump-if-false
		7:  cmp,      // less-than
		8:  cmp,      // equals
	}
	halt = false
)

const (
	ModePositional = 0
	ModeImmediate  = 1
)

/*
	Format:

	instruction:
	1, A, B, C		Add, *C = *A + *B (A, B, C Positions)
	2, A, B, C		Mult *C = *A * *B
	99 				Finished
	else			Something went wrong
*/

func main() {
	file, err := ioutil.ReadFile("./2019/day5/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	contents := string(file)
	//contents := "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99"
	//contents := "3,9,8,9,10,9,4,9,99,-1,8"
	programStr := strings.Split(contents, ",")
	var memory []int
	for _, s := range programStr {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal("failed to convert " + s)
		}
		memory = append(memory, n)
	}

	compute(memory)

	//noun := 0
	//stop := false
	//for !stop {
	//	verb := 0
	//	for verb <= 99 {
	//		p := append(memory[:0:0], memory...) // clone memory
	//		p[1] = noun
	//		p[2] = verb
	//		compute(p)
	//		result := p[0]
	//		if result == 19690720 {
	//			fmt.Printf("FOUND: noun: %d, verb: %d", noun, verb)
	//			stop = true
	//		}
	//		verb++
	//	}
	//	noun++
	//
	//}
}

func compute(program []int) {
	//halt := false
	i := 0
	for !halt {
		op := program[i]
		opcode, modes, err := parseInstrOpCode(op)
		if err != nil {
			log.Fatal(err.Error())
		}
		if f, ok := opcodes[opcode]; ok {
			i = f(opcode, i, modes, program)
		} else {
			log.Printf("invalid instruction at pos %d: %d", i, op)
			halt = true
		}
		if i > len(program) {
			log.Fatal("i exceeds len(program): " + strconv.Itoa(i))
		}
	}
}

// returns opcode, [paramsModes 0-n]
func parseInstrOpCode(instrOp int) (int, []int, error) {
	modes := []int{0, 0, 0}

	if instrOp < 100 {
		return instrOp, modes, nil
	}
	op := instrOp % 100
	instrOp /= 100
	modes[0] = instrOp % 10
	instrOp /= 10
	modes[1] = instrOp % 10
	instrOp /= 10
	modes[2] = instrOp % 10

	//
	//c := strconv.Itoa(instrOp)
	//op, err := strconv.Atoi(c[len(c)-2:])
	//if err != nil {
	//	return 0, nil, err
	//}
	//// 1002
	//// len(c)-1 because array starts at 0 and -2 because we want to skip the last 2 because they are opcode
	//for i := len(c)-1-2; i >= 0; i-- {
	//	n, _ := strconv.Atoi(string(c[i]))
	//	modes[len(c)-1-2-1] = n
	//}
	return op, modes, nil
}

func haltF(_, ic int, _ []int, _ []int) int {
	halt = true
	return ic
}

func add(_, ic int, modes []int, program []int) int {
	calcArgs(modes, ic, program, 3)

	program[modes[2]] = program[modes[0]] + program[modes[1]]
	return ic + 4
}

func mult(_, ic int, modes []int, program []int) int {
	calcArgs(modes, ic, program, 3)

	program[modes[2]] = program[modes[0]] * program[modes[1]]
	return ic + 4
}

func read(_, ic int, modes []int, program []int) int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number: ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	in, err := strconv.Atoi(text)
	if err != nil {
		log.Printf("not a valid number %s", text)
		return -1
	}
	calcArgs(modes, ic, program, 1)
	program[modes[0]] = in
	return ic + 2
}

func printVal(_, ic int, modes []int, program []int) int {
	calcArgs(modes, ic, program, 1)

	n := strconv.Itoa(program[modes[0]])
	fmt.Printf("%s", n)
	return ic + 2
}

// instructionlength: 3
func jumpCond(opcode, ic int, modes []int, program []int) int {
	calcArgs(modes, ic, program, 2)

	isZero := program[modes[0]] == 0
	addr := program[modes[1]]

	switch opcode {
	case 5:
		if !isZero {
			return addr
		}
	case 6:
		if isZero {
			return addr
		}
	default:
		log.Printf("invlaid opcode in jmp: %d", opcode)
	}
	//if opcode == 5 && !isZero || opcode == 6 && isZero {
	//	return addr
	//}

	return ic + 3
}

func cmp(opcode, ic int, modes []int, program []int) int {
	calcArgs(modes, ic, program, 3)

	arg0 := program[modes[0]]
	arg1 := program[modes[1]]

	v := 0
	switch opcode {
	case 7:
		if arg0 < arg1 {
			v = 1
		}
	case 8:
		if arg0 == arg1 {
			v = 1
		}
	default:
		log.Printf("invlaid opcode in cmp: %d", opcode)
	}
	program[modes[2]] = v
	return ic + 4 // op + 3 params
}

func calcArgs(modes []int, ic int, program []int, readCount int) {
	for i := 0; i < readCount; i++ {
		insPtr := ic + 1 + i
		switch modes[i] {
		case ModePositional:
			modes[i] = program[insPtr]
		case ModeImmediate:
			modes[i] = insPtr
		}
	}
}
