package main

import (
	"fmt"
	"gitea.curunir/Pokerkoffer/gitealoggertest/log"
	"io/ioutil"
	"os"
	"strings"
)

var (
	orbits map[string]string
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	contents := string(input)

	//	contents := `COM)B
	//B)C
	//C)D
	//D)E
	//E)F
	//B)G
	//G)H
	//D)I
	//E)J
	//J)K
	//K)L`
	orbits = make(map[string]string)

	in := strings.Split(contents, "\n")
	for _, o := range in {
		objects := strings.Split(o, ")")
		orbitee := objects[0] // the object orbited by some other object A)B => A
		orbiter := objects[1] // the object orbiting some other object A)B => B
		orbits[orbiter] = orbitee
	}

	totalOrbits := 0
	// iterate over all objects that orbit another object
	for orbiter := range orbits {
		dist := distToCOM(orbiter)
		fmt.Printf("Dist of %s to COM: %d\n", orbiter, dist)
		totalOrbits += dist
	}
	fmt.Printf("Total orbits: %d\n", totalOrbits)
}

func distToCOM(orbiter string) int {
	// will not happen because every object orbits exactly one other object according to problem statement
	//if _, ok := orbitMap[orbiter]; !ok {
	//	return -1
	//}
	dist := 0
	currObj := orbiter
	for currObj != "COM" {
		orbitee, ok := orbits[currObj]
		if !ok {
			return 0
		}
		currObj = orbitee
		dist++
	}
	return dist
}
