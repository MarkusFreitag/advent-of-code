package day4

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

var (
	rgxHairColor = regexp.MustCompile(`^#[a-z0-9]{6}$`)
	rgxPID       = regexp.MustCompile(`^[0-9]{9}$`)
)

func Part1(input string) (string, error) {
	var valid int
	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, block := range strings.Split(input, "\n\n") {
		v := true
		for _, field := range fields {
			if !strings.Contains(block, field+":") {
				v = false
				break
			}
		}
		if v {
			valid++
		}
	}

	return strconv.Itoa(valid), nil
}

func Part2(input string) (string, error) {
	var valid int
	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, block := range strings.Split(input, "\n\n") {
		v := true
		for _, field := range fields {
			if !strings.Contains(block, field+":") {
				v = false
				break
			}
		}
		if !v {
			continue
		}

		block := strings.ReplaceAll(block, "\n", " ")
		parts := strings.Split(block, " ")
		for _, part := range parts {
			b := strings.Split(part, ":")
			switch b[0] {
			case "byr":
				yr, _ := strconv.Atoi(b[1])
				if yr < 1920 || yr > 2002 {
					v = false
				}
			case "iyr":
				yr, _ := strconv.Atoi(b[1])
				if yr < 2010 || yr > 2020 {
					v = false
				}
			case "eyr":
				yr, _ := strconv.Atoi(b[1])
				if yr < 2020 || yr > 2030 {
					v = false
				}
			case "hgt":
				if strings.HasSuffix(b[1], "cm") {
					s := strings.TrimSuffix(b[1], "cm")
					h, _ := strconv.Atoi(s)
					if h < 150 || h > 193 {
						v = false
					}
				} else if strings.HasSuffix(b[1], "in") {
					s := strings.TrimSuffix(b[1], "in")
					h, _ := strconv.Atoi(s)
					if h < 59 || h > 76 {
						v = false
					}
				} else {
					v = false
				}
			case "hcl":
				if !rgxHairColor.MatchString(b[1]) {
					v = false
				}
			case "ecl":
				colors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
				if !util.StrInSlice(b[1], colors) {
					v = false
				}
			case "pid":
				if !rgxPID.MatchString(b[1]) {
					v = false
				}
			}
		}
		if v {
			valid++
		}
	}

	return strconv.Itoa(valid), nil
}
