package day9

import (
	"slices"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
	"github.com/MarkusFreitag/advent-of-code/util/sliceutil"
)

func predict(seq []int) int {
	seqs := [][]int{seq}

	for !sliceutil.All(seqs[len(seqs)-1], 0) {
		next := make([]int, 0)
		last := len(seqs) - 1
		for i := 0; i < len(seqs[last]); i++ {
			if i+1 == len(seqs[last]) {
				break
			}
			next = append(next, seqs[last][i+1]-seqs[last][i])
		}
		seqs = append(seqs, next)
	}

	seqs[len(seqs)-1] = append(seqs[len(seqs)-1], 0)

	for i := len(seqs) - 1; i > 0; i-- {
		seqs[i-1] = append(seqs[i-1], seqs[i][len(seqs[i])-1]+seqs[i-1][len(seqs[i-1])-1])
	}

	return seqs[0][len(seqs[0])-1]
}

func Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")
	nums := make([]int, len(lines))
	for idx, line := range lines {
		nums[idx] = predict(util.StringsToInts(strings.Fields(line)))
	}
	return strconv.Itoa(numbers.Sum(nums...)), nil
}

func Part2(input string) (string, error) {
	lines := strings.Split(input, "\n")
	nums := make([]int, len(lines))
	for idx, line := range lines {
		seq := util.StringsToInts(strings.Fields(line))
		slices.Reverse(seq)
		nums[idx] = predict(seq)
	}
	return strconv.Itoa(numbers.Sum(nums...)), nil
}
