package day18

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	size = 6
	firstBytes = 12
	input := `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "22", solution)
}

func TestPart2(t *testing.T) {
	size = 6
	firstBytes = 12
	input := `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "6,1", solution)
}
