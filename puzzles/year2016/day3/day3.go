package day3

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

type triangle [3]int

func (t triangle) check() bool {
	return t[0]+t[1] > t[2] && t[1]+t[2] > t[0] && t[0]+t[2] > t[1]
}

type triangles []triangle

func (t triangles) countValid() int {
	var counter int
	for _, tri := range t {
		if tri.check() {
			counter++
		}
	}
	return counter
}

func flatten(strs []string) []string {
	buf := make([]string, 0)
	for _, str := range strs {
		if str != "" {
			buf = append(buf, str)
		}
	}
	return buf
}

func Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")
	angles := make(triangles, len(lines))
	for idx, line := range lines {
		nums := util.StrsToInts(flatten(strings.Split(line, " ")))
		angles[idx] = triangle{nums[0], nums[1], nums[2]}
	}
	return strconv.Itoa(angles.countValid()), nil
}

func Part2(input string) (string, error) {
	lines := strings.Split(input, "\n")
	numbers := make([][]int, len(lines))
	for idx, line := range lines {
		numbers[idx] = util.StrsToInts(flatten(strings.Split(line, " ")))
	}
	angles := make(triangles, len(lines))
	for i := 0; i < len(numbers); i += 3 {
		for j := 0; j < len(numbers[i]); j++ {
			angles = append(angles, triangle{numbers[i][j], numbers[i+1][j], numbers[i+2][j]})
		}
	}
	return strconv.Itoa(angles.countValid()), nil
}
