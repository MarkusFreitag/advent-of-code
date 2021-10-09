package day15

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDisk(t *testing.T) {
	disk := NewDisk(5, 2)
	positions := []int{2, 3, 4, 0, 1, 2, 3}
	for idx, pos := range positions {
		require.Equal(t, pos, disk.CalculatePos(idx))
	}
}

func TestPart1(t *testing.T) {
	input := `Disc #1 has 5 positions; at time=0, it is at position 4.
Disc #2 has 2 positions; at time=0, it is at position 1.`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "5", solution)
}

func TestPart2(t *testing.T) {
	testcases := map[string]string{}
	for input, expected := range testcases {
		solution, err := Part2(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}
