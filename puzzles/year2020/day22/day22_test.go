package day22

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "306", solution)
}

func TestPart2(t *testing.T) {
	input := `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "291", solution)
}
