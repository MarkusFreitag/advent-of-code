package day9

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	p := &Part1{}

	solution, err := p.Solve("109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99")
	require.Nil(t, err)
	require.Equal(t, "[109 1 204 -1 1001 100 1 100 1008 100 16 101 1006 101 0 99]", solution)

	solution, err = p.Solve("1102,34915192,34915192,7,4,7,99,0")
	require.Nil(t, err)
	require.Equal(t, 18, len(solution))

	solution, err = p.Solve("104,1125899906842624,99")
	require.Nil(t, err)
	require.Equal(t, 18, len(solution))
}
