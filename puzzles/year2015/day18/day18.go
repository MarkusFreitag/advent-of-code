package day18

import (
	"strconv"
	"strings"
)

type Grid [][]bool

func NewGrid(size int) Grid {
	grid := make(Grid, size)
	for idx := range grid {
		grid[idx] = make([]bool, size)
	}
	return grid
}

func NewGridFromString(str string) Grid {
	lines := strings.Split(str, "\n")
	grid := make(Grid, len(lines))
	for idx, line := range lines {
		grid[idx] = make([]bool, len(line))
		for i, r := range line {
			var state bool
			if r == '#' {
				state = true
			}
			grid[idx][i] = state
		}
	}
	return grid
}

func (g Grid) Next() Grid {
	next := NewGrid(len(g))
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			neighsOn := g.Neighbours(i, j, true)
			if g[i][j] {
				if neighsOn == 2 || neighsOn == 3 {
					next[i][j] = true
				}
			} else {
				if neighsOn == 3 {
					next[i][j] = true
				}
			}
		}
	}
	return next
}

func (g Grid) Neighbours(y, x int, state bool) int {
	var count int
	for i := -1; i <= 1; i++ {
		row := y + i
		if row < 0 || row >= len(g) {
			continue
		}
		for j := -1; j <= 1; j++ {
			col := x + j
			if col < 0 || col >= len(g[row]) {
				continue
			}
			if row == y && col == x {
				continue
			}
			if g[row][col] == state {
				count++
			}
		}
	}
	return count
}

func (g Grid) CountState(state bool) int {
	var count int
	for _, row := range g {
		for _, col := range row {
			if col == state {
				count++
			}
		}
	}
	return count
}

func (g Grid) String() string {
	lines := make([]string, 0)
	for _, row := range g {
		var line string
		for _, col := range row {
			if col {
				line += "#"
			} else {
				line += "."
			}
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}

func Part1(input string) (string, error) {
	grid := NewGridFromString(input)
	for i := 0; i < 100; i++ {
		grid = grid.Next()
	}
	return strconv.Itoa(grid.CountState(true)), nil
}

func Part2(input string) (string, error) {
	grid := NewGridFromString(input)
	grid[0][0] = true
	grid[0][len(grid[0])-1] = true
	grid[len(grid)-1][0] = true
	grid[len(grid)-1][len(grid[0])-1] = true
	for i := 0; i < 100; i++ {
		grid = grid.Next()
		grid[0][0] = true
		grid[0][len(grid[0])-1] = true
		grid[len(grid)-1][0] = true
		grid[len(grid)-1][len(grid[0])-1] = true
	}
	return strconv.Itoa(grid.CountState(true)), nil
}
