package util

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInputFilename(t *testing.T) {
	require.Equal(t, "inputs/2019_1.txt", inputFilename(2019, 1))
}

func TestStringsToInts(t *testing.T) {
	require.Equal(t, []int{1, 2, 3, 4, -5}, StringsToInts([]string{"1", "2", "3", "+4", "-5"}))

	require.Panics(t, func() { StringsToInts([]string{"1", "a", "3"}) })
}

func TestIntsToStrings(t *testing.T) {
	require.Equal(t, []string{"1", "2", "3", "4", "-5"}, IntsToStrings([]int{1, 2, 3, +4, -5}))
}

func TestParseInt(t *testing.T) {
	testcases := map[string]int{
		"1":  1,
		"+2": 2,
		"-3": -3,
	}
	for input, expected := range testcases {
		require.Equal(t, expected, ParseInt(input))
	}
}

func TestParseFloat(t *testing.T) {
	testcases := map[string]float64{
		"1.2":  1.2,
		"+2.1": 2.1,
		"-3.0": -3.0,
	}
	for input, expected := range testcases {
		require.Equal(t, expected, ParseFloat(input))
	}
}

func TestBinStringToDecInt(t *testing.T) {
	testcases := map[string]int{
		"1":     1,
		"1010":  10,
		"11111": 31,
	}
	for input, expected := range testcases {
		require.Equal(t, expected, BinStringToDecInt(input))
	}
}

func TestDecIntToBinString(t *testing.T) {
	testcases := map[int]string{
		1:  "1",
		10: "1010",
		31: "11111",
	}
	for input, expected := range testcases {
		require.Equal(t, expected, DecIntToBinString(input))
	}
}

func TestRangeInt(t *testing.T) {
	require.Equal(t, []int{1}, slices.Collect(RangeInt(1, 1, 1)))
	require.Equal(t, []int{1, 2, 3}, slices.Collect(RangeInt(1, 3, 1)))
	require.Equal(t, []int{3, 2, 1}, slices.Collect(RangeInt(3, 1, 1)))
	require.Equal(t, []int{1, 3, 5}, slices.Collect(RangeInt(1, 6, 2)))
	require.Equal(t, []int{6, 4, 2}, slices.Collect(RangeInt(6, 1, 2)))
}

func TestOnLineInt(t *testing.T) {
	require.True(t, OnLineInt(0, 0, 5, 5, 1, 1))
	require.True(t, OnLineInt(5, 5, 0, 0, 1, 1))
	require.True(t, OnLineInt(0, 0, 0, 5, 0, 1))
	require.True(t, OnLineInt(0, 5, 0, 0, 0, 1))
	require.True(t, OnLineInt(0, 0, 5, 0, 1, 0))
	require.True(t, OnLineInt(5, 0, 0, 0, 1, 0))
	require.False(t, OnLineInt(0, 0, 5, 5, 2, 3))
}

func TestFlatten(t *testing.T) {
	nestedSlice := []any{"a", []any{1, true}, []any{[]any{5}, "b"}, 1.1}
	flattenedSlice := []any{"a", 1, true, 5, "b", 1.1}
	require.Equal(t, flattenedSlice, Flatten(nestedSlice))
}

func TestRepeat(t *testing.T) {
	require.Equal(t, []string{"a"}, Repeat("a", 1))
	require.Equal(t, []string{"a", "a"}, Repeat("a", 2))
	require.Equal(t, [][]string{{"a", "b"}, {"a", "b"}}, Repeat([]string{"a", "b"}, 2))
}
