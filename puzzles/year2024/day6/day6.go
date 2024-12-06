package day6

import (
	"image"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util/directions"
)

func turn(dir directions.Move) directions.Move {
	switch dir {
	case directions.Up:
		return directions.Left
	case directions.Right:
		return directions.Up
	case directions.Down:
		return directions.Right
	case directions.Left:
		return directions.Down
	}
	return dir
}

func guardsPath(grid [][]rune, pos image.Point, dir directions.Move) map[image.Point]struct{} {
	onMap := func(p image.Point) bool {
		if p.Y < 0 || p.Y >= len(grid) {
			return false
		}
		if p.X < 0 || p.X >= len(grid[0]) {
			return false
		}
		return true
	}
	visited := make(map[image.Point]struct{})
	visited[pos] = struct{}{}

	for {
		next := pos.Add(dir.Point())

		if !onMap(next) {
			break
		}

		if grid[next.Y][next.X] == '#' {
			dir = turn(dir)
			continue
		}

		pos = next
		visited[pos] = struct{}{}
	}

	return visited
}

func Part1(input string) (string, error) {
	grid := make([][]rune, 0)
	var pos image.Point
	for y, line := range strings.Split(input, "\n") {
		if x := strings.Index(line, "^"); x != -1 {
			pos = image.Pt(x, y)
			line = strings.ReplaceAll(line, "^", ".")
		}
		grid = append(grid, []rune(line))
	}

	return strconv.Itoa(len(guardsPath(grid, pos, directions.Down))), nil
}

func Part2(input string) (string, error) {
	grid := make([][]rune, 0)
	var start image.Point
	for y, line := range strings.Split(input, "\n") {
		if x := strings.Index(line, "^"); x != -1 {
			start = image.Pt(x, y)
			line = strings.ReplaceAll(line, "^", ".")
		}
		grid = append(grid, []rune(line))
	}

	visited := guardsPath(grid, start, directions.Down)

	var counter int
	for p := range visited {
		if p == start {
			continue
		}
		dup := make([][]rune, len(grid))
		for idx, row := range grid {
			dup[idx] = make([]rune, len(row))
			copy(dup[idx], grid[idx])
		}
		dup[p.Y][p.X] = '#'
		if inLoop(dup, start) {
			counter++
		}
	}

	return strconv.Itoa(counter), nil
}

func inLoop(grid [][]rune, pos image.Point) bool {
	onMap := func(p image.Point) bool {
		if p.Y < 0 || p.Y >= len(grid) {
			return false
		}
		if p.X < 0 || p.X >= len(grid[0]) {
			return false
		}
		return true
	}
	dir := directions.Down
	seen := make(map[string]struct{})
	for {
		nextPos := pos.Add(dir.Point())
		if !onMap(nextPos) {
			break
		}
		if grid[nextPos.Y][nextPos.X] == '#' {
			dir = turn(dir)
		} else {
			pos = nextPos
		}
		k := pos.String() + dir.String()
		if _, ok := seen[k]; ok {
			return true
		}
		seen[k] = struct{}{}
	}
	return false
}
