package day13

import (
	"strconv"
	"strings"
)

func parseBuses(str string, filter bool) []int {
	buses := make([]int, 0)
	for _, b := range strings.Split(str, ",") {
		if b == "x" {
			if filter {
				continue
			} else {
				b = "1"
			}
		}
		num, err := strconv.Atoi(b)
		if err != nil {
			panic(err)
		}
		buses = append(buses, num)
	}
	return buses
}

func Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")
	tsp, _ := strconv.Atoi(lines[0])
	buses := parseBuses(lines[1], true)

	var offset int
	for {
		for _, bus := range buses {
			if (tsp+offset)%bus == 0 {
				return strconv.Itoa(offset * bus), nil
			}
		}
		offset++
	}
}

func Part2(input string) (string, error) {
	lines := strings.Split(input, "\n")
	buses := parseBuses(lines[1], false)

	tsp := 1
	for {
		offset := 1
		valid := true

		for idx, bus := range buses {
			if (tsp+idx)%bus != 0 {
				valid = false
				break
			}
			offset *= bus
		}

		if valid {
			return strconv.Itoa(tsp), nil
		}

		tsp += offset
	}
}
