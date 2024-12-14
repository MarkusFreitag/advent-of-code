package day14

import (
	"fmt"
	"image"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util/directions"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

var (
	spaceWidth  = 101
	spaceHeight = 103
)

type Robot struct {
	Pos, Vel image.Point
}

func (r *Robot) Move() {
	r.Pos = r.Pos.Add(r.Vel)
	if r.Pos.X < 0 {
		r.Pos.X = spaceWidth + r.Pos.X
	} else if r.Pos.X >= spaceWidth {
		r.Pos.X = r.Pos.X - spaceWidth
	}
	if r.Pos.Y < 0 {
		r.Pos.Y = spaceHeight + r.Pos.Y
	} else if r.Pos.Y >= spaceHeight {
		r.Pos.Y = r.Pos.Y - spaceHeight
	}
}

func printGrid(robots []*Robot) {
	positions := make(map[image.Point]int)
	for _, robot := range robots {
		positions[robot.Pos] = positions[robot.Pos] + 1
	}

	for y := 0; y < spaceHeight; y++ {
		var line string
		for x := 0; x < spaceWidth; x++ {
			if v, ok := positions[image.Pt(x, y)]; ok {
				line += strconv.Itoa(v)
			} else {
				line += "."
			}
		}
		fmt.Println(line)
	}
}

func Part1(input string) (string, error) {
	robots := make([]*Robot, 0)
	for _, line := range strings.Split(input, "\n") {
		var robot Robot
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &robot.Pos.X, &robot.Pos.Y, &robot.Vel.X, &robot.Vel.Y)
		robots = append(robots, &robot)
	}

	for range 100 {
		for _, robot := range robots {
			robot.Move()
		}
	}

	grid := make([][]int, 0)
	for range spaceHeight {
		grid = append(grid, make([]int, spaceWidth))
	}
	for _, robot := range robots {
		grid[robot.Pos.Y][robot.Pos.X] = grid[robot.Pos.Y][robot.Pos.X] + 1
	}

	middleRow := spaceHeight / 2
	middleCol := spaceWidth / 2

	quads := make([]int, 4)
	for y, row := range grid {
		for x, col := range row {
			switch {
			case y < middleRow && x < middleCol:
				quads[0] += col
			case y < middleRow && x > middleCol:
				quads[1] += col
			case y > middleRow && x < middleCol:
				quads[2] += col
			case y > middleRow && x > middleCol:
				quads[3] += col
			}
		}
	}
	return strconv.Itoa(numbers.Multiply(quads...)), nil
}

func Part2(input string) (string, error) {
	robots := make([]*Robot, 0)
	for _, line := range strings.Split(input, "\n") {
		var robot Robot
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &robot.Pos.X, &robot.Pos.Y, &robot.Vel.X, &robot.Vel.Y)
		robots = append(robots, &robot)
	}

	for i := 1; true; i++ {
		m := make(map[image.Point]struct{})
		for _, robot := range robots {
			robot.Move()
			m[robot.Pos] = struct{}{}
		}

		/*
			Assuming that the christmas tree consists of points connected to each other
			we look for a high density of points.
		*/
		var density int
		for pt := range m {
			if _, ok := m[pt.Add(directions.Right.Point())]; ok {
				density++
			}
		}
		if density > 100 {
			printGrid(robots)
			return strconv.Itoa(i), nil
		}
	}
	return "", nil
}
