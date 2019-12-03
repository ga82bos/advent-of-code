package main

import (
	"fmt"
	"gitea.curunir/Pokerkoffer/gitealoggertest/log"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	OpCodeHALT = 99
	OpCodeADD  = 1
	OpCodeMULT = 2
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
	file, err := ioutil.ReadFile("./input_original.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	contents := string(file)
	programStr := strings.Split(contents, ",")
	var memory []int
	for _, s := range programStr {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal("failed to convert " + s)
		}
		memory = append(memory, n)
	}

	noun := 0
	stop := false
	for !stop {
		verb := 0
		for verb <= 99 {
			p := append(memory[:0:0], memory...) // clone memory
			p[1] = noun
			p[2] = verb
			compute(p)
			result := p[0]
			if result == 19690720 {
				fmt.Printf("FOUND: noun: %d, verb: %d", noun, verb)
				stop = true
			}
			verb++
		}
		noun++

		if noun > 100 || verb > 100 {
			log.Fatal("err")
		}
	}

	//compute(memory)
	//fmt.Println(memory)
}

func compute(program []int) {
	halt := false
	i := 0
	for !halt {
		op := program[i]
		switch op {
		case OpCodeHALT:
			halt = true
		case OpCodeADD:
			processInstruction(program[i], program[i+1], program[i+2], program[i+3], program)
			i += 4
		case OpCodeMULT:
			processInstruction(program[i], program[i+1], program[i+2], program[i+3], program)
			i += 4
		default:
			log.Error(fmt.Sprintf("invalid instruction at pos %d: %d", i, op))
			halt = true
		}
		if i > len(program) {
			log.Fatal("i exceeds len(program): " + strconv.Itoa(i))
		}
	}
}

func processInstruction(opcode, aPtr, bPtr, storePtr int, program []int) {

	aVal := program[aPtr]
	bVal := program[bPtr]

	switch opcode {
	case 1:
		program[storePtr] = aVal + bVal
	case 2:
		program[storePtr] = aVal * bVal
	default:
		log.Fatal(fmt.Sprintf("wrong opcode: %d", opcode))
	}

}
