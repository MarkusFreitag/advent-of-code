package day1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	solution, err := Part1("12")
	require.Nil(t, err)
	require.Equal(t, "2", solution)
	solution, err = Part1("14")
	require.Nil(t, err)
	require.Equal(t, "2", solution)
	solution, err = Part1("1969")
	require.Nil(t, err)
	require.Equal(t, "654", solution)
	solution, err = Part1("100756")
	require.Nil(t, err)
	require.Equal(t, "33583", solution)
	solution, err = Part1("12\n14\n1969\n100756")
	require.Nil(t, err)
	require.Equal(t, "34241", solution)
}

func TestPart2(t *testing.T) {
	solution, err := Part2("14")
	require.Nil(t, err)
	require.Equal(t, "2", solution)
	solution, err = Part2("1969")
	require.Nil(t, err)
	require.Equal(t, "966", solution)
	solution, err = Part2("100756")
	require.Nil(t, err)
	require.Equal(t, "50346", solution)
	solution, err = Part2("14\n1969\n100756")
	require.Nil(t, err)
	require.Equal(t, "51314", solution)
}
