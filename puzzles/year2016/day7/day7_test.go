package day7

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	testcases := map[string]string{
		"abba[mnop]qrst":       "1",
		"abcd[bddb]xyyx":       "0",
		"aaaa[qwer]tyui":       "0",
		"ioxxoj[asdfgh]zxcvbn": "1",
	}
	for input, expected := range testcases {
		solution, err := Part1(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}

func TestPart2(t *testing.T) {
	testcases := map[string]string{
		"aba[bab]xyz":   "1",
		"xyx[xyx]xyx":   "0",
		"aaa[kek]eke":   "1",
		"zazbz[bzb]cdb": "1",
	}
	for input, expected := range testcases {
		solution, err := Part2(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}
