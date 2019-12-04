package main

import "fmt"

func main() {
	start := 172_851
	end := 675_869

	count := 0
	for i := start; i <= end; i++ {
		if meetsCriteria(i) {
			count++
		}
	}
	fmt.Printf("count of numbers meeting criteria: %d\n", count)
}

func meetsCriteria(num int) bool {
	hasAdjacentDigits := false
	isSorted := true

	digits := digits(num)
	// note: iterate until i < len(digits)-1 because we access last element with i+1
	for i := 0; i < len(digits)-1; i++ {
		if digits[i] > digits[i+1] {
			isSorted = false
			break
		}
		if digits[i] == digits[i+1] {
			hasAdjacentDigits = true
		}
	}
	return hasAdjacentDigits && isSorted
}

// digits takes a number and returns a slice of its digits, e.g. 123_456 => [1 2 3 4 5 6]
func digits(num int) []int {
	if num < 10 {
		return []int{num}
	}
	result := []int{num % 10}
	return append(digits(num/10), result...)
}
