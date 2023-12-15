package day15

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "1320", solution)
}

func TestPart2(t *testing.T) {
	input := `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "145", solution)
}
