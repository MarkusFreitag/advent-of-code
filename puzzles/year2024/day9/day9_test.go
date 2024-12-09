package day9

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `2333133121414131402`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "1928", solution)
}

func TestPart2(t *testing.T) {
	input := `2333133121414131402`
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "2858", solution)
}
