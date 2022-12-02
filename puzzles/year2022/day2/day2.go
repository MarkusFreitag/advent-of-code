package day2

import (
	"strconv"
	"strings"
)

var (
	scores = map[string]int{
		"A":    1,
		"B":    2,
		"C":    3,
		"X":    1,
		"Y":    2,
		"Z":    3,
		"draw": 3,
		"win":  6,
	}
	win = map[string]string{
		"A": "Y",
		"B": "Z",
		"C": "X",
	}
	draw = map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}
	defeat = map[string]string{
		"A": "Z",
		"B": "X",
		"C": "Y",
	}
)

func Part1(input string) (string, error) {
	var total int
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Fields(line)

		total += scores[parts[1]]

		if v, ok := draw[parts[0]]; ok && v == parts[1] {
			total += scores["draw"]
		}

		if v, ok := win[parts[0]]; ok && v == parts[1] {
			total += scores["win"]
		}
	}

	return strconv.Itoa(total), nil
}

func Part2(input string) (string, error) {
	var total int
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Fields(line)

		switch parts[1] {
		case "X":
			total += scores[defeat[parts[0]]]
		case "Y":
			total += scores["draw"]
			total += scores[draw[parts[0]]]
		case "Z":
			total += scores["win"]
			total += scores[win[parts[0]]]
		}
	}

	return strconv.Itoa(total), nil
}
