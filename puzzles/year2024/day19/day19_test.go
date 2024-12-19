package day19

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "6", solution)
}

func TestPart2(t *testing.T) {
	input := `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "16", solution)
}
