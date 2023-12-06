package day6

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `Time:      7  15   30
Distance:  9  40  200`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "288", solution)
}

func TestPart2(t *testing.T) {
	input := `Time:      7  15   30
Distance:  9  40  200`
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "71503", solution)
}
