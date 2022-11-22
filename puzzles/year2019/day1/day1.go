package day1

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

func Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")
	values := make([]int, len(lines))
	for idx, line := range lines {
		mass := util.ParseInt(line)
		values[idx] = mass/3 - 2
	}
	return strconv.Itoa(numbers.Sum(values...)), nil
}

func Part2(input string) (string, error) {
	lines := strings.Split(input, "\n")
	values := make([]int, 0)
	for _, line := range lines {
		mass := util.ParseInt(line)
		fuel := mass/3 - 2
		for fuel > 0 {
			values = append(values, fuel)
			fuel = fuel/3 - 2
		}
	}
	return strconv.Itoa(numbers.Sum(values...)), nil
}
