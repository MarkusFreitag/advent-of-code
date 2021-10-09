package day2

import (
	"strings"
)

func enterCode(kp keypad, pos position, instructions []string) string {
	var code string
	for _, instruction := range instructions {
		for _, str := range instruction {
			switch str {
			case 'U':
				if kp.check(pos.row-1, pos.col) {
					pos.row--
				}
			case 'R':
				if kp.check(pos.row, pos.col+1) {
					pos.col++
				}
			case 'D':
				if kp.check(pos.row+1, pos.col) {
					pos.row++
				}
			case 'L':
				if kp.check(pos.row, pos.col-1) {
					pos.col--
				}
			}
		}
		code += kp[pos.row][pos.col]
	}
	return code
}

type keypad [][]string

func (k keypad) check(row, col int) bool {
	if row >= 0 && row < len(k) {
		if col >= 0 && col < len(k[row]) {
			return bool(k[row][col] != "")
		}
	}
	return false
}

type position struct {
	row, col int
}

func Part1(input string) (string, error) {
	kp := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
	return enterCode(
		kp,
		position{row: 1, col: 1},
		strings.Split(input, "\n"),
	), nil
}

func Part2(input string) (string, error) {
	kp := [][]string{
		{"", "", "1", "", ""},
		{"", "2", "3", "4", ""},
		{"5", "6", "7", "8", "9"},
		{"", "A", "B", "C", ""},
		{"", "", "D", "", ""},
	}
	return enterCode(
		kp,
		position{row: 2, col: 0},
		strings.Split(input, "\n"),
	), nil
}
