package day1

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

func Part1(input string) (string, error) {
	var count int
	nums := util.StringsToInts(strings.Fields(input))
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			count++
		}
	}
	return strconv.Itoa(count), nil
}

func Part2(input string) (string, error) {
	var count int
	nums := util.StringsToInts(strings.Fields(input))
	lastSum := numbers.Sum(nums[:3]...)
	for i := 1; i < len(nums); i++ {
		if i+2 >= len(nums) {
			break
		}
		newSum := numbers.Sum(nums[i : i+3]...)
		if lastSum < newSum {
			count++
		}
		lastSum = newSum
	}
	return strconv.Itoa(count), nil
}
