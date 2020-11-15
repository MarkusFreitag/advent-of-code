package day3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "4", solution)
}

func TestPart2(t *testing.T) {
	input := `#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2`
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "3", solution)
}
