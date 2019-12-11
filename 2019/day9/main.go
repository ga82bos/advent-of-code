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

	//mem := "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99"
	//mem := "1102,34915192,34915192,7,4,7,99,0"
	//mem :="104,1125899906842624,99"
	c := shared.NewComputer(0, mem, func(cid int, err error) {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	})
	c.QueueInput(2)
	c.Run()
	c.WaitUntilHalted()
	outputs, _ := c.ReadAllOutputs()
	fmt.Printf("%v\n", outputs)
}
