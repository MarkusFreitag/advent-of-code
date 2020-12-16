package day16

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "71", solution)
}

func TestPart2(t *testing.T) {
	input := `class: 1-3 or 5-7
departure row: 6-11 or 33-44
departure seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "98", solution)
}
