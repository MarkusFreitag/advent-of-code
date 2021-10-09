package day19

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	solution, err := Part1("5")
	require.Nil(t, err)
	require.Equal(t, "3", solution)
}

func TestPart2(t *testing.T) {
	solution, err := Part2("5")
	require.Nil(t, err)
	require.Equal(t, "2", solution)
}
