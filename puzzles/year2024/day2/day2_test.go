package day2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "2", solution)
}

func TestPart2(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "4", solution)
}
