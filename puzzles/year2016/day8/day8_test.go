package day8

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	displayWidth, displayHeight = 6, 3
	input := `rect 3x2
rotate column x=1 by 1
rotate row y=0 by 4
rotate column x=1 by 1`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "6", solution)
}

func TestPart2(t *testing.T) {
	testcases := map[string]string{}
	for input, expected := range testcases {
		solution, err := Part2(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}
