package day10

import (
	"image"
	"iter"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/directions"
)

func parseInput(input string) ([][]int, []image.Point, []image.Point, image.Rectangle) {
	grid := make([][]int, 0)
	for _, line := range strings.Split(input, "\n") {
		row := make([]int, 0)
		for _, char := range line {
			if char == '.' {
				row = append(row, -1)
			} else {
				row = append(row, util.ParseInt(string(char)))
			}
		}
		grid = append(grid, row)
	}
	trailheads := make([]image.Point, 0)
	destinations := make([]image.Point, 0)
	for y, row := range grid {
		for x, col := range row {
			if col == 0 {
				trailheads = append(trailheads, image.Pt(x, y))
			}
			if col == 9 {
				destinations = append(destinations, image.Pt(x, y))
			}
		}
	}

	return grid, trailheads, destinations, image.Rect(0, 0, len(grid[0]), len(grid))
}

func neighbours(grid [][]int, border image.Rectangle) util.NeighboursFunc[image.Point] {
	return func(p image.Point) iter.Seq[image.Point] {
		return func(yield func(image.Point) bool) {
			for neighbour := range directions.Moves() {
				np := p.Add(neighbour.Point())
				if np.In(border) && grid[np.Y][np.X] != -1 && grid[np.Y][np.X]-grid[p.Y][p.X] == 1 {
					if !yield(np) {
						return
					}
				}
			}
		}
	}
}

func Part1(input string) (string, error) {
	grid, trailheads, destinations, border := parseInput(input)

	var sum int
	for _, head := range trailheads {
		for _, dest := range destinations {
			if path := util.BFS(head, neighbours(grid, border), util.GoalPt(dest)); path != nil {
				sum++
			}
		}
	}
	return strconv.Itoa(sum), nil
}

func Part2(input string) (string, error) {
	grid, trailheads, destinations, border := parseInput(input)

	var sum int
	for _, head := range trailheads {
		for _, dest := range destinations {
			sum += len(util.AllPathsBFS(head, neighbours(grid, border), util.GoalPt(dest)))
		}
	}
	return strconv.Itoa(sum), nil
}
