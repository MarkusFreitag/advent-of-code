package day14

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

var cycleCache = make(map[[2]string]string)

type point [2]int

func pt(r, c int) point {
	return point{r, c}
}

func (p point) add(other point) point {
	return pt(p[0]+other[0], p[1]+other[1])
}

func tilt(input string, dir point) string {
	m, size := stringToMap(input)
	moved := true
	for moved {
		moved = false
		for r := 0; r < size[0]; r++ {
			if r+dir[0] < 0 || r+dir[0] >= size[0] {
				continue
			}
			for c := 0; c < size[1]; c++ {
				if c+dir[1] < 0 || c+dir[1] >= size[1] {
					continue
				}
				p := pt(r, c)
				if m[p] != 'O' {
					continue
				}
				if np := p.add(dir); m[np] == '.' {
					moved = true
					m[np] = 'O'
					m[p] = '.'
				}
			}
		}

	}
	return mapToString(m, size)
}

func mapToString(m map[point]rune, size point) string {
	rows := make([]string, size[0])
	for r := 0; r < size[0]; r++ {
		row := make([]rune, size[1])
		for c := 0; c < size[1]; c++ {
			row[c] = m[pt(r, c)]
		}
		rows[r] = string(row)
	}
	return strings.Join(rows, "\n")
}

func stringToMap(input string) (map[point]rune, point) {
	m := make(map[point]rune)
	var p point
	for r, row := range strings.Split(input, "\n") {
		p[0] = numbers.Max(p[0], r+1)
		for c, cell := range row {
			p[1] = numbers.Max(p[1], c+1)
			m[pt(r, c)] = rune(cell)
		}
	}
	return m, p
}

func calcLoad(input string) int {
	lines := strings.Fields(input)
	var total int
	for idx, row := range lines {
		total += strings.Count(string(row), "O") * (len(lines) - idx)
	}
	return total
}

func cycle(input string) string {
	for _, dir := range []point{pt(-1, 0), pt(0, -1), pt(1, 0), pt(0, 1)} {
		key := [2]string{input, fmt.Sprintf("%v", dir)}
		if v, ok := cycleCache[key]; ok {
			input = v
		} else {
			input = tilt(input, dir)
			cycleCache[key] = input
		}
	}
	return input
}

func Part1(input string) (string, error) {
	for {
		tmp := tilt(input, pt(-1, 0))
		if input == tmp {
			break
		}
		input = tmp
	}
	return strconv.Itoa(calcLoad(input)), nil
}

func Part2(input string) (string, error) {
	cache := make(map[string]int)

	maxCycle := 1000000000
	for i := 0; i < maxCycle; i++ {
		if prev, ok := cache[input]; ok {
			length := i - prev
			target := prev + ((maxCycle - prev) % length)

			for k, v := range cache {
				if v == target {
					input = k
				}
			}
			break
		}

		cache[input] = i
		input = cycle(input)
	}

	return strconv.Itoa(calcLoad(input)), nil
}
