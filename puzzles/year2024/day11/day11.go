package day11

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

func play(input string, times int) int {
	stones := make(map[int]int)
	for _, stone := range util.StringsToInts(strings.Fields(input)) {
		stones[stone] = 1
	}
	for range times {
		dup := make(map[int]int)
		for stone, count := range stones {
			if stone == 0 {
				dup[1] = dup[1] + count
			} else if digits := strconv.Itoa(stone); len(digits)%2 == 0 {
				left := util.ParseInt(digits[:len(digits)/2])
				right := util.ParseInt(digits[len(digits)/2:])
				dup[left] = dup[left] + count
				dup[right] = dup[right] + count
			} else {
				dup[stone*2024] = dup[stone*2024] + count
			}
		}
		stones = dup
	}
	var total int
	for _, count := range stones {
		total += count
	}
	return total
}

func Part1(input string) (string, error) {
	return strconv.Itoa(play(input, 25)), nil
}

func Part2(input string) (string, error) {
	return strconv.Itoa(play(input, 75)), nil
}
