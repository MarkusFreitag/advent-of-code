package day3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "4361", solution)
}

func TestPart2(t *testing.T) {
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "467835", solution)
}
