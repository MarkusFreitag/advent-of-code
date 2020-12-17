package day17

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `.#.
..#
###`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "112", solution)
}

func TestPart2(t *testing.T) {
	input := `.#.
..#
###`
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "848", solution)
}
