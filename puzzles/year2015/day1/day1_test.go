package day1

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/MarkusFreitag/advent-of-code/util"
)

func TestPart1(t *testing.T) {
	solution, err := Part1("dummy")
	require.NotNil(t, err)
	require.Equal(t, util.ErrNotSolved, err)
	require.Equal(t, "dummy", solution)
}

func TestPart2(t *testing.T) {
	solution, err := Part2("dummy")
	require.NotNil(t, err)
	require.Equal(t, util.ErrNotSolved, err)
	require.Equal(t, "dummy", solution)
}
