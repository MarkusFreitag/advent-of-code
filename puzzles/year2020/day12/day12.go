package day12

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

type Action struct {
	Cmd   rune
	Steps int
}

func turn(now, dir rune) rune {
	if dir == 'R' {
		switch now {
		case 'N':
			return 'E'
		case 'E':
			return 'S'
		case 'S':
			return 'W'
		case 'W':
			return 'N'
		}
	}

	switch now {
	case 'N':
		return 'W'
	case 'E':
		return 'N'
	case 'S':
		return 'E'
	case 'W':
		return 'S'
	}
	return '+'
}

func Part1(input string) (string, error) {
	actions := make([]Action, 0)
	for _, line := range strings.Split(input, "\n") {
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			return "", err
		}
		actions = append(actions, Action{
			Cmd:   rune(line[0]),
			Steps: num,
		})
	}

	var x, y int
	facing := 'E'
	for _, action := range actions {
		switch action.Cmd {
		case 'N':
			y += action.Steps
		case 'S':
			y -= action.Steps
		case 'E':
			x += action.Steps
		case 'W':
			x -= action.Steps
		case 'L', 'R':
			count := action.Steps / 90
			for i := 0; i < count; i++ {
				facing = turn(facing, action.Cmd)
			}
		case 'F':
			switch facing {
			case 'N':
				y += action.Steps
			case 'S':
				y -= action.Steps
			case 'E':
				x += action.Steps
			case 'W':
				x -= action.Steps
			}
		}
	}
	return strconv.Itoa(util.Abs(x) + util.Abs(y)), nil
}

func Part2(input string) (string, error) {
	actions := make([]Action, 0)
	for _, line := range strings.Split(input, "\n") {
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			return "", err
		}
		actions = append(actions, Action{
			Cmd:   rune(line[0]),
			Steps: num,
		})
	}

	wx, wy := 10, 1
	var sx, sy int
	for _, action := range actions {
		switch action.Cmd {
		case 'N':
			wy += action.Steps
		case 'S':
			wy -= action.Steps
		case 'E':
			wx += action.Steps
		case 'W':
			wx -= action.Steps
		case 'L', 'R':
			count := action.Steps / 90
			for i := 0; i < count; i++ {
				wx, wy = turnWP(wx, wy, action.Cmd)
			}
		case 'F':
			sx += action.Steps * wx
			sy += action.Steps * wy
		}
	}
	return strconv.Itoa(util.Abs(sx) + util.Abs(sy)), nil
}

func turnWP(wx, wy int, dir rune) (int, int) {
	if dir == 'R' {
		return wy, wx * -1
	}
	return wy * -1, wx
}
