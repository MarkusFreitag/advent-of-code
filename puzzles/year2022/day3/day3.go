package day3

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util/numbers"
	"github.com/MarkusFreitag/advent-of-code/util/slice"
)

func convert(r rune) int {
	val := int(r)
	if val < 97 {
		return val - 38
	}
	return val - 96
}

func Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")
	prios := make([]int, 0)
	for _, line := range lines {
		length := len(line)
		left, right := line[:length/2], line[length/2:]
		for _, item := range left {
			if strings.ContainsRune(right, item) {
				prios = append(prios, convert(item))
				break
			}
		}
	}
	return strconv.Itoa(numbers.Sum(prios...)), nil
}

func shortest(strs []string) string {
	var str string
	length := numbers.MaxInteger
	for _, s := range strs {
		l := len(s)
		if l < length {
			length = l
			str = s
		}
	}
	return str
}

func Part2(input string) (string, error) {
	lines := strings.Split(input, "\n")
	prios := make([]int, 0)
	chunks := slice.Chunks(lines, 3)
	for _, chunk := range chunks {
		for _, item := range shortest(chunk) {
			found := make([]bool, 3)
			for idx, line := range chunk {
				found[idx] = strings.ContainsRune(line, item)
			}
			if slice.All(found, true) {
				prios = append(prios, convert(item))
				break
			}
		}
	}
	return strconv.Itoa(numbers.Sum(prios...)), nil
}
