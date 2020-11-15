package day1

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	testcases := map[string]string{
		"+1, -2, +3, +1": "3",
		"+1, +1, +1":     "3",
		"+1, +1, -2":     "0",
		"-1, -2, -3":     "-6",
	}
	for input, expected := range testcases {
		solution, err := Part1(strings.ReplaceAll(input, ",", "\n"))
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}

func TestPart2(t *testing.T) {
	testcases := map[string]string{
		"+1, -1":             "0",
		"+3, +3, +4, -2, -4": "10",
		"-6, +3, +8, +5, -6": "5",
		"+7, +7, -2, -7, -4": "14",
	}
	for input, expected := range testcases {
		solution, err := Part2(strings.ReplaceAll(input, ",", "\n"))
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}
