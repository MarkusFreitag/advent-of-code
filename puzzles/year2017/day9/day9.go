package day9

import (
	"strconv"
	"strings"
)

func checkIfEscaped(str string, pos int) bool {
	if pos == 0 {
		return false
	}
	var escaped bool
	for i := pos - 1; i >= 0; i-- {
		if str[i] != '!' {
			break
		}
		escaped = !escaped
	}
	return escaped
}

func cleanupGarbage(str string) string {
	chars := []rune(str)
	clean := make([]rune, 0)
	for i := 0; i < len(chars); i++ {
		if chars[i] == '!' {
			i++
			continue
		}
		clean = append(clean, chars[i])
	}
	return string(clean)
}

func eval(line string) ([]int, int) {
	lvls := make([]int, 0)
	var currentLevel int
	garbageStart := -1
	garbageBlocks := make([]string, 0)
	for idx, char := range line {
		if char != '>' && garbageStart != -1 {
			continue
		}
		if char == '{' {
			if checkIfEscaped(line, idx) {
				continue
			}
			currentLevel++
		} else if char == '}' {
			if checkIfEscaped(line, idx) {
				continue
			}
			lvls = append(lvls, currentLevel)
			currentLevel--
		} else if char == '<' {
			if checkIfEscaped(line, idx) {
				continue
			}
			garbageStart = idx
		} else if char == '>' {
			if checkIfEscaped(line, idx) {
				continue
			}
			garbageBlocks = append(garbageBlocks, line[garbageStart+1:idx])
			garbageStart = -1
		}
	}
	if garbageStart == -1 && currentLevel > 0 {
		lvls = append(lvls, currentLevel)
	}
	var total int
	for _, block := range garbageBlocks {
		total += len(cleanupGarbage(block))
	}
	return lvls, total
}

func sum(nums []int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}

func Part1(input string) (string, error) {
	var total int
	for _, line := range strings.Split(input, "\n") {
		groupLvls, _ := eval(line)
		total += sum(groupLvls)
	}
	return strconv.Itoa(total), nil
}

func Part2(input string) (string, error) {
	var total int
	for _, line := range strings.Split(input, "\n") {
		_, garbageCount := eval(line)
		total += garbageCount
	}
	return strconv.Itoa(total), nil
}
