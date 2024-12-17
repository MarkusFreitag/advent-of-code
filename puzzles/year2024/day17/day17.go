package day17

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

func run(registers map[string]int, program []int) (string, map[string]int) {
	evalComboOperand := func(op int) int {
		if op >= 0 && op <= 3 {
			return op
		}
		if op == 4 {
			return registers["A"]
		}
		if op == 5 {
			return registers["B"]
		}
		if op == 6 {
			return registers["C"]
		}
		if op == 7 {
			panic("reserved opcode")
		}
		panic("invalid opcode")
	}

	output := make([]string, 0)
	var pos int
	for {
		instruction := program[pos]
		operand := program[pos+1]

		var jumped bool
		switch instruction {
		case 0:
			registers["A"] = registers["A"] / numbers.Pow(2, evalComboOperand(operand))
		case 1:
			registers["B"] = registers["B"] ^ operand
		case 2:
			registers["B"] = evalComboOperand(operand) % 8
		case 3:
			if registers["A"] != 0 {
				pos = operand
				jumped = true
			}
		case 4:
			registers["B"] = registers["B"] ^ registers["C"]
		case 5:
			output = append(output, strconv.Itoa(evalComboOperand(operand)%8))
		case 6:
			registers["B"] = registers["A"] / numbers.Pow(2, evalComboOperand(operand))
		case 7:
			registers["C"] = registers["A"] / numbers.Pow(2, evalComboOperand(operand))
		}

		if !jumped {
			pos += 2
		}
		if pos >= len(program) {
			break
		}
	}

	return strings.Join(output, ","), registers
}

func Part1(input string) (string, error) {
	blocks := strings.Split(input, "\n\n")
	registers := make(map[string]int)
	for _, line := range strings.Split(blocks[0], "\n") {
		parts := strings.Fields(line)
		registers[strings.TrimSuffix(parts[1], ":")] = util.ParseInt(parts[2])
	}
	progStr := strings.Fields(blocks[1])[1]
	program := util.StringsToInts(strings.Split(progStr, ","))

	output, _ := run(registers, program)

	return output, nil
}

func Part2(input string) (string, error) {
	blocks := strings.Split(input, "\n\n")
	registers := make(map[string]int)
	for _, line := range strings.Split(blocks[0], "\n") {
		parts := strings.Fields(line)
		registers[strings.TrimSuffix(parts[1], ":")] = util.ParseInt(parts[2])
	}
	progStr := strings.Fields(blocks[1])[1]
	program := util.StringsToInts(strings.Split(progStr, ","))

	a := 1
	for {
		registers["A"] = a
		var output string
		output, registers = run(registers, program)

		if progStr == output {
			fmt.Printf("%s %s found A=%d\n", progStr, output, a)
			break
		}
		if strings.HasSuffix(progStr, output) {
			fmt.Printf("%s %s good A=%d\n", progStr, output, a)
			a = a * 8
		} else {
			a++
		}
	}

	return strconv.Itoa(a), nil
}
