package day15

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	testcases := map[string]string{
		"0,3,6": "436",
		"1,3,2": "1",
		"2,1,3": "10",
		"1,2,3": "27",
		"2,3,1": "78",
		"3,2,1": "438",
		"3,1,2": "1836",
	}
	for input, expected := range testcases {
		solution, err := Part1(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}

func TestPart2(t *testing.T) {
	testcases := map[string]string{
		"0,3,6": "175594",
		"1,3,2": "2578",
		"2,1,3": "3544142",
		"1,2,3": "261214",
		"2,3,1": "6895259",
		"3,2,1": "18",
		"3,1,2": "362",
	}
	for input, expected := range testcases {
		solution, err := Part2(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}
