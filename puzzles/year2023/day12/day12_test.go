package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	testcases := map[string]string{
		"???.### 1,1,3":             "1",
		".??..??...?##. 1,1,3":      "4",
		"?#?#?#?#?#?#?#? 1,3,1,6":   "1",
		"????.#...#... 4,1,1":       "1",
		"????.######..#####. 1,6,5": "4",
		"?###???????? 3,2,1":        "10",
	}
	for input, expected := range testcases {
		solution, err := Part1(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}

	input := `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "21", solution)
}

func TestPart2(t *testing.T) {
	testcases := map[string]string{
		"???.### 1,1,3":             "1",
		".??..??...?##. 1,1,3":      "16384",
		"?#?#?#?#?#?#?#? 1,3,1,6":   "1",
		"????.#...#... 4,1,1":       "16",
		"????.######..#####. 1,6,5": "2500",
		"?###???????? 3,2,1":        "506250",
	}
	for input, expected := range testcases {
		solution, err := Part2(input)
		require.Nil(t, err)
		assert.Equal(t, expected, solution)
	}

	input := `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "525152", solution)
}
