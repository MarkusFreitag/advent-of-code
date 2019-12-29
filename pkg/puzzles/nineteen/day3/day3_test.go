package day3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	solution, err := Part1("R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83")
	require.Nil(t, err)
	require.Equal(t, "159", solution)
	solution, err = Part1("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7")
	require.Nil(t, err)
	require.Equal(t, "135", solution)
}

func TestPart2(t *testing.T) {
	solution, err := Part2("R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83")
	require.Nil(t, err)
	require.Equal(t, "610", solution)
	solution, err = Part2("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7")
	require.Nil(t, err)
	require.Equal(t, "410", solution)
}
