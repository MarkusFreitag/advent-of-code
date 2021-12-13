package day13

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

func parseInput(input string) ([][]bool, []string) {
	blocks := strings.Split(input, "\n\n")

	lines := strings.Fields(blocks[0])
	xs, ys := make([]int, len(lines)), make([]int, len(lines))
	var maxX, maxY int
	for idx, line := range lines {
		parts := strings.Split(line, ",")
		num := util.ParseInt(parts[0])
		if num > maxX {
			maxX = num
		}
		xs[idx] = num
		num = util.ParseInt(parts[1])
		if num > maxY {
			maxY = num
		}
		ys[idx] = num
	}

	paper := make([][]bool, maxY+1)
	for i := 0; i <= maxY; i++ {
		paper[i] = make([]bool, maxX+1)
	}

	for idx, y := range ys {
		paper[y][xs[idx]] = true
	}

	return paper, strings.Split(blocks[1], "\n")
}

func show(paper [][]bool) string {
	lines := make([]string, len(paper))
	for idx, row := range paper {
		var str string
		for _, col := range row {
			if col {
				str += "#"
			} else {
				str += " "
			}
		}
		lines[idx] = str
	}
	return strings.Join(lines, "\n")
}

func fold(paper [][]bool, instr string) [][]bool {
	fields := strings.Fields(instr)
	parts := strings.Split(fields[2], "=")
	num := util.ParseInt(parts[1])
	if strings.HasPrefix(fields[2], "y") {
		newPaper := paper[:num]
		for y, row := range paper[num+1:] {
			for x, col := range row {
				if col {
					newPaper[len(newPaper)-1-y][x] = true
				}
			}
		}
		return newPaper
	}
	newPaper := make([][]bool, len(paper))
	for y, row := range paper {
		newPaper[y] = row[:num]
		for x, col := range row[num+1:] {
			if col {
				newPaper[y][len(newPaper[y])-1-x] = true
			}
		}
	}
	return newPaper
}

func Part1(input string) (string, error) {
	paper, instructions := parseInput(input)

	paper = fold(paper, instructions[0])

	var dots int
	for _, row := range paper {
		for _, col := range row {
			if col {
				dots++
			}
		}
	}

	return strconv.Itoa(dots), nil
}

func Part2(input string) (string, error) {
	paper, instructions := parseInput(input)

	for _, instr := range instructions {
		paper = fold(paper, instr)
	}

	return show(paper), nil
}
