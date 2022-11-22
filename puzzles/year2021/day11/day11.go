package day11

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/slice"
)

type Octopus struct {
	Y, X    int
	Energy  int
	Flashed bool
}

func (o *Octopus) Ready() bool {
	return o.Energy > 9 && !o.Flashed
}

func (o *Octopus) Incr() {
	o.Energy++
}

func (o *Octopus) Reset() {
	o.Energy = 0
	o.Flashed = false
}

func Part1(input string) (string, error) {
	lines := strings.Fields(input)
	grid := make([][]*Octopus, len(lines))
	for idx, line := range lines {
		row := make([]*Octopus, len(line))
		for i, num := range util.StringsToInts(util.StringToStrings(line)) {
			row[i] = &Octopus{Y: idx, X: i, Energy: num}
		}
		grid[idx] = row
	}

	var flashes int
	for i := 0; i < 100; i++ {
		flasher := make([]*Octopus, 0)
		for _, row := range grid {
			for _, octo := range row {
				octo.Incr()
				if octo.Ready() {
					flasher = append(flasher, octo)
				}
			}
		}

		for len(flasher) > 0 {
			for _, octo := range flasher {
				octo.Flashed = true
				for i := -1; i <= 1; i++ {
					y := octo.Y + i
					if y < 0 || y >= len(grid) {
						continue
					}
					for j := -1; j <= 1; j++ {
						x := octo.X + j
						if x < 0 || x >= len(grid[y]) {
							continue
						}
						grid[y][x].Incr()
					}
				}
			}

			flasher = make([]*Octopus, 0)
			for _, row := range grid {
				for _, octo := range row {
					if octo.Ready() {
						flasher = append(flasher, octo)
					}
				}
			}
		}

		for _, row := range grid {
			for _, octo := range row {
				if octo.Flashed {
					flashes++
					octo.Reset()
				}
			}
		}
	}

	return strconv.Itoa(flashes), nil
}

func Part2(input string) (string, error) {
	lines := strings.Fields(input)
	grid := make([][]*Octopus, len(lines))
	for idx, line := range lines {
		row := make([]*Octopus, len(line))
		for i, num := range util.StringsToInts(util.StringToStrings(line)) {
			row[i] = &Octopus{Y: idx, X: i, Energy: num}
		}
		grid[idx] = row
	}

	var steps int
	for ; true; steps++ {
		flasher := make([]*Octopus, 0)
		for _, row := range grid {
			for _, octo := range row {
				octo.Incr()
				if octo.Ready() {
					flasher = append(flasher, octo)
				}
			}
		}

		for len(flasher) > 0 {
			for _, octo := range flasher {
				octo.Flashed = true
				for i := -1; i <= 1; i++ {
					y := octo.Y + i
					if y < 0 || y >= len(grid) {
						continue
					}
					for j := -1; j <= 1; j++ {
						x := octo.X + j
						if x < 0 || x >= len(grid[0]) {
							continue
						}
						grid[y][x].Incr()
					}
				}
			}

			flasher = make([]*Octopus, 0)
			for _, row := range grid {
				for _, octo := range row {
					if octo.Ready() {
						flasher = append(flasher, octo)
					}
				}
			}
		}

		b := make([]bool, 0)
		for _, row := range grid {
			for _, octo := range row {
				b = append(b, octo.Flashed)
				if octo.Flashed {
					octo.Reset()
				}
			}
		}
		if slice.All(b, true) {
			break
		}
	}

	return strconv.Itoa(steps + 1), nil
}
