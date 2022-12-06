package day6

import (
	"strconv"

	"github.com/MarkusFreitag/advent-of-code/util/slice"
)

func check(buff string, window int) int {
	chars := []rune(buff)
	for slide := range slice.SlidingWindow(chars, window) {
		if len(slide.Values) < window {
			break
		}
		m := make(map[rune]bool)
		for _, char := range slide.Values {
			m[char] = true
		}
		if len(m) == window {
			return slide.Index + window
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
