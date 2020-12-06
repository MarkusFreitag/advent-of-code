package day6

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `abc

a
b
c

ab
ac

a
a
a
a

b`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "11", solution)
}

func TestPart2(t *testing.T) {
	input := `abc

a
b
c

ab
ac

a
a
a
a

b`
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "6", solution)
}
