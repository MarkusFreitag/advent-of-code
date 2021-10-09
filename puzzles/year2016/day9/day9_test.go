package day9

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	testcases := map[string]string{
		"ADVENT":            "6",  // ADVENT
		"A(1x5)BC":          "7",  // ABBBBBC
		"(3x3)XYZ":          "9",  // XYZXYZXYZ
		"A(2x2)BCD(2x2)EFG": "11", // ABCBCDEFEFG
		"(6x1)(1x3)A":       "6",  // (1x3)A
		"X(8x2)(3x3)ABCY":   "18", // X(3x3)ABC(3x3)ABCY
	}
	for input, expected := range testcases {
		solution, err := Part1(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}

func TestPart2(t *testing.T) {
	testcases := map[string]string{
		"(3x3)XYZ":                           "9",
		"X(8x2)(3x3)ABCY":                    "20",
		"(27x12)(20x12)(13x14)(7x10)(1x12)A": "241920",
		"(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN": "445",
	}
	for input, expected := range testcases {
		solution, err := Part2(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}
