package main

type Coord struct {
	X   int
	Y   int
	Len int
}

type Intersection struct {
	P1 Coord
	P2 Coord
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
