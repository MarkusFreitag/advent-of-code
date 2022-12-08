package day8

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testInput = `30373
25512
65332
33549
35390`

func TestPart1(t *testing.T) {
	solution, err := Part1(testInput)
	require.Nil(t, err)
	require.Equal(t, "21", solution)
}

func TestPart2(t *testing.T) {
	solution, err := Part2(testInput)
	require.Nil(t, err)
	require.Equal(t, "8", solution)
}
