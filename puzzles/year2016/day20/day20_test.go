package day20

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	max = 9
	input := `5-8
0-2
4-7`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "3", solution)
}

func TestPart2(t *testing.T) {
	max = 9
	input := `5-8
0-2
4-7`
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "2", solution)
}
