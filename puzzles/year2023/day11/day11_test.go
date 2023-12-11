package day11

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	input := `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`
	require.Equal(t, 374, solution(input, 2))
	require.Equal(t, 1030, solution(input, 10))
	require.Equal(t, 8410, solution(input, 100))
}
