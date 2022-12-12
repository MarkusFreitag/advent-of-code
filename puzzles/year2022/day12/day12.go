package day12

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util/numbers"
	"github.com/MarkusFreitag/advent-of-code/util/slice"
)

func parseInput(input string) ([][]int, [2]int, [2]int) {
	var start, dest [2]int
	lines := strings.Split(input, "\n")
	grid := make([][]int, len(lines))
	for idx, line := range lines {
		if v := strings.Index(line, "S"); v != -1 {
			start[0], start[1] = idx, v
			line = strings.ReplaceAll(line, "S", "a")
		}
		if v := strings.Index(line, "E"); v != -1 {
			dest[0], dest[1] = idx, v
			line = strings.ReplaceAll(line, "E", "z")
		}
		bs := []byte(line)
		grid[idx] = make([]int, len(bs))
		for i, b := range bs {
			grid[idx][i] = int(b)
		}
	}
	return grid, start, dest
}

type item struct {
	pos  [2]int
	dist int
}

func bfs(g [][]int, start, end [2]int) int {
	q := make([]item, 0)
	q = append(q, item{start, 0})
	seen := make(map[[2]int]bool)
	for len(q) > 0 {
		var i item
		i, q = slice.PopFront(q)
		if i.pos == end {
			return i.dist
		}
		if _, ok := seen[i.pos]; ok {
			continue
		}
		seen[i.pos] = true
		y, x := i.pos[0], i.pos[1]
		for _, off := range [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			ny, nx := y+off[0], x+off[1]
			if 0 <= nx && nx < len(g[0]) && 0 <= ny && ny < len(g) && g[ny][nx]-g[y][x] <= 1 {
				q = append(q, item{[2]int{ny, nx}, i.dist + 1})
			}
		}

	}
	return numbers.MaxInteger
}

func Part1(input string) (string, error) {
	grid, pos, dest := parseInput(input)
	steps := bfs(grid, pos, dest)
	return strconv.Itoa(steps), nil
}

func Part2(input string) (string, error) {
	grid, _, dest := parseInput(input)
	starts := make([][2]int, 0)
	for y, row := range grid {
		for x, col := range row {
			if col == 97 {
				starts = append(starts, [2]int{y, x})
			}
		}
	}

	paths := make([]int, len(starts))
	for idx, start := range starts {
		paths[idx] = bfs(grid, start, dest)
	}

	return strconv.Itoa(numbers.Min(paths...)), nil
}
