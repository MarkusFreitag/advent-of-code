package day8

import (
	"math"
	"strconv"
	"strings"
)

type Part1 struct{}

func (p *Part1) Solve(input string) (string, error) {
	width := 25
	height := 6
	layerSize := width * height

	zeroCount := math.MaxInt64
	var layer string
	for i := 0; i < len(input); i += layerSize {
		l := input[i : i+layerSize]
		if zeros := strings.Count(l, "0"); zeros < zeroCount {
			layer = l
			zeroCount = zeros
		}
	}

	return strconv.Itoa(strings.Count(layer, "1") * strings.Count(layer, "2")), nil
}

type Part2 struct{}

func (p *Part2) Solve(input string) (string, error) {
	width := 25
	height := 6
	layerSize := width * height

	layers := make([][][]int, 0)
	for i := 0; i < len(input); i += layerSize {
		block := input[i : i+layerSize]
		layer := make([][]int, 0)
		for w := 0; w < len(block); w += width {
			row := make([]int, 0)
			for _, c := range block[w : w+width] {
				d, _ := strconv.Atoi(string(c))
				row = append(row, d)
			}
			layer = append(layer, row)
		}
		layers = append(layers, layer)
	}

	var output []string
	for h := 0; h < height; h++ {
		var row string
		for w := 0; w < width; w++ {
			for l := 0; l < len(layers); l++ {
				p := layers[l][h][w]
				if p == 2 {
					continue
				}
				if p == 0 {
					row += " "
				} else {
					row += "#"
				}
				break
			}
		}
		output = append(output, row)
	}

	return "\n" + strings.Join(output, "\n"), nil
}

func IntToStrSlice(nums []int) []string {
	s := make([]string, len(nums))
	for idx, num := range nums {
		s[idx] = strconv.Itoa(num)
	}
	return s
}
