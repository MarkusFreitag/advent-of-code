package day1

import (
	"strconv"
	"strings"
)

func Part1(input string) (string, error) {
	var freq int
	for _, change := range parseChanges(input) {
		freq += change
	}
	return strconv.Itoa(freq), nil
}

func Part2(input string) (string, error) {
	freqs := make([]int, 1)
	freqs[0] = 0
	changes := parseChanges(input)
	for {
		for _, change := range changes {
			lastFreq := freqs[len(freqs)-1]
			lastFreq += change
			if intInSlice(lastFreq, freqs) {
				return strconv.Itoa(lastFreq), nil
			}
			freqs = append(freqs, lastFreq)
		}
	}
}

func parseChanges(input string) []int {
	changes := make([]int, 0)
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		num, _ := strconv.Atoi(line[1:])
		if line[:1] == "-" {
			num = num * -1
		}
		changes = append(changes, num)
	}
	return changes
}

func intInSlice(num int, slice []int) bool {
	for _, i := range slice {
		if i == num {
			return true
		}
	}
	return false
}
