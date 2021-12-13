package day13

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var input = `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`

func TestPart1(t *testing.T) {
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "17", solution)
}

func TestPart2(t *testing.T) {
	expected := `#####
#   #
#   #
#   #
#####
     
     `
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, expected, solution)
}
