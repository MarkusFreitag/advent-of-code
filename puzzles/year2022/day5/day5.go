package day5

import (
	"slices"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
	"github.com/MarkusFreitag/advent-of-code/util/sliceutil"
)

func parseInput(input string) ([][]string, []string) {
	blocks := strings.Split(input, "\n\n")
	stackLines := strings.Split(blocks[0], "\n")
	instrLines := strings.Split(strings.TrimSpace(blocks[1]), "\n")

	slices.Reverse(stackLines)

	labels := sliceutil.Chunks([]rune(stackLines[0]), 4)
	stacks := make([][]string, len(labels))
	for idx := range stacks {
		stacks[idx] = make([]string, 0)
	}

	for _, line := range stackLines[1:] {
		groups := sliceutil.Chunks([]rune(line), 4)
		for idx, group := range groups {
			if numbers.Between(group[1], 65, 90) {
				stacks[idx] = append(stacks[idx], string(rune(group[1])))
			}
		}
	}

	return stacks, instrLines
}

func parseInstruction(instr string) (int, int, int) {
	fields := strings.Fields(instr)
	return util.ParseInt(fields[1]),
		util.ParseInt(fields[3]) - 1,
		util.ParseInt(fields[5]) - 1
}

func Part1(input string) (string, error) {
	stacks, instructions := parseInput(input)
	for _, instr := range instructions {
		count, from, to := parseInstruction(instr)
		fromPile := stacks[from]
		toPile := stacks[to]
		crates, fromPile := sliceutil.PopN(fromPile, count)
		slices.Reverse(crates)
		toPile = append(toPile, crates...)
		stacks[from] = fromPile
		stacks[to] = toPile
	}
	var top string
	for _, pile := range stacks {
		top += sliceutil.Tail(pile, 1)[0]
	}
	return top, nil
}

func Part2(input string) (string, error) {
	stacks, instructions := parseInput(input)
	for _, instr := range instructions {
		count, from, to := parseInstruction(instr)
		fromPile := stacks[from]
		toPile := stacks[to]
		crates, fromPile := sliceutil.PopN(fromPile, count)
		toPile = append(toPile, crates...)
		stacks[from] = fromPile
		stacks[to] = toPile
	}
	var top string
	for _, pile := range stacks {
		top += sliceutil.Tail(pile, 1)[0]
	}
	return top, nil
}
