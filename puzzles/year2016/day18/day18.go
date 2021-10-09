package day18

import (
	"strconv"
)

func count(row []bool, state bool) int {
	var i int
	for _, f := range row {
		if f == state {
			i++
		}
	}
	return i
}

func parseRow(str string) []bool {
	row := make([]bool, len(str))
	for i, r := range str {
		if r == '^' {
			row[i] = true
		}
	}
	return row
}

func trap(p []bool) bool {
	// ^ ^ .
	if p[0] && p[1] && !p[2] {
		return true
	}
	// . ^ ^
	if !p[0] && p[1] && p[2] {
		return true
	}
	// ^ . .
	if p[0] && !p[1] && !p[2] {
		return true
	}
	// . . ^
	if !p[0] && !p[1] && p[2] {
		return true
	}
	return false
}

func generateRow(row []bool) []bool {
	r := make([]bool, len(row))
	for i := 0; i < len(row); i++ {
		part := make([]bool, 3)
		if i == 0 {
			part[1] = row[i]
			part[2] = row[i+1]
		} else if i == len(row)-1 {
			part[0] = row[i-1]
			part[1] = row[i]
		} else {
			part = row[i-1 : i+2]
		}
		r[i] = trap(part)
	}
	return r
}

func run(row []bool) int {
	var safe int
	for i := 0; i < totalRows; i++ {
		safe += count(row, false)
		row = generateRow(row)
	}
	return safe
}

var totalRows = 40

func Part1(input string) (string, error) {
	return strconv.Itoa(run(parseRow(input))), nil
}

func Part2(input string) (string, error) {
	totalRows = 400000
	return strconv.Itoa(run(parseRow(input))), nil
}
