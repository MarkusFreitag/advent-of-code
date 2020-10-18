package day9

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "605", solution)
}

func TestPart2(t *testing.T) {
	input := `London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "982", solution)
}
