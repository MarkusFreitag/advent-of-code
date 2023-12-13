package day13

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

func check(pattern []string, smudge bool) int {
	for i := 1; i < len(pattern); i++ {
		tmp := make([]string, len(pattern))
		copy(tmp, pattern)

		left, right := tmp[:i], tmp[i:]

		min := numbers.Min(len(left), len(right))
		left = left[len(left)-min:]
		right = right[:min]
		slices.Reverse(right)

		if !smudge && strings.Join(left, "\n") == strings.Join(right, "\n") {
			return i
		}
		if smudge && util.StringDiff(strings.Join(left, "\n"), strings.Join(right, "\n")) == 1 {
			return i
		}
	}
	return -1
}

func rotate(pattern []string) []string {
	rotated := make([]string, len(pattern[0]))
	for i := 0; i < len(pattern[0]); i++ {
		for j := range pattern {
			rotated[i] += string(pattern[len(pattern)-1-j][i])
		}
	}
	return rotated
}

func Part1(input string) (string, error) {
	blocks := strings.Split(input, "\n\n")
	patterns := make([][]string, len(blocks))
	for idx, block := range blocks {
		patterns[idx] = strings.Split(block, "\n")
	}

	var total int
	for idx, pattern := range patterns {
		if r := check(pattern, false); r != -1 {
			total += r * 100
			continue
		}
		if c := check(rotate(pattern), false); c != -1 {
			total += c
			continue
		}
		fmt.Printf("couldn't find reflection line in pattern %d\n", idx)
	}

	return strconv.Itoa(total), nil
}

func Part2(input string) (string, error) {
	blocks := strings.Split(input, "\n\n")
	patterns := make([][]string, len(blocks))
	for idx, block := range blocks {
		patterns[idx] = strings.Split(block, "\n")
	}

	var total int
	for idx, pattern := range patterns {
		if r := check(pattern, true); r != -1 {
			total += r * 100
			continue
		}
		if c := check(rotate(pattern), true); c != -1 {
			total += c
			continue
		}
		fmt.Printf("couldn't find reflection line in pattern %d\n", idx)
	}

	return strconv.Itoa(total), nil
}
