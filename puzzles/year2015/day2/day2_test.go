package day2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToNums(t *testing.T) {
	nums, err := toNums("1x2x3")
	require.Nil(t, err)
	require.Equal(t, []int{1, 2, 3}, nums)
}

func TestMulNums(t *testing.T) {
	require.Equal(t, 6, mulNums([]int{1, 2, 3}))
}

func TestPart1(t *testing.T) {
	testcases := map[string]string{
		"2x3x4":  "58",
		"1x1x10": "43",
	}
	for input, expected := range testcases {
		solution, err := Part1(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}

func TestPart2(t *testing.T) {
	testcases := map[string]string{
		"2x3x4":  "34",
		"1x1x10": "14",
	}
	for input, expected := range testcases {
		solution, err := Part2(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}
