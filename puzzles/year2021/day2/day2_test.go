package day2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var input = `forward 5
down 5
forward 8
up 3
down 8
forward 2`

func TestPart1(t *testing.T) {
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "150", solution)
}

func TestPart2(t *testing.T) {
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "900", solution)
}
