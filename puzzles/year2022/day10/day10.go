package day10

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

func values(input string) []int {
	vals := make([]int, 0)
	vals = append(vals, 0)
	for _, line := range strings.Split(input, "\n") {
		fields := strings.Fields(line)
		vals = append(vals, 0)
		if fields[0] == "addx" {
			vals = append(vals, util.ParseInt(fields[1]))
		}
	}
	return vals
}

func Part1(input string) (string, error) {
	x := 1
	var sum int
	for idx, val := range values(input) {
		if idx == 20 || idx == 60 || idx == 100 || idx == 140 || idx == 180 || idx == 220 {
			sum += idx * x
		}
		x += val
	}

	return strconv.Itoa(sum), nil
}

func Part2(input string) (string, error) {
	x := 1
	var out string
	for idx, val := range values(input) {
		x += val

		if idx%40 == 0 {
			out += "\n"
		}
		if v := idx % 40; v == x-1 || v == x || v == x+1 {
			out += "#"
		} else {
			out += " "
		}
	}
	fmt.Println(out)

	return "", nil
}
