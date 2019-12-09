package main

import (
	"../../shared"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type LayerHistogram struct {
	LayerID     int
	digitCounts map[int]int
}

func main() {
	file, err := ioutil.ReadFile("./2019/day8/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	//
	input := string(file)
	//input := "0222112222120000"
	//input := "123456789012"
	width := 25
	height := 6

	code, err := shared.StrToDigitArr(input)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	layers := parseImageCode(width, height, code)
	layerHists := computeDigitHistogram(layers)
	layerHists = sortDescByDigit(layerHists, 0)

	// part 1
	answer := layerHists[0].digitCounts[1] * layerHists[0].digitCounts[2]
	fmt.Println(answer)

	// part 2
	imgCode := render(layers)

	var decodedCode strings.Builder
	for i := range imgCode {
		for j := range imgCode[i] {
			decodedCode.WriteString(strconv.Itoa(imgCode[i][j]))
		}
	}
	fmt.Printf("Answer 2: decoded image: %s\n", decodedCode.String())

	var pixels []uint8
	for i := range imgCode {
		for j := range imgCode[i] {
			pixels = append(pixels, blackOrWhite(imgCode[i][j]))
		}
	}
	rendered := image.NewGray(image.Rect(0, 0, width, height))
	rendered.Pix = pixels
	f, _ := os.Create("./2019/day8/result.png")
	defer f.Close()

	png.Encode(f, rendered)
	//
	//coutLayerZeros := countDigits(layers)
	//layerWithFewestZeros := layers[coutLayerZeros[0]]
	//num1Digits :=
	//
	//	fmt.Printf("%+v", layerWithFewestZeros)
}

func blackOrWhite(i int) uint8 {
	return uint8(i * 255)
}

// result is in layers[0]
func render(layers [][][]int) [][]int {
	// skip first layer because nothing to compare
	for layerID := 1; layerID <= len(layers)-1; layerID++ {
		layer := layers[layerID]
		for h := 0; h <= len(layer)-1; h++ {
			for w := 0; w <= len(layer[h])-1; w++ {
				pixelInImage := layers[0][h][w]
				if pixelInImage != 2 {
					continue
				}
				currPixel := layers[layerID][h][w]
				layers[0][h][w] = currPixel
			}
		}
	}
	return layers[0]
}

func sortDescByDigit(hists []LayerHistogram, digit int) []LayerHistogram {
	arr := clone(hists)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].digitCounts[digit] < arr[j].digitCounts[digit]
	})
	return arr
}

func clone(hists []LayerHistogram) []LayerHistogram {
	var out []LayerHistogram

	for _, s := range hists {
		newMap := make(map[int]int)
		for k, v := range s.digitCounts {
			newMap[k] = v
		}
		out = append(out, LayerHistogram{
			LayerID:     s.LayerID,
			digitCounts: newMap,
		})
	}
	return out
}

func computeDigitHistogram(layers [][][]int) []LayerHistogram {
	var hists []LayerHistogram
	for layerID, layer := range layers {
		digitCounts := make(map[int]int)
		for _, row := range layer {
			for _, digit := range row {
				digitCounts[digit]++
			}
		}
		hists = append(hists, LayerHistogram{
			LayerID:     layerID,
			digitCounts: digitCounts,
		})
	}
	return hists
}

func parseImageCode(width, height int, code []int) [][][]int {
	// [layers][rows][cols]
	numLayers := len(code) / (width * height)
	layerSize := width * height
	var layers [][][]int

	for l := 0; l <= numLayers-1; l++ {
		var layer [][]int
		for h := 0; h <= height-1; h++ {
			var row []int
			for w := 0; w <= width-1; w++ {
				pos := (l * layerSize) + h*width + w
				row = append(row, code[pos])
			}
			layer = append(layer, row)
		}
		layers = append(layers, layer)
	}
	return layers
}
