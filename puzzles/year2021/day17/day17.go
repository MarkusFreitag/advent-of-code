package day17

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

type Point struct {
	X, Y int
}

func (p Point) String() string { return fmt.Sprintf("x=%d|y=%d", p.X, p.Y) }

func (p Point) Within(a *Area) bool {
	return a.A.X <= p.X && p.X <= a.B.X && a.A.Y >= p.Y && p.Y >= a.B.Y
}

type Probe struct {
	Point
	Vel Point
}

func NewProbe(vel Point) *Probe {
	return &Probe{
		Vel: vel,
	}
}

func (p *Probe) String() string { return fmt.Sprintf("probe: x=%d|y=%d vel[%s]", p.X, p.Y, p.Vel) }

func (p *Probe) Move() {
	p.X += p.Vel.X
	p.Y += p.Vel.Y

	if p.Vel.X > 0 {
		p.Vel.X--
	} else if p.Vel.X < 0 {
		p.Vel.X++
	}

	p.Vel.Y--
}

type Area struct {
	A, B Point
}

func NewArea(str string) *Area {
	var area Area
	fields := strings.Fields(str)

	field := strings.TrimSuffix(fields[2], ",")
	parts := strings.Split(field, "=")
	nums := util.StringsToInts(strings.Split(parts[1], ".."))
	area.A.X, area.B.X = numbers.MinMax(nums[0], nums[1])

	field = strings.TrimSuffix(fields[3], ",")
	parts = strings.Split(field, "=")
	nums = util.StringsToInts(strings.Split(parts[1], ".."))
	area.B.Y, area.A.Y = numbers.MinMax(nums[0], nums[1])

	return &area
}

func (a *Area) String() string { return fmt.Sprintf("area: a[%s] => b[%s]", a.A, a.B) }

func simulateProbes(area *Area) map[string]int {
	probes := make(map[string]int)
	for x := 1; x <= area.B.X; x++ {
		for y := area.B.Y; y <= 1000; y++ {
			initVel := Point{X: x, Y: y}
			probe := NewProbe(initVel)
			var reachedArea bool
			var highest int
			for {
				probe.Move()

				if probe.Within(area) {
					reachedArea = true
				}
				if probe.X > area.B.X || probe.Y < area.B.Y {
					break
				}
				if probe.Y > highest {
					highest = probe.Y
				}
			}

			if reachedArea {
				probes[initVel.String()] = highest
			}
		}
	}
	return probes
}

var probes map[string]int

func Part1(input string) (string, error) {
	if probes == nil {
		area := NewArea(input)
		probes = simulateProbes(area)
	}

	var highest int
	for _, probe := range probes {
		if probe > highest {
			highest = probe
		}
	}
	return strconv.Itoa(highest), nil
}

func Part2(input string) (string, error) {
	if probes == nil {
		area := NewArea(input)
		probes = simulateProbes(area)
	}

	return strconv.Itoa(len(probes)), nil
}
