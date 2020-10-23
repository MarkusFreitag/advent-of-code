package day16

import (
	"strconv"
	"strings"
)

var auntSue = desc{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

type desc map[string]int

func newDesc(str string) desc {
	d := make(desc)
	attrs := strings.Split(str, ",")
	for _, attr := range attrs {
		parts := strings.Split(attr, ":")
		num, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		d[strings.TrimSpace(parts[0])] = num
	}
	return d
}

func Part1(input string) (string, error) {
	var auntID string
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimPrefix(line, "Sue ")
		idx := strings.Index(line, ":")
		match := true
		for key, value := range newDesc(line[idx+1:]) {
			if v, ok := auntSue[key]; !ok || v != value {
				match = false
			}
		}
		if match {
			auntID = line[:idx]
		}
	}
	return auntID, nil
}

func Part2(input string) (string, error) {
	var auntID string
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimPrefix(line, "Sue ")
		idx := strings.Index(line, ":")
		match := true
		for key, value := range newDesc(line[idx+1:]) {
			switch key {
			case "cats", "trees":
				if v, ok := auntSue[key]; !ok || v > value {
					match = false
				}
			case "pomeranians", "goldfish":
				if v, ok := auntSue[key]; !ok || v < value {
					match = false
				}
			default:
				if v, ok := auntSue[key]; !ok || v != value {
					match = false
				}
			}
		}
		if match {
			auntID = line[:idx]
		}
	}
	return auntID, nil
}
