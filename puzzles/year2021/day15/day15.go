package day15

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

type Point struct {
	X, Y int
}

func neighbours(g [][]int, p Point) []Point {
	neighs := make([]Point, 0)

	if p.Y-1 > 0 {
		neighs = append(neighs, Point{X: p.X, Y: p.Y - 1})
	}

	if p.X+1 < len(g[p.Y]) {
		neighs = append(neighs, Point{X: p.X + 1, Y: p.Y})
	}

	if p.Y+1 < len(g) {
		neighs = append(neighs, Point{X: p.X, Y: p.Y + 1})
	}

	if p.X-1 > 0 {
		neighs = append(neighs, Point{X: p.X - 1, Y: p.Y})
	}

	return neighs
}

func minPathScore(grid [][]int) int {
	costGrid := make([][]int, len(grid))
	for y, row := range grid {
		costGrid[y] = make([]int, len(row))
		for x := range row {
			costGrid[y][x] = util.MaxInteger
		}
	}
	costGrid[0][0] = 0

	todo := make([]Point, 1)
	todo[0] = Point{}

	for len(todo) > 0 {
		var p Point
		p, todo = todo[0], todo[1:]
		for _, neigh := range neighbours(grid, p) {
			neighRisk := costGrid[neigh.Y][neigh.X]
			newNeighRisk := costGrid[p.Y][p.X] + grid[neigh.Y][neigh.X]
			if neighRisk > newNeighRisk {
				costGrid[neigh.Y][neigh.X] = newNeighRisk
				todo = append(todo, neigh)
			}
		}
	}
	return costGrid[len(costGrid)-1][len(costGrid[0])-1]
}

func Part1(input string) (string, error) {
	lines := strings.Fields(input)
	grid := make([][]int, len(lines))
	for idx, line := range lines {
		grid[idx] = util.StrsToInts(util.StrToStrs(line))
	}

	return strconv.Itoa(minPathScore(grid)), nil
}

func Part2(input string) (string, error) {
	lines := strings.Fields(input)
	grid := make([][]int, len(lines))
	for idx, line := range lines {
		grid[idx] = util.StrsToInts(util.StrToStrs(line))
	}

	// expand the grid 5 times in each direction
	height, width := len(grid), len(grid[0])
	bigGrid := make([][]int, height*5)
	for y := range bigGrid {
		bigGrid[y] = make([]int, width*5)
	}

	for y, row := range bigGrid {
		for x := range row {
			risk := grid[y%height][x%width]
			offset := y/height + x/width
			bigGrid[y][x] = (risk+offset-1)%9 + 1
		}
	}

	return strconv.Itoa(minPathScore(bigGrid)), nil
}
