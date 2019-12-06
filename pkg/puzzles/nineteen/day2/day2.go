package day2

import (
	"strconv"
	"strings"
)

type intcode []int

func (icode intcode) Interpret() {
	for i := 0; i < len(icode); i += 4 {
		if icode[i] == 1 {
			icode[icode[i+3]] = icode[icode[i+1]] + icode[icode[i+2]]
		} else if icode[i] == 2 {
			icode[icode[i+3]] = icode[icode[i+1]] * icode[icode[i+2]]
		} else if icode[i] == 99 {
			break
		}
	}
}

type Part1 struct {
	intcode intcode
}

func (p *Part1) Solve(input string) (string, error) {
	items := strings.Split(input, ",")
	p.intcode = make(intcode, len(items))
	for idx, item := range items {
		num, err := strconv.Atoi(item)
		if err != nil {
			return "", err
		}
		p.intcode[idx] = num
	}
	p.intcode[1] = 12
	p.intcode[2] = 2
	p.intcode.Interpret()
	return strconv.Itoa(p.intcode[0]), nil
}

type Part2 struct {
	intcode intcode
}

func (p *Part2) Solve(input string) (string, error) {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			items := strings.Split(input, ",")
			p.intcode = make(intcode, len(items))
			for idx, item := range items {
				num, err := strconv.Atoi(item)
				if err != nil {
					return "", err
				}
				p.intcode[idx] = num
			}
			p.intcode[1] = noun
			p.intcode[2] = verb
			p.intcode.Interpret()
			if p.intcode[0] == 19690720 {
				return strconv.Itoa(100*noun + verb), nil
			}
		}
	}
	return "", nil
}
