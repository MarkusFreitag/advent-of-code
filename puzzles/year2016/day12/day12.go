package day12

import (
	"strconv"
	"strings"
)

func runProgram(lines []string, registers map[string]int) map[string]int {
	for i := 0; i < len(lines); i++ {
		parts := strings.Fields(lines[i])
		switch parts[0] {
		case "cpy":
			v, err := strconv.Atoi(parts[1])
			if err != nil {
				v = registers[parts[1]]
			}
			registers[parts[2]] = v
		case "inc":
			registers[parts[1]] = registers[parts[1]] + 1
		case "dec":
			registers[parts[1]] = registers[parts[1]] - 1
		case "jnz":
			p, err := strconv.Atoi(parts[1])
			if err != nil {
				p = registers[parts[1]]
			}
			if p != 0 {
				v, _ := strconv.Atoi(parts[2])
				i += v - 1
			}
		}
	}
	return registers
}

func Part1(input string) (string, error) {
	registers := make(map[string]int)
	registers = runProgram(strings.Split(input, "\n"), registers)
	return strconv.Itoa(registers["a"]), nil
}

func Part2(input string) (string, error) {
	registers := make(map[string]int)
	registers["c"] = 1
	registers = runProgram(strings.Split(input, "\n"), registers)
	return strconv.Itoa(registers["a"]), nil
}
