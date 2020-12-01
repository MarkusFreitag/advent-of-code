package day19

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `H => HO
H => OH
O => HH

HOH`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "4", solution)

	input = `H => HO
H => OH
O => HH

HOHOHO`
	solution, err = Part1(input)
	require.Nil(t, err)
	require.Equal(t, "7", solution)
}

func TestPart2(t *testing.T) {
	input := `e => H
e => O
H => HO
H => OH
O => HH

HOH`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "3", solution)

	input = `e => H
e => O
H => HO
H => OH
O => HH

HOHOHO`
	solution, err = Part1(input)
	require.Nil(t, err)
	require.Equal(t, "6", solution)
}
