package day21

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `029A
980A
179A
456A
379A`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "126384", solution)
}
