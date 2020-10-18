package day12

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	testcases := map[string]string{
		`[1,2,3]`:              "6",
		`{"a":2,"b":4}`:        "6",
		`[[[3]]]`:              "3",
		`{"a":{"b":4},"c":-1}`: "3",
		`{"a":[-1,1]}`:         "0",
		`[-1,{"a":1}]`:         "0",
		`[]`:                   "0",
		`{}`:                   "0",
	}
	for input, expected := range testcases {
		solution, err := Part1(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}

func TestPart2(t *testing.T) {
	testcases := map[string]string{
		`[1,2,3]`:                         "6",
		`[1,{"c":"red","b":2},3]`:         "4",
		`{"d":"red","e":[1,2,3,4],"f":5}`: "0",
		`[1,"red",5]`:                     "6",
	}
	for input, expected := range testcases {
		solution, err := Part2(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}
