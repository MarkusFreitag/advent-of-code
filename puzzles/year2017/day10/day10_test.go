package day10

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProcess(t *testing.T) {

	list := []int{0, 1, 2, 3, 4}
	steps := []int{3, 4, 1, 5}
	var pos, skip int
	pos, skip = process(list, steps, pos, skip)
	require.Equal(t, []int{3, 4, 2, 1, 0}, list)
	require.Equal(t, 4, pos)
	require.Equal(t, 4, skip)
}

func TestXOR(t *testing.T) {
	require.Equal(t, 64, xor(65, 27, 9, 1, 4, 3, 40, 50, 91, 7, 6, 0, 2, 5, 68, 22))
}
