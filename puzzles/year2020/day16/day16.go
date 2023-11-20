package day16

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
	"github.com/MarkusFreitag/advent-of-code/util/sliceutil"
)

var rgxCond = regexp.MustCompile(`^(.*):\s(\d+)-(\d+)\sor\s(\d+)-(\d+)$`)

type Condition struct {
	Name string
	Min1 int
	Max1 int
	Min2 int
	Max2 int
}

func (c Condition) CheckValue(v int) bool {
	return numbers.Between(v, c.Min1, c.Max1) || numbers.Between(v, c.Min2, c.Max2)
}

func parseInput(input string) (map[string]Condition, []int, [][]int) {
	blocks := strings.Split(input, "\n\n")

	conds := make(map[string]Condition)
	for _, line := range strings.Split(blocks[0], "\n") {
		matches := rgxCond.FindAllStringSubmatch(line, -1)[0]
		min1, _ := strconv.Atoi(matches[2])
		max1, _ := strconv.Atoi(matches[3])
		min2, _ := strconv.Atoi(matches[4])
		max2, _ := strconv.Atoi(matches[5])
		cond := Condition{
			Name: matches[1],
			Min1: min1,
			Max1: max1,
			Min2: min2,
			Max2: max2,
		}
		conds[cond.Name] = cond
	}

	lines := strings.Split(blocks[1], "\n")
	myTicket := util.StringsToInts(strings.Split(lines[1], ","))

	lines = strings.Split(blocks[2], "\n")
	otherTickets := make([][]int, 0)
	for _, line := range lines[1:] {
		otherTickets = append(
			otherTickets,
			util.StringsToInts(strings.Split(line, ",")),
		)
	}

	return conds, myTicket, otherTickets
}

func Part1(input string) (string, error) {
	conds, _, otherTickets := parseInput(input)

	invalidValues := make([]int, 0)
	for _, t := range otherTickets {
		for _, v := range t {
			var valid bool
			for _, c := range conds {
				if numbers.Between(v, c.Min1, c.Max1) || numbers.Between(v, c.Min2, c.Max2) {
					valid = true
				}
			}
			if !valid {
				invalidValues = append(invalidValues, v)
			}
		}
	}

	return strconv.Itoa(numbers.Sum(invalidValues...)), nil
}

func Part2(input string) (string, error) {
	conds, myTicket, otherTickets := parseInput(input)

	// get valid tickets
	validTickets := make([][]int, 0)
	for _, ticket := range otherTickets {
		bools := make([]bool, len(ticket))
		for idx, field := range ticket {
			for _, cond := range conds {
				if cond.CheckValue(field) {
					bools[idx] = true
				}
			}
		}
		if sliceutil.All(bools, true) {
			validTickets = append(validTickets, ticket)
		}
	}

	// get possible positions per condition
	poss := make(map[string][]int)
	for _, cond := range conds {
		for i := 0; i < len(myTicket); i++ {
			valid := true
			for _, ticket := range validTickets {
				if !cond.CheckValue(ticket[i]) {
					valid = false
				}
			}
			if valid {
				posis, ok := poss[cond.Name]
				if !ok {
					posis = make([]int, 0)
				}
				poss[cond.Name] = append(posis, i)
			}
		}
	}

	// sort conditions into order
	condOrder := make([]string, len(myTicket))
	for len(poss) > 0 {
		for cond, positions := range poss {
			if len(positions) == 1 {
				condOrder[positions[0]] = cond
				delPos(poss, positions[0])
			}
		}
	}

	// get fields for conditions having the 'departure' prefix
	values := make([]int, 0)
	for idx, name := range condOrder {
		if strings.HasPrefix(name, "departure") {
			values = append(values, myTicket[idx])
		}
	}

	// multiply those fields
	return strconv.Itoa(numbers.Multiply(values...)), nil
}

func delPos(hash map[string][]int, pos int) map[string][]int {
	for name, positions := range hash {
		if len(positions) == 0 {
			continue
		} else if len(positions) == 1 && positions[0] == pos {
			delete(hash, name)
		} else {
			index := -1
			for idx, item := range positions {
				if item == pos {
					index = idx
				}
			}
			if index == -1 {
				continue
			}
			copy(positions[index:], positions[index+1:])
			positions[len(positions)-1] = -1
			positions = positions[:len(positions)-1]
			hash[name] = positions
		}
	}
	return hash
}
