package day2

import (
	"strconv"
	"strings"
)

func Part1(input string) (string, error) {
	var doubles, tripples int
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		letters := make(map[rune]int)
		for _, letter := range line {
			if _, ok := letters[letter]; ok {
				letters[letter] = letters[letter] + 1
			} else {
				letters[letter] = 1
			}
		}
		var doubleLetters, trippleLetters bool
		for _, count := range letters {
			if count == 2 {
				doubleLetters = true
			}
			if count == 3 {
				trippleLetters = true
			}
		}
		if doubleLetters {
			doubles++
		}
		if trippleLetters {
			tripples++
		}
	}
	return strconv.Itoa(doubles * tripples), nil
}

func Part2(input string) (string, error) {
	ids := strings.Split(input, "\n")
	for idx, id := range ids {
		for _, i := range ids[idx:] {
			this, next := id, i
			var differ int
			var same string
			for idx, letter := range this {
				if letter != rune(next[idx]) {
					differ++
				} else {
					same += string(letter)
				}
			}
			if differ == 1 {
				return same, nil
			}
		}
	}
	return "", nil
}
