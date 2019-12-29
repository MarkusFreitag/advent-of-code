package day24

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

type Grid [][]string

func (g Grid) String() string {
	lines := make([]string, len(g))
	for idx, row := range g {
		lines[idx] = strings.Join(row, "")
	}
	return strings.Join(lines, "\n")
}

func (g Grid) CountNeighbours(y, x int) int {
	var counter int
	for i := -1; i <= 1; i++ {
		if y+i < 0 || y+i >= len(g) {
			continue
		}
		for j := -1; j <= 1; j++ {
			if x+j < 0 || x+j >= len(g[0]) {
				continue
			}
			if abs(i)+abs(j) == 2 || abs(i)+abs(j) == 0 {
				continue
			}
			if g[y+i][x+j] == "#" {
				counter++
			}
		}
	}

	return counter
}

func (g Grid) Copy() Grid {
	buf := make(Grid, len(g))
	for idx, row := range g {
		buf[idx] = make([]string, len(row))
		for i, c := range row { // nolint: gosimple
			buf[idx][i] = c
		}
	}
	return buf
}

func (g Grid) Rating() int {
	var rating int
	for idx, r := range strings.ReplaceAll(g.String(), "\n", "") {
		if string(r) == "#" {
			rating += int(math.Pow(2, float64(idx)))
		}
	}
	return rating
}

func Part1(input string) (string, error) {
	grid := make(Grid, 0)
	for _, line := range strings.Split(input, "\n") {
		row := make([]string, len(line))
		for i, c := range line {
			row[i] = string(c)
		}
		grid = append(grid, row)
	}

	fmt.Printf("Initial state:\n%s\n\n", grid.String())

	states := make(map[string]bool)
	states[grid.String()] = true
	round := 1
	for {
		buf := grid.Copy()
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				count := grid.CountNeighbours(y, x)
				if grid[y][x] == "#" {
					if count != 1 {
						buf[y][x] = "."
					}
				}
				if grid[y][x] == "." {
					if count == 1 || count == 2 {
						buf[y][x] = "#"
					}
				}
			}
		}
		grid = buf
		fmt.Printf("After %d minutes:\n%s\n\n", round, grid.String())
		if _, ok := states[grid.String()]; ok {
			return strconv.Itoa(grid.Rating()), nil
		}
		states[grid.String()] = true
		round++
	}

}

func Part2(input string) (string, error) {
	return "n/a", nil
}
