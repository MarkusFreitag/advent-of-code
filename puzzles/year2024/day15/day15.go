package day15

import (
	"image"
	"slices"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util/directions"
)

func parseInput(input string, wide bool) ([][]rune, image.Point, string) {
	blocks := strings.Split(input, "\n\n")
	var pos image.Point
	grid := make([][]rune, 0)
	for y, line := range strings.Split(blocks[0], "\n") {
		if wide {
			line = strings.ReplaceAll(line, "#", "##")
			line = strings.ReplaceAll(line, "O", "[]")
			line = strings.ReplaceAll(line, ".", "..")
			line = strings.ReplaceAll(line, "@", "@.")
		}
		row := []rune(line)
		grid = append(grid, row)
		for x, char := range row {
			if char == '@' {
				pos = image.Pt(x, y)
			}
		}
	}

	var moves string
	for _, line := range strings.Split(blocks[1], "\n") {
		moves += line
	}
	return grid, pos, moves
}

var moveTable = map[rune]image.Point{
	'^': directions.Down.Point(),
	'>': directions.Right.Point(),
	'v': directions.Up.Point(),
	'<': directions.Left.Point(),
}

func movePossible(grid [][]rune, pos, dir image.Point) bool {
	if dir == directions.Right.Point() || dir == directions.Left.Point() {
		chars := make([]rune, 0)
		for p := pos; grid[p.Y][p.X] != '#'; p = p.Add(dir) {
			chars = append(chars, grid[p.Y][p.X])
		}
		return slices.Contains(chars, '.')
	}

	nPos := pos.Add(dir)
	nChar := grid[nPos.Y][nPos.X]
	if nChar == '.' {
		return true
	}
	if nChar == 'O' {
		return movePossible(grid, nPos, dir)
	}
	if nChar == '[' {
		return movePossible(grid, nPos, dir) && movePossible(grid, nPos.Add(directions.Right.Point()), dir)
	}
	if nChar == ']' {
		return movePossible(grid, nPos, dir) && movePossible(grid, nPos.Add(directions.Left.Point()), dir)
	}

	return false
}

func moveCell(grid [][]rune, from, to, dir image.Point) [][]rune {
	cell := grid[from.Y][from.X]
	nextCell := grid[to.Y][to.X]
	if cell == '.' {
		return grid
	}

	if nextCell == '.' {
		grid[from.Y][from.X] = '.'
		grid[to.Y][to.X] = cell
		return grid
	}

	if nextCell == 'O' {
		grid = moveCell(grid, to, to.Add(dir), dir)
		grid[from.Y][from.X] = '.'
		grid[to.Y][to.X] = cell
		return grid
	}

	if dir == directions.Right.Point() || dir == directions.Left.Point() {
		if nextCell == '[' {
			grid = moveCell(grid, to, to.Add(dir), dir)
			grid[from.Y][from.X] = '.'
			grid[to.Y][to.X] = cell
		} else if nextCell == ']' {
			grid = moveCell(grid, to, to.Add(dir), dir)
			grid[from.Y][from.X] = '.'
			grid[to.Y][to.X] = cell
		}
		return grid
	}

	if nextCell == '[' {
		grid = moveCell(grid, to, to.Add(dir), dir)
		grid = moveCell(grid, to.Add(directions.Right.Point()), to.Add(dir).Add(directions.Right.Point()), dir)
		grid[from.Y][from.X] = '.'
		grid[to.Y][to.X] = cell
		return grid
	}
	if nextCell == ']' {
		grid = moveCell(grid, to, to.Add(dir), dir)
		grid = moveCell(grid, to.Add(directions.Left.Point()), to.Add(dir).Add(directions.Left.Point()), dir)
		grid[from.Y][from.X] = '.'
		grid[to.Y][to.X] = cell
		return grid
	}

	return grid
}

func simulation(input string, wide bool) int {
	grid, pos, moves := parseInput(input, wide)

	for _, move := range moves {
		dir := moveTable[move]
		np := pos.Add(dir)
		nextChar := grid[np.Y][np.X]
		if nextChar == '#' {
			continue
		}
		if nextChar == '.' || movePossible(grid, pos, dir) {
			grid = moveCell(grid, pos, np, dir)
			pos = np
		}
	}

	box := 'O'
	if wide {
		box = '['
	}
	var sum int
	for y, row := range grid {
		for x, col := range row {
			if col == box {
				sum += 100*y + x
			}
		}
	}
	return sum
}

func Part1(input string) (string, error) {
	return strconv.Itoa(simulation(input, false)), nil
}

func Part2(input string) (string, error) {
	return strconv.Itoa(simulation(input, true)), nil
}
