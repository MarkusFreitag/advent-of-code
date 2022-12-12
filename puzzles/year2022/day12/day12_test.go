package day12

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testInput = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

func TestPart1(t *testing.T) {
	solution, err := Part1(testInput)
	require.Nil(t, err)
	require.Equal(t, "31", solution)
}

func TestPart2(t *testing.T) {
	solution, err := Part2(testInput)
	require.Nil(t, err)
	require.Equal(t, "29", solution)
}
