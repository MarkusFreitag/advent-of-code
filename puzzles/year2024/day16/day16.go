package day16

import (
	"image"
	"iter"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/directions"
)

func parseInput(input string) ([][]rune, image.Point, image.Point) {
	grid := make([][]rune, 0)
	var start, dest image.Point
	for y, line := range strings.Split(input, "\n") {
		for x, char := range line {
			if char == 'E' {
				dest = image.Pt(x, y)
			}
			if char == 'S' {
				start = image.Pt(x, y)
			}
		}
		line = strings.ReplaceAll(line, "E", ".")
		line = strings.ReplaceAll(line, "S", ".")
		grid = append(grid, []rune(line))
	}
	return grid, start, dest
}

func neighboursCost(grid [][]rune) util.NeighboursCostFunc[Reindeer] {
	border := image.Rect(0, 0, len(grid[0]), len(grid))
	return func(rd Reindeer) iter.Seq2[Reindeer, int] {
		return func(yield func(Reindeer, int) bool) {
			for neighbour := range directions.Moves() {
				nrd := Reindeer{Pos: rd.Pos.Add(neighbour.Point()), Dir: neighbour}
				if directions.Opposite(rd.Dir, nrd.Dir) {
					continue
				}
				if nrd.Pos.In(border) && grid[nrd.Pos.Y][nrd.Pos.X] == '.' {
					cost := 1
					if rd.Dir != nrd.Dir {
						cost += 1000
					}
					if !yield(nrd, cost) {
						return
					}
				}
			}
		}
	}
}

type Reindeer struct {
	Pos image.Point
	Dir directions.Move
}

func goal(dest image.Point) util.GoalFunc[Reindeer] {
	return func(rd Reindeer) bool { return rd.Pos.Eq(dest) }
}

func Part1(input string) (string, error) {
	grid, start, dest := parseInput(input)

	reindeer := Reindeer{Pos: start, Dir: directions.Right}

	shortestPath := util.Dijkstra(reindeer, neighboursCost(grid), goal(dest))

	return strconv.Itoa(shortestPath.Cost), nil
}

func Part2(input string) (string, error) {
	grid, start, dest := parseInput(input)

	reindeer := Reindeer{Pos: start, Dir: directions.Right}

	shortestPath := util.Dijkstra(reindeer, neighboursCost(grid), goal(dest))

	paths := util.AllPathsWithScore(reindeer, neighboursCost(grid), goal(dest), shortestPath.Cost)

	paths = append(paths, shortestPath)

	seenPts := make(map[image.Point]struct{})
	for _, p := range paths {
		for rd := range p.Seq() {
			seenPts[rd.Pos] = struct{}{}
		}
	}
	return strconv.Itoa(len(seenPts)), nil
}
