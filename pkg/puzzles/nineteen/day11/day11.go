package day11

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/pkg/util"
)

func numFromStr(str string, pos int) int {
	n, _ := strconv.Atoi(string(str[pos]))
	return n
}

func parseOpCode(code int) (int, int, int, int) {
	codeStr := strconv.Itoa(code)
	var opcode, firstMode, secondMode, targetMode int
	opcode = code % 100
	switch len(codeStr) {
	case 3:
		firstMode = numFromStr(codeStr, 0)
	case 4:
		firstMode = numFromStr(codeStr, 1)
		secondMode = numFromStr(codeStr, 0)
	case 5:
		firstMode = numFromStr(codeStr, 2)
		secondMode = numFromStr(codeStr, 1)
		targetMode = numFromStr(codeStr, 0)
	}
	return opcode, firstMode, secondMode, targetMode
}

type intcode []int

func (icode intcode) Interpret(input, output chan int) {
	var relBase int
	var counter int
	for counter < len(icode) {
		opcode, firstMode, secondMode, targetMode := parseOpCode(icode[counter])
		var firstValue, secondValue int
		if util.IntInSlice(opcode, []int{1, 2, 5, 6, 7, 8}) {
			switch firstMode {
			case 0:
				firstValue = icode[icode[counter+1]]
			case 1:
				firstValue = icode[counter+1]
			case 2:
				firstValue = icode[relBase+icode[counter+1]]
			}
			switch secondMode {
			case 0:
				secondValue = icode[icode[counter+2]]
			case 1:
				secondValue = icode[counter+2]
			case 2:
				secondValue = icode[relBase+icode[counter+2]]
			}
		}
		switch opcode {
		case 1:
			var idx int
			switch targetMode {
			case 0:
				idx = icode[counter+3]
			case 2:
				idx = relBase + icode[counter+3]
			}
			icode[idx] = firstValue + secondValue
			counter += 4
		case 2:
			var idx int
			switch targetMode {
			case 0:
				idx = icode[counter+3]
			case 2:
				idx = relBase + icode[counter+3]
			}
			icode[idx] = firstValue * secondValue
			counter += 4
		case 3:
			var idx int
			switch firstMode {
			case 0:
				idx = icode[counter+1]
			case 2:
				idx = relBase + icode[counter+1]
			}
			icode[idx] = <-input
			counter += 2
		case 4:
			var val int
			switch firstMode {
			case 0:
				val = icode[icode[counter+1]]
			case 1:
				val = icode[counter+1]
			case 2:
				val = icode[relBase+icode[counter+1]]
			}
			output <- val
			counter += 2
		case 5:
			if firstValue != 0 {
				counter = secondValue
			} else {
				counter += 3
			}
		case 6:
			if firstValue == 0 {
				counter = secondValue
			} else {
				counter += 3
			}
		case 7:
			var val int
			if firstValue < secondValue {
				val = 1
			}
			var idx int
			switch targetMode {
			case 0:
				idx = icode[counter+3]
			case 2:
				idx = relBase + icode[counter+3]
			}
			icode[idx] = val
			counter += 4
		case 8:
			var val int
			if firstValue == secondValue {
				val = 1
			}
			var idx int
			switch targetMode {
			case 0:
				idx = icode[counter+3]
			case 2:
				idx = relBase + icode[counter+3]
			}
			icode[idx] = val
			counter += 4
		case 9:
			var val int
			switch firstMode {
			case 0:
				val = icode[icode[counter+1]]
			case 1:
				val = icode[counter+1]
			case 2:
				val = icode[relBase+icode[counter+1]]
			}
			relBase += val
			counter += 2
		case 99:
			counter = len(icode)
			close(output)
		}
	}
}

func newIntcode(input []string) (intcode, error) {
	icode := make(intcode, 10000)
	for idx, item := range input {
		num, err := strconv.Atoi(item)
		if err != nil {
			return nil, err
		}
		icode[idx] = num
	}
	return icode, nil
}

type Pos struct {
	X, Y int
	Dir  string
}

func (p *Pos) Move(i int) {
	switch p.Dir {
	case "^":
		if i == 0 {
			p.Dir = "<"
			p.X -= 1
		} else {
			p.Dir = ">"
			p.X += 1
		}
	case "v":
		if i == 0 {
			p.Dir = ">"
			p.X += 1
		} else {
			p.Dir = "<"
			p.X -= 1
		}
	case ">":
		if i == 0 {
			p.Dir = "^"
			p.Y += 1
		} else {
			p.Dir = "v"
			p.Y -= 1
		}
	case "<":
		if i == 0 {
			p.Dir = "v"
			p.Y -= 1
		} else {
			p.Dir = "^"
			p.Y += 1
		}
	}
}

func (p *Pos) String() string {
	return fmt.Sprintf("%d|%d", p.X, p.Y)
}

type Part1 struct{}

func (p *Part1) Solve(input string) (string, error) {
	items := strings.Split(input, ",")
	icode, err := newIntcode(items)
	if err != nil {
		return "", err
	}
	inp := make(chan int, 10)
	out := make(chan int, 100)

	go icode.Interpret(inp, out)

	grid := make(map[string]int)
	currentPos := &Pos{X: 0, Y: 0, Dir: "^"}
	for {
		pos := currentPos.String()
		if v, ok := grid[pos]; ok {
			inp <- v
		} else {
			inp <- 0
		}
		grid[pos] = <-out
		move, ok := <-out
		if !ok {
			break
		}
		currentPos.Move(move)
	}
	return strconv.Itoa(len(grid)), nil
}

type Part2 struct{}

func (p *Part2) Solve(input string) (string, error) {
	items := strings.Split(input, ",")
	icode, err := newIntcode(items)
	if err != nil {
		return "", err
	}
	inp := make(chan int, 10)
	out := make(chan int, 100)

	go icode.Interpret(inp, out)

	grid := make(map[string]int)
	currentPos := &Pos{X: 0, Y: 0, Dir: "^"}
	grid[currentPos.String()] = 1
	for {
		pos := currentPos.String()
		if v, ok := grid[pos]; ok {
			inp <- v
		} else {
			inp <- 0
		}
		grid[pos] = <-out
		move, ok := <-out
		if !ok {
			break
		}
		currentPos.Move(move)
	}

	rows := make([]string, 0)
	for y := -10; y < 10; y++ {
		row := make([]string, 0)
		for x := -50; x < 50; x++ {
			val := "."
			if v, ok := grid[fmt.Sprintf("%d|%d", x, y)]; ok {
				if v == 1 {
					val = "#"
				}
			}
			row = append(row, val)
		}
		rows = append(rows, strings.Join(row, ""))
	}
	return "\n" + strings.Join(rows, "\n"), nil
}
