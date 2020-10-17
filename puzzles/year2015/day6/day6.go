package day6

import (
	"strconv"
	"strings"
)

type grid [][]int

func newGrid(rows, cols int) grid {
	g := make(grid, rows)
	for idx := range g {
		g[idx] = make([]int, cols)
	}
	return g
}

func (g grid) Count(state int) int {
	var counter int
	for _, r := range g {
		for _, c := range r {
			if c == state {
				counter++
			}
		}
	}
	return counter
}

func (g grid) Sum() int {
	var sum int
	for _, r := range g {
		for _, c := range r {
			sum += c
		}
	}
	return sum
}

func splitNums(s string) []int {
	parts := strings.Split(s, ",")
	nums := make([]int, 2)
	nums[0], _ = strconv.Atoi(parts[0])
	nums[1], _ = strconv.Atoi(parts[1])
	return nums
}

func parseLine(str string) (string, []int, []int) {
	str = strings.TrimPrefix(str, "turn ")
	parts := strings.Split(str, " ")
	return parts[0], splitNums(parts[1]), splitNums(parts[3])
}

func Part1(input string) (string, error) {
	g := newGrid(1000, 1000)
	for _, line := range strings.Split(input, "\n") {
		instr, start, end := parseLine(line)
		for i := start[0]; i <= end[0]; i++ {
			for j := start[1]; j <= end[1]; j++ {
				switch instr {
				case "on":
					g[i][j] = 1
				case "off":
					g[i][j] = 0
				case "toggle":
					if g[i][j] == 0 {
						g[i][j] = 1
					} else {
						g[i][j] = 0
					}
				}
			}
		}
	}
	return strconv.Itoa(g.Count(1)), nil
}

func Part2(input string) (string, error) {
	g := newGrid(1000, 1000)
	for _, line := range strings.Split(input, "\n") {
		instr, start, end := parseLine(line)
		for i := start[0]; i <= end[0]; i++ {
			for j := start[1]; j <= end[1]; j++ {
				switch instr {
				case "on":
					g[i][j] = g[i][j] + 1
				case "off":
					if g[i][j] > 0 {
						g[i][j] = g[i][j] - 1
					}
				case "toggle":
					g[i][j] = g[i][j] + 2
				}
			}
		}
	}
	return strconv.Itoa(g.Sum()), nil
}
