package day7

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

func concat(a, b int) int {
	return util.ParseInt(strconv.Itoa(a) + strconv.Itoa(b))
}

func check(expected int, nums []int, cc bool) bool {
	if len(nums) == 2 {
		if (nums[0]+nums[1] == expected) || (nums[0]*nums[1] == expected) {
			return true
		}
		if cc && (concat(nums[0], nums[1]) == expected) {
			return true
		}
		return false
	}
	if check(expected, append([]int{nums[0] + nums[1]}, nums[2:]...), cc) {
		return true
	}
	if check(expected, append([]int{nums[0] * nums[1]}, nums[2:]...), cc) {
		return true
	}
	if cc && check(expected, append([]int{concat(nums[0], nums[1])}, nums[2:]...), cc) {
		return true
	}
	return false
}
func Part1(input string) (string, error) {
	var sum int
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ":")
		nums := util.StringsToInts(strings.Fields(parts[1]))
		val := util.ParseInt(parts[0])
		if check(val, nums, false) {
			sum += val
		}
	}
	return strconv.Itoa(sum), nil
}

func Part2(input string) (string, error) {
	var sum int
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ":")
		nums := util.StringsToInts(strings.Fields(parts[1]))
		val := util.ParseInt(parts[0])
		if check(val, nums, true) {
			sum += val
		}
	}
	return strconv.Itoa(sum), nil
}
