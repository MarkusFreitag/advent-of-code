package day23

import (
	"fmt"
	"strconv"
	"time"
)

func strToInts(str string) []int {
	nums := make([]int, len(str))
	for idx, char := range str {
		num, _ := strconv.Atoi(string(char))
		nums[idx] = num
	}
	return nums
}

func intsToStr(nums []int) string {
	var str string
	for _, num := range nums {
		str += strconv.Itoa(num)
	}
	return str
}

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
	circle := strToInts(input)

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

	return intsToStr(solution[1:]), nil
}

func Part2(input string) (string, error) {
	nums := strToInts(input)

	circle := make([]int, 1000000)
	_, max := maxInts(circle)
	for i := len(nums); i < len(circle); i++ {
		max++
		circle[i] = max
	}

	for i := 0; i < 10000000; i++ {
		start := time.Now()
		three := circle[1:4]
		sub := append([]int{circle[0]}, circle[4:]...)

		dest := destination(sub) + 1

		circle = append(sub[:dest], append(three, sub[dest:]...)...)

		var first int
		first, circle = circle[0], circle[1:]
		circle = append(circle, first)
		if i%1000 == 0 {
			fmt.Println(i, time.Since(start))
		}
	}

	var pos int
	for idx, num := range circle {
		if num == 1 {
			pos = idx
		}
	}

	return strconv.Itoa(circle[pos+1] * circle[pos+2]), nil
}
