package day11

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `125 17`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "55312", solution)
}
