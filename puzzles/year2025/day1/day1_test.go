package day1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "3", solution)
}

func TestPart2(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "6", solution)
}
