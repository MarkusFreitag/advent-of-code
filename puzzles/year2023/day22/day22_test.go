package day22

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "5", solution)
}

func TestPart2(t *testing.T) {
	input := `1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9`
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "7", solution)
}
