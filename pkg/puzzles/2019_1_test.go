package puzzles

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestY2019D1P1(t *testing.T) {
	p := &y2019d1p1{}
	solution, err := p.Solve("12")
	require.Nil(t, err)
	require.Equal(t, "2", solution)
	solution, err = p.Solve("14")
	require.Nil(t, err)
	require.Equal(t, "2", solution)
	solution, err = p.Solve("1969")
	require.Nil(t, err)
	require.Equal(t, "654", solution)
	solution, err = p.Solve("100756")
	require.Nil(t, err)
	require.Equal(t, "33583", solution)
	solution, err = p.Solve("12\n14\n1969\n100756")
	require.Nil(t, err)
	require.Equal(t, "34241", solution)
}

func TestY2019D1P2(t *testing.T) {
	p := &y2019d1p2{}
	solution, err := p.Solve("14")
	require.Nil(t, err)
	require.Equal(t, "2", solution)
	solution, err = p.Solve("1969")
	require.Nil(t, err)
	require.Equal(t, "966", solution)
	solution, err = p.Solve("100756")
	require.Nil(t, err)
	require.Equal(t, "50346", solution)
	solution, err = p.Solve("14\n1969\n100756")
	require.Nil(t, err)
	require.Equal(t, "51314", solution)
}
