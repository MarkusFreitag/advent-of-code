package day3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPointAdd(t *testing.T) {
	a := point{X: 0, Y: 1}
	b := point{X: 1, Y: 2}
	c := a.Add(b)
	require.Equal(t, 1, c.X)
	require.Equal(t, 3, c.Y)
}

func TestPointsIncludes(t *testing.T) {
	pts := points{point{X: 0, Y: 1}}
	require.True(t, pts.Includes(point{X: 0, Y: 1}))
	require.False(t, pts.Includes(point{X: 0, Y: 2}))
}

func TestPart1(t *testing.T) {
	testcases := map[string]string{
		">":          "2",
		"^>v<":       "4",
		"^v^v^v^v^v": "2",
	}
	for input, expected := range testcases {
		solution, err := Part1(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}

func TestPart2(t *testing.T) {
	testcases := map[string]string{
		"^v":         "3",
		"^>v<":       "3",
		"^v^v^v^v^v": "11",
	}
	for input, expected := range testcases {
		solution, err := Part2(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}
