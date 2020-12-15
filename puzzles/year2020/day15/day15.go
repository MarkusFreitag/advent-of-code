package day15

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

func playMemory(startingNums []int, rounds int) int {
	numbers := make(map[int][]int)
	var turn, lastNum int
	for {
		var num int
		if turn < len(startingNums) {
			num = startingNums[turn]
		} else {
			n := lastNum
			v, ok := numbers[n]
			if !ok || len(v) == 1 {
				num = 0
			} else {
				num = v[len(v)-1] - v[len(v)-2]
			}
		}

		if turn == rounds-1 {
			return num
		}
		if v, ok := numbers[num]; ok {
			numbers[num] = append(v, turn)
		} else {
			slice := make([]int, 0)
			numbers[num] = append(slice, turn)
		}
		lastNum = num
		turn++
	}
}

func Part1(input string) (string, error) {
	nums := util.StrsToInts(strings.Split(input, ","))
	return strconv.Itoa(playMemory(nums, 2020)), nil
}

func Part2(input string) (string, error) {
	nums := util.StrsToInts(strings.Split(input, ","))
	return strconv.Itoa(playMemory(nums, 30000000)), nil
}
