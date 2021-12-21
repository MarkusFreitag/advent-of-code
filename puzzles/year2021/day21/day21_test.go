package day21

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlayerMove(t *testing.T) {
	p := &Player{Pos: 6, Score: 742}
	p.Move(88 + 89 + 90)
	require.Equal(t, 3, p.Pos)
	require.Equal(t, 745, p.Score)

	p.Move(10)
	require.Equal(t, 3, p.Pos)
	require.Equal(t, 748, p.Score)

	p.Move(11)
	require.Equal(t, 4, p.Pos)
	require.Equal(t, 752, p.Score)
}

var input = `Player 1 starting position: 4
Player 2 starting position: 8`

func TestPart1(t *testing.T) {
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "739785", solution)
}

func TestPart2(t *testing.T) {
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "444356092776315", solution)
}
