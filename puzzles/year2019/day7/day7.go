package day7

import (
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

func (icode intcode) Interpret(phase int, input, output chan int) {
	firstInput := true
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
			var val int
			if firstInput {
				val = phase
				firstInput = false
			} else {
				val = <-input
			}
			icode[icode[counter+1]] = val
			counter += 2
		case 4:
			var val int
			if firstMode == 0 {
				val = icode[icode[counter+1]]
			} else {
				val = icode[counter+1]
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
}

func newIntcode(input []string) (intcode, error) {
	icode := make(intcode, len(input))
	for idx, item := range input {
		num, err := strconv.Atoi(item)
		if err != nil {
			return nil, err
		}
		icode[idx] = num
	}
	return icode, nil
}

func ampController1(icodeStr string, seq []int) int {
	items := strings.Split(icodeStr, ",")

	ampA, _ := newIntcode(items)
	ampB, _ := newIntcode(items)
	ampC, _ := newIntcode(items)
	ampD, _ := newIntcode(items)
	ampE, _ := newIntcode(items)

	out := make(chan int, 10)
	out <- 0
	ampA.Interpret(seq[0], out, out)
	ampB.Interpret(seq[1], out, out)
	ampC.Interpret(seq[2], out, out)
	ampD.Interpret(seq[3], out, out)
	ampE.Interpret(seq[4], out, out)

	return <-out
}

func ampController2(icodeStr string, seq []int) int {
	items := strings.Split(icodeStr, ",")

	ampA, _ := newIntcode(items)
	ampB, _ := newIntcode(items)
	ampC, _ := newIntcode(items)
	ampD, _ := newIntcode(items)
	ampE, _ := newIntcode(items)

	atob := make(chan int, 10)
	btoc := make(chan int, 10)
	ctod := make(chan int, 10)
	dtoe := make(chan int, 10)
	etoa := make(chan int, 10)

	etoa <- 0
	go ampA.Interpret(seq[0], etoa, atob)
	go ampB.Interpret(seq[1], atob, btoc)
	go ampC.Interpret(seq[2], btoc, ctod)
	go ampD.Interpret(seq[3], ctod, dtoe)
	ampE.Interpret(seq[4], dtoe, etoa)

	return <-etoa
}

func Perm(a []int, f func([]int)) {
	perm(a, f, 0)
}

func perm(a []int, f func([]int), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func Part1(input string) (string, error) {
	var highest int
	Perm([]int{0, 1, 2, 3, 4}, func(a []int) {
		val := ampController1(input, a)
		if val > highest {
			highest = val
		}
	})
	return strconv.Itoa(highest), nil
}

func Part2(input string) (string, error) {
	var highest int
	Perm([]int{5, 6, 7, 8, 9}, func(a []int) {
		val := ampController2(input, a)
		if val > highest {
			highest = val
		}
	})
	return strconv.Itoa(highest), nil
}
