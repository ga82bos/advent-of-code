package main

import (
	"../../shared"
	"fmt"
	"io/ioutil"
	"os"
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
	file, err := ioutil.ReadFile("./2019/day2/input_original.txt")
	if err != nil {
		fmt.Printf("ERROR: %s", err.Error())
		os.Exit(1)
	}
	mem := string(file)
	res := Day2Task1(mem)
	shared.PrintSolution(2, 1, "%d", res)

	noun, verb := Day2Task2(mem)
	shared.PrintSolution(2, 2, "noun: %d verb: %d", noun, verb)
}

func Day2Task2(mem string) (int, int) {
	c := shared.NewComputer(0, mem, func(cid int, err error) {
		fmt.Printf("ERROR: %s", err.Error())
		os.Exit(1)
	})

	found := false
	for !found {
		verb := 0
		for verb < 100 {
			noun := 0
			for noun < 100 {
				c.Reset()
				c.WriteMem(1, noun)
				c.WriteMem(2, verb)
				c.Run()
				c.WaitUntilHalted()
				if c.ReadMem(0) == 19690720 {
					found = true
					return noun, verb
				}
				noun++
			}
			verb++
		}
	}
	return -1, -1
}

func Day2Task1(mem string) int {
	c := shared.NewComputer(0, mem, func(cid int, err error) {
		fmt.Printf("ERROR: %s", err.Error())
		os.Exit(1)
	})

	c.Run()
	c.WaitUntilHalted()
	return c.ReadMem(0)
}
