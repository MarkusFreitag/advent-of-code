package day5

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

func (icode intcode) Interpret(input int) []int {
	out := make([]int, 0)
	var counter int
	for counter < len(icode) {
		opcode, firstMode, secondMode, _ := parseOpCode(icode[counter])
		var firstValue, secondValue int
		if slice.Contains([]int{1, 2, 5, 6, 7, 8}, opcode) {
			if firstMode == 0 {
				firstValue = icode[icode[counter+1]]
			} else {
				firstValue = icode[counter+1]
			}
			if secondMode == 0 {
				secondValue = icode[icode[counter+2]]
			} else {
				secondValue = icode[counter+2]
			}
		}
		switch opcode {
		case 1:
			icode[icode[counter+3]] = firstValue + secondValue
			counter += 4
		case 2:
			icode[icode[counter+3]] = firstValue * secondValue
			counter += 4
		case 3:
			icode[icode[counter+1]] = input
			counter += 2
		case 4:
			if firstMode == 0 {
				out = append(out, icode[icode[counter+1]])
			} else {
				out = append(out, icode[counter+1])
			}
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
			icode[icode[counter+3]] = val
			counter += 4
		case 8:
			var val int
			if firstValue == secondValue {
				val = 1
			}
			icode[icode[counter+3]] = val
			counter += 4
		case 99:
			counter = len(icode)
		}
	}
	return out
}

func Part1(input string) (string, error) {
	items := strings.Split(input, ",")
	icode := make(intcode, len(items))
	for idx, item := range items {
		num, err := strconv.Atoi(item)
		if err != nil {
			return "", err
		}
		icode[idx] = num
	}
	out := icode.Interpret(1)
	var failed bool
	for idx, item := range out {
		if item != 0 && idx != len(out)-1 {
			failed = true
			break
		}
	}
	if failed {
		return "", fmt.Errorf("running intcode not successfully: %v", out)
	}
	return strconv.Itoa(out[len(out)-1]), nil
}

func Part2(input string) (string, error) {
	items := strings.Split(input, ",")
	icode := make(intcode, len(items))
	for idx, item := range items {
		num, err := strconv.Atoi(item)
		if err != nil {
			return "", err
		}
		icode[idx] = num
	}
	out := icode.Interpret(5)
	var failed bool
	for idx, item := range out {
		if item != 0 && idx != len(out)-1 {
			failed = true
			break
		}
	}
	if failed {
		return "", fmt.Errorf("running intcode not successfully: %v", out)
	}
	return strconv.Itoa(out[len(out)-1]), nil
}
