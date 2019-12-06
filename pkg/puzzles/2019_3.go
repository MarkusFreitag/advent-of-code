package puzzles

import (
	"math"
	"strconv"
	"strings"
)

type point struct {
	X, Y int
}

func newPoint(x, y int) *point {
	return &point{X: x, Y: y}
}

type points []*point

func (pts points) contains(pt *point) int {
	for idx, item := range pts {
		if item.X == pt.X && item.Y == pt.Y {
			return idx
		}
	}
	return -1
}

func calcPath(pathStr string) (points, error) {
	path := make(points, 1)
	path[0] = newPoint(0, 0)
	for _, item := range strings.Split(pathStr, ",") {
		dir := string(item[0])
		count, err := strconv.Atoi(item[1:])
		if err != nil {
			return nil, err
		}
		last := path[len(path)-1]
		for i := 1; i <= count; i++ {
			switch strings.ToUpper(dir) {
			case "R":
				path = append(path, newPoint(last.X+i, last.Y))
			case "L":
				path = append(path, newPoint(last.X-i, last.Y))
			case "U":
				path = append(path, newPoint(last.X, last.Y+i))
			case "D":
				path = append(path, newPoint(last.X, last.Y-i))
			}
		}
	}
	return path, nil
}

type y2019d3p1 struct{}

func (p *y2019d3p1) Solve(input string) (string, error) {
	items := strings.Split(input, "\n")
	wireA, err := calcPath(items[0])
	if err != nil {
		return "", err
	}
	wireB, err := calcPath(items[1])
	if err != nil {
		return "", err
	}
	crossing := math.MaxInt64
	for _, pt := range wireA {
		if pt.Y == 0 && pt.X == 0 {
			continue
		}
		if pos := wireB.contains(pt); pos != -1 {
			length := int(math.Abs(float64(pt.Y)) + math.Abs(float64(pt.X)))
			if length < crossing {
				crossing = length
			}
		}
	}
	return strconv.Itoa(crossing), nil
}

type y2019d3p2 struct{}

func (p *y2019d3p2) Solve(input string) (string, error) {
	items := strings.Split(input, "\n")
	wireA, err := calcPath(items[0])
	if err != nil {
		return "", err
	}
	wireB, err := calcPath(items[1])
	if err != nil {
		return "", err
	}
	crossing := math.MaxInt64
	for idx, pt := range wireA {
		if pt.Y == 0 && pt.X == 0 {
			continue
		}
		if pos := wireB.contains(pt); pos != -1 {
			length := idx + pos
			if length < crossing {
				crossing = length
			}
		}
	}
	return strconv.Itoa(crossing), nil
}
