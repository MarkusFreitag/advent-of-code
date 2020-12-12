package day12

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTurn(t *testing.T) {
	require.Equal(t, 'E', turn('N', 'R'))
	require.Equal(t, 'S', turn('E', 'R'))
	require.Equal(t, 'W', turn('S', 'R'))
	require.Equal(t, 'N', turn('W', 'R'))

	require.Equal(t, 'W', turn('N', 'L'))
	require.Equal(t, 'N', turn('E', 'L'))
	require.Equal(t, 'E', turn('S', 'L'))
	require.Equal(t, 'S', turn('W', 'L'))
}

func TestTurnWP(t *testing.T) {
	x, y := turnWP(10, 4, 'R')
	require.Equal(t, 4, x)
	require.Equal(t, -10, y)

	x, y = turnWP(10, 4, 'L')
	require.Equal(t, -4, x)
	require.Equal(t, 10, y)
}

func TestPart1(t *testing.T) {
	input := `F10
N3
F7
R90
F11`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "25", solution)
}

func TestPart2(t *testing.T) {
	input := `F10
N3
F7
R90
F11`
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "286", solution)
}
