package day6

import (
	"strconv"

	"github.com/MarkusFreitag/advent-of-code/util/sliceutil"
)

func check(buff string, window int) int {
	chars := []rune(buff)
	for index, values := range sliceutil.SlidingWindow(chars, window) {
		if len(values) < window {
			break
		}
		m := make(map[rune]bool)
		for _, char := range values {
			m[char] = true
		}
		if len(m) == window {
			return index + window
		}
	}
	return -1
}

func Part1(input string) (string, error) {
	return strconv.Itoa(check(input, 4)), nil
}

func Part2(input string) (string, error) {
	return strconv.Itoa(check(input, 14)), nil
}
