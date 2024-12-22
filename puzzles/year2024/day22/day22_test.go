package day22

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `1
10
100
2024`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "37327623", solution)
}

func TestPart2(t *testing.T) {
	input := `1
2
3
2024`
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "23", solution)
}
