package day23

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	solution, err := Part1("389125467")
	require.Nil(t, err)
	require.Equal(t, "67384529", solution)
}

func TestPart2(t *testing.T) {
	solution, err := Part2("389125467")
	require.Nil(t, err)
	require.Equal(t, "149245887792", solution)
}
