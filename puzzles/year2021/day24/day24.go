package day24

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

func intAppendDigit(num, digit int) int { return num*10 + digit }

func value(reg map[string]int, s string) int {
	if v, err := strconv.Atoi(s); err == nil {
		return v
	}
	return reg[s]
}

func runBlock(lines []string, digit, z int) int {
	reg := map[string]int{"w": digit, "x": 0, "y": 0, "z": z}
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) != 3 {
			continue
		}
		a, b := value(reg, fields[1]), value(reg, fields[2])
		switch fields[0] {
		case "add":
			reg[fields[1]] = a + b
		case "mul":
			reg[fields[1]] = a * b
		case "div":
			if b == 0 {
				return numbers.MaxInteger
			}
			reg[fields[1]] = a / b
		case "mod":
			if a < 0 || b <= 0 {
				return numbers.MaxInteger
			}
			reg[fields[1]] = a % b
		case "eql":
			if a == b {
				reg[fields[1]] = 1
			} else {
				reg[fields[1]] = 0
			}
		}
	}
	return reg["z"]
}

func parseInput(input string) [][]string {
	blocks := make([][]string, 0)
	block := make([]string, 0)
	for _, line := range strings.Split(input, "\n") {
		if len(block) > 0 && strings.HasPrefix(line, "inp") {
			blocks = append(blocks, block)
			block = make([]string, 0)
		}
		block = append(block, line)
	}
	blocks = append(blocks, block)
	return blocks
}

func tryDigits(block []string, cache map[int]int, maxZ int, minNum bool) map[int]int {
	newCache := make(map[int]int)
	for prevZ, num := range cache {
		for d := 1; d <= 9; d++ {
			z := runBlock(block, d, prevZ)
			if z > maxZ {
				continue
			}
			newNum := intAppendDigit(num, d)
			if _, ok := newCache[z]; !ok {
				newCache[z] = newNum
			} else {
				if minNum {
					newCache[z] = numbers.Min(newCache[z], newNum)
				} else {
					newCache[z] = numbers.Max(newCache[z], newNum)
				}
			}
		}
	}
	return newCache
}

func Part1(input string) (string, error) {
	blocks := parseInput(input)

	zCache := make(map[int]int)
	zCache[0] = 0
	for i := 1; i < 15; i++ {
		zCache = tryDigits(blocks[i-1], zCache, numbers.Pow(26, 14-i), false)
	}

	return strconv.Itoa(zCache[0]), nil
}

func Part2(input string) (string, error) {
	blocks := parseInput(input)

	zCache := make(map[int]int)
	zCache[0] = 0
	for i := 1; i < 15; i++ {
		zCache = tryDigits(blocks[i-1], zCache, numbers.Pow(26, 14-i), true)
	}

	return strconv.Itoa(zCache[0]), nil
}
