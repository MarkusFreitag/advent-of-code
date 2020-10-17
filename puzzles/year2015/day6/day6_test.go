package day6

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGrid(t *testing.T) {
	g := newGrid(10, 10)
	require.Equal(t, 0, g.Count(1))
	require.Equal(t, 100, g.Count(0))
	require.Equal(t, 0, g.Sum())

	g[5][5] = 1
	require.Equal(t, 1, g.Count(1))
	require.Equal(t, 99, g.Count(0))
	require.Equal(t, 1, g.Sum())

	g[0][0] = 1
	g[5][5] = 2
	require.Equal(t, 3, g.Sum())
}

func TestSplitNums(t *testing.T) {
	require.Equal(t, []int{1, 2}, splitNums("1,2"))
}

func TestParseLine(t *testing.T) {
	instr, start, end := parseLine("toggle 0,0 through 1,2")
	require.Equal(t, "toggle", instr)
	require.Equal(t, []int{0, 0}, start)
	require.Equal(t, []int{1, 2}, end)

	instr, start, end = parseLine("turn on 0,0 through 1,2")
	require.Equal(t, "on", instr)
	require.Equal(t, []int{0, 0}, start)
	require.Equal(t, []int{1, 2}, end)
}

func TestPart1(t *testing.T) {
	testcases := map[string]string{
		"turn on 0,0 through 999,999":      "1000000",
		"toggle 0,0 through 999,0":         "1000",
		"turn off 499,499 through 500,500": "0",
	}
	for input, expected := range testcases {
		solution, err := Part1(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}

func TestPart2(t *testing.T) {
	testcases := map[string]string{
		"turn on 0,0 through 0,0":    "1",
		"toggle 0,0 through 999,999": "2000000",
	}
	for input, expected := range testcases {
		solution, err := Part2(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}
