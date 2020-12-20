package day20

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

const (
	TOP = iota
	RIGHT
	BOTTOM
	LEFT
)

type Grid [][]rune

func parseGrid(layout string) Grid {
	grid := make(Grid, 0)
	for _, line := range strings.Split(layout, "\n") {
		row := make([]rune, len(line))
		for idx, char := range line {
			row[idx] = rune(char)
		}
		grid = append(grid, row)
	}
	return grid
}

func (g Grid) Border(dir int) string {
	switch dir {
	case TOP:
		return string(g[0])
	case RIGHT:
		var border string
		for _, row := range g {
			border += string(row[len(row)-1])
		}
		return border
	case BOTTOM:
		return string(g[len(g)-1])
	case LEFT:
		var border string
		for _, row := range g {
			border += string(row[0])
		}
		return border
	}
	return ""
}

var rgxHead = regexp.MustCompile(`^Tile\s(\d+):`)

func Part1(input string) (string, error) {
	tiles := make(map[int]Grid)
	for _, block := range strings.Split(input, "\n\n") {
		head := rgxHead.FindAllStringSubmatch(block, -1)[0]
		id, _ := strconv.Atoi(head[1])
		tiles[id] = parseGrid(strings.Split(block, ":")[1])
	}
	size := int(math.Sqrt(float64(len(tiles))))
	layout := make([][]int, size)
	for idx := range layout {
		layout[idx] = make([]int, size)
	}

	return strconv.Itoa(layout[0][0] * layout[0][len(layout[0])-1] * layout[len(layout)-1][0] * layout[len(layout)-1][len(layout[0])-1]), nil
}

func Part2(input string) (string, error) {
	return "not solved yet", nil
}
