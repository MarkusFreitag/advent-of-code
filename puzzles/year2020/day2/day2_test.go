package day2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "2", solution)
}

func TestPart2(t *testing.T) {
	input := `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "1", solution)
}
