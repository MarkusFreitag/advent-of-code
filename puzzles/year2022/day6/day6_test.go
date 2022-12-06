package day6

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	testcases := map[string]string{
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      "5",
		"nppdvjthqldpwncqszvftbrmjlhg":      "6",
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": "10",
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  "11",
	}
	for input, expected := range testcases {
		solution, err := Part1(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}

func TestPart2(t *testing.T) {
	testcases := map[string]string{
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      "23",
		"nppdvjthqldpwncqszvftbrmjlhg":      "23",
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": "29",
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  "26",
	}
	for input, expected := range testcases {
		solution, err := Part2(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}
