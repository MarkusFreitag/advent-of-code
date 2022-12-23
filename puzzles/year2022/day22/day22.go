package day22

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

var (
	turnLeft = map[string]string{
		"R": "U",
		"U": "L",
		"L": "D",
		"D": "R",
	}
	turnRight = map[string]string{
		"R": "D",
		"D": "L",
		"L": "U",
		"U": "R",
	}
	directions = map[string]point{
		"U": point{X: 0, Y: -1},
		"R": point{X: 1, Y: 0},
		"D": point{X: 0, Y: 1},
		"L": point{X: -1, Y: 0},
	}
)

type point struct {
	X, Y int
}

func (p point) String() string {
	return fmt.Sprintf("y=%d|x=%d", p.Y, p.X)
}

func (p point) Add(pt point) point {
	return point{
		X: p.X + pt.X,
		Y: p.Y + pt.Y,
	}
}

func nextSteps(instructions string) (int, string) {
	if i, err := strconv.Atoi(instructions); err == nil {
		return i, ""
	}
	var num string
	for instructions[0] != 'R' && instructions[0] != 'L' {
		num += instructions[:1]
		instructions = instructions[1:]
	}
	return util.ParseInt(num), instructions
}

func nextTurn(instructions string) (string, string) {
	return instructions[:1], instructions[1:]
}

func parseInput(input string) (map[point]bool, map[point]point, string, point) {
	blocks := strings.Split(input, "\n\n")
	board := make(map[point]bool)
	jumps := make(map[point]point)
	var start point
	var max point
	for rowIdx, row := range strings.Split(blocks[0], "\n") {
		firstPt := point{Y: rowIdx}
		lastPt := point{Y: rowIdx}
		first := true
		for colIdx, char := range row {
			if char == ' ' {
				continue
			}
			if first {
				firstPt.X = colIdx
				first = false
				if rowIdx == 0 {
					start = firstPt
				}
			}
			board[point{X: colIdx, Y: rowIdx}] = char == '.'
			lastPt.X = colIdx
			if colIdx > max.X {
				max.X = colIdx
			}
		}
		jumps[firstPt] = lastPt
		jumps[lastPt] = firstPt
		if rowIdx > max.Y {
			max.Y = rowIdx
		}
	}

	for x := 0; x < max.X; x++ {
		firstPt := point{X: x}
		lastPt := point{X: x}
		first := true
		for y := 0; y < max.Y; y++ {
			if _, ok := board[point{X: x, Y: y}]; !ok {
				if !first {
					lastPt.Y = y - 1
				}
				continue
			}
			if first {
				firstPt.Y = y
				first = false
			}
		}
		jumps[firstPt] = lastPt
		jumps[lastPt] = firstPt
	}

	return board, jumps, strings.TrimSpace(blocks[1]), start
}

func Part1(input string) (string, error) {
	board, jumps, instructions, pos := parseInput(input)

	dir := "R"

	for {
		var steps int
		steps, instructions = nextSteps(instructions)
		for s := 0; s < steps; s++ {
			newPos := pos.Add(directions[dir])

			boardState, boardOk := board[newPos]
			if boardOk {
				if boardState {
					pos = newPos
					continue
				} else {
					break
				}
			}

			if !boardOk {
				jumpTo, jumpOk := jumps[pos]
				if jumpOk {
					if board[jumpTo] {
						pos = jumpTo
					} else {
						break
					}
				}
			}

			if v, ok := board[newPos]; ok && !v {
				break
			} else {
				if j, ok := jumps[pos]; ok {
					if v, ok := board[j]; ok && !v {
						break
					} else {
						newPos = j
					}
				}
			}
			pos = newPos
		}

		if instructions == "" {
			break
		}

		var turn string
		turn, instructions = nextTurn(instructions)
		switch turn {
		case "R":
			dir = turnRight[dir]
		case "L":
			dir = turnLeft[dir]
		}

		if instructions == "" {
			break
		}
	}

	pos = pos.Add(point{1, 1})
	var dirVal int
	switch dir {
	case "R":
	case "D":
		dirVal = 1
	case "L":
		dirVal = 2
	case "U":
		dirVal = 3
	}

	return strconv.Itoa(1000*pos.Y + 8*pos.X + dirVal), nil
}

func Part2(input string) (string, error) {
	return "not solved yet", nil
}
