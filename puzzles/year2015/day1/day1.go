package day1

import "strconv"

func Part1(input string) (string, error) {
	var floor int
	for _, char := range input {
		switch char {
		case '(':
			floor++
		case ')':
			floor--
		}
	}
	return strconv.Itoa(floor), nil
}

func Part2(input string) (string, error) {
	var floor int
	for idx, char := range input {
		switch char {
		case '(':
			floor++
		case ')':
			floor--
		}
		if floor < 0 {
			return strconv.Itoa(idx + 1), nil
		}
	}
	return "", nil
}
