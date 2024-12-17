package day17

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "4,6,3,5,6,3,5,2,1,0", solution)
}

func TestPart2(t *testing.T) {
	input := `Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "117440", solution)
}
