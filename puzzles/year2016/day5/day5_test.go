package day5

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var input = "abc"

func TestPart1(t *testing.T) {
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "18f47a30", solution)
}

func TestPart2(t *testing.T) {
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "05ace8e3", solution)
}
