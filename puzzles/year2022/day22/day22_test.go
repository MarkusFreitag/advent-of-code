package day22

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testInput = `        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5
`

func TestPart1(t *testing.T) {
	solution, err := Part1(testInput)
	require.Nil(t, err)
	require.Equal(t, "6032", solution)
}

func TestPart2(t *testing.T) {
	solution, err := Part2(testInput)
	require.Nil(t, err)
	require.Equal(t, "", solution)
}
