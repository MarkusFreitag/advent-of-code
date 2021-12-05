package day5

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var input = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

func TestPart1(t *testing.T) {
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "5", solution)
}

func TestPart2(t *testing.T) {
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "12", solution)
}
