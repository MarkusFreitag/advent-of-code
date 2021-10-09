package day21

import (
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

func swap_pos(str string, i, j int) string {
	runes := []rune(str)
	runes[i], runes[j] = runes[j], runes[i]
	return string(runes)
}

func swap_letters(str, a, b string) string {
	i := strings.Index(str, a)
	j := strings.Index(str, b)
	return swap_pos(str, i, j)
}

func rotate(str, dir string, steps int) string {
	if steps < 0 || len(str) == 0 {
		return str
	}

	steps = len(str) - steps%len(str)

	var result string
	switch dir {
	case "right":
		result = str[steps:] + str[:steps]
	case "left":
		steps = len(str) - steps
		result = str[steps:] + str[:steps]
	}
	return result
}

func rotate_on_letter(str, letter string) string {
	index := strings.Index(str, letter)
	if index >= 4 {
		index++
	}
	return rotate(str, "right", index+1)
}

func rotate_on_letter_back(str, letter string) string {
	index := strings.Index(str, letter)
	if index%2 != 0 {
		return rotate(str, "left", index/2+1)
	}
	var steps int
	switch index {
	case 2:
		steps = 2
	case 4:
		steps = 1
	case 6:
		steps = 0
	case 0:
		steps = 7
	}
	return rotate(str, "right", steps)
}

func reverse_positions(str string, i, j int) string {
	if i > j {
		i, j = j, i
	}
	return str[:i] + util.Reverse(str[i:j+1]) + str[j+1:]
}

func move_pos(str string, i, j int) string {
	r := rune(str[i])
	runes := []rune(str)
	runes = append(runes[:i], runes[i+1:]...)
	runes = append(runes[:j], append([]rune{r}, runes[j:]...)...)
	return string(runes)
}

var pass = "abcdefgh"

func Part1(input string) (string, error) {
	result := pass
	for _, line := range strings.Split(input, "\n") {
		fields := strings.Fields(line)
		switch fields[0] {
		case "swap":
			if fields[1] == "position" {
				result = swap_pos(result, util.ParseInt(fields[2]), util.ParseInt(fields[5]))
			} else if fields[1] == "letter" {
				result = swap_letters(result, fields[2], fields[5])
			}
		case "reverse":
			result = reverse_positions(result, util.ParseInt(fields[2]), util.ParseInt(fields[4]))
		case "rotate":
			if fields[1] == "based" {
				result = rotate_on_letter(result, fields[6])
			} else {
				result = rotate(result, fields[1], util.ParseInt(fields[2]))
			}
		case "move":
			result = move_pos(result, util.ParseInt(fields[2]), util.ParseInt(fields[5]))
		}
	}
	return result, nil
}

func Part2(input string) (string, error) {
	result := "fbgdceah"
	lines := strings.Split(input, "\n")
	for idx := len(lines) - 1; idx >= 0; idx-- {
		fields := strings.Fields(lines[idx])
		switch fields[0] {
		case "swap":
			if fields[1] == "position" {
				result = swap_pos(result, util.ParseInt(fields[2]), util.ParseInt(fields[5]))
			} else if fields[1] == "letter" {
				result = swap_letters(result, fields[2], fields[5])
			}
		case "reverse":
			result = reverse_positions(result, util.ParseInt(fields[2]), util.ParseInt(fields[4]))
		case "rotate":
			if fields[1] == "based" {
				result = rotate_on_letter_back(result, fields[6])
			} else {
				if fields[1] == "right" {
					result = rotate(result, "left", util.ParseInt(fields[2]))
				} else {
					result = rotate(result, "right", util.ParseInt(fields[2]))
				}
			}
		case "move":
			result = move_pos(result, util.ParseInt(fields[5]), util.ParseInt(fields[2]))
		}
	}
	return result, nil
}
