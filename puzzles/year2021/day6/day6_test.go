package day6

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var input = "3,4,3,1,2"

func TestPart1(t *testing.T) {
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "5934", solution)
}

func TestPart2(t *testing.T) {
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "26984457539", solution)
}
