package util_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/MarkusFreitag/advent-of-code/util"
)

func TestStringReverse(t *testing.T) {
	testcases := map[string]string{
		"abab": "baba",
		"abba": "abba",
	}
	for input, expected := range testcases {
		require.Equal(t, expected, util.StringReverse(input))
	}
}

func TestStringPadLeft(t *testing.T) {
	testcases := map[string]string{
		"1":     "0001",
		"1010":  "1010",
		"11111": "1111",
	}
	for input, expected := range testcases {
		require.Equal(t, expected, util.StringPadLeft(input, "0", 4))
	}
}

func TestStringPadRight(t *testing.T) {
	testcases := map[string]string{
		"1":     "1000",
		"1010":  "1010",
		"11111": "1111",
	}
	for input, expected := range testcases {
		require.Equal(t, expected, util.StringPadRight(input, "0", 4))
	}
}

func TestStringToStrings(t *testing.T) {
	slice := util.StringToStrings("a")
	require.Equal(t, 1, len(slice))
	require.Equal(t, "a", slice[0])

	slice = util.StringToStrings("abc")
	require.Equal(t, 3, len(slice))
	require.Equal(t, "a", slice[0])
	require.Equal(t, "b", slice[1])
	require.Equal(t, "c", slice[2])
}

func TestStringSorter(t *testing.T) {
	testcases := map[string]string{
		"a":   "a",
		"abc": "abc",
		"zoa": "aoz",
	}
	for input, expected := range testcases {
		require.Equal(t, expected, util.StringSorter(input))
	}
}

func TestStringContainsAny(t *testing.T) {
	require.True(t, util.StringContainsAny("abc", "a"))
	require.True(t, util.StringContainsAny("abc", "a", "b"))
	require.True(t, util.StringContainsAny("abc", "d", "b"))
	require.False(t, util.StringContainsAny("abc", "d"))
}

func TestStringTally(t *testing.T) {
	require.Equal(t, map[string]int{"a": 2, "b": 2, "c": 1, "d": 2, "E": 1, "A": 1}, util.StringTally("abcddEaAb"))
}
