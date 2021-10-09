package day12

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var input = `cpy 41 a
inc a
inc a
dec a
jnz a 2
dec a`

func TestPart1(t *testing.T) {
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "42", solution)
}

func TestPart2(t *testing.T) {
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "42", solution)
}
