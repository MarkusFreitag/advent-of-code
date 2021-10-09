package day17

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	testcases := map[string]string{
		"ihgpwlah": "DDRRRD",
		"kglvqrro": "DDUDRLRRUDRD",
		"ulqzkmiv": "DRURDRUDDLLDLUURRDULRLDUUDDDRR",
	}
	for input, expected := range testcases {
		solution, err := Part1(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}

func TestPart2(t *testing.T) {
	testcases := map[string]string{
		"ihgpwlah": "370",
		"kglvqrro": "492",
		"ulqzkmiv": "830",
	}
	for input, expected := range testcases {
		solution, err := Part2(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}
