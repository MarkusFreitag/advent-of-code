package day25

import (
	"strconv"
	"strings"
)

func gridString(g [][]rune) string {
	lines := make([]string, len(g))
	for idx, row := range g {
		lines[idx] = string(row)
	}
	return strings.Join(lines, "\n")
}

func newGrid(width, height int) [][]rune {
	g := make([][]rune, height)
	for idx := range g {
		g[idx] = []rune(strings.Repeat(".", width))
	}
	return g
}

func Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")
	grid := make([][]rune, len(lines))
	for idx, line := range lines {
		grid[idx] = []rune(line)
	}

	var steps int
	for {
		nGrid := newGrid(len(grid[0]), len(grid))

		// >
		for y, row := range grid {
			for x, col := range row {
				if col == '>' {
					nx := x + 1
					if nx == len(row) {
						nx = 0
					}
					if grid[y][nx] == '.' {
						nGrid[y][nx] = '>'
					} else {
						nGrid[y][x] = '>'
					}
				}
			}
		}

		// v
		for y, row := range grid {
			for x, col := range row {
				if col == 'v' {
					ny := y + 1
					if ny == len(grid) {
						ny = 0
					}
					if grid[ny][x] == '.' && nGrid[ny][x] == '.' {
						nGrid[ny][x] = 'v'
					} else if grid[ny][x] == '>' && nGrid[ny][x] == '.' {
						nGrid[ny][x] = 'v'
					} else {
						nGrid[y][x] = 'v'
					}
				}
			}
		}

		steps++
		if gridString(grid) == gridString(nGrid) {
			break
		}
		grid = nGrid
	}

	return strconv.Itoa(steps), nil
}

func Part2(input string) (string, error) {
	return "not solved yet", nil
}
