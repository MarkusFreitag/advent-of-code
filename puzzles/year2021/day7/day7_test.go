package day7

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var input = "16,1,2,0,4,2,7,1,2,14"

func TestPart1(t *testing.T) {
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "37", solution)
}

func TestPart2(t *testing.T) {
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "168", solution)
}
