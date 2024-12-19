package day19

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util/sliceutil"
)

func isPossible(design string, patterns []string) bool {
	for _, pattern := range patterns {
		if design == pattern {
			return true
		}
		if rest, ok := strings.CutPrefix(design, pattern); ok {
			if isPossible(rest, patterns) {
				return true
			}
		}
	}
	return false
}

func countCombinations(design string, patterns []string, cache map[string]int) int {
	if val, ok := cache[design]; ok {
		return val
	}
	var total int
	for _, pattern := range patterns {
		if pattern == design {
			total++
			continue
		}
		if rest, ok := strings.CutPrefix(design, pattern); ok {
			total += countCombinations(rest, patterns, cache)
		}
	}
	cache[design] = total
	return total
}

func Part1(input string) (string, error) {
	blocks := strings.Split(input, "\n\n")
	patterns := sliceutil.Map(
		strings.Fields(blocks[0]),
		func(s string) (string, bool) {
			return strings.TrimSuffix(s, ","), true
		},
	)
	var total int
	for _, line := range strings.Fields(blocks[1]) {
		if isPossible(line, patterns) {
			total++
		}
	}
	return strconv.Itoa(total), nil
}

func Part2(input string) (string, error) {
	blocks := strings.Split(input, "\n\n")
	patterns := sliceutil.Map(
		strings.Fields(blocks[0]),
		func(s string) (string, bool) {
			return strings.TrimSuffix(s, ","), true
		},
	)
	var total int
	for _, line := range strings.Fields(blocks[1]) {
		total += countCombinations(line, patterns, make(map[string]int))
	}
	return strconv.Itoa(total), nil
}
