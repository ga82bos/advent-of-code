package main

// Advent of Code 2019 - Day 3

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := []string{
		//"R8,U5,L5,D3",
		//"U7,R6,D4,L4",
		"R999,U626,R854,D200,R696,D464,R54,D246,L359,U57,R994,D813,L889,U238,L165,U970,L773,D904,L693,U512,R126,D421,R732,D441,R453,D349,R874,D931,R103,D794,R934,U326,L433,D593,L984,U376,R947,U479,R533,U418,R117,D395,L553,D647,R931,D665,L176,U591,L346,D199,L855,D324,L474,U251,R492,D567,L97,D936,L683,U192,R198,U706,L339,U66,R726,D102,R274,U351,R653,D602,L695,U921,R890,D654,R981,U351,R15,U672,R856,D319,R102,D234,R248,U169,L863,U375,L412,U75,L511,U298,L303,U448,R445,U638,L351,D312,R768,D303,L999,D409,L746,U266,L16,U415,L951,D763,L976,U342,L505,U770,L228,D396,L992,U3,R243,D794,L496,U611,R587,U772,L306,D119,L470,D490,L336,U518,L734,D654,L150,U581,L874,U691,L243,U94,L9,D582,L402,U563,R468,U96,L311,D10,R232,U762,R630,D1,L674,U685,R240,D907,R394,U703,L64,U397,L810,D272,L996,D954,R797,U789,R790,D526,R103,D367,R143,D41,L539,D735,R51,D172,L33,U241,R814,D981,R748,D699,L716,U647,L381,D351,L381,D121,L52,U601,R515,U713,L404,U45,R362,U670,L235,U102,R373,U966,L362,U218,R280,U951,R371,U378,L10,U670,R958,D423,L740,U888,R235,U899,L387,U167,R392,D19,L330,D916,R766,D471,L708,D83,R749,D696,L50,D159,R828,U479,L980,D613,L182,D875,L307,U472,L317,U999,R435,D364,R737,U550,L233,U190,L501,U610,R433,U470,L801,U52,L393,D596,L378,U220,L967,D807,R357,D179,L731,D54,L804,D865,L994,D151,L181,U239,R794,D378,L487,U408,R817,U809,R678,D599,L564,U480,R525,D189,L641,D771,L514,U72,L248,D334,L859,D318,R590,D571,R453,U732,R911,U632,R992,D80,R490,D234,L710,U816,L585,U180,L399,D238,L103,U605,R993,D539,R330",
		"L996,U383,L962,U100,L836,D913,R621,U739,R976,D397,L262,D151,L12,U341,R970,U123,L713,U730,L52,D223,L190,D81,R484,D777,R374,U755,R640,D522,R603,D815,R647,U279,R810,U942,R314,D19,L938,U335,R890,U578,R273,U338,R186,D271,L230,U90,R512,U672,R666,D328,L970,U17,R368,D302,L678,D508,L481,U12,L783,D409,L315,D579,L517,D729,R961,D602,R253,D746,R418,D972,R195,D270,L46,D128,L124,U875,R632,D788,L576,U695,R159,U704,R599,D597,R28,D703,L18,D879,L417,U633,L56,U302,R289,U916,R820,D55,R213,U712,R250,D265,L935,D171,L680,U738,L361,D939,R547,D606,L255,U880,R968,U255,R902,D624,L251,U452,L412,D60,L996,D140,L971,U196,R796,D761,L54,U54,L98,D758,L521,U578,L861,U365,L901,D495,L234,D124,L121,D329,L38,U481,L491,D938,L840,D311,L993,D954,R654,U925,L528,D891,L994,D681,L879,D476,L933,U515,L292,U626,R348,D963,L145,U230,L114,D11,R651,D929,R318,D672,R125,D827,L590,U338,L755,D925,L577,D52,R131,D465,R657,D288,R22,D363,R162,D545,L904,D457,R987,D389,L566,D931,L773,D53,R162,U271,L475,U666,L594,U733,R279,D847,R359,U320,R450,D704,L698,D173,R35,D267,L165,D66,L301,U879,R862,U991,R613,D489,L326,D393,R915,U718,R667,U998,R554,U199,R300,U693,R753,U938,R444,U12,L844,D912,R297,D668,R366,U710,L821,U384,R609,D493,R233,U898,R407,U683,R122,U790,L1,U834,L76,U572,R220,U752,L728,D85,L306,D805,R282,U507,R414,D687,L577,U174,L211,U308,L15,U483,R741,D828,L588,D192,L409,D605,L931,U260,L239,D424,L846,U429,L632,U122,L266,D544,R248,U188,R465,U721,R621,U3,L884,U361,L322,U504,R999,U381,R327,U555,L467,D849,R748,U175,R356",
	}

	dir2Vec := map[string]Coord{
		"R": {
			X: 1,
			Y: 0,
		},
		"L": {
			X: -1,
			Y: 0,
		},
		"U": {
			X: 0,
			Y: 1,
		},
		"D": {
			X: 0,
			Y: -1,
		},
	}

	paths := make([][]string, len(input))
	for i, p := range input {
		paths[i] = strings.Split(p, ",")
	}

	var pathCoords [][]Coord
	var intersections []Intersection

	for _, path := range paths {
		coords := pathToCoordList(path, dir2Vec)
		//fmt.Printf("%+v", coords)

		// find coords of other paths that have the same x and y value => intersection
		for _, cP1 := range coords {
			for _, p := range pathCoords {
				for _, cP2 := range p {
					if cP1.X != cP2.X || cP1.Y != cP2.Y {
						continue
					}
					intersections = append(intersections, Intersection{
						P1: cP1,
						P2: cP2,
					})
				}
			}
		}
		pathCoords = append(pathCoords, coords)
	}

	// Solution Part1: sort by manhattan distance
	sort.Slice(intersections, func(i, j int) bool {
		return distToCenter(intersections[i].P1.X, intersections[i].P1.Y) < distToCenter(intersections[j].P1.X, intersections[j].P1.Y)
	})

	fmt.Println("Part1: manhattan distance (first is solution):")
	for i, is := range intersections {
		fmt.Printf("%d. %d\n", i+1, distToCenter(is.P1.X, is.P1.Y))
	}

	fmt.Println()

	// Solution Part2: sort by length of path to intersection
	sort.Slice(intersections, func(i, j int) bool {
		return intersections[i].P1.Len+intersections[i].P2.Len < intersections[j].P1.Len+intersections[j].P2.Len
	})

	fmt.Println("Part2: smallest path to intersection (first is solution):")
	for i, is := range intersections {
		fmt.Printf("%d. %d\n", i+1, is.P1.Len+is.P2.Len)
	}
}

func distToCenter(x, y int) int {
	return abs(x) + abs(y)
}

func pathToCoordList(path []string, dir2Vec map[string]Coord) []Coord {
	var pathCoords []Coord

	x, y, l := 0, 0, 0
	for _, v := range path {
		dirVec := dir2Vec[string(v[0])]
		length, _ := strconv.Atoi(v[1:])

		for i := 0; i < length; i++ {
			x += dirVec.X
			y += dirVec.Y
			l += 1
			pathCoords = append(pathCoords, Coord{
				X:   x,
				Y:   y,
				Len: l,
			})
		}

	}
	return pathCoords
}

func printCoords(coords []Coord) {
	type maxDistCenter struct {
		up    int
		down  int
		left  int
		right int
	}

	dis := maxDistCenter{}

	// get max distances from center in each direction
	for _, c := range coords {
		dis.up = max(dis.up, c.Y)
		dis.down = min(dis.down, c.Y)

		dis.right = max(dis.right, c.X)
		dis.left = min(dis.left, c.X)
	}
	dis.up = abs(dis.up)
	dis.down = abs(dis.down)

	dis.right = abs(dis.right)
	dis.left = abs(dis.left)

	rows := max(dis.down+dis.up, 1)    // we must have at least 1 row
	cols := max(dis.left+dis.right, 1) // same as row

	arr := make([][]rune, rows+1) // add 1 to rows because x=1 means we need 2 rows (0,1)
	for i := range arr {
		arr[i] = make([]rune, cols+1) // same as rows
		for j := range arr[i] {
			arr[i][j] = '.'
		}
	}

	// arr[rows][cols]
	// arr[Y][X]
	centerX := dis.left
	centerY := dis.down

	for _, c := range coords {
		arr[centerY+c.Y][centerX+c.X] = '*' // coords are relative to center
	}

	arr[centerY][centerX] = 'O'

	fmt.Println()
	for i := range arr {
		for j := range arr[i] {
			val := arr[len(arr)-1-i][j]
			fmt.Printf("%s", string(val))
		}
		fmt.Println()
	}
}
