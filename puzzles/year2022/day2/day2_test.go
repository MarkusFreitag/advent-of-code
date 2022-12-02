package day2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testInput = `A Y
B X
C Z`

func TestPart1(t *testing.T) {
	solution, err := Part1(testInput)
	require.Nil(t, err)
	require.Equal(t, "15", solution)
}

func TestPart2(t *testing.T) {
	solution, err := Part2(testInput)
	require.Nil(t, err)
	require.Equal(t, "12", solution)
}
