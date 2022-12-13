package day13

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

var (
	RIGHT = true
	WRONG = false
)

type Pair struct {
	Left  []any
	Right []any
}

func compare(left, right any) *bool {
	if lVal, ok := left.(float64); ok {
		if rVal, ok := right.(float64); ok {
			if lVal < rVal {
				return &RIGHT
			} else if lVal > rVal {
				return &WRONG
			}
			return nil
		}
		return compare([]any{left}, right)
	}
	if _, ok := right.(float64); ok {
		return compare(left, []any{right})
	}
	lVal := left.([]any)
	rVal := right.([]any)
	for i := 0; i < numbers.Min(len(lVal), len(rVal)); i++ {
		r := compare(lVal[i], rVal[i])
		if r != nil {
			return r
		}
	}

	if len(lVal) < len(rVal) {
		return &RIGHT
	}
	if len(lVal) > len(rVal) {
		return &WRONG
	}
	return nil
}

func Part1(input string) (string, error) {
	blocks := strings.Split(input, "\n\n")
	pairs := make([]*Pair, len(blocks))
	for idx, block := range blocks {
		lines := strings.Split(block, "\n")
		var left, right []any
		json.Unmarshal([]byte(lines[0]), &left)
		json.Unmarshal([]byte(lines[1]), &right)
		pairs[idx] = &Pair{
			Left:  left,
			Right: right,
		}
	}

	rightOrder := make([]int, 0)

	for pIdx, pair := range pairs {
		if compare(pair.Left, pair.Right) == &RIGHT {
			rightOrder = append(rightOrder, pIdx+1)
		}
	}

	return strconv.Itoa(numbers.Sum(rightOrder...)), nil
}

func Part2(input string) (string, error) {
	lines := strings.Split(input, "\n")
	packets := make([]any, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		var p any
		json.Unmarshal([]byte(line), &p)
		packets = append(packets, p)
	}
	startDivider := []any{2.0}
	endDivider := []any{6.0}
	packets = append(packets, []any{startDivider})
	packets = append(packets, []any{endDivider})

	sort.SliceStable(packets, func(i, j int) bool {
		return compare(packets[i], packets[j]) == &RIGHT
	})

	var start, end int
	for idx, packet := range packets {
		if fmt.Sprintf("%v", packet) == "[[2]]" {
			start = idx + 1
		}
		if fmt.Sprintf("%v", packet) == "[[6]]" {
			end = idx + 1
		}
	}

	return strconv.Itoa(start * end), nil
}
