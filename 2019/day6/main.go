package main

import (
	"fmt"
	"gitea.curunir/Pokerkoffer/gitealoggertest/log"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

var (
	orbits    map[string]string
	orbitedBy map[string][]string // part2
	visited   map[string]struct{}
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	contents := string(input)
	//	contents := `COM)B
	//B)YOU
	//YOU)D
	//D)E
	//E)F
	//B)G
	//G)H
	//D)I
	//E)J
	//J)K
	//K)L
	//I)SAN`
	orbits = make(map[string]string)
	orbitedBy = make(map[string][]string)
	visited = make(map[string]struct{})

	in := strings.Split(contents, "\n")
	for _, o := range in {
		objects := strings.Split(o, ")")
		orbitee := objects[0] // the object orbited by some other object A)B => A
		orbiter := objects[1] // the object orbiting some other object A)B => B
		orbits[orbiter] = orbitee

		orbitedBy[orbitee] = append(orbitedBy[orbitee], orbiter)
	}

	totalOrbits := 0
	// iterate over all objects that orbit another object
	for orbiter := range orbits {
		dist := distToCOM(orbiter)
		//fmt.Printf("Dist of %s to COM: %d\n", orbiter, dist)
		totalOrbits += dist
	}
	fmt.Printf("Total orbits: %d\n", totalOrbits)

	dist := walk("YOU", "SAN", 0)
	fmt.Printf("Dist to SAN: %d", dist-2)
}

func distToCOM(orbiter string) int {
	return distTo("COM", orbiter)
}

func distTo(obj, orbiter string) int {
	// will not happen because every object orbits exactly one other object according to problem statement
	//if _, ok := orbitMap[orbiter]; !ok {
	//	return -1
	//}
	dist := 0
	currObj := orbiter
	for currObj != obj {
		orbitee, ok := orbits[currObj]
		if !ok {
			return 0
		}
		currObj = orbitee
		dist++
	}
	return dist
}

func walk(curr, end string, currLengthFromStart int) int {
	if _, ok := visited[curr]; ok {
		return math.MaxInt32
	}

	if curr == end {
		return currLengthFromStart
	}

	// mark planet as visited
	visited[curr] = struct{}{}

	// walk through all objects that orbit curr
	distFromOrbiters := math.MaxInt32
	distFromOrbitees := math.MaxInt32
	if orbiters, hasOrbiters := orbitedBy[curr]; hasOrbiters {
		for _, orbiter := range orbiters {
			tmp := walk(orbiter, end, currLengthFromStart+1)
			distFromOrbiters = min(distFromOrbiters, tmp)
		}
	}
	// go to orbitee
	if orbitee, hasOrbitee := orbits[curr]; hasOrbitee {
		tmp := walk(orbitee, end, currLengthFromStart+1)
		distFromOrbitees = min(distFromOrbitees, tmp)
	}
	return min(distFromOrbitees, distFromOrbiters)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//func search(curr, end string) int {
//	if curr == end {
//		return 0
//	}
//	return 1 + search()
//}
