package day20

import (
	"image"
	"iter"
	"slices"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/directions"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

func neighbours(grid [][]rune, char rune) util.NeighboursFunc[image.Point] {
	border := image.Rect(0, 0, len(grid[0]), len(grid))
	return func(p image.Point) iter.Seq[image.Point] {
		return func(yield func(image.Point) bool) {
			for neighbour := range directions.Moves() {
				np := p.Add(neighbour.Point())
				if np.In(border) && grid[np.Y][np.X] == char {
					if !yield(np) {
						return
					}
				}
			}
		}
	}
}

func parseInput(input string) ([][]rune, image.Point, image.Point) {
	grid := make([][]rune, 0)
	var start, end image.Point
	for y, row := range strings.Fields(input) {
		for x, col := range row {
			if col == 'S' {
				start = image.Pt(x, y)
			}
			if col == 'E' {
				end = image.Pt(x, y)
			}
		}
		grid = append(grid, []rune(row))
	}
	grid[start.Y][start.X] = '.'
	grid[end.Y][end.X] = '.'
	return grid, start, end
}

func solveMaze(grid [][]rune, start, end image.Point) []image.Point {
	path := slices.Collect(util.BFS(start, neighbours(grid, '.'), util.GoalPt(end)).Seq())
	slices.Reverse(path)
	return path
}

func findCheats(grid [][]rune, path []image.Point, maxCheatDist int) map[int]int {
	border := image.Rect(0, 0, len(grid[0]), len(grid))
	saves := make(map[int]int)
	for idx, pt := range path {
		for diffX := -maxCheatDist; diffX <= maxCheatDist; diffX++ {
			for diffY := -maxCheatDist; diffY <= maxCheatDist; diffY++ {
				possibleCheatTarget := pt.Add(image.Pt(diffX, diffY))
				if !possibleCheatTarget.In(border) {
					continue
				}
				cheatDist := numbers.Abs(diffX) + numbers.Abs(diffY)
				if cheatDist == 0 || cheatDist > maxCheatDist {
					continue
				}
				if i := slices.Index(path, possibleCheatTarget); i != -1 {
					if saved := i - idx - cheatDist; saved > 0 {
						saves[saved]++
					}
				}
			}
		}
	}
	return saves
}

func Part1(input string) (string, error) {
	grid, start, end := parseInput(input)
	path := solveMaze(grid, start, end)
	cheats := findCheats(grid, path, 2)
	var total int
	for saves, count := range cheats {
		if saves >= 100 {
			total += count
		}
	}
	return strconv.Itoa(total), nil
}

func Part2(input string) (string, error) {
	grid, start, end := parseInput(input)
	path := solveMaze(grid, start, end)
	cheats := findCheats(grid, path, 20)
	var total int
	for saves, count := range cheats {
		if saves >= 100 {
			total += count
		}
	}
	return strconv.Itoa(total), nil
}
