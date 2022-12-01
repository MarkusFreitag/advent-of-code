package day1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testInput = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func TestPart1(t *testing.T) {
	solution, err := Part1(testInput)
	require.Nil(t, err)
	require.Equal(t, "24000", solution)
}

func TestPart2(t *testing.T) {
	solution, err := Part2(testInput)
	require.Nil(t, err)
	require.Equal(t, "45000", solution)
}
