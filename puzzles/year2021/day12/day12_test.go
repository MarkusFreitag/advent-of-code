package day12

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var input = `start-A
start-b
A-c
A-b
b-d
A-end
b-end`

func TestPart1(t *testing.T) {
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "10", solution)
}

func TestPart2(t *testing.T) {
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "36", solution)
}
