package day16

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	solution, err := Part1("12345678")
	require.Nil(t, err)
	require.Equal(t, "23845678", solution)
}
