package day9

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

func decompressedLength(str string, incr bool) int {
	var length int
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			closingBracket := i + strings.Index(str[i:], ")")
			marker := str[i+1 : closingBracket]
			nums := util.StrsToInts(strings.Split(marker, "x"))
			if incr {
				length += nums[1] * decompressedLength(
					str[closingBracket+1:closingBracket+nums[0]+1],
					true,
				)
			} else {
				length += nums[1] * len(str[closingBracket+1:closingBracket+nums[0]+1])
			}
			i = closingBracket + nums[0]
		} else {
			length++
		}
	}
	return length
}

func Part1(input string) (string, error) {
	var length int
	for _, line := range strings.Split(input, "\n") {
		length += decompressedLength(line, false)
	}
	return strconv.Itoa(length), nil
}

func Part2(input string) (string, error) {
	var length int
	for _, line := range strings.Split(input, "\n") {
		length += decompressedLength(line, true)
	}
	return strconv.Itoa(length), nil
}
