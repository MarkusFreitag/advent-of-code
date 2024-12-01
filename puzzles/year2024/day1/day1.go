package day1

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
	"github.com/MarkusFreitag/advent-of-code/util/sliceutil"
)

func parseInput(input string) ([]int, []int) {
	left := make([]int, 0)
	right := make([]int, 0)
	for _, line := range strings.Split(input, "\n") {
		nums := strings.Fields(line)
		left = append(left, util.ParseInt(nums[0]))
		right = append(right, util.ParseInt(nums[len(nums)-1]))
	}
	return left, right
}

func Part1(input string) (string, error) {
	left, right := parseInput(input)
	sliceutil.SortAsc(left)
	sliceutil.SortAsc(right)
	var dist int
	for idx, num := range left {
		dist += numbers.Abs(num - right[idx])
	}
	return strconv.Itoa(dist), nil
}

func Part2(input string) (string, error) {
	left, right := parseInput(input)
	var score int
	for _, num := range left {
		score += num * sliceutil.Count(right, num)
	}
	return strconv.Itoa(score), nil
}
