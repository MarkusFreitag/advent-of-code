package day1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "11", solution)
}

func TestPart2(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "31", solution)
}
