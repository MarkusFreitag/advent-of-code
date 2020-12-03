package day3

import (
	"strconv"
	"strings"
)

func countTrees(rows []string, right, down int) int {
	var row, col, trees int
	for {
		col += right
		row += down
		if col >= len(rows[row]) {
			col = col % len(rows[row])
		}
		if string(rows[row][col]) == "#" {
			trees++
		}
		if row == len(rows)-1 {
			break
		}
	}
	return trees
}

func Part1(input string) (string, error) {
	rows := strings.Split(input, "\n")
	return strconv.Itoa(countTrees(rows, 3, 1)), nil
}

func Part2(input string) (string, error) {
	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	rows := strings.Split(input, "\n")
	var totalTrees int
	for _, slope := range slopes {
		trees := countTrees(rows, slope[0], slope[1])
		if totalTrees == 0 {
			totalTrees = trees
		} else {
			totalTrees *= trees
		}
	}
	return strconv.Itoa(totalTrees), nil
}
