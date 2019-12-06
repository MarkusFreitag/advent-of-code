package puzzles

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestY2019D6P1(t *testing.T) {
	p := &y2019d6p1{}
	solution, err := p.Solve("COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L")
	require.Nil(t, err)
	require.Equal(t, "42", solution)
}

func TestY2019D6P2(t *testing.T) {
	p := &y2019d6p2{}
	solution, err := p.Solve("COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L\nK)YOU\nI)SAN")
	require.Nil(t, err)
	require.Equal(t, "4", solution)
}
