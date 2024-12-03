package day3

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

func Part1(input string) (string, error) {
	var total int
	rgx := regexp.MustCompile(`(?:mul\((\d{1,3}),(\d{1,3})\))`)
	for _, matches := range rgx.FindAllStringSubmatch(input, -1) {
		total += numbers.Multiply(util.ParseInt(matches[1]), util.ParseInt(matches[2]))
	}
	return strconv.Itoa(total), nil
}

func Part2(input string) (string, error) {
	var total int
	rgx := regexp.MustCompile(`(?:mul\((\d{1,3}),(\d{1,3})\))|(?:do\(\))|(?:don't\(\))`)
	enabled := true
	for _, matches := range rgx.FindAllStringSubmatch(input, -1) {
		if strings.HasPrefix(matches[0], "don't") {
			enabled = false
			continue
		}
		if strings.HasPrefix(matches[0], "do") {
			enabled = true
			continue
		}
		if !enabled {
			continue
		}
		total += numbers.Multiply(util.ParseInt(matches[1]), util.ParseInt(matches[2]))
	}
	return strconv.Itoa(total), nil
}
