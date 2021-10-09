package day16

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	testcases := map[string]string{
		"1":            "100",
		"0":            "001",
		"11111":        "11111000000",
		"111100001010": "1111000010100101011110000",
	}
	for str, expected := range testcases {
		require.Equal(t, expected, string(generate([]byte(str))))
	}
}

func TestFillDisk(t *testing.T) {
	require.Equal(t, "10000011110010000111110", string(fillDisk([]byte("10000"), 23)))

	require.Equal(t, "10000011110010000111", string(fillDisk([]byte("10000"), 20)))
}

func TestChecksum(t *testing.T) {
	require.Equal(t, "100", string(checksum([]byte("110010110100"))))
}

func TestPart1(t *testing.T) {
	diskSize = 20
	solution, err := Part1("10000")
	require.Nil(t, err)
	require.Equal(t, "01100", solution)
}

func TestPart2(t *testing.T) {
	testcases := map[string]string{}
	for input, expected := range testcases {
		solution, err := Part2(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}
