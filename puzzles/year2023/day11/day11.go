package day11

import (
	"slices"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util/maputil"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

func scaledDistance(a, b [2]int, emptyRows, emptyCols []int, scale int) int {
	var dist int

	minRow, maxRow := numbers.MinMax(a[0], b[0])
	for i := minRow; i < maxRow; i++ {
		if slices.Contains(emptyRows, i) {
			dist += scale
		} else {
			dist++
		}
	}

	minCol, maxCol := numbers.MinMax(a[1], b[1])
	for i := minCol; i < maxCol; i++ {
		if slices.Contains(emptyCols, i) {
			dist += scale
		} else {
			dist++
		}
	}

	return dist
}

func solution(input string, scale int) int {
	lines := strings.Split(input, "\n")
	grid := make([][]rune, 0)
	emptyRows := make([]int, 0)
	for idx, line := range lines {
		grid = append(grid, []rune(line))
		if strings.Count(line, ".") == len(line) {
			emptyRows = append(emptyRows, idx)
		}
	}

	emptyCols := make([]int, 0)
	for x := 0; x < len(grid[0]); x++ {
		empty := true
		for y := 0; y < len(grid); y++ {
			if grid[y][x] != '.' {
				empty = false
				break
			}
		}
		if empty {
			emptyCols = append(emptyCols, x)
		}
	}

	galaxies := make([][2]int, 0)
	for y, row := range grid {
		for x, cell := range row {
			if cell == '#' {
				galaxies = append(galaxies, [2]int{y, x})
			}
		}
	}

	pairs := make(map[[2]int]int)
	for index := range galaxies {
		for idx := range galaxies {
			if index == idx {
				continue
			}

			dist := scaledDistance(galaxies[index], galaxies[idx], emptyRows, emptyCols, scale)

			min, max := numbers.MinMax(index, idx)
			key := [2]int{min, max}

			if d, ok := pairs[key]; ok {
				pairs[key] = numbers.Min(d, dist)
			} else {
				pairs[key] = dist
			}
		}
	}

	return numbers.Sum(maputil.Values(pairs)...)
}

func Part1(input string) (string, error) {
	return strconv.Itoa(solution(input, 2)), nil
}

func Part2(input string) (string, error) {
	return strconv.Itoa(solution(input, 1000000)), nil
}
