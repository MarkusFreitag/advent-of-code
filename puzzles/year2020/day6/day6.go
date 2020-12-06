package day6

import (
	"strconv"
	"strings"
)

func uniqueQuestions(block string) map[string]bool {
	qs := make(map[string]bool)
	for _, line := range strings.Split(block, "\n") {
		for _, s := range line {
			qs[string(s)] = true
		}
	}
	return qs
}

func Part1(input string) (string, error) {
	var count int
	for _, block := range strings.Split(input, "\n\n") {
		count += len(uniqueQuestions(block))
	}
	return strconv.Itoa(count), nil
}

func Part2(input string) (string, error) {
	var count int
	for _, block := range strings.Split(input, "\n\n") {
		for q := range uniqueQuestions(block) {
			all := true
			for _, line := range strings.Split(block, "\n") {
				if !strings.Contains(line, q) {
					all = false
				}
			}
			if all {
				count++
			}
		}
	}
	return strconv.Itoa(count), nil
}
