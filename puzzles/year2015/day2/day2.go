package day2

import (
	"sort"
	"strconv"
	"strings"
)

func toNums(input string) ([]int, error) {
	parts := strings.Split(input, "x")
	nums := make([]int, len(parts))
	for idx, part := range parts {
		var err error
		nums[idx], err = strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
	}
	return nums, nil
}

func mulNums(nums []int) int {
	total := 1
	for _, num := range nums {
		total = total * num
	}
	return total
}

func Part1(input string) (string, error) {
	var total int
	for _, line := range strings.Split(input, "\n") {
		nums, err := toNums(line)
		if err != nil {
			return "", err
		}
		sort.Ints(nums)
		total += nums[0]*nums[1] + 2*nums[0]*nums[1] + 2*nums[1]*nums[2] + 2*nums[2]*nums[0]
	}
	return strconv.Itoa(total), nil
}

func Part2(input string) (string, error) {
	var total int
	for _, line := range strings.Split(input, "\n") {
		nums, err := toNums(line)
		if err != nil {
			return "", err
		}
		sort.Ints(nums)
		total += mulNums(nums) + 2*nums[0] + 2*nums[1]
	}
	return strconv.Itoa(total), nil
}
