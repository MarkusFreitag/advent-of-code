package day18

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseRow(t *testing.T) {
	testcases := map[string]string{
		"..^^.": "[false false true true false]",
		".^^^^": "[false true true true true]",
		"^^..^": "[true true false false true]",
	}
	for input, expected := range testcases {
		require.Equal(t, expected, fmt.Sprintf("%v", parseRow(input)))
	}
}

func TestGenerateRow(t *testing.T) {
	testcases := map[string][]bool{
		"[false true true true true]":  {false, false, true, true, false},
		"[true true false false true]": {false, true, true, true, true},
	}
	for expected, input := range testcases {
		require.Equal(t, expected, fmt.Sprintf("%v", generateRow(input)))
	}
}

func TestPart1(t *testing.T) {
	totalRows = 10
	solution, err := Part1(".^^.^.^^^^")
	require.Nil(t, err)
	require.Equal(t, "38", solution)
}

func TestPart2(t *testing.T) {
	testcases := map[string]string{}
	for input, expected := range testcases {
		solution, err := Part2(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}
