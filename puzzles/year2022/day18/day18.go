package day18

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
	"github.com/MarkusFreitag/advent-of-code/util/slice"
)

var (
	neighbors = [6]point{
		{-1, 0, 0},
		{1, 0, 0},
		{0, -1, 0},
		{0, 1, 0},
		{0, 0, -1},
		{0, 0, 1},
	}
	exist struct{}
)

type point [3]int

func (pt point) Add(p point) point {
	return point{
		pt[0] + p[0], pt[1] + p[1], pt[2] + p[2],
	}
}

func parseInput(input string) map[point]struct{} {
	lines := strings.Split(input, "\n")
	cubes := make(map[point]struct{})
	for _, line := range lines {
		parts := strings.Split(line, ",")
		cubes[point{
			util.ParseInt(parts[0]),
			util.ParseInt(parts[1]),
			util.ParseInt(parts[2]),
		}] = exist
	}
	return cubes
}

func Part1(input string) (string, error) {
	lava := parseInput(input)

	var total int
	for cube := range lava {
		sides := 6
		for _, neigh := range neighbors {
			if _, ok := lava[cube.Add(neigh)]; ok {
				sides--
			}
		}
		total += sides
	}

	return strconv.Itoa(total), nil
}

func Part2(input string) (string, error) {
	lava := parseInput(input)
	minCube := point{numbers.MaxInteger, numbers.MaxInteger, numbers.MaxInteger}
	maxCube := point{numbers.MinInteger, numbers.MinInteger, numbers.MinInteger}
	for cube := range lava {
		minCube[0] = numbers.Min(minCube[0], cube[0])
		minCube[1] = numbers.Min(minCube[1], cube[1])
		minCube[2] = numbers.Min(minCube[2], cube[2])
		maxCube[0] = numbers.Max(maxCube[0], cube[0])
		maxCube[1] = numbers.Max(maxCube[1], cube[1])
		maxCube[2] = numbers.Max(maxCube[2], cube[2])
	}

	minCube = minCube.Add(point{-1, -1, -1})

	air := make(map[point]struct{})
	air[minCube] = exist
	queue := make([]point, 0)
	queue = append(queue, minCube)
	var total int
	for len(queue) > 0 {
		var cube point
		cube, queue = slice.PopFront(queue)

		for _, neigh := range neighbors {
			newCube := cube.Add(neigh)

			if newCube[0] < minCube[0]-1 || newCube[1] < minCube[1]-1 || newCube[2] < minCube[2]-1 {
				continue
			}
			if newCube[0] > maxCube[0]+1 || newCube[1] > maxCube[1]+1 || newCube[2] > maxCube[2]+1 {
				continue
			}

			_, lavaOk := lava[newCube]
			if lavaOk {
				total++
			}

			_, airOk := air[newCube]
			if airOk || lavaOk {
				continue
			}

			queue = append(queue, newCube)
			air[newCube] = exist
		}
	}

	return strconv.Itoa(total), nil
}
