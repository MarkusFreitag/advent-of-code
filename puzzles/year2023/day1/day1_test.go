package day1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	testcases := map[string]string{
		"1abc2":       "12",
		"pqr3stu8vwx": "38",
		"a1b2c3d4e5f": "15",
		"treb7uchet":  "77",
	}
	for input, expected := range testcases {
		solution, err := Part1(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}

func TestPart2(t *testing.T) {
	testcases := map[string]string{
		"two1nine":         "29",
		"eightwothree":     "83",
		"abcone2threexyz":  "13",
		"xtwone3four":      "24",
		"4nineeightseven2": "42",
		"zoneight234":      "14",
		"7pqrstsixteen":    "76",
	}
	for input, expected := range testcases {
		solution, err := Part2(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}
