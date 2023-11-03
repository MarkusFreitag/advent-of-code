package day10

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/slice"
)

func sub(list []int, idx, offset int) []int {
	s := make([]int, 0)
	for o := 0; o < offset; o++ {
		s = append(s, list[(idx+o)%len(list)])
	}
	return s
}

func process(list, steps []int, pos, skip int) (int, int) {
	for _, step := range steps {
		sublist := sub(list, pos, step)
		slice.Reverse(sublist)
		for si, n := range sublist {
			list[(pos+si)%len(list)] = n
		}
		pos = (pos + step + skip) % len(list)
		skip++
	}
	return pos, skip
}

func Part1(input string) (string, error) {
	steps := util.StringsToInts(strings.Split(input, ","))

	list := make([]int, 256)
	for i := 0; i < 256; i++ {
		list[i] = i
	}

	process(list, steps, 0, 0)

	return strconv.Itoa(list[0] * list[1]), nil
}

func Part2(input string) (string, error) {
	steps := make([]int, 0)
	for _, b := range []byte(input) {
		steps = append(steps, int(b))
	}
	steps = append(steps, []int{17, 31, 73, 47, 23}...)

	list := make([]int, 256)
	for i := 0; i < 256; i++ {
		list[i] = i
	}

	var pos int
	var skip int
	for r := 0; r < 64; r++ {
		pos, skip = process(list, steps, pos, skip)
	}

	var result string
	for _, chunk := range slice.Chunks(list, 16) {
		result += fmt.Sprintf("%02x", xor(chunk...))
	}

	return result, nil
}

func xor(nums ...int) int {
	var total int
	for _, n := range nums {
		total = total ^ n
	}
	return total
}
