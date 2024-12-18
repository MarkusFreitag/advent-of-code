package day18

import (
	"image"
	"iter"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/directions"
)

var (
	size       = 70
	firstBytes = 1024
)

func neighbours(grid [][]rune) util.NeighboursFunc[image.Point] {
	border := image.Rect(0, 0, size+1, size+1)
	return func(p image.Point) iter.Seq[image.Point] {
		return func(yield func(image.Point) bool) {
			for neighbour := range directions.Moves() {
				np := p.Add(neighbour.Point())
				if np.In(border) && grid[np.Y][np.X] != '#' {
					if !yield(np) {
						return
					}
				}
			}
		}
	}
}

func Part1(input string) (string, error) {
	grid := make([][]rune, size+1)
	for idx := range grid {
		grid[idx] = []rune(strings.Repeat(".", size+1))
	}
	for _, line := range strings.Fields(input)[:firstBytes] {
		parts := strings.Split(line, ",")
		grid[util.ParseInt(parts[1])][util.ParseInt(parts[0])] = '#'
	}

	goalFn := func(p image.Point) bool { return p.Y == size && p.X == size }
	return strconv.Itoa(util.BFS(image.Pt(0, 0), neighbours(grid), goalFn).Dist()), nil
}

func Part2(input string) (string, error) {
	grid := make([][]rune, size+1)
	for idx := range grid {
		grid[idx] = []rune(strings.Repeat(".", size+1))
	}
	for _, line := range strings.Fields(input)[:firstBytes] {
		parts := strings.Split(line, ",")
		grid[util.ParseInt(parts[1])][util.ParseInt(parts[0])] = '#'
	}

	for _, line := range strings.Fields(input)[firstBytes:] {
		parts := strings.Split(line, ",")
		grid[util.ParseInt(parts[1])][util.ParseInt(parts[0])] = '#'

		goalFn := func(p image.Point) bool { return p.Y == size && p.X == size }
		if util.BFS(image.Pt(0, 0), neighbours(grid), goalFn) == nil {
			return line, nil
		}
	}
	return "", nil
}
