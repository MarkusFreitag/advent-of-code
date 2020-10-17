package day3

import "strconv"

var dir = map[rune]point{
	'>': {X: 1, Y: 0},
	'<': {X: -1, Y: 0},
	'^': {X: 0, Y: -1},
	'v': {X: 0, Y: 1},
}

type point struct {
	X, Y int
}

func (pt point) Add(p point) point {
	return point{X: pt.X + p.X, Y: pt.Y + p.Y}
}

type points []point

func (pts points) Includes(p point) bool {
	for _, pt := range pts {
		if pt.X == p.X && pt.Y == p.Y {
			return true
		}
	}
	return false
}

func Part1(input string) (string, error) {
	santa := point{X: 0, Y: 0}
	houses := make(points, 1)
	houses[0] = santa
	for _, move := range input {
		santa = santa.Add(dir[move])
		if !houses.Includes(santa) {
			houses = append(houses, santa)
		}
	}
	return strconv.Itoa(len(houses)), nil
}

func Part2(input string) (string, error) {
	santa := point{X: 0, Y: 0}
	robot := point{X: 0, Y: 0}
	houses := make(points, 1)
	houses[0] = santa
	for idx, move := range input {
		if idx%2 == 0 {
			santa = santa.Add(dir[move])
			if !houses.Includes(santa) {
				houses = append(houses, santa)
			}
		} else {
			robot = robot.Add(dir[move])
			if !houses.Includes(robot) {
				houses = append(houses, robot)
			}
		}
	}
	return strconv.Itoa(len(houses)), nil
}
