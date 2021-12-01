package day1

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

func Part1(input string) (string, error) {
	var count int
	nums := util.StrsToInts(strings.Fields(input))
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			count++
		}
	}
	return strconv.Itoa(count), nil
}

func Part2(input string) (string, error) {
	var count int
	nums := util.StrsToInts(strings.Fields(input))
	lastSum := util.SumInts(nums[:3]...)
	for i := 1; i < len(nums); i++ {
		if i+2 >= len(nums) {
			break
		}
		newSum := util.SumInts(nums[i : i+3]...)
		if lastSum < newSum {
			count++
		}
		lastSum = newSum
	}
	return strconv.Itoa(count), nil
}
