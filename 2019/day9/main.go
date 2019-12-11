package main

import (
	"../../shared"
	"fmt"
	"os"
)

func main() {
	mem, err := shared.ReadFile("./2019/day9/input.txt")
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	res := Day9Task(mem, 1)
	shared.PrintSolution(9, 1, "BOOST keycode: %d", res)
	res = Day9Task(mem, 2)
	shared.PrintSolution(9, 2, "coordinates of the distress signal: %d", res)
}

func Day9Task(mem string, in int) int {
	c := shared.NewComputer(0, mem, func(cid int, err error) {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	})
	c.QueueInput(in)
	c.Run()
	c.WaitUntilHalted()
	outputs, _ := c.ReadAllOutputs()
	return outputs[0]
}
