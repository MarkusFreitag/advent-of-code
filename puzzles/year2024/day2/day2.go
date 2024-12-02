package day2

import (
	"slices"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

func Part1(input string) (string, error) {
	seqs := make([][]int, 0)
	for _, line := range strings.Split(input, "\n") {
		seqs = append(seqs, util.StringsToInts(strings.Fields(line)))
	}
	var total int
	for _, seq := range seqs {
		if checkSeq(seq) {
			total++
		}
	}
	return strconv.Itoa(total), nil
}

func Part2(input string) (string, error) {
	seqs := make([][]int, 0)
	for _, line := range strings.Split(input, "\n") {
		seqs = append(seqs, util.StringsToInts(strings.Fields(line)))
	}
	var total int
	for _, seq := range seqs {
		for index := range seq {
			dup := slices.Clone(seq)
			if checkSeq(append(dup[:index], dup[index+1:]...)) {
				total++
				break
			}
		}
	}
	return strconv.Itoa(total), nil
}

func checkSeq(seq []int) bool {
	sign := numbers.Sign(seq[1] - seq[0])
	for i := 1; i < len(seq); i++ {
		diff := seq[i] - seq[i-1]
		if numbers.Sign(diff) != sign || !numbers.Between(numbers.Abs(diff), 1, 3) {
			return false
		}
	}
	return true
}
