package day15

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var input = `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`

func TestPart1(t *testing.T) {
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "40", solution)
}

func TestPart2(t *testing.T) {
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "315", solution)
}
