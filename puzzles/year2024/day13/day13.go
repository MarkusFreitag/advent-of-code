package day13

import (
	"image"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

type Machine struct {
	ButtonA image.Point
	ButtonB image.Point
	Prize   image.Point
}

func (m Machine) Tokens() int {
	/*
		ax,ay -> ButtonA coordinates
		bx,by -> ButtonB coordinates
		px,py -> Prize coordinates

		a -> amount of ButtonA presses
		b -> amount of ButtonB presses

		base formulas to calculate the prize point
			px = ax * a + bx * b
			py = ay * a + by * b

		resolve both base formulas for a
			px = ax * a + bx * b	| - bx*b
			px - bx * b = ax * a	| / ax
			(px - bx * b) / ax = a

			py = ay * a + by * b	| - bx*b
			py - bx * b = ay * a	| / ay
			(py - bx * b) / ay = a

		combine both resolved formulas and resolve for b
			(px - bx * b) / ax = (py - by * b) / ay	| *ax and *ay
			ay * (px - bx * b) = ax * (py - by * b)	| multiply to get rid of the parentheses
			ay*px - ay*bx*b = ax*py - ax*by*b				| +(ay*bx*b) and -(ax*py) to bring b to one side
			ay*px - ax*py = ay*bx*b - ax*by*b				| isolate b
			ay*px - ax*py = b * (ay*bx - ax*by)			| / (ay*bx - ax*by)
			(ay*px - ax*py) / (ay*bx - ax*by) = b

		calculate b and use it with one of the base formulas to get a

		using a and b calculate the prize point and check whether it is equal to the given one
	*/
	ax, ay := m.ButtonA.X, m.ButtonA.Y
	bx, by := m.ButtonB.X, m.ButtonB.Y
	px, py := m.Prize.X, m.Prize.Y

	b := (ay*px - ax*py) / (ay*bx - ax*by)
	a := (px - bx*b) / ax

	if m.Prize.Eq(image.Pt((ax*a)+(bx*b), (ay*a)+(by*b))) {
		return 3*a + b
	}
	return 0
}

func parseInput(input string) []Machine {
	machines := make([]Machine, 0)
	for _, block := range strings.Split(input, "\n\n") {
		lines := strings.Split(block, "\n")
		var machine Machine
		fields := strings.Fields(lines[0])
		machine.ButtonA = image.Pt(
			util.ParseInt(strings.TrimPrefix(strings.TrimSuffix(fields[2], ","), "X+")),
			util.ParseInt(strings.TrimPrefix(fields[3], "Y+")),
		)
		fields = strings.Fields(lines[1])
		machine.ButtonB = image.Pt(
			util.ParseInt(strings.TrimPrefix(strings.TrimSuffix(fields[2], ","), "X+")),
			util.ParseInt(strings.TrimPrefix(fields[3], "Y+")),
		)
		fields = strings.Fields(lines[2])
		machine.Prize = image.Pt(
			util.ParseInt(strings.TrimPrefix(strings.TrimSuffix(fields[1], ","), "X=")),
			util.ParseInt(strings.TrimPrefix(fields[2], "Y=")),
		)
		machines = append(machines, machine)
	}
	return machines
}

func Part1(input string) (string, error) {
	machines := parseInput(input)
	var total int
	for _, machine := range machines {
		total += machine.Tokens()
	}
	return strconv.Itoa(total), nil
}

func Part2(input string) (string, error) {
	machines := parseInput(input)
	const off = 10000000000000
	for idx, machine := range machines {
		machine.Prize.X += off
		machine.Prize.Y += off
		machines[idx] = machine
	}
	var total int
	for _, machine := range machines {
		total += machine.Tokens()
	}
	return strconv.Itoa(total), nil
}
