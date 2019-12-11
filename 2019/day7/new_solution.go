package main

import (
	"../../shared"
	"fmt"
	"io/ioutil"
	"log"
	"modernc.org/mathutil"
	"sort"
)

var input []int
var phaseSeqWritten []bool
var phaseSequence =

const numComputers = 5

func main() {
	file, err := ioutil.ReadFile("./2019/day7/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	mem := string(file)

	maxSig, maxSeq := Day7Task1(mem, sort.IntSlice{0, 1, 2, 3, 4})
	fmt.Printf("Result Day7-Task1: maxSeq: %v maxSignal: %d\n", maxSeq, maxSig)

	maxSig, maxSeq = Day7Task2(mem, sort.IntSlice{5, 6, 7, 8, 9})
	fmt.Printf("Result Day7-Task2: maxSeq: %v maxSignal: %d\n", maxSeq, maxSig)

}

func Day7Task1(memory string, phaseSequence sort.IntSlice) (int, []int) {
	var amps []*shared.Computer
	for i := 0; i < numComputers; i++ {
		amp := shared.NewComputer(i, memory, onError)
		amps = append(amps, amp)
	}

	maxThrusterSignal := 0
	var maxPermutation []int
	mathutil.PermutationFirst(phaseSequence)
	stop := false
	for !stop {
		// set AMP a
		ampA := amps[0]
		ampA.QueueInput(phaseSequence[0])
		ampA.QueueInput(0)
		res := ampA.Run()

		for i := 1; i < len(amps); i++ {
			amp := amps[i]
			amp.QueueInput(phaseSequence[i])
			amp.QueueInput(res)
			res = amp.Run()
			//fmt.Printf("RESULT: %d: %d\n\n", i, res)
		}
		if res > maxThrusterSignal {
			maxThrusterSignal = res
			maxPermutation = shared.CloneSlice(phaseSequence)
		}
		for _, amp := range amps {
			amp.Reset()
		}
		stop = !mathutil.PermutationNext(phaseSequence)
	}
	return maxThrusterSignal, maxPermutation
}

func Day7Task2(memory string, phaseSequence sort.IntSlice) (int, []int) {
	var amps []*shared.Computer
	for i := 0; i < numComputers; i++ {
		amp := shared.NewComputer(i, memory, onError)
		amps = append(amps, amp)
	}

	maxThrusterSignal := 0
	var maxPermutation []int
	mathutil.PermutationFirst(phaseSequence)
	stop := false
	for !stop {
		// set AMP a
		ampA := amps[0]
		ampA.QueueInput(phaseSequence[0])
		ampA.QueueInput(0)
		res := ampA.Run()

		for i := 1; i < len(amps); i++ {
			amp := amps[i]
			amp.QueueInput(phaseSequence[i])
			amp.QueueInput(res)
			//res = amp.Run()

			status := amp.Run()
			if status == StatusWaiting {
				amp.GetResult()
			}
			//fmt.Printf("RESULT: %d: %d\n\n", i, res)
		}
		if res > maxThrusterSignal {
			maxThrusterSignal = res
			maxPermutation = shared.CloneSlice(phaseSequence)
		}
		for _, amp := range amps {
			amp.Reset()
		}
		stop = !mathutil.PermutationNext(phaseSequence)
	}
	return maxThrusterSignal, maxPermutation
}

func onError(cid int, err error) {
	fmt.Printf("Error from %d received: %v", cid, err)
}
