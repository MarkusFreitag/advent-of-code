package day3

import (
	"regexp"
	"strconv"
	"strings"
)

type point struct {
	X, Y int
}

func (p point) Add(pt point) point {
	return point{
		X: p.X + pt.X,
		Y: p.Y + pt.Y,
	}
}

type claim struct {
	ID    int
	Start point
	Size  point
}

func parseClaim(line string) claim {
	rgx := regexp.MustCompile(`#(\d)+\s@\s(\d+),(\d+):\s(\d+)x(\d+)`)
	matches := rgx.FindAllStringSubmatch(line, -1)[0]
	id, _ := strconv.Atoi(matches[1])
	startX, _ := strconv.Atoi(matches[2])
	startY, _ := strconv.Atoi(matches[3])
	sizeX, _ := strconv.Atoi(matches[4])
	sizeY, _ := strconv.Atoi(matches[5])

	return claim{
		ID:    id,
		Start: point{X: startX, Y: startY},
		Size:  point{X: sizeX, Y: sizeY},
	}
}

type fabric [][][]int

func newFabric(size int) fabric {
	grid := make([][][]int, size)
	for idx := range grid {
		grid[idx] = make([][]int, size)
		for j := range grid[idx] {
			grid[idx][j] = make([]int, 0)
		}
	}
	return grid
}

func (f fabric) addClaim(c claim) {
	for y := 0; y < c.Size.Y; y++ {
		for x := 0; x < c.Size.X; x++ {
			f[c.Start.Y+y][c.Start.X+x] = append(f[c.Start.Y+y][c.Start.X+x], c.ID)
		}
	}
}

func Part1(input string) (string, error) {
	fabric := newFabric(1000)
	for _, line := range strings.Split(input, "\n") {
		fabric.addClaim(parseClaim(line))
	}
	var count int
	for _, row := range fabric {
		for _, col := range row {
			if len(col) > 1 {
				count++
			}
		}
	}
	return strconv.Itoa(count), nil
}

func Part2(input string) (string, error) {
	claims := make([]claim, 0)
	for _, line := range strings.Split(input, "\n") {
		claims = append(claims, parseClaim(line))
	}

	fabric := newFabric(1000)
	for _, claim := range claims {
		fabric.addClaim(claim)
	}

	for _, c := range claims {
		save := true
		for y := 0; y < c.Size.Y; y++ {
			for x := 0; x < c.Size.X; x++ {
				if len(fabric[c.Start.Y+y][c.Start.X+x]) > 1 {
					save = false
				}
			}
		}

		if save {
			return strconv.Itoa(c.ID), nil
		}
	}
	return "", nil
}
