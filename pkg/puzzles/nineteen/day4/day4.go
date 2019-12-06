package day4

import (
	"strconv"
	"strings"
)

func includeDouble(num string) bool {
	for i := 1; i < 10; i++ {
		if strings.Contains(num, strings.Repeat(strconv.Itoa(i), 2)) {
			return true
		}
	}
	return false
}

func onlyDoubles(num string) bool {
	var ok bool
	for i := 1; i < 10; i++ {
		if strings.Contains(num, strings.Repeat(strconv.Itoa(i), 2)) && !strings.Contains(num, strings.Repeat(strconv.Itoa(i), 3)) {
			ok = true
		}
	}
	return ok
}

func onlyIncreasing(num string) bool {
	for idx, i := range num {
		if idx == len(num)-1 {
			break
		}
		n, _ := strconv.Atoi(string(i))
		m, _ := strconv.Atoi(string(num[idx+1]))
		if n > m {
			return false
		}
	}
	return true
}

type Part1 struct{}

func (p *Part1) Solve(input string) (string, error) {
	parts := strings.Split(input, "-")
	min, err := strconv.Atoi(parts[0])
	if err != nil {
		return "", err
	}
	max, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", err
	}
	numbers := make([]string, 0)
	for i := min; i <= max; i++ {
		num := strconv.Itoa(i)
		if includeDouble(num) && onlyIncreasing(num) {
			numbers = append(numbers, num)
		}
	}
	return strconv.Itoa(len(numbers)), nil
}

type Part2 struct{}

func (p *Part2) Solve(input string) (string, error) {
	parts := strings.Split(input, "-")
	min, err := strconv.Atoi(parts[0])
	if err != nil {
		return "", err
	}
	max, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", err
	}
	numbers := make([]string, 0)
	for i := min; i <= max; i++ {
		num := strconv.Itoa(i)
		if onlyDoubles(num) && onlyIncreasing(num) {
			numbers = append(numbers, num)
		}
	}
	return strconv.Itoa(len(numbers)), nil
}
