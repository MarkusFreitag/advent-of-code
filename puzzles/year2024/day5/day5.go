package day5

import (
	"slices"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

func filterPageUpdates(input string) (map[int][][2]int, [][]int, [][]int) {
	blocks := strings.Split(input, "\n\n")
	rules := make(map[int][][2]int, 0)
	for _, line := range strings.Split(blocks[0], "\n") {
		parts := strings.Split(line, "|")
		rule := [2]int{util.ParseInt(parts[0]), util.ParseInt(parts[1])}
		if val, ok := rules[rule[0]]; ok {
			rules[rule[0]] = append(val, rule)
		} else {
			rules[rule[0]] = [][2]int{rule}
		}
	}
	validUpdates := make([][]int, 0)
	invalidUpdates := make([][]int, 0)
	for _, line := range strings.Split(blocks[1], "\n") {
		update := util.StringsToInts(strings.Split(line, ","))
		if slices.IsSortedFunc(update, sortByRulesFunc(rules)) {
			validUpdates = append(validUpdates, update)
		} else {
			invalidUpdates = append(invalidUpdates, update)
		}
	}
	return rules, validUpdates, invalidUpdates
}

func sortByRulesFunc(rules map[int][][2]int) func(int, int) int {
	return func(a, b int) int {
		if vals, ok := rules[a]; ok {
			for _, rule := range vals {
				if rule[0] == a && rule[1] == b {
					return -1
				}
				if rule[0] == b && rule[1] == a {
					return +1
				}
			}
		}
		return 0
	}
}

func Part1(input string) (string, error) {
	_, validUpdates, _ := filterPageUpdates(input)
	var sum int
	for _, update := range validUpdates {
		sum += update[len(update)/2]
	}
	return strconv.Itoa(sum), nil
}

func Part2(input string) (string, error) {
	rules, _, invalidUpdates := filterPageUpdates(input)
	var sum int
	for _, update := range invalidUpdates {
		slices.SortFunc(update, sortByRulesFunc(rules))
		sum += update[len(update)/2]
	}
	return strconv.Itoa(sum), nil
}
