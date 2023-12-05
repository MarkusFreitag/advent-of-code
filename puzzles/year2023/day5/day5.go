package day5

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
	"github.com/MarkusFreitag/advent-of-code/util/sliceutil"
)

type mappingFn func(int) int

func newMappingFn(ranges [][]int) mappingFn {
	return func(num int) int {
		for _, r := range ranges {
			if num >= r[1] && num < r[1]+r[2] {
				return r[0] + num - r[1]
			}
		}
		return num
	}
}

func parseBlocks(blocks []string) mappingFn {
	fns := make([]mappingFn, len(blocks))
	for idx, block := range blocks {
		lines := strings.Split(block, "\n")
		ranges := make([][]int, 0)
		for _, line := range lines[1:] {
			ranges = append(ranges, util.StringsToInts(strings.Fields(line)))
		}
		fns[idx] = newMappingFn(ranges)
	}

	return func(n int) int {
		for _, fn := range fns {
			n = fn(n)
		}
		return n
	}
}

func Part1(input string) (string, error) {
	blocks := strings.Split(input, "\n\n")
	seeds := util.StringsToInts(strings.Fields(strings.Split(blocks[0], ":")[1]))
	mapper := parseBlocks(blocks[1:])

	min := numbers.MaxInteger
	for _, seed := range seeds {
		min = numbers.Min(min, mapper(seed))
	}

	return strconv.Itoa(min), nil
}

func Part2(input string) (string, error) {
	blocks := strings.Split(input, "\n\n")
	seeds := util.StringsToInts(strings.Fields(strings.Split(blocks[0], ":")[1]))
	mapper := parseBlocks(blocks[1:])

	min := numbers.MaxInteger
	for _, seed := range sliceutil.Chunks(seeds, 2) {
		for s := seed[0]; s < seed[0]+seed[1]; s++ {
			min = numbers.Min(min, mapper(s))
		}
	}

	return strconv.Itoa(min), nil
}
