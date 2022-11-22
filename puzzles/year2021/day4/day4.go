package day4

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/slice"
)

type Num struct {
	Value   int
	Checked bool
}

func parseNum(str string) *Num {
	return &Num{
		Value: util.ParseInt(str),
	}
}

type Board struct {
	b [][]*Num
}

func parseBoard(str string) *Board {
	board := Board{
		b: make([][]*Num, 0),
	}
	lines := strings.Split(str, "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		nums := make([]*Num, len(fields))
		for idx, field := range fields {
			nums[idx] = parseNum(field)
		}
		board.b = append(board.b, nums)
	}
	return &board
}

func (b *Board) AddNum(i int) {
	for _, row := range b.b {
		for _, num := range row {
			if num.Value == i {
				num.Checked = true
			}
		}
	}
}

func (b *Board) Bingo() bool {
	// check rows
	for _, row := range b.b {
		bingo := make([]bool, 0)
		for _, num := range row {
			bingo = append(bingo, num.Checked)
		}
		if slice.All(bingo, true) {
			return true
		}
	}
	// check cols
	for c := 0; c < len(b.b); c++ {
		bingo := make([]bool, 0)
		for r := 0; r < len(b.b); r++ {
			bingo = append(bingo, b.b[r][c].Checked)
		}
		if slice.All(bingo, true) {
			return true
		}
	}

	return false
}

func (b *Board) Score() int {
	var score int
	for _, row := range b.b {
		for _, num := range row {
			if !num.Checked {
				score += num.Value
			}
		}
	}
	return score
}

func parseInput(input string) ([]int, []*Board) {
	parts := strings.Split(input, "\n\n")
	nums := util.StringsToInts(strings.Split(strings.TrimSpace(parts[0]), ","))

	boards := make([]*Board, len(parts[1:]))
	for idx, block := range parts[1:] {
		boards[idx] = parseBoard(block)
	}

	return nums, boards
}

func Part1(input string) (string, error) {
	nums, boards := parseInput(input)

	for _, num := range nums {
		for _, board := range boards {
			board.AddNum(num)
			if board.Bingo() {
				unmarkedScore := board.Score()
				return strconv.Itoa(unmarkedScore * num), nil
			}
		}
	}

	return "not solved yet", nil
}

func Part2(input string) (string, error) {
	nums, boards := parseInput(input)

	winners := make(map[int]bool)
	for _, num := range nums {
		for boardIdx, board := range boards {
			board.AddNum(num)
			if board.Bingo() {
				unmarkedScore := board.Score()
				finalScore := unmarkedScore * num
				if _, ok := winners[boardIdx]; ok {
					continue
				}
				winners[boardIdx] = true
				if len(winners) == len(boards) {
					return strconv.Itoa(finalScore), nil
				}
			}
		}
	}

	return "not solved yet", nil
}
