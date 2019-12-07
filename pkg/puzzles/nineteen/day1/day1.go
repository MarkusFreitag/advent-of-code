package day1

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/pkg/util"
)

type Part1 struct {
	values []int
}

func (p *Part1) Solve(input string) (string, error) {
	lines := strings.Split(input, "\n")
	p.values = make([]int, len(lines))
	for idx, line := range lines {
		mass, err := strconv.Atoi(line)
		if err != nil {
			return "", err
		}
		p.values[idx] = mass/3 - 2
	}
	return strconv.Itoa(util.Sum(p.values...)), nil
}

type Part2 struct {
	values []int
}

func (p *Part2) Solve(input string) (string, error) {
	lines := strings.Split(input, "\n")
	p.values = make([]int, 0)
	for _, line := range lines {
		mass, err := strconv.Atoi(line)
		if err != nil {
			return "", err
		}
		fuel := mass/3 - 2
		for fuel > 0 {
			p.values = append(p.values, fuel)
			fuel = fuel/3 - 2
		}
	}
	return strconv.Itoa(util.Sum(p.values...)), nil
}