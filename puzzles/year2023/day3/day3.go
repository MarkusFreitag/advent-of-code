package day3

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

type number struct {
	row         int
	cols        []int
	val         string
	specialChar string
	specialPos  string
}

func (n *number) Int() int { return util.ParseInt(n.val) }

func (n *number) Check(grid [][]string) bool {
	for y := n.row - 1; y <= n.row+1; y++ {
		if y < 0 || y >= len(grid) {
			continue
		}
		for x := n.cols[0] - 1; x <= n.cols[len(n.cols)-1]+1; x++ {
			if x < 0 || x >= len(grid[y]) {
				continue
			}

			if !special.MatchString(grid[y][x]) {
				n.specialChar = grid[y][x]
				n.specialPos = fmt.Sprintf("%d|%d", y, x)
				return true
			}
		}
	}
	return false
}

var special = regexp.MustCompile(`^[\.0-9]$`)

func getValidNumbers(input string) []*number {
	grid := make([][]string, 0)
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, util.StringToStrings(line))
	}

	valid := make([]*number, 0)
	for y, row := range grid {
		var current *number
		for x, col := range row {
			if _, err := strconv.Atoi(col); err == nil {
				if current == nil {
					current = &number{row: y, cols: []int{x}, val: col}
				} else {
					current.cols = append(current.cols, x)
					current.val += col
				}
			} else {
				if current != nil {
					if current.Check(grid) {
						valid = append(valid, current)
					}
					current = nil
				}
			}
		}

		if current != nil {
			if current.Check(grid) {
				valid = append(valid, current)
			}
			current = nil
		}
	}
	return valid
}

func Part1(input string) (string, error) {
	var sum int
	for _, num := range getValidNumbers(input) {
		sum += num.Int()
	}

	return strconv.Itoa(sum), nil
}

func Part2(input string) (string, error) {
	validNumbers := getValidNumbers(input)

	var sum int
	seen := make(map[string]bool)
	for index, num := range validNumbers {
		for idx, n := range validNumbers {
			if index == idx {
				continue
			}
			_, found := seen[num.specialPos]
			if num.specialChar == "*" && n.specialChar == "*" && num.specialPos == n.specialPos && !found {
				sum += num.Int() * n.Int()
				seen[num.specialPos] = true
			}
		}
	}

	return strconv.Itoa(sum), nil
}
