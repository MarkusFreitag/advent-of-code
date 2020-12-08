package day8

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "5", solution)
}

func TestPart2(t *testing.T) {
	input := `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "8", solution)
}
