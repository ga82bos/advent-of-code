package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("failed to open file: %w", err)
	}
	contents := string(file)
	lines := strings.Split(contents, "\n")

	var fuelTotal int
	for _, line := range lines {
		mass, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal("failed to convert string to int: %w", err)
		}
		fuelTotal += calcRequiredFuel(mass)
	}

	fmt.Println(fuelTotal)
}

func calcRequiredFuel(mass int) int {
	requiredFuel := int(math.Floor(float64(mass)/3) - 2)

	if requiredFuel <= 0 {
		return 0
	}

	return requiredFuel + calcRequiredFuel(requiredFuel)
}
