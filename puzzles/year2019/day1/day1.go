package day1

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

func Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")
	values := make([]int, len(lines))
	for idx, line := range lines {
		mass, err := strconv.Atoi(line)
		if err != nil {
			return "", err
		}
		values[idx] = mass/3 - 2
	}
	return strconv.Itoa(util.Sum(values...)), nil
}

func Part2(input string) (string, error) {
	lines := strings.Split(input, "\n")
	values := make([]int, 0)
	for _, line := range lines {
		mass, err := strconv.Atoi(line)
		if err != nil {
			return "", err
		}
		fuel := mass/3 - 2
		for fuel > 0 {
			values = append(values, fuel)
			fuel = fuel/3 - 2
		}
	}
	return strconv.Itoa(util.Sum(values...)), nil
}
