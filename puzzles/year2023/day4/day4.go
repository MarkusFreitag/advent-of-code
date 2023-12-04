package day4

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/maputil"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

func lineScore(line string) (int, int, int) {
	parts := strings.Split(line, ":")
	lists := strings.Split(parts[1], "|")
	winning := make(map[int]bool)
	for _, num := range strings.Fields(lists[0]) {
		winning[util.ParseInt(num)] = true
	}
	having := util.StringsToInts(strings.Fields(lists[1]))
	var score, hits int
	for _, have := range having {
		if _, ok := winning[have]; ok {
			if score == 0 {
				score = 1
			} else {
				score = score * 2
			}
			hits++
		}
	}
	return util.ParseInt(strings.Fields(parts[0])[1]), hits, score
}

func Part1(input string) (string, error) {
	cards := make([]int, 0)
	for _, line := range strings.Split(input, "\n") {
		_, _, score := lineScore(line)
		cards = append(cards, score)
	}
	return strconv.Itoa(numbers.Sum(cards...)), nil
}

func Part2(input string) (string, error) {
	copies := make(map[int]int)
	for _, line := range strings.Split(input, "\n") {
		id, hits, _ := lineScore(line)

		v, ok := copies[id]
		if ok {
			copies[id] = v + 1
		} else {
			copies[id] = 1
		}

		for i := 1; i <= hits; i++ {
			copies[id+i] = copies[id+i] + copies[id]
		}
	}

	return strconv.Itoa(numbers.Sum(maputil.Values(copies)...)), nil
}
