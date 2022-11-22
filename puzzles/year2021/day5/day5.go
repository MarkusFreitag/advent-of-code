package day5

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

type Points map[string]int

func (pts Points) Add(x, y int) {
	pt := fmt.Sprintf("%d|%d", x, y)
	pts[pt] = pts[pt] + 1
}

func (pts Points) Count(limit int) int {
	var count int
	for _, v := range pts {
		if v >= limit {
			count++
		}
	}
	return count
}

func parseLine(str string) (int, int, int, int) {
	fields := strings.Fields(str)
	start := strings.Split(fields[0], ",")
	end := strings.Split(fields[2], ",")
	return util.ParseInt(start[0]), util.ParseInt(start[1]),
		util.ParseInt(end[0]), util.ParseInt(end[1])
}

func Part1(input string) (string, error) {
	points := make(Points)
	for _, line := range strings.Split(input, "\n") {
		startX, startY, endX, endY := parseLine(line)

		if startX == endX {
			min, max := numbers.MinMax(startY, endY)
			for y := min; y <= max; y++ {
				points.Add(startX, y)
			}
		} else if startY == endY {
			min, max := numbers.MinMax(startX, endX)
			for x := min; x <= max; x++ {
				points.Add(x, startY)
			}
		}
	}
	return strconv.Itoa(points.Count(2)), nil
}

func Part2(input string) (string, error) {
	points := make(Points)
	for _, line := range strings.Split(input, "\n") {
		startX, startY, endX, endY := parseLine(line)

		if startX == endX {
			min, max := numbers.MinMax(startY, endY)
			for y := min; y <= max; y++ {
				points.Add(startX, y)
			}
		} else if startY == endY {
			min, max := numbers.MinMax(startX, endX)
			for x := min; x <= max; x++ {
				points.Add(x, startY)
			}
		} else {
			minX, maxX := numbers.MinMax(startX, endX)
			minY, maxY := numbers.MinMax(startY, endY)
			for y := minY; y <= maxY; y++ {
				for x := minX; x <= maxX; x++ {
					if util.OnLineInt(startX, startY, endX, endY, x, y) {
						points.Add(x, y)
					}
				}
			}
		}
	}
	return strconv.Itoa(points.Count(2)), nil
}
