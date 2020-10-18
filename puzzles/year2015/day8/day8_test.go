package day8

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	testcases := map[string]string{
		`""`:         "2",
		`"abc"`:      "2",
		`"aaa\"aaa"`: "3",
		`"\x27"`:     "5",
	}
	for input, expected := range testcases {
		solution, err := Part1(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}

func TestPart2(t *testing.T) {
	testcases := map[string]string{
		`""`:         "4",
		`"abc"`:      "4",
		`"aaa\"aaa"`: "6",
		`"\x27"`:     "5",
	}
	for input, expected := range testcases {
		solution, err := Part2(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}
