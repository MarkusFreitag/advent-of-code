package day6

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	solution, err := Part1("COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L")
	require.Nil(t, err)
	require.Equal(t, "42", solution)
}

func TestPart2(t *testing.T) {
	solution, err := Part2("COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L\nK)YOU\nI)SAN")
	require.Nil(t, err)
	require.Equal(t, "4", solution)
}
