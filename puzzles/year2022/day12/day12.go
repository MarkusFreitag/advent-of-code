package day12

import (
	"image"
	"iter"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/directions"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

func parseInput(input string) ([][]int, image.Point, image.Rectangle, util.GoalFunc[image.Point]) {
	var start, dest image.Point
	lines := strings.Split(input, "\n")
	grid := make([][]int, len(lines))
	for idx, line := range lines {
		if v := strings.Index(line, "S"); v != -1 {
			start = image.Pt(v, idx)
			line = strings.ReplaceAll(line, "S", "a")
		}
		if v := strings.Index(line, "E"); v != -1 {
			dest = image.Pt(v, idx)
			line = strings.ReplaceAll(line, "E", "z")
		}
		bs := []byte(line)
		grid[idx] = make([]int, len(bs))
		for i, b := range bs {
			grid[idx][i] = int(b)
		}
	}
	goalFn := func(p image.Point) bool { return p.Eq(dest) }
	return grid, start, image.Rect(0, 0, len(grid[0]), len(grid)), goalFn
}

func neighbours(grid [][]int, border image.Rectangle) util.NeighboursFunc[image.Point] {
	return func(p image.Point) iter.Seq[image.Point] {
		return func(yield func(image.Point) bool) {
			for neighbour := range directions.Moves() {
				np := p.Add(neighbour.Point())
				if np.In(border) && grid[np.Y][np.X]-grid[p.Y][p.X] <= 1 {
					if !yield(np) {
						return
					}
				}
			}
		}
	}
}

func Part1(input string) (string, error) {
	grid, pos, border, goalFn := parseInput(input)
	return strconv.Itoa(util.BFS(pos, neighbours(grid, border), goalFn).Dist()), nil
}

func Part2(input string) (string, error) {
	grid, _, border, goalFn := parseInput(input)
	dist := numbers.MaxInteger
	for y, row := range grid {
		for x, col := range row {
			if col == 97 {
				if sn := util.BFS(image.Pt(x, y), neighbours(grid, border), goalFn); sn != nil {
					dist = min(dist, sn.Dist())
				}
			}
		}
	}
	return strconv.Itoa(dist), nil
}
