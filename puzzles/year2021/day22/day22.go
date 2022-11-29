package day22

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

type Vector3 struct {
	X, Y, Z int
}

func (v Vector3) String() string {
	return fmt.Sprintf("x=%d,y=%d,z=%d", v.X, v.Y, v.Z)
}

type Range struct {
	From, To Vector3
}

func parseRange(x, y, z string) Range {
	var from, to Vector3

	parts := strings.Split(x, "..")
	from.X, to.X = util.ParseInt(parts[0]), util.ParseInt(parts[1])

	parts = strings.Split(y, "..")
	from.Y, to.Y = util.ParseInt(parts[0]), util.ParseInt(parts[1])

	parts = strings.Split(z, "..")
	from.Z, to.Z = util.ParseInt(parts[0]), util.ParseInt(parts[1])

	return Range{From: from, To: to}
}

func (r Range) String() string {
	return fmt.Sprintf("%s..%s", r.From, r.To)
}

func (r Range) Count() int {
	return numbers.Abs((r.To.X - r.From.X + 1) * (r.To.Y - r.From.Y + 1) * (r.To.Z - r.From.Z + 1))
}

func between(a, b, c int) bool {
	return a <= b && b <= c
}

func (r Range) Intersect(other Range) (*Range, bool) {
	var invalid bool
	if other.To.X < r.From.X || r.To.X < other.From.X {
		invalid = true
	}
	if other.To.Y < r.From.Y || r.To.Y < other.From.Y {
		invalid = true
	}
	if other.To.Z < r.From.Z || r.To.Z < other.From.Z {
		invalid = true
	}
	if invalid {
		return nil, false
	}

	var from Vector3
	if between(r.From.X, other.From.X, r.To.X) {
		from.X = other.From.X
	} else {
		from.X = r.From.X
	}
	if between(r.From.Y, other.From.Y, r.To.Y) {
		from.Y = other.From.Y
	} else {
		from.Y = r.From.Y
	}
	if between(r.From.Z, other.From.Z, r.To.Z) {
		from.Z = other.From.Z
	} else {
		from.Z = r.From.Z
	}

	var to Vector3
	if between(r.From.X, other.To.X, r.To.X) {
		to.X = other.To.X
	} else {
		to.X = r.To.X
	}
	if between(r.From.Y, other.To.Y, r.To.Y) {
		to.Y = other.To.Y
	} else {
		to.Y = r.To.Y
	}
	if between(r.From.Z, other.To.Z, r.To.Z) {
		to.Z = other.To.Z
	} else {
		to.Z = r.To.Z
	}

	return &Range{
		From: from,
		To:   to,
	}, true
}

type Op struct {
	R     Range
	State bool
}

func parseOp(str string) Op {
	rgx := regexp.MustCompile(`^(on|off)\sx=(-?\d+\.\.-?\d+),y=(-?\d+\.\.-?\d+),z=(-?\d+\.\.-?\d+)$`)
	matches := rgx.FindAllStringSubmatch(str, -1)
	op := Op{
		R: parseRange(matches[0][2], matches[0][3], matches[0][4]),
	}
	if matches[0][1] == "on" {
		op.State = true
	}
	return op
}

type Cube struct {
	Range Range
	Holes []*Cube
}

func NewCube(r Range) *Cube {
	return &Cube{
		Range: r,
		Holes: make([]*Cube, 0),
	}
}

func (c *Cube) Remove(other Range) {
	if intersectRange, ok := c.Range.Intersect(other); ok {
		for _, hole := range c.Holes {
			hole.Remove(*intersectRange)
		}
		c.Holes = append(c.Holes, NewCube(*intersectRange))
	}
}

func (c *Cube) Volume() int {
	var holeSum int
	for _, hole := range c.Holes {
		holeSum += hole.Volume()
	}
	return c.Range.Count() - holeSum
}

func Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")
	ops := make([]Op, len(lines))
	for idx, line := range lines {
		ops[idx] = parseOp(line)
	}

	restricted := Range{
		From: Vector3{X: -50, Y: -50, Z: -50},
		To:   Vector3{X: 50, Y: 50, Z: 50},
	}
	cubes := make([]*Cube, 0)
	for _, op := range ops {
		if _, ok := op.R.Intersect(restricted); !ok {
			continue
		}

		for _, cube := range cubes {
			cube.Remove(op.R)
		}
		if op.State {
			cubes = append(cubes, NewCube(op.R))
		}
	}

	var total int
	for _, cube := range cubes {
		total += cube.Volume()
	}

	return strconv.Itoa(total), nil
}

func Part2(input string) (string, error) {
	lines := strings.Split(input, "\n")
	ops := make([]Op, len(lines))
	for idx, line := range lines {
		ops[idx] = parseOp(line)
	}

	cubes := make([]*Cube, 0)
	for _, op := range ops {
		for _, cube := range cubes {
			cube.Remove(op.R)
		}
		if op.State {
			cubes = append(cubes, NewCube(op.R))
		}
	}

	var total int
	for _, cube := range cubes {
		total += cube.Volume()
	}

	return strconv.Itoa(total), nil
}
