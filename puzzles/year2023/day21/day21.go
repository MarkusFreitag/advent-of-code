package day21

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util/maputil"
	"github.com/MarkusFreitag/advent-of-code/util/sliceutil"
)

var steps = 64

func Part1(input string) (string, error) {
	lines := strings.Fields(input)
	grid := make([][]rune, len(lines))
	var start [2]int
	var found bool
	for idx, line := range lines {
		grid[idx] = []rune(line)
		if !found {
			for c, cell := range grid[idx] {
				if cell == 'S' {
					start[0] = idx
					start[1] = c
					found = true
				}
			}
		}
	}

	poss := make(map[[2]int]bool)
	poss[start] = true
	lastRound := make([][2]int, 0)
	lastRound = append(lastRound, start)
	for step := 0; step < steps; step++ {
		newPoss := make(map[[2]int]bool)
		queue := maputil.Keys(poss)
		for len(queue) > 0 {
			var pos [2]int
			pos, queue = sliceutil.PopFront(queue)
			for _, neigh := range [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
				pt := [2]int{pos[0] + neigh[0], pos[1] + neigh[1]}
				if pt[0] < 0 || pt[0] >= len(grid) || pt[1] < 0 || pt[1] >= len(grid[0]) {
					continue
				}
				if grid[pt[0]][pt[1]] == '#' {
					continue
				}
				newPoss[pt] = true
			}
		}
		poss = newPoss
	}
	return strconv.Itoa(len(poss)), nil
}

func Part2(input string) (string, error) {
	lines := strings.Fields(input)
	grid := make([][]rune, len(lines))
	var start [2]int
	var found bool
	for idx, line := range lines {
		grid[idx] = []rune(line)
		if !found {
			for c, cell := range grid[idx] {
				if cell == 'S' {
					start[0] = idx
					start[1] = c
					found = true
				}
			}
		}
	}

	poss := make(map[[2]int]bool)
	poss[start] = true
	lastRound := make([][2]int, 0)
	lastRound = append(lastRound, start)
	for step := 0; step < steps; step++ {
		newPoss := make(map[[2]int]bool)
		queue := maputil.Keys(poss)
		for len(queue) > 0 {
			var pos [2]int
			pos, queue = sliceutil.PopFront(queue)
			for _, neigh := range [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
				pt := [2]int{pos[0] + neigh[0], pos[1] + neigh[1]}
				if stone(grid, pt) {
					continue
				}
				newPoss[pt] = true
			}
		}
		poss = newPoss
	}
	return strconv.Itoa(len(poss)), nil
}

func stone(g [][]rune, pt [2]int) bool {
	if pt[0] < 0 {
		pt[0] = len(g) + (pt[0] % len(g))
	} else if pt[0] >= len(g) {
		pt[0] = pt[0] % len(g)
	}
	if pt[1] < 0 {
		pt[1] = len(g[0]) + (pt[1] % len(g[0])) - 1
	} else if pt[1] >= len(g[0]) {
		pt[1] = pt[1] % len(g[0])
	}
	return g[pt[0]][pt[1]] == '#'
}
