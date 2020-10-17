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

func Part1(input string) (string, error) {
	items := strings.Split(input, ",")
	icode := make(intcode, len(items))
	for idx, item := range items {
		num, err := strconv.Atoi(item)
		if err != nil {
			return "", err
		}
		icode[idx] = num
	}
	icode[1] = 12
	icode[2] = 2
	icode.Interpret()
	return strconv.Itoa(icode[0]), nil
}

func Part2(input string) (string, error) {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			items := strings.Split(input, ",")
			icode := make(intcode, len(items))
			for idx, item := range items {
				num, err := strconv.Atoi(item)
				if err != nil {
					return "", err
				}
				icode[idx] = num
			}
			icode[1] = noun
			icode[2] = verb
			icode.Interpret()
			if icode[0] == 19690720 {
				return strconv.Itoa(100*noun + verb), nil
			}
		}
	}
	return "", nil
}
