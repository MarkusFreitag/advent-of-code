package day1

import (
	"errors"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

func Part1(input string) (string, error) {
	err := errors.New("couldn't find valid number pair")
	expenses := util.StringsToInts(strings.Split(input, "\n"))
	for i, a := range expenses {
		if i+1 >= len(expenses) {
			return "", err
		}
		for _, b := range expenses[i+1:] {
			if a+b == 2020 {
				return strconv.Itoa(a * b), nil
			}
		}
	}
	return "", err
}

func Part2(input string) (string, error) {
	err := errors.New("couldn't find valid number triplet")
	expenses := util.StringsToInts(strings.Split(input, "\n"))
	for i, a := range expenses {
		if i+1 >= len(expenses) || i+2 >= len(expenses) {
			return "", err
		}
		for _, b := range expenses[i+1:] {
			for _, c := range expenses[i+2:] {
				if a+b+c == 2020 {
					return strconv.Itoa(a * b * c), nil
				}
			}
		}
	}
	return "", err
}
