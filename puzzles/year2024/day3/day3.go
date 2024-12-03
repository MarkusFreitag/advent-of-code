package day3

import (
	"regexp"
	"strconv"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

func evalMemory(mem string, rgx *regexp.Regexp) int {
	var total int
	enabled := true
	for _, matches := range rgx.FindAllStringSubmatch(mem, -1) {
		switch matches[0] {
		case "don't()":
			enabled = false
		case "do()":
			enabled = true
		default:
			if enabled {
				total += numbers.Multiply(util.ParseInt(matches[1]), util.ParseInt(matches[2]))
			}
		}
	}
	return total
}

func Part1(input string) (string, error) {
	return strconv.Itoa(evalMemory(
		input,
		regexp.MustCompile(`(?:mul\((\d{1,3}),(\d{1,3})\))`),
	)), nil
}

func Part2(input string) (string, error) {
	return strconv.Itoa(evalMemory(
		input,
		regexp.MustCompile(`(?:mul\((\d{1,3}),(\d{1,3})\))|(?:do\(\))|(?:don't\(\))`),
	)), nil
}
