package main

import "fmt"

func main() {
	start := 172_851
	end := 675_869

	//test := []int{112_222, 123_456, 111_122, 112_233, 223_333, 122_333, 111_111, 111_112, 111_122, 111_333}
	//
	//for _, i := range test {
	//	fmt.Printf("%d: %t\n", i, meetsCriteria(i))
	//}

	//return
	count := 0
	for i := start; i <= end; i++ {
		if meetsCriteria(i) {
			count++
		}
	}
	fmt.Printf("count of numbers meeting criteria: %d\n", count)
}

func meetsCriteria(num int) bool {
	hasTuple := false
	possibleTuple := false

	inAdjacentGroup := false

	digits := digits(num)
	// note: iterate until i < len(digits)-1 because we access last element with i+1
	for i := 0; i < len(digits)-1; i++ {
		if digits[i] > digits[i+1] {
			return false
		}
		match := digits[i] == digits[i+1]
		possibleTuple = !match && possibleTuple || match && !inAdjacentGroup
		hasTuple = !match && (hasTuple || inAdjacentGroup && possibleTuple) || match && hasTuple
		inAdjacentGroup = match

		/*
			// more readable version below ;) Propositional calculus ftw
			if digits[i] == digits[i+1] {
				possibleTuple = !atAdjacentStart
				atAdjacentStart = true

				if !atAdjacentStart {
					atAdjacentStart = true
					possibleTuple = true
					continue
				}
				possibleTuple = false
			} else {
				hasTuple = hasTuple || atAdjacentStart && possibleTuple
				atAdjacentStart = false
			}
		*/
	}
	hasTuple = hasTuple || inAdjacentGroup && possibleTuple // edge case for last 2 digits are tuple
	return hasTuple
}

// digits takes a number and returns a slice of its digits, e.g. 123_456 => [1 2 3 4 5 6]
func digits(num int) []int {
	if num < 10 {
		return []int{num}
	}
	result := []int{num % 10}
	return append(digits(num/10), result...)
}
