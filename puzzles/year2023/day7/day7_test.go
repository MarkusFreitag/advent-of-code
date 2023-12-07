package day7

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "6440", solution)
}

func TestPart2(t *testing.T) {
	input := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "5905", solution)
}
