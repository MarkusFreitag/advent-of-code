package day15

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

var (
	interest = 2000000
	area     = 4000000
)

type point struct {
	x, y int
}

func (pt point) add(p point) point {
	return point{
		x: pt.x + p.x,
		y: pt.y + p.y,
	}
}

type sensor struct {
	pt     point
	beacon point
	dist   int
}

func parseNum(s string) int {
	r := strings.NewReplacer(":", "", ",", "")
	s = r.Replace(s)
	parts := strings.Split(s, "=")
	return util.ParseInt(parts[1])
}

func parseInput(input string) []*sensor {
	lines := strings.Split(input, "\n")
	sensors := make([]*sensor, len(lines))
	for idx, line := range strings.Split(input, "\n") {
		fields := strings.Fields(line)
		s := &sensor{
			pt: point{
				x: parseNum(fields[2]),
				y: parseNum(fields[3]),
			},
			beacon: point{
				x: parseNum(fields[8]),
				y: parseNum(fields[9]),
			},
		}
		s.dist = manhattan(s.pt, s.beacon)
		sensors[idx] = s
	}
	return sensors
}

func manhattan(a, b point) int {
	return numbers.Abs(a.x-b.x) + numbers.Abs(a.y-b.y)
}

var exist struct{}

func Part1(input string) (string, error) {
	sensors := parseInput(input)

	grid := make(map[point]string)
	var count int
	for _, sensor := range sensors {
		grid[sensor.pt] = "S"
		grid[sensor.beacon] = "B"

		y := interest

		for x := sensor.pt.x - sensor.dist - 1; x < sensor.pt.x+sensor.dist+1; x++ {
			if (y == sensor.pt.y && x == sensor.pt.x) || (y == sensor.beacon.y && x == sensor.beacon.x) {
				continue
			}
			if y != interest {
				continue
			}
			pt := point{y: y, x: x}
			if manhattan(sensor.pt, pt) <= sensor.dist {
				if _, ok := grid[pt]; !ok {
					count++
				}
				grid[pt] = "#"
			}
		}
	}

	return strconv.Itoa(count), nil
}

func free(sensors []*sensor, p point) int {
	for _, sensor := range sensors {
		if manhattan(sensor.pt, p) <= sensor.dist {
			return (sensor.dist - numbers.Abs(sensor.pt.y-p.y)) + sensor.pt.x - p.x
		}
	}
	return -1
}

func Part2(input string) (string, error) {
	sensors := parseInput(input)

	for y := 0; y <= area; y++ {
		for x := 0; x <= area; x++ {
			if skip := free(sensors, point{x: x, y: y}); skip >= 0 {
				x += skip
			} else {
				return strconv.Itoa(x*4000000 + y), nil
			}
		}
	}

	return "n/a", nil
}
