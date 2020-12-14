package day13

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `939
7,13,x,x,59,x,31,19`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "295", solution)
}

func TestPart2(t *testing.T) {
	testcases := map[string]string{
		"#\n17,x,13,19":      "3417",
		"#\n67,7,59,61":      "754018",
		"#\n67,x,7,59,61":    "779210",
		"#\n67,7,x,59,61":    "1261476",
		"#\n1789,37,47,1889": "1202161486",
	}
	for input, expected := range testcases {
		solution, err := Part2(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}
