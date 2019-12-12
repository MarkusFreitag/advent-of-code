package day12

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	p := &Part1{}
	solution, err := p.Solve("12")
	require.Nil(t, err)
	require.Equal(t, "2", solution)
}
/*
func TestPart2(t *testing.T) {
	p := &Part2{}
	solution, err := p.Solve("14")
	require.Nil(t, err)
	require.Equal(t, "2", solution)
}
*/
