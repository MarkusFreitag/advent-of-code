package day21

import (
	"strconv"
	"strings"
)

func parseInput(input string) map[string]string {
	monkeys := make(map[string]string)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ": ")
		monkeys[parts[0]] = parts[1]
	}
	return monkeys
}

func calc(monkeys map[string]string, name string) int {
	op := monkeys[name]
	num, err := strconv.Atoi(op)
	if err == nil {
		return num
	}

	fields := strings.Fields(op)
	switch fields[1] {
	case "+":
		return calc(monkeys, fields[0]) + calc(monkeys, fields[2])
	case "-":
		return calc(monkeys, fields[0]) - calc(monkeys, fields[2])
	case "*":
		return calc(monkeys, fields[0]) * calc(monkeys, fields[2])
	case "/":
		return calc(monkeys, fields[0]) / calc(monkeys, fields[2])
	}

	return 0
}

func Part1(input string) (string, error) {
	monkeys := parseInput(input)
	return strconv.Itoa(calc(monkeys, "root")), nil
}

func Part2(input string) (string, error) {
	monkeys := parseInput(input)

	fields := strings.Fields(monkeys["root"])
	left, right := fields[0], fields[2]

	min, max := 0, int(1e15)

	for min < max {
		middle := (min + max) / 2
		monkeys["humn"] = strconv.Itoa(middle)
		leftVal, rightVal := calc(monkeys, left), calc(monkeys, right)
		if leftVal == rightVal {
			return strconv.Itoa(middle), nil
		}
		if leftVal > rightVal {
			min = middle + 1
		} else {
			max = middle
		}
	}
	return "", nil
}
