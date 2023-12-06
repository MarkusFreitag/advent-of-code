package day6

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

func simulateRace(time, recordDistance int) int {
	var count int
	for t := 0; t < time; t++ {
		speed := t
		if (time-t)*speed > recordDistance {
			count++
		}
	}
	return count
}

func Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")
	times := util.StringsToInts(strings.Fields(strings.TrimPrefix(lines[0], "Time:")))
	distances := util.StringsToInts(strings.Fields(strings.TrimPrefix(lines[1], "Distance:")))

	holds := make([]int, len(times))
	for idx, time := range times {
		holds[idx] = simulateRace(time, distances[idx])
	}

	return strconv.Itoa(numbers.Multiply(holds...)), nil
}

func Part2(input string) (string, error) {
	lines := strings.Split(input, "\n")
	time := util.ParseInt(strings.Join(strings.Fields(strings.TrimPrefix(lines[0], "Time:")), ""))
	distance := util.ParseInt(strings.Join(strings.Fields(strings.TrimPrefix(lines[1], "Distance:")), ""))

	return strconv.Itoa(simulateRace(time, distance)), nil
}
