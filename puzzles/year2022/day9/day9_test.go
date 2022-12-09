package day9

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	testInput := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`
	solution, err := Part1(testInput)
	require.Nil(t, err)
	require.Equal(t, "13", solution)
}

func TestPart2(t *testing.T) {
	testInput := `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`
	solution, err := Part2(testInput)
	require.Nil(t, err)
	require.Equal(t, "36", solution)
}
