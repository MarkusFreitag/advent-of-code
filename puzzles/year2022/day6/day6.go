package day6

import (
	"strconv"
)

func check(buff string, window int) int {
	chars := []rune(buff)
	for idx := range chars {
		m := make(map[rune]bool)
		for i := 0; i < window; i++ {
			m[chars[idx+i]] = true
		}
		if len(m) == window {
			return idx + window
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
