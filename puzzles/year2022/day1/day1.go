package day1

import (
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

func parseCalories(inventories []string) []int {
	elves := make([]int, 0)
	var total int
	for _, line := range inventories {
		if line == "" {
			elves = append(elves, total)
			total = 0
			continue
		}
		total += util.ParseInt(line)
	}
	elves = append(elves, total)
	sort.Ints(elves)
	slices.Reverse(elves)
	return elves
}

func Part1(input string) (string, error) {
	elves := parseCalories(strings.Split(input, "\n"))
	return strconv.Itoa(elves[0]), nil
}

func Part2(input string) (string, error) {
	elves := parseCalories(strings.Split(input, "\n"))
	return strconv.Itoa(numbers.Sum(elves[:3]...)), nil
}
