package day2

import (
	"regexp"
	"strconv"
	"strings"
)

var rgx = regexp.MustCompile(`(\d+)-(\d+)\s([a-z]):\s([a-z]+)`)

func Part1(input string) (string, error) {
	var valid int
	for _, line := range strings.Split(input, "\n") {
		matches := rgx.FindStringSubmatch(line)

		char := matches[3]
		pass := matches[4]

		min, _ := strconv.Atoi(matches[1])
		max, _ := strconv.Atoi(matches[2])

		charCount := strings.Count(pass, char)
		if charCount >= min && charCount <= max {
			valid++
		}
	}
	return strconv.Itoa(valid), nil
}

func Part2(input string) (string, error) {
	var valid int
	for _, line := range strings.Split(input, "\n") {
		matches := rgx.FindStringSubmatch(line)

		char := matches[3]
		pass := matches[4]
		firstPos, _ := strconv.Atoi(matches[1])
		secondPos, _ := strconv.Atoi(matches[2])

		a := string(pass[firstPos-1])
		b := string(pass[secondPos-1])
		if a != b && (a == char || b == char) {
			valid++
		}
	}
	return strconv.Itoa(valid), nil
}
