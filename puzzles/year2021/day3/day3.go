package day3

import (
	"strconv"
	"strings"
)

func countBits(nums []string) []map[string]int {
	bits := make([]map[string]int, len(nums[0]))
	for nIdx, line := range nums {
		for cIdx, char := range line {
			var m map[string]int
			if nIdx == 0 {
				m = make(map[string]int)
				m["0"] = 0
				m["1"] = 0
			} else {
				m = bits[cIdx]
			}
			m[string(char)] = m[string(char)] + 1
			bits[cIdx] = m
		}
	}
	return bits
}

func result(a, b string) string {
	aDec, _ := strconv.ParseInt(a, 2, 64)
	bDec, _ := strconv.ParseInt(b, 2, 64)
	return strconv.FormatInt(aDec*bDec, 10)
}

func Part1(input string) (string, error) {
	bits := countBits(strings.Fields(input))

	var gammaRate, epsilonRate string
	for _, bit := range bits {
		if bit["0"] > bit["1"] {
			gammaRate += "0"
			epsilonRate += "1"
		} else {
			gammaRate += "1"
			epsilonRate += "0"
		}
	}

	return result(gammaRate, epsilonRate), nil
}

func calculateRate(nums []string, criteria func(map[string]int) string) string {
	for pos := 0; ; pos++ {
		bits := countBits(nums)
		newNums := make([]string, 0)
		commonBit := criteria(bits[pos])
		for _, num := range nums {
			if string(num[pos]) == commonBit {
				newNums = append(newNums, num)
			}
		}
		if len(newNums) == 1 {
			return newNums[0]
		}
		nums = newNums
	}
}

func Part2(input string) (string, error) {
	oxygenRate := calculateRate(
		strings.Fields(input),
		func(m map[string]int) string {
			if m["0"] <= m["1"] {
				return "0"
			}
			return "1"
		},
	)
	co2Rate := calculateRate(
		strings.Fields(input),
		func(m map[string]int) string {
			if m["0"] > m["1"] {
				return "0"
			}
			return "1"
		},
	)

	return result(oxygenRate, co2Rate), nil
}
