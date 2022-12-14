package day14

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

type pt struct {
	x, y int
}

func newPt(str string) pt {
	parts := strings.Split(str, ",")
	return pt{
		x: util.ParseInt(parts[0]),
		y: util.ParseInt(parts[1]),
	}
}

func parseLine(str string) []pt {
	line := make([]pt, 0)
	parts := strings.Fields(str)
	last := newPt(parts[0])
	for _, part := range parts[2:] {
		if part == "->" {
			continue
		}
		cur := newPt(part)
		if last.x == cur.x {
			if last.y > cur.y {
				for y := last.y; y >= cur.y; y-- {
					line = append(line, pt{x: last.x, y: y})
				}
			} else {
				for y := last.y; y <= cur.y; y++ {
					line = append(line, pt{x: last.x, y: y})
				}
			}
		} else if last.y == cur.y {
			if last.x > cur.x {
				for x := last.x; x >= cur.x; x-- {
					line = append(line, pt{x: x, y: last.y})
				}
			} else {
				for x := last.x; x <= cur.x; x++ {
					line = append(line, pt{x: x, y: last.y})
				}
			}
		}

		last = cur
	}
	return line
}

func parseInput(input string) (map[pt]struct{}, int) {
	grid := make(map[pt]struct{})
	var lowest int
	for _, line := range strings.Split(input, "\n") {
		for _, p := range parseLine(line) {
			grid[p] = exist
			if p.y > lowest {
				lowest = p.y
			}
		}
	}
	return grid, lowest
}

func checkNext(grid map[pt]struct{}, current pt) pt {
	next := pt{x: current.x, y: current.y + 1}
	if _, ok := grid[next]; !ok {
		return next
	}

	next = pt{x: current.x - 1, y: current.y + 1}
	if _, ok := grid[next]; !ok {
		return next
	}

	next = pt{x: current.x + 1, y: current.y + 1}
	if _, ok := grid[next]; !ok {
		return next
	}

	return current
}

var exist struct{}

func Part1(input string) (string, error) {
	grid, lowest := parseInput(input)

	var units int
	for {
		sand := pt{x: 500, y: 0}
		for {
			if _, ok := grid[pt{x: sand.x, y: sand.y + 1}]; !ok && sand.y+1 >= lowest {
				return strconv.Itoa(units), nil
			}
			next := checkNext(grid, sand)

			if next.x == sand.x && next.y == sand.y {
				grid[sand] = exist
				units++
				break
			}

			sand = next
		}
	}

	return "n/a", nil
}

func Part2(input string) (string, error) {
	grid, lowest := parseInput(input)

	var units int
	for {
		sand := pt{x: 500, y: 0}
		for {
			if _, ok := grid[sand]; ok {
				return strconv.Itoa(units), nil
			}
			if sand.y+1 == lowest+2 {
				grid[sand] = exist
				units++
				break
			}
			next := checkNext(grid, sand)

			if next.x == sand.x && next.y == sand.y {
				grid[sand] = exist
				units++
				break
			}

			sand = next
		}
	}

	return "n/a", nil
}
