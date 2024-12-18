package day18

import (
	"image"
	"iter"
	"slices"
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

	return strconv.Itoa(
		util.BFS(image.Pt(0, 0), neighbours(grid), util.GoalPt(image.Pt(size, size))).Dist(),
	), nil
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

	path := util.BFS(image.Pt(0, 0), neighbours(grid), util.GoalPt(image.Pt(size, size)))
	pts := slices.Collect(path.Seq())

	for _, line := range strings.Fields(input)[firstBytes:] {
		parts := strings.Split(line, ",")
		pt := image.Pt(util.ParseInt(parts[0]), util.ParseInt(parts[1]))
		grid[pt.Y][pt.X] = '#'

		/*
			We need to only check whether the exit is still accessible when a byte falls onto the
			current path. This way we cut down the number of searches significantly.
		*/
		if slices.Contains(pts, pt) {
			path := util.BFS(image.Pt(0, 0), neighbours(grid), util.GoalPt(image.Pt(size, size)))
			if path == nil {
				return line, nil
			}
			pts = slices.Collect(path.Seq())
		}
	}
	return "", nil
}
