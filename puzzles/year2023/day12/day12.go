package day12

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/maputil"
)

func countArrangements(left string, nums []int) int {
	arrangement := "."
	for _, num := range nums {
		arrangement += strings.Repeat("#", num)
		arrangement += "."
	}

	states := make(map[int]int)
	states[0] = 1
	newStates := make(map[int]int)
	for _, char := range left {
		for key, value := range states {
			if char == '?' {
				if key+1 < len(arrangement) {
					newStates[key+1] = maputil.Get(newStates, key+1, 0) + value
				}
				if arrangement[key] == '.' {
					newStates[key] = maputil.Get(newStates, key, 0) + value
				}
			} else if char == '.' {
				if key+1 < len(arrangement) && arrangement[key+1] == '.' {
					newStates[key+1] = maputil.Get(newStates, key+1, 0) + value
				}
				if arrangement[key] == '.' {
					newStates[key] = maputil.Get(newStates, key, 0) + value
				}
			} else if char == '#' {
				if key+1 < len(arrangement) && arrangement[key+1] == '#' {
					newStates[key+1] = maputil.Get(newStates, key+1, 0) + value
				}
			}
		}

		states = newStates
		newStates = make(map[int]int)
	}

	return maputil.Get(states, len(arrangement)-1, 0) + maputil.Get(states, len(arrangement)-2, 0)
}

func Part1(input string) (string, error) {
	var total int
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		total += countArrangements(parts[0], util.StringsToInts(strings.Split(parts[1], ",")))
	}
	return strconv.Itoa(total), nil
}

func Part2(input string) (string, error) {
	var total int
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")

		left := strings.Join(util.Repeat(parts[0], 5), "?")
		nums := util.StringsToInts(strings.Split(strings.Join(util.Repeat(parts[1], 5), ","), ","))

		total += countArrangements(left, nums)
	}
	return strconv.Itoa(total), nil
}
