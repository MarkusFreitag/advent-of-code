package day2

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/maputil"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

func Part1(input string) (string, error) {
	var total int
	cubeCounts := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ":")
		p := true

		for _, set := range strings.Split(parts[1], ";") {
			for _, cube := range strings.Split(set, ",") {
				if fields := strings.Fields(cube); cubeCounts[fields[1]] < util.ParseInt(fields[0]) {
					p = false
					break
				}
			}
			if !p {
				break
			}
		}

		if p {
			total += util.ParseInt(strings.Fields(parts[0])[1])
		}
	}
	return strconv.Itoa(total), nil
}

func Part2(input string) (string, error) {
	var total int
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ":")
		cubeMax := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, set := range strings.Split(parts[1], ";") {
			for _, cube := range strings.Split(set, ",") {
				fields := strings.Fields(cube)
				cubeMax[fields[1]] = numbers.Max(cubeMax[fields[1]], util.ParseInt(fields[0]))
			}
		}

		total += numbers.Multiply(maputil.Values(cubeMax)...)
	}
	return strconv.Itoa(total), nil
}
