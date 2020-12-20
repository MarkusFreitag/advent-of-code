package day20

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

const (
	TOP = iota
	RIGHT
	BOTTOM
	LEFT
)

var SIDES = []int{TOP, RIGHT, BOTTOM, LEFT}

type Grid [][]rune

func parseGrid(layout string) Grid {
	layout = strings.TrimSpace(layout)
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

func (g Grid) Rotate() Grid {
	grid := make(Grid, 0)
	for i := 0; i < len(g[0]); i++ {
		row := make([]rune, 0)
		for j := len(g) - 1; j >= 0; j-- {
			row = append(row, g[j][i])
		}
		grid = append(grid, row)
	}
	return grid
}

func (g Grid) FlipVertical() Grid {
	grid := make(Grid, 0)
	for idx := len(g) - 1; idx >= 0; idx-- {
		grid = append(grid, g[idx])
	}
	return grid
}

func (g Grid) FlipHorizontal() Grid {
	grid := make(Grid, 0)
	for _, row := range g {
		grid = append(grid, []rune(util.ReverseStr(string(row))))
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

func (g Grid) Show() {
	for _, row := range g {
		fmt.Println(string(row))
	}
}

var rgxHead = regexp.MustCompile(`^Tile\s(\d+):`)

func Part1(input string) (string, error) {
	grids := make(map[int]Grid)
	for _, block := range strings.Split(input, "\n\n") {
		head := rgxHead.FindAllStringSubmatch(block, -1)[0]
		id, _ := strconv.Atoi(head[1])
		grids[id] = parseGrid(strings.Split(block, ":")[1])
	}

	allSides := make([]string, 0)
	for _, grid := range grids {
		for _, dir := range SIDES {
			side := grid.Border(dir)
			allSides = append(allSides, side)
			allSides = append(allSides, util.ReverseStr(side))
		}
	}

	edges := make([]int, 0)
	for id, grid := range grids {
		var count int
		for _, dir := range SIDES {
			if util.CountStrInSlice(grid.Border(dir), allSides) == 1 {
				count++
			}
		}
		if count == 2 {
			edges = append(edges, id)
		}
	}

	return strconv.Itoa(util.MulInts(edges...)), nil
}

func Part2(input string) (string, error) {
	return "not solved yet", nil
}
