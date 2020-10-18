package day8

import (
	"strconv"
	"strings"
)

func Part1(input string) (string, error) {
	var code, memory int
	for _, line := range strings.Split(input, "\n") {
		code += len(line)
		str, _ := strconv.Unquote(line)
		memory += len(str)
	}
	return strconv.Itoa(code - memory), nil
}

func Part2(input string) (string, error) {
	var newStr, origStr int
	for _, line := range strings.Split(input, "\n") {
		origStr += len(line)
		newStr += len(strconv.Quote(line))
	}
	return strconv.Itoa(newStr - origStr), nil
}
