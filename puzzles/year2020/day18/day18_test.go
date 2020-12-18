package day18

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	testcases := map[string]string{
		"1 + 2 * 3 + 4 * 5 + 6":                           "71",
		"1 + (2 * 3) + (4 * (5 + 6))":                     "51",
		"2 * 3 + (4 * 5)":                                 "26",
		"5 + (8 * 3 + 9 + 3 * 4 * 3)":                     "437",
		"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))":       "12240",
		"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2": "13632",
	}
	for input, expected := range testcases {
		solution, err := Part1(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}

func TestPart2(t *testing.T) {
	testcases := map[string]string{
		"1 + 2 * 3 + 4 * 5 + 6":                           "231",
		"1 + (2 * 3) + (4 * (5 + 6))":                     "51",
		"2 * 3 + (4 * 5)":                                 "46",
		"5 + (8 * 3 + 9 + 3 * 4 * 3)":                     "1445",
		"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))":       "669060",
		"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2": "23340",
	}
	for input, expected := range testcases {
		solution, err := Part2(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}
