package day1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	testcases := map[string]string{
		"(())":    "0",
		"()()":    "0",
		"(((":     "3",
		"(()(()(": "3",
		"))(((((": "3",
		"())":     "-1",
		"))(":     "-1",
		")))":     "-3",
		")())())": "-3",
	}
	for input, expected := range testcases {
		solution, err := Part1(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}

func TestPart2(t *testing.T) {
	testcases := map[string]string{
		")":     "1",
		"()())": "5",
	}
	for input, expected := range testcases {
		solution, err := Part2(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}
