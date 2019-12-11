package main

import (
	"../../shared"
	"fmt"
	"io/ioutil"
	"log"
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
	file, err := ioutil.ReadFile("./2019/day5/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	mem := string(file)
	res := Day5Task12(mem, 1)
	shared.PrintSolution(5, 1, "Diagnostic code: %d", res)

	res = Day5Task12(mem, 5)
	shared.PrintSolution(5, 2, "Diagnostic code: %d", res)

}

func Day5Task12(mem string, systemID int) int {
	c := shared.NewComputer(0, mem, func(cid int, err error) {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	})

	c.QueueInput(systemID)
	go c.Run()
	c.WaitUntilHalted()
	result, err := c.ReadAllOutputs()
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		os.Exit(1)
	}
	for _, res := range result {
		if res != 0 {
			return res
		}
	}
	return -1
}
