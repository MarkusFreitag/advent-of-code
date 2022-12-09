package day9

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

var exists = struct{}{}

type Point struct {
	X, Y int
}

func moveTail(head, tail Point) Point {
	distY := head.Y - tail.Y
	distX := head.X - tail.X
	if numbers.Abs(distY) >= 2 || numbers.Abs(distX) >= 2 {
		tail.Y += numbers.Sign(distY)
		tail.X += numbers.Sign(distX)
	}
	return tail
}

func ropeSimulation(moves []string, head Point, tail []Point) int {
	visited := make(map[Point]struct{})
	tailLength := len(tail)
	visited[tail[tailLength-1]] = exists
	for _, move := range moves {
		fields := strings.Fields(move)
		num := util.ParseInt(fields[1])
		for i := 0; i < num; i++ {
			switch fields[0] {
			case "U":
				head.Y++
			case "D":
				head.Y--
			case "R":
				head.X++
			case "L":
				head.X--
			}
			for idx, tailPiece := range tail {
				if idx == 0 {
					tail[idx] = moveTail(head, tailPiece)
				} else {
					tail[idx] = moveTail(tail[idx-1], tailPiece)
				}
			}
			visited[tail[tailLength-1]] = exists
		}
	}
	return len(visited)
}

func Part1(input string) (string, error) {
	head := Point{}
	tail := make([]Point, 1)
	for idx := range tail {
		tail[idx] = Point{}
	}
	return strconv.Itoa(ropeSimulation(strings.Split(input, "\n"), head, tail)), nil
}

func Part2(input string) (string, error) {
	head := Point{}
	tail := make([]Point, 9)
	for idx := range tail {
		tail[idx] = Point{}
	}
	return strconv.Itoa(ropeSimulation(strings.Split(input, "\n"), head, tail)), nil
}
