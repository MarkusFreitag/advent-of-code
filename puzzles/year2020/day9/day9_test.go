package day9

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	preambleSize = 5
	input := `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "127", solution)
}

func TestPart2(t *testing.T) {
	preambleSize = 5
	input := `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "62", solution)
}
