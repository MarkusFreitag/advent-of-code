package day1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `1721
979
366
299
675
1456`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "514579", solution)
}

func TestPart2(t *testing.T) {
	input := `1721
979
366
299
675
1456`
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "241861950", solution)
}
