package day18

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	ADD = "+"
	MUL = "*"
)

func findClosingBracket(task string, fromIndex int) int {
	var open int
	for i := fromIndex; i < len(task); i++ {
		switch string(task[i]) {
		case "(":
			open++
		case ")":
			if open > 0 {
				open--
			} else {
				return i
			}
		}
	}
	return -1
}

func calc(task string, precedence bool) int {
	if num, err := strconv.Atoi(task); err == nil {
		return num
	}
	for strings.Contains(task, "(") {
		oBracket := strings.Index(task, "(") + 1
		cBracket := findClosingBracket(task, oBracket)
		subtask := task[oBracket:cBracket]
		task = strings.Replace(task, "("+subtask+")", strconv.Itoa(calc(subtask, precedence)), -1)
	}

	if precedence {
		for strings.Contains(task, "+") && strings.Contains(task, "*") {
			rgx := regexp.MustCompile(`(\d+\s\+\s\d+)`)
			matches := rgx.FindAllStringSubmatch(task, 1)[0]
			task = strings.Replace(task, matches[1], strconv.Itoa(calc(matches[1], precedence)), 1)
		}
	}

	op := ADD
	var result int
	for _, field := range strings.Fields(task) {
		if num, err := strconv.Atoi(field); err == nil {
			switch op {
			case ADD:
				result += num
			case MUL:
				result *= num
			}
		} else {
			op = field
		}
	}
	return result
}

func Part1(input string) (string, error) {
	var sum int
	for _, line := range strings.Split(input, "\n") {
		sum += calc(line, false)
	}
	return strconv.Itoa(sum), nil
}

func Part2(input string) (string, error) {
	var sum int
	for _, line := range strings.Split(input, "\n") {
		sum += calc(line, true)
	}
	return strconv.Itoa(sum), nil
}
