package day10

import (
	"strconv"
	"strings"
)

type astroid struct {
	X, Y, Others int
}

func newAstroid(x, y int) *astroid {
	return &astroid{X: x, Y: y}
}

func (a *astroid) CalcWaypointsTo(target *astroid) []astroid {
	sX, tX := a.X, target.X
	if a.X > target.X {
		sX, tX = target.X, a.X
	}
	sY, tY := a.Y, target.Y
	if a.Y > target.Y {
		sY, tY = target.Y, a.Y
	}
	wps := make([]astroid, 0)
	if sX == tX && sY == tY {
		return wps
	}
	if sX == tX {
		for y := sY + 1; y < tY; y++ {
			wps = append(wps, astroid{X: sX, Y: y})
		}
	} else if sY == tY {
		for x := sX + 1; x < tX; x++ {
			wps = append(wps, astroid{X: x, Y: sY})
		}
	} else {
		for x := sX + 1; x < tX; x++ {
			for y := sY + 1; y < tY; y++ {
				if (float64(tY)-float64(sY))/(float64(tX)-float64(sX)) == (float64(tY)-float64(y))/(float64(tX)-float64(x)) {
					wps = append(wps, astroid{X: x, Y: y})
				}
			}
		}
	}
	return wps
}

func Part1(input string) (string, error) {
	items := strings.Split(input, "\n")
	aField := make([][]bool, len(items))
	astroids := make([]*astroid, 0)
	for y, line := range items {
		row := make([]bool, len(line))
		for x, char := range line {
			if string(char) == "." {
				continue
			}
			row[x] = true
			astroids = append(astroids, newAstroid(x, y))
		}
		aField[y] = row
	}

	bestPos := &astroid{X: -1, Y: -1, Others: -1}
	for _, astr := range astroids {
		for _, a := range astroids {
			if astr.X == a.X && astr.Y == a.Y {
				continue
			}
			waypoints := astr.CalcWaypointsTo(a)
			vision := true
			for _, wp := range waypoints {
				if aField[wp.Y][wp.X] {
					vision = false
				}
			}
			if vision {
				astr.Others++
			}
		}
		if astr.Others > bestPos.Others {
			bestPos = astr
		}
	}
	return strconv.Itoa(bestPos.Others), nil
}

func Part2(input string) (string, error) {
	return "n/a", nil
}
