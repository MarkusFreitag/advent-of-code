package day2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	testcases := map[string]string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green":                   "1",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue":         "2",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red": "0",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red": "0",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green":                   "5",
	}
	for input, expected := range testcases {
		solution, err := Part1(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}

	input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "8", solution)
}

func TestPart2(t *testing.T) {
	testcases := map[string]string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green":                   "48",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue":         "12",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red": "1560",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red": "630",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green":                   "36",
	}
	for input, expected := range testcases {
		solution, err := Part2(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}

	input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "2286", solution)
}
