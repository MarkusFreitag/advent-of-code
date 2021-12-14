package day14

import (
	"strconv"
	"strings"
)

func solution(input string, steps int) int {
	blocks := strings.Split(input, "\n\n")

	polymer := blocks[0]

	lines := strings.Split(blocks[1], "\n")
	rules := make(map[string]string)
	for _, line := range lines {
		fields := strings.Fields(line)
		rules[fields[0]] = fields[2]
	}

	pairs := make(map[string]int)
	for i := 0; i < len(polymer)-1; i++ {
		pair := polymer[i : i+2]
		pairs[pair] = pairs[pair] + 1
	}

	for i := 1; i <= steps; i++ {
		newPairs := make(map[string]int)
		for pair, count := range pairs {
			insert := rules[pair]
			left := string(pair[0]) + insert
			right := insert + string(pair[1])

			newPairs[left] = newPairs[left] + count
			newPairs[right] = newPairs[right] + count
		}
		pairs = newPairs
	}

	chars := make(map[rune]int)
	chars[rune(polymer[0])] = chars[rune(polymer[0])] + 1
	for pair, count := range pairs {
		chars[rune(pair[1])] = chars[rune(pair[1])] + count
	}

	var most, least int
	for _, count := range chars {
		if count > most {
			most = count
		}
		if least == 0 || count < least {
			least = count
		}
	}

	return most - least
}

func Part1(input string) (string, error) {
	return strconv.Itoa(solution(input, 10)), nil
}

func Part2(input string) (string, error) {
	return strconv.Itoa(solution(input, 40)), nil
}
