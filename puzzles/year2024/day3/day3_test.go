package day3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	solution, err := Part1("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")
	require.Nil(t, err)
	require.Equal(t, "161", solution)
}

func TestPart2(t *testing.T) {
	solution, err := Part2("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))")
	require.Nil(t, err)
	require.Equal(t, "48", solution)
}
