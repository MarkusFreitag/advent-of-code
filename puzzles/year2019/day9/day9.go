package day9

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util/slice"
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
		if slice.Contains([]int{1, 2, 5, 6, 7, 8}, opcode) {
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

func Part1(input string) (string, error) {
	items := strings.Split(input, ",")
	icode, err := newIntcode(items)
	if err != nil {
		return "", err
	}
	inp := make(chan int, 10)
	inp <- 1
	out := make(chan int, 100)

	icode.Interpret(inp, out)
	result := make([]int, 0)
	for i := range out {
		fmt.Printf("output: %d\n", i)
		result = append(result, i)
	}
	return fmt.Sprintf("%v", result), nil
}

func Part2(input string) (string, error) {
	items := strings.Split(input, ",")
	icode, err := newIntcode(items)
	if err != nil {
		return "", err
	}
	inp := make(chan int, 10)
	inp <- 2
	out := make(chan int, 100)

	icode.Interpret(inp, out)
	result := make([]int, 0)
	for i := range out {
		fmt.Printf("output: %d\n", i)
		result = append(result, i)
	}
	return fmt.Sprintf("%v", result), nil
}
