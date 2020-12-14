package day14

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "165", solution)
}

func TestPart2(t *testing.T) {
	input := `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "208", solution)
}
