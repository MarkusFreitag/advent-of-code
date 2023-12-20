package day20

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "32000000", solution)

	input = `broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output`
	solution, err = Part1(input)
	require.Nil(t, err)
	require.Equal(t, "11687500", solution)
}
