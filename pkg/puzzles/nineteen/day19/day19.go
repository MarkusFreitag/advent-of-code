package day19

import (
	"strconv"

	"github.com/MarkusFreitag/advent-of-code/pkg/intcode"
)

func checkPoint(icode intcode.IntCode, y, x int) bool {
	in := make(chan int64)
	out := make(chan intcode.Message)
	go intcode.RunAsync(icode, in, out)
	in <- int64(x)
	in <- int64(y)
	msg := <-out
	return int(msg.Value) == 1
}

type Part1 struct{}

func (p *Part1) Solve(input string) (string, error) {
	icode, err := intcode.New(input)
	if err != nil {
		return "", err
	}

	var fields int
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			if checkPoint(icode, y, x) {
				fields++
			}
		}
	}
	return strconv.Itoa(fields), nil
}

type Part2 struct{}

func (p *Part2) Solve(input string) (string, error) {
	icode, err := intcode.New(input)
	if err != nil {
		return "", err
	}

	y := 99
	var x int
	for {
		for {
			if checkPoint(icode, y, x) {
				if checkPoint(icode, y-99, x+99) {
					return strconv.Itoa(x*10000 + y - 99), nil
				}
				break
			}
			x++
		}
		y++
	}
}
