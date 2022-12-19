package day18

import (
  "testing"

	"github.com/stretchr/testify/require"
)

var testInput = `2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5`

func TestPart1(t *testing.T) {
    solution, err := Part1(testInput)
    require.Nil(t, err)
    require.Equal(t, "64", solution)
}

func TestPart2(t *testing.T) {
    solution, err := Part2(testInput)
    require.Nil(t, err)
    require.Equal(t, "58", solution)
}
