package day4

import (
	"regexp"
	"strconv"
	"strings"
)

func requiredFields(block string) bool {
	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, field := range fields {
		if !strings.Contains(block, field+":") {
			return false
		}
	}
	return true
}

func inBetween(s string, min, max int) bool {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num >= min && num <= max
}

func Part1(input string) (string, error) {
	var valid int
	for _, block := range strings.Split(input, "\n\n") {
		if requiredFields(block) {
			valid++
		}
	}
	return strconv.Itoa(valid), nil
}

func Part2(input string) (string, error) {
	rgxHairColor := regexp.MustCompile(`^#[a-z0-9]{6}$`)
	rgxValidColor := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	rgxPID := regexp.MustCompile(`^[0-9]{9}$`)

	var valid int
	for _, block := range strings.Split(input, "\n\n") {
		if !requiredFields(block) {
			continue
		}
		v := true
		block = strings.ReplaceAll(block, "\n", " ")
		for _, part := range strings.Split(block, " ") {
			pair := strings.Split(part, ":")
			switch pair[0] {
			case "byr":
				if !inBetween(pair[1], 1920, 2002) {
					v = false
				}
			case "iyr":
				if !inBetween(pair[1], 2010, 2020) {
					v = false
				}
			case "eyr":
				if !inBetween(pair[1], 2020, 2030) {
					v = false
				}
			case "hgt":
				if strings.HasSuffix(pair[1], "cm") {
					if !inBetween(strings.TrimSuffix(pair[1], "cm"), 150, 193) {
						v = false
					}
				} else if strings.HasSuffix(pair[1], "in") {
					if !inBetween(strings.TrimSuffix(pair[1], "in"), 59, 76) {
						v = false
					}
				} else {
					v = false
				}
			case "hcl":
				if !rgxHairColor.MatchString(pair[1]) {
					v = false
				}
			case "ecl":
				if !rgxValidColor.MatchString(pair[1]) {
					v = false
				}
			case "pid":
				if !rgxPID.MatchString(pair[1]) {
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
