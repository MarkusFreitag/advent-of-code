package day16

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

const (
	up int = iota
	right
	down
	left
)

var dirs = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func add(a, b [2]int) [2]int {
	return [2]int{a[0] + b[0], a[1] + b[1]}
}

type Beam struct {
	pos [2]int
	dir int
}

func (b Beam) Move(dir int) Beam {
	return Beam{
		pos: add(b.pos, dirs[dir]),
		dir: dir,
	}
}

func followBeam(grid [][]rune, energized map[[2]int]bool, visited map[Beam]bool, beam Beam) {
	if beam.pos[0] < 0 || beam.pos[0] >= len(grid) {
		return
	}
	if beam.pos[1] < 0 || beam.pos[1] >= len(grid[0]) {
		return
	}

	energized[beam.pos] = true
	if _, ok := visited[beam]; ok {
		return
	}
	visited[beam] = true

	for _, b := range check(grid[beam.pos[0]][beam.pos[1]], beam) {
		followBeam(grid, energized, visited, b)
	}
}

func check(char rune, beam Beam) []Beam {
	if char == '.' {
		return []Beam{beam.Move(beam.dir)}
	}

	if beam.dir == up {
		if char == '\\' {
			return []Beam{beam.Move(left)}
		}
		if char == '/' {
			return []Beam{beam.Move(right)}
		}
		if char == '-' {
			return []Beam{
				beam.Move(left),
				beam.Move(right),
			}
		}
		return []Beam{beam.Move(beam.dir)}
	}
	if beam.dir == right {
		if char == '\\' {
			return []Beam{beam.Move(down)}
		}
		if char == '/' {
			return []Beam{beam.Move(up)}
		}
		if char == '|' {
			return []Beam{
				beam.Move(up),
				beam.Move(down),
			}
		}
		return []Beam{beam.Move(beam.dir)}
	}
	if beam.dir == down {
		if char == '\\' {
			return []Beam{beam.Move(right)}
		}
		if char == '/' {
			return []Beam{beam.Move(left)}
		}
		if char == '-' {
			return []Beam{
				beam.Move(right),
				beam.Move(left),
			}
		}
		return []Beam{beam.Move(beam.dir)}
	}
	if beam.dir == left {
		if char == '\\' {
			return []Beam{beam.Move(up)}
		}
		if char == '/' {
			return []Beam{beam.Move(down)}
		}
		if char == '|' {
			return []Beam{
				beam.Move(down),
				beam.Move(up),
			}
		}
		return []Beam{beam.Move(beam.dir)}
	}

	panic("unreachable code with check")
}

func Part1(input string) (string, error) {
	lines := strings.Fields(input)
	grid := make([][]rune, len(lines))
	for idx, row := range lines {
		grid[idx] = []rune(row)
	}
	energized := make(map[[2]int]bool)
	followBeam(grid, energized, make(map[Beam]bool), Beam{pos: [2]int{0, 0}, dir: right})
	return strconv.Itoa(len(energized)), nil
}

func Part2(input string) (string, error) {
	lines := strings.Fields(input)
	grid := make([][]rune, len(lines))
	for idx, row := range lines {
		grid[idx] = []rune(row)
	}
	var max int
	// from top to bottom
	for i := 0; i < len(grid[0]); i++ {
		energized := make(map[[2]int]bool)
		followBeam(grid, energized, make(map[Beam]bool), Beam{pos: [2]int{0, i}, dir: down})
		max = numbers.Max(max, len(energized))
	}
	// from bottom to top
	for i := 0; i < len(grid[0]); i++ {
		energized := make(map[[2]int]bool)
		followBeam(grid, energized, make(map[Beam]bool), Beam{pos: [2]int{len(grid) - 1, i}, dir: up})
		max = numbers.Max(max, len(energized))
	}
	// from left to right
	for i := 0; i < len(grid[0]); i++ {
		energized := make(map[[2]int]bool)
		followBeam(grid, energized, make(map[Beam]bool), Beam{pos: [2]int{i, 0}, dir: right})
		max = numbers.Max(max, len(energized))
	}
	// from right to left
	for i := 0; i < len(grid[0]); i++ {
		energized := make(map[[2]int]bool)
		followBeam(grid, energized, make(map[Beam]bool), Beam{pos: [2]int{i, len(grid[0]) - 1}, dir: left})
		max = numbers.Max(max, len(energized))
	}
	return strconv.Itoa(max), nil

}
