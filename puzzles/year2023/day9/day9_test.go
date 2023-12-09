package day9

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `0 3 6 9 12 15`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "18", solution)

	input = `1 3 6 10 15 21`
	solution, err = Part1(input)
	require.Nil(t, err)
	require.Equal(t, "28", solution)

	input = `10 13 16 21 30 45`
	solution, err = Part1(input)
	require.Nil(t, err)
	require.Equal(t, "68", solution)

	input = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`
	solution, err = Part1(input)
	require.Nil(t, err)
	require.Equal(t, "114", solution)
}

func TestPart2(t *testing.T) {
	input := `0 3 6 9 12 15`
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "-3", solution)

	input = `1 3 6 10 15 21`
	solution, err = Part2(input)
	require.Nil(t, err)
	require.Equal(t, "0", solution)

	input = `10 13 16 21 30 45`
	solution, err = Part2(input)
	require.Nil(t, err)
	require.Equal(t, "5", solution)

	input = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`
	solution, err = Part2(input)
	require.Nil(t, err)
	require.Equal(t, "2", solution)
}
