package day3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `5 10 25
5 10 14`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "1", solution)
}
