package shared

import (
	"strconv"
	"strings"
)

// intToDigitArr takes a number and returns an array of digits (e.g. 12345 => [1 2 3 4 5])
func IntToDigitArr(num int) []int {
	if num < 10 {
		return []int{num}
	}
	result := []int{num % 10}
	return append(IntToDigitArr(num/10), result...)
}

// intToDigitArr takes a number string and returns an array of digits (e.g. "0012345" => [1 2 3 4 5])
func StrToDigitArr(num string) ([]int, error) {
	split := strings.Split(num, "")
	var nums []int
	for _, num := range split {
		n, err := strconv.Atoi(num)
		if err != nil {
			return []int{}, err
		}
		nums = append(nums, n)
	}
	return nums, nil
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
