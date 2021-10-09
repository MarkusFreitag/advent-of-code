package day2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var input = `ULL
RRDDD
LURDL
UUUUD`

func TestPart1(t *testing.T) {
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "1985", solution)
}

func TestPart2(t *testing.T) {
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "5DB3", solution)
}
