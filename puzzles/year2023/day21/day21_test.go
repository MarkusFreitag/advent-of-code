package day21

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........`
	steps = 6
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "16", solution)
}

func TestPart2(t *testing.T) {
	input := `...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........`
	testcases := [][2]int{
		{6, 16},
		{10, 50},
		{50, 1594},
		{100, 6536},
		{500, 167004},
		{1000, 668697},
		{5000, 16733044},
	}
	for _, tc := range testcases {
		steps = tc[0]
		solution, err := Part2(input)
		require.Nil(t, err)
		require.Equal(t, strconv.Itoa(tc[1]), solution)
	}
}
