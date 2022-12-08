package day8

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

func parseInput(input string) [][]int {
	lines := strings.Split(input, "\n")
	grid := make([][]int, len(lines))
	for oIdx, line := range lines {
		row := make([]int, len(line))
		for iIdx, char := range line {
			row[iIdx] = util.ParseInt(string(char))
		}
		grid[oIdx] = row
	}
	return grid
}

func Part1(input string) (string, error) {
	grid := parseInput(input)

	rowLen, colLen := len(grid[0]), len(grid)
	visible := 2*rowLen + 2*colLen - 4

	for y := 1; y < colLen-1; y++ {
		for x := 1; x < rowLen-1; x++ {
			tree := grid[y][x]

			{ // look left
				if numbers.Max(grid[y][:x]...) < tree {
					visible++
					continue
				}
			}

			{ // look up
				trees := make([]int, 0)
				for i := 0; i < y; i++ {
					trees = append(trees, grid[i][x])
				}
				max := numbers.Max(trees...)
				if max < tree {
					visible++
					continue
				}
			}

			{ // look right
				if numbers.Max(grid[y][x+1:]...) < tree {
					visible++
					continue
				}
			}

			{ // look down
				trees := make([]int, 0)
				for i := y + 1; i < colLen; i++ {
					trees = append(trees, grid[i][x])
				}
				max := numbers.Max(trees...)
				if max < tree {
					visible++
					continue
				}
			}
		}
	}

	return strconv.Itoa(visible), nil
}

func Part2(input string) (string, error) {
	grid := parseInput(input)

	var highest int
	rowLen, colLen := len(grid[0]), len(grid)

	for y := 1; y < colLen-1; y++ {
		for x := 1; x < rowLen-1; x++ {
			tree := grid[y][x]
			scores := make([]int, 4)

			{ // look up
				var count int
				for i := y - 1; i >= 0; i-- {
					count++
					if grid[i][x] >= tree {
						break
					}
				}
				scores[0] = count
			}

			{ // look left
				var count int
				for i := x - 1; i >= 0; i-- {
					count++
					if grid[y][i] >= tree {
						break
					}
				}
				scores[1] = count
			}

			{ // look down
				var count int
				for i := y + 1; i < colLen; i++ {
					count++
					if grid[i][x] >= tree {
						break
					}
				}
				scores[2] = count
			}

			{ // look right
				var count int
				for i := x + 1; i < rowLen; i++ {
					count++
					if grid[y][i] >= tree {
						break
					}
				}
				scores[3] = count
			}

			if score := numbers.Multiply(scores...); score > highest {
				highest = score
			}
		}
	}

	return strconv.Itoa(highest), nil
}
