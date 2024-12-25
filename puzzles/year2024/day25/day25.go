package day25

import (
	"slices"
	"strconv"
	"strings"
)

func parsePins(lines []string) []int {
	pins := make([]int, len(lines[0]))
	for _, line := range lines {
		for idx, col := range line {
			if col == '#' {
				pins[idx] = pins[idx] + 1
			}
		}
	}
	return pins
}

func Part1(input string) (string, error) {
	locks := make([][]int, 0)
	keys := make([][]int, 0)
	for _, block := range strings.Split(input, "\n\n") {
		lines := strings.Fields(block)
		if lines[0] == strings.Repeat("#", len(lines[0])) {
			//lock
			locks = append(locks, parsePins(lines[1:]))
		}
		if lines[0] == strings.Repeat(".", len(lines[0])) {
			//key
			slices.Reverse(lines)
			keys = append(keys, parsePins(lines[1:]))
		}
	}
	var fits int
	for _, lock := range locks {
		for _, key := range keys {
			fit := true
			for p := 0; p < len(lock); p++ {
				if lock[p]+key[p] > 5 {
					fit = false
					break
				}
			}
			if fit {
				fits++
			}
		}
	}
	return strconv.Itoa(fits), nil
}

func Part2(input string) (string, error) {
	return "get all stars", nil
}
