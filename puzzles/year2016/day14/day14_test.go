package day14

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var input = "abc"

func TestPart1(t *testing.T) {
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "22728", solution)
}

func TestPart2(t *testing.T) {
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "22551", solution)
}
