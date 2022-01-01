package day11

import (
	"regexp"
	"strconv"
	"strings"
)

func parseInput(input string) []map[string]bool {
	rgxGenerator := regexp.MustCompile(`a\s([a-z]+)\sgenerator`)
	rgxMicrochip := regexp.MustCompile(`a\s([a-z]+)-compatible\smicrochip`)
	lines := strings.Split(input, "\n")
	floors := make([]map[string]bool, len(lines))
	for idx, line := range lines {
		floor := make(map[string]bool)
		matches := rgxGenerator.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			floor[match[1][:1]+"g"] = true
		}
		matches = rgxMicrochip.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			floor[match[1][:1]+"m"] = true
		}
		floors[idx] = floor
	}
	return floors
}

func move(floors []map[string]bool, elevator int) (int, bool) {
	return 1, true
}

func Part1(input string) (string, error) {
	floors := parseInput(input)
	steps, _ := move(floors, 0)
	return strconv.Itoa(steps), nil
}

func Part2(input string) (string, error) {
	return "n/a", nil
}
