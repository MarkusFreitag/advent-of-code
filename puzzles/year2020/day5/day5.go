package day5

import (
	"sort"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

func seatID(row, col int) int {
	return row*8 + col
}

func calcRow(str string) int {
	rows := makeSlice(0, 127)
	for _, s := range str {
		switch string(s) {
		case "B":
			rows = makeSlice(rows[0]+len(rows)/2, rows[len(rows)-1])
		case "F":
			rows = makeSlice(rows[0], rows[0]+len(rows)/2)
		}
	}
	return rows[0]
}

func calcCol(str string) int {
	cols := makeSlice(0, 7)
	for _, s := range str {
		switch string(s) {
		case "R":
			cols = makeSlice(cols[0]+len(cols)/2, cols[len(cols)-1])
		case "L":
			cols = makeSlice(cols[0], cols[0]+len(cols)/2)
		}
	}
	return cols[0]
}

func makeSlice(min, max int) []int {
	slice := make([]int, 0)
	for i := min; i <= max; i++ {
		slice = append(slice, i)
	}
	return slice
}

func Part1(input string) (string, error) {
	ids := make([]int, 0)
	for _, line := range strings.Split(input, "\n") {
		row := calcRow(line[:7])
		col := calcCol(line[7:])
		ids = append(ids, seatID(row, col))
	}
	sort.Ints(ids)
	return strconv.Itoa(ids[len(ids)-1]), nil
}

func Part2(input string) (string, error) {
	ids := make([]int, 0)
	for _, line := range strings.Split(input, "\n") {
		row := calcRow(line[:7])
		col := calcCol(line[7:])
		ids = append(ids, seatID(row, col))
	}
	sort.Ints(ids)
	for i := ids[0]; i <= ids[len(ids)-1]; i++ {
		if !util.IntInSlice(i, ids) {
			return strconv.Itoa(i), nil
		}
	}
	return "no missing seat id", nil
}
