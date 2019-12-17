package intcode

import (
	"math"
	"strconv"
	"strings"
)

const (
	MessageWantsInput = iota
	MessageOutput
	MessageHalt
)

type Message struct {
	Type  int
	Value int64
}

type IntCode []int64

func New(code string) (IntCode, error) {
	items := strings.Split(code, ",")
	icode := make(IntCode, len(items))
	for idx, item := range items {
		num, err := strconv.ParseInt(item, 10, 64)
		if err != nil {
			return nil, err
		}
		icode[idx] = num
	}
	return icode, nil
}

func RunSync(icode IntCode, input <-chan int64, output chan<- Message) {
	run(icode, false, input, output)
}

func RunAsync(icode IntCode, input <-chan int64, output chan<- Message) {
	run(icode, true, input, output)
}

func run(icode IntCode, async bool, input <-chan int64, output chan<- Message) {
	memory := make(IntCode, len(icode))
	copy(memory, icode)

	getMemoryPointer := func(index int64) *int64 {
		for int64(len(memory)) <= index {
			memory = append(memory, 0)
		}
		return &memory[index]
	}

	var relBase int64
	var pos int64
	for {
		instruction := memory[pos]
		opcode := instruction % 100

		getParameter := func(offset int64) *int64 {
			parameter := memory[pos+offset]
			mode := instruction / int64(math.Pow10(int(offset+1))) % 10
			switch mode {
			case 0: // position mode
				return getMemoryPointer(parameter)
			case 1: // immediate mode
				return &parameter
			case 2: // relative mode
				return getMemoryPointer(relBase + parameter)
			}
			return nil
		}

		switch opcode {
		case 1:
			a, b, c := getParameter(1), getParameter(2), getParameter(3)
			*c = *a + *b
			pos += 4
		case 2:
			a, b, c := getParameter(1), getParameter(2), getParameter(3)
			*c = *a * *b
			pos += 4
		case 3:
			if !async {
				output <- Message{Type: MessageWantsInput}
			}
			a := getParameter(1)
			*a = <-input
			pos += 2
		case 4:
			a := getParameter(1)
			output <- Message{Type: MessageOutput, Value: *a}
			pos += 2
		case 5:
			a, b := getParameter(1), getParameter(2)
			if *a != 0 {
				pos = *b
			} else {
				pos += 3
			}
		case 6:
			a, b := getParameter(1), getParameter(2)
			if *a == 0 {
				pos = *b
			} else {
				pos += 3
			}
		case 7:
			a, b, c := getParameter(1), getParameter(2), getParameter(3)
			var val int64
			if *a < *b {
				val = 1
			}
			*c = val
			pos += 4
		case 8:
			a, b, c := getParameter(1), getParameter(2), getParameter(3)
			var val int64
			if *a == *b {
				val = 1
			}
			*c = val
			pos += 4
		case 9:
			a := getParameter(1)
			relBase += *a
			pos += 2
		case 99:
			output <- Message{Type: MessageHalt}
			return
		}
	}
}
