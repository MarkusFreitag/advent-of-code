package day1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var input = `199
200
208
210
200
207
240
269
260
263`

func TestPart1(t *testing.T) {
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "7", solution)
}

func TestPart2(t *testing.T) {
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "5", solution)
}
