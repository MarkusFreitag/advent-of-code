package day2

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

func Part1(input string) (string, error) {
	var hori, vert int
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Fields(line)
		num := util.ParseInt(parts[1])
		switch parts[0] {
		case "forward":
			hori += num
		case "up":
			vert -= num
		case "down":
			vert += num
		}
	}
	return strconv.Itoa(hori * vert), nil
}

func Part2(input string) (string, error) {
	var hori, vert, aim int
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Fields(line)
		num := util.ParseInt(parts[1])
		switch parts[0] {
		case "forward":
			hori += num
			vert += aim * num
		case "up":
			aim -= num
		case "down":
			aim += num
		}
	}
	return strconv.Itoa(hori * vert), nil
}
