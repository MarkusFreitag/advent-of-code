package day17

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var input = "target area: x=20..30, y=-10..-5"

func TestPart1(t *testing.T) {
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "45", solution)
}

func TestPart2(t *testing.T) {
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "112", solution)
}
