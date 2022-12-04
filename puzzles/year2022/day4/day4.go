package day4

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

func parseLine(line string) (int, int, int, int) {
	pairs := strings.Split(line, ",")
	leftPair, rightPair := strings.Split(pairs[0], "-"), strings.Split(pairs[1], "-")
	return util.ParseInt(leftPair[0]),
		util.ParseInt(leftPair[1]),
		util.ParseInt(rightPair[0]),
		util.ParseInt(rightPair[1])
}

func fullOverlap(leftMin, leftMax, rightMin, rightMax int) bool {
	if numbers.Between(rightMin, leftMin, leftMax) && numbers.Between(rightMax, leftMin, leftMax) {
		return true
	}

	if numbers.Between(leftMin, rightMin, rightMax) && numbers.Between(leftMax, rightMin, rightMax) {
		return true
	}

	return false
}

func partialOverlap(leftMin, leftMax, rightMin, rightMax int) bool {
	if numbers.Between(rightMin, leftMin, leftMax) && rightMax >= leftMax {
		return true
	}

	if numbers.Between(leftMin, rightMin, rightMax) && leftMax >= rightMax {
		return true
	}

	return false
}

func Part1(input string) (string, error) {
	var total int
	for _, line := range strings.Split(input, "\n") {
		leftMin, leftMax, rightMin, rightMax := parseLine(line)

		if fullOverlap(leftMin, leftMax, rightMin, rightMax) {
			total++
		}
	}
	return strconv.Itoa(total), nil
}

func Part2(input string) (string, error) {
	var total int
	for _, line := range strings.Split(input, "\n") {
		leftMin, leftMax, rightMin, rightMax := parseLine(line)

		if fullOverlap(leftMin, leftMax, rightMin, rightMax) {
			total++
			continue
		}

		if partialOverlap(leftMin, leftMax, rightMin, rightMax) {
			total++
			continue
		}
	}
	return strconv.Itoa(total), nil
}
