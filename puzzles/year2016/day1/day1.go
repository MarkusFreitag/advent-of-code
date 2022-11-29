package day1

import (
	"container/ring"
	"fmt"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

const (
	NORTH = iota
	EAST
	SOUTH
	WEST
)

var direction = ring.New(4)

func init() {
	for i := 0; i < 4; i++ {
		direction.Value = i
		direction = direction.Next()
	}
}

func resetDirection() {
	for direction.Value != NORTH {
		direction = direction.Next()
	}
}

func turn(str string) {
	if str == "R" {
		direction = direction.Next()
	} else {
		direction = direction.Prev()
	}
}

func Part1(input string) (string, error) {
	resetDirection()
	var x, y int
	for _, part := range strings.Split(input, ",") {
		part = strings.TrimSpace(part)

		turn(part[:1])
		steps, _ := strconv.Atoi(part[1:])

		switch direction.Value {
		case NORTH:
			y += steps
		case EAST:
			x += steps
		case SOUTH:
			y -= steps
		case WEST:
			x -= steps
		}
	}
	return strconv.Itoa(numbers.Abs(x + y)), nil
}

func Part2(input string) (string, error) {
	resetDirection()
	var x, y int
	locs := make(map[string]bool)
	for _, part := range strings.Split(input, ",") {
		part = strings.TrimSpace(part)

		turn(part[:1])
		steps, _ := strconv.Atoi(part[1:])

		for i := 0; i < steps; i++ {
			switch direction.Value {
			case NORTH:
				y++
			case EAST:
				x++
			case SOUTH:
				y--
			case WEST:
				x--
			}

			pos := fmt.Sprintf("%d|%d", x, y)
			if _, ok := locs[pos]; ok {
				return strconv.Itoa(numbers.Abs(x + y)), nil
			}
			locs[pos] = true
		}
	}
	return "-1", nil
}
