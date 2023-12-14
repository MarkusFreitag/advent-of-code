package day14

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "136", solution)
}

func TestPart2(t *testing.T) {
	input := `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "64", solution)
}
