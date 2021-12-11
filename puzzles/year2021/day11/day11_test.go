package day11

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var input = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

func TestPart1(t *testing.T) { /*
			small := `11111
		19991
		19191
		19991
		11111`
			solution, err := Part1(small)
			require.Nil(t, err)
			require.Equal(t, "1656", solution)
	*/

	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "1656", solution)
}

func TestPart2(t *testing.T) {
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "195", solution)
}
