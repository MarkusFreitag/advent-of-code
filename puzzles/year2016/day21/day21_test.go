package day21

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSwapPos(t *testing.T) {
	require.Equal(t, "ebcda", swap_pos("abcde", 4, 0))
	require.Equal(t, "ebcda", swap_pos("abcde", 0, 4))
	require.Equal(t, "acbde", swap_pos("abcde", 1, 2))
	require.Equal(t, "abdce", swap_pos("abcde", 3, 2))
}

func TestSwapLetters(t *testing.T) {
	require.Equal(t, "ebcda", swap_letters("abcde", "a", "e"))
	require.Equal(t, "adcbe", swap_letters("abcde", "b", "d"))
}

func TestReversePositions(t *testing.T) {
	require.Equal(t, "cbade", reverse_positions("abcde", 0, 2))
	require.Equal(t, "abedc", reverse_positions("abcde", 2, 4))
	require.Equal(t, "adcbe", reverse_positions("abcde", 1, 3))
	require.Equal(t, "adcbe", reverse_positions("abcde", 3, 1))
}

func TestRotate(t *testing.T) {
	require.Equal(t, "eabcd", rotate("abcde", "right", 1))
	require.Equal(t, "deabc", rotate("abcde", "right", 2))
	require.Equal(t, "cdeab", rotate("abcde", "right", 3))
	require.Equal(t, "bcdea", rotate("abcde", "right", 4))
	require.Equal(t, "abcde", rotate("abcde", "right", 5))
	require.Equal(t, "eabcd", rotate("abcde", "right", 6))
	require.Equal(t, "deabc", rotate("abcde", "right", 7))
	require.Equal(t, "bcdea", rotate("abcde", "left", 1))
	require.Equal(t, "cdeab", rotate("abcde", "left", 2))
	require.Equal(t, "deabc", rotate("abcde", "left", 3))
	require.Equal(t, "eabcd", rotate("abcde", "left", 4))
	require.Equal(t, "abcde", rotate("abcde", "left", 5))
	require.Equal(t, "bcdea", rotate("abcde", "left", 6))
	require.Equal(t, "cdeab", rotate("abcde", "left", 7))
}

func TestRotateOnLetter(t *testing.T) {
	require.Equal(t, "habcdefg", rotate_on_letter("abcdefgh", "a"))
	require.Equal(t, "ghabcdef", rotate_on_letter("abcdefgh", "b"))
	require.Equal(t, "fghabcde", rotate_on_letter("abcdefgh", "c"))
	require.Equal(t, "efghabcd", rotate_on_letter("abcdefgh", "d"))
	require.Equal(t, "cdefghab", rotate_on_letter("abcdefgh", "e"))
	require.Equal(t, "bcdefgha", rotate_on_letter("abcdefgh", "f"))
	require.Equal(t, "abcdefgh", rotate_on_letter("abcdefgh", "g"))
	require.Equal(t, "habcdefg", rotate_on_letter("abcdefgh", "h"))
}

func TestRotateOnLetterBack(t *testing.T) {
	require.Equal(t, "abcdefgh", rotate_on_letter_back("habcdefg", "a"))
	require.Equal(t, "abcdefgh", rotate_on_letter_back("ghabcdef", "b"))
	require.Equal(t, "abcdefgh", rotate_on_letter_back("fghabcde", "c"))
	require.Equal(t, "abcdefgh", rotate_on_letter_back("efghabcd", "d"))
	require.Equal(t, "abcdefgh", rotate_on_letter_back("cdefghab", "e"))
	require.Equal(t, "abcdefgh", rotate_on_letter_back("bcdefgha", "f"))
	require.Equal(t, "abcdefgh", rotate_on_letter_back("abcdefgh", "g"))
	require.Equal(t, "abcdefgh", rotate_on_letter_back("habcdefg", "h"))
}

func TestMovePos(t *testing.T) {
	require.Equal(t, "abcde", move_pos("abcde", 0, 0))
}

func TestPart1(t *testing.T) {
	input := `swap position 4 with position 0
swap letter d with letter b
reverse positions 0 through 4
rotate left 1 step
move position 1 to position 4
move position 3 to position 0
rotate based on position of letter b
rotate based on position of letter d`
	pass = "abcde"
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "decab", solution)
}

func TestPart2(t *testing.T) {
	testcases := map[string]string{}
	for input, expected := range testcases {
		solution, err := Part2(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}
