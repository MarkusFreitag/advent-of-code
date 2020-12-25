package day25

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

func Part1(input string) (string, error) {
	nums := util.StrsToInts(strings.Fields(input))

	var loopCounter int
	for k := 1; k != nums[0]; loopCounter++ {
		k = k * 7 % 20201227
	}

	key := 1
	for i := 0; i < loopCounter; i++ {
		key = key * nums[1] % 20201227
	}

	return strconv.Itoa(key), nil
}

func Part2(input string) (string, error) {
	return "not solved yet", nil
}
