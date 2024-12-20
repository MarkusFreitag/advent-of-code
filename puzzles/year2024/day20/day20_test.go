package day20

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

func prepareTest() ([][]rune, []image.Point) {
	input := `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`
	grid, start, end := parseInput(input)
	path := solveMaze(grid, start, end)
	return grid, path
}

func TestPart1(t *testing.T) {
	grid, path := prepareTest()
	cheats := findCheats(grid, path, 2)

	expectedCheats := map[int]int{
		64: 1,
		40: 1,
		38: 1,
		36: 1,
		20: 1,
		12: 3,
		10: 2,
		8:  4,
		6:  2,
		4:  14,
		2:  14,
	}
	assert.Equal(t, expectedCheats, cheats)
}

func TestPart2(t *testing.T) {
	grid, path := prepareTest()
	cheats := findCheats(grid, path, 20)

	filteredCheats := make(map[int]int)
	for saved, count := range cheats {
		if saved < 50 {
			continue
		}
		filteredCheats[saved] = count
	}

	expectedCheats := map[int]int{
		76: 3,
		74: 4,
		72: 22,
		70: 12,
		68: 14,
		66: 12,
		64: 19,
		62: 20,
		60: 23,
		58: 25,
		56: 39,
		54: 29,
		52: 31,
		50: 32,
	}
	assert.Equal(t, expectedCheats, filteredCheats)
}
