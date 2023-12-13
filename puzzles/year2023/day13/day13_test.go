package day13

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "5", solution)

	input = `#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`
	solution, err = Part1(input)
	require.Nil(t, err)
	require.Equal(t, "400", solution)

	input = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`
	solution, err = Part1(input)
	require.Nil(t, err)
	require.Equal(t, "405", solution)
}

func TestPart2(t *testing.T) {
	input := `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.`
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "300", solution)

	input = `#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`
	solution, err = Part2(input)
	require.Nil(t, err)
	require.Equal(t, "100", solution)

	input = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`
	solution, err = Part2(input)
	require.Nil(t, err)
	require.Equal(t, "400", solution)
}
