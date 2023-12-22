package day18

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

type Range [2][2]int

func (r Range) Include(p [2]int) bool {
	return (p[0] >= r[0][0] && p[0] <= r[1][0] && p[1] >= r[0][1] && p[1] <= r[1][1]) || (p[0] >= r[1][0] && p[0] <= r[0][0] && p[1] >= r[1][1] && p[1] <= r[0][1])
}

func (r Range) Len() int {
	if r[0][0] == r[1][0] {
		return numbers.Abs(r[0][1] - r[1][1])
	}
	return numbers.Abs(r[0][0] - r[1][0])
}

func (r Range) String() string {
	return fmt.Sprintf("%d|%d => %d|%d", r[0][0], r[0][1], r[1][0], r[1][1])
}

type Ranges []Range

func (r Ranges) Include(p [2]int) int {
	for idx, rng := range r {
		if rng.Include(p) {
			return idx
		}
	}
	return -1
}
func (r Ranges) Len() int {
	var total int
	for _, rng := range r {
		total += rng.Len()
	}
	return total
}

/*
func count(minY, minX, maxY, maxX int, ranges Ranges, pt [2]int) int {
	queue := make([][2]int, 0)
	queue = append(queue, pt)
	seen := make(map[[2]int]bool)
	var count int
	for len(queue) > 0 {
		var item [2]int
		item, queue = sliceutil.PopFront(queue)

		if _, ok := seen[item]; ok {
			continue
		}
		seen[item] = true

		if ranges.Include(item) {
			continue
		}

		count++

		for _, dir := range [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
			p := [2]int{item[0] + dir[0], item[1] + dir[1]}
			if p[0] < minY || p[1] < minX || p[0] > maxY || p[1] > maxX {
				continue
			}
			if _, ok := seen[p]; ok {
				continue
			}
			if ranges.Include(p) {
				continue
			}
			queue = append(queue, p)
		}
	}
	return count
}

func startingPoint(minY, minX, maxY, maxX int, ranges Ranges, pt [2]int) [2]int {
	for _, dir := range [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}} {
		p := [2]int{pt[0] + dir[0], pt[1] + dir[1]}
		if ranges.Include(p) {
			continue
		}
		if p[0] < minY || p[1] < minX || p[0] > maxY || p[1] > maxX {
			continue
		}
		return p
	}
	return pt
}
*/

func parseInput(input string, fn func(string) (string, int)) (int, int, int, int, Ranges, Ranges) {
	minY, minX := numbers.MaxInteger, numbers.MaxInteger
	maxY, maxX := 0, 0
	pos := [2]int{0, 0}
	lines := strings.Split(input, "\n")
	vertRanges, horiRanges := make(Ranges, 0), make(Ranges, 0)
	for _, line := range lines {
		dir, meter := fn(line)
		var newPos [2]int
		switch dir {
		case "U":
			newPos = [2]int{pos[0] - meter, pos[1]}
		case "R":
			newPos = [2]int{pos[0], pos[1] + meter}
		case "D":
			newPos = [2]int{pos[0] + meter, pos[1]}
		case "L":
			newPos = [2]int{pos[0], pos[1] - meter}
		}
		if dir == "U" || dir == "D" {
			vertRanges = append(vertRanges, Range{pos, newPos})
		}
		if dir == "R" || dir == "L" {
			horiRanges = append(horiRanges, Range{pos, newPos})
		}
		pos = newPos
		minY = numbers.Min(minY, pos[0])
		minX = numbers.Min(minX, pos[1])
		maxY = numbers.Max(maxY, pos[0])
		maxX = numbers.Max(maxX, pos[1])
	}
	return minY, minX, maxY, maxX, vertRanges, horiRanges
}

/*
#######
#.....#
###...#
..#...#
..#...#
###.###
#...#..
##..###
.#....#
.######
*/

func count(minY, minX, maxY, maxX int, vertRanges, horiRanges Ranges) int {
	var count int
	for y := minY; y <= maxY; y++ {
		var inside bool
		for x := minX; x <= maxX; x++ {
			if vertRanges.Include([2]int{y, x}) != -1 {
				inside = !inside

				/*
					if x-1 >= minX && vertRanges.Include([2]int{y, x - 1}) != -1 {
						inside = true
					}
				*/

				continue
			} else if idx := horiRanges.Include([2]int{y, x}); idx != -1 {
				x = numbers.Max(horiRanges[idx][0][1], horiRanges[idx][1][1])
				continue
			}

			if inside {
				count++
			}
		}
	}
	return count
}

func Part1(input string) (string, error) {
	minY, minX, maxY, maxX, vertRanges, horiRanges := parseInput(input, func(line string) (string, int) {
		fields := strings.Fields(line)
		meter := util.ParseInt(fields[1])
		return fields[0], meter
	})

	return strconv.Itoa(count(minY, minX, maxY, maxX, vertRanges, horiRanges) + vertRanges.Len() + horiRanges.Len()), nil
	/*
		start := startingPoint(minY, minX, maxY, maxX, ranges, [2]int{0, 0})
		return strconv.Itoa(count(minY, minX, maxY, maxX, ranges, start) + ranges.Len()), nil
	*/
}

func Part2(input string) (string, error) {
	minY, minX, maxY, maxX, vertRanges, horiRanges := parseInput(input, func(line string) (string, int) {
		dirMap := map[int]string{3: "U", 0: "R", 1: "D", 2: "L"}
		fields := strings.Fields(line)
		dir, _ := strconv.ParseInt(fields[2][7:8], 16, 64)
		meter, _ := strconv.ParseInt(fields[2][2:7], 16, 64)
		return dirMap[int(dir)], int(meter)
	})

	return strconv.Itoa(count(minY, minX, maxY, maxX, vertRanges, horiRanges)), nil
	/*
		start := startingPoint(minY, minX, maxY, maxX, ranges, [2]int{0, 0})
		return strconv.Itoa(count(minY, minX, maxY, maxX, ranges, start) + ranges.Len()), nil
	*/
}
