package day23

import (
	"strconv"

	"github.com/MarkusFreitag/advent-of-code/util"
)

func maxInts(nums []int) (int, int) {
	var max int
	var pos int
	for idx, num := range nums {
		if num > max {
			max = num
			pos = idx
		}
	}
	return pos, max
}

func destination(nums []int) int {
	goal := nums[0] - 1
	for {
		if goal < 0 {
			idx, _ := maxInts(nums)
			return idx
		}
		for idx, num := range nums {
			if num == goal {
				return idx
			}
		}
		goal--
	}
}

func Part1(input string) (string, error) {
	circle := util.DigitStrToInts(input)

	for i := 0; i < 100; i++ {
		three := circle[1:4]
		sub := append([]int{circle[0]}, circle[4:]...)

		dest := destination(sub) + 1

		circle = append(sub[:dest], append(three, sub[dest:]...)...)

		var first int
		first, circle = circle[0], circle[1:]
		circle = append(circle, first)
	}

	var pos int
	for idx, num := range circle {
		if num == 1 {
			pos = idx
		}
	}
	solution := circle[pos:]
	solution = append(solution, circle[:pos]...)

	return util.IntsToDigitStr(solution[1:]), nil
}

func Part2(input string) (string, error) {
	nums := util.DigitStrToInts(input)

	circle := make([]int, 1000000)
	_, max := maxInts(circle)
	for i := len(nums); i < len(circle); i++ {
		max++
		circle[i] = max
	}

	for i := 0; i < 10000000; i++ {
		three := circle[1:4]
		sub := append([]int{circle[0]}, circle[4:]...)

		dest := destination(sub) + 1

		circle = append(sub[:dest], append(three, sub[dest:]...)...)

		var first int
		first, circle = circle[0], circle[1:]
		circle = append(circle, first)
	}

	var pos int
	for idx, num := range circle {
		if num == 1 {
			pos = idx
		}
	}

	return strconv.Itoa(circle[pos+1] * circle[pos+2]), nil
}
