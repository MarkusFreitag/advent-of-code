package day12

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func toNum(s string) int {
	s = strings.TrimSpace(s)
	num, _ := strconv.Atoi(s[2:])
	return num
}

func parseLine(line string) (int, int, int) {
	line = line[1 : len(line)-1]
	parts := strings.Split(line, ",")
	x := toNum(parts[0])
	y := toNum(parts[1])
	z := toNum(parts[2])
	return x, y, z
}

type Point struct {
	X, Y, Z int
}

func (p *Point) Energy() int {
	return int(math.Abs(float64(p.X)) + math.Abs(float64(p.Y)) + math.Abs(float64(p.Z)))
}

func (p *Point) String() string {
	return fmt.Sprintf("<x=%3d,y=%3d,z=%3d>", p.X, p.Y, p.Z)
}

type Object struct {
	Pos *Point
	Vel *Point
}

func NewObject(x, y, z int) *Object {
	return &Object{
		Pos: &Point{X: x, Y: y, Z: z},
		Vel: &Point{X: 0, Y: 0, Z: 0},
	}
}

func (o *Object) Accelerate() {
	o.Pos.X += o.Vel.X
	o.Pos.Y += o.Vel.Y
	o.Pos.Z += o.Vel.Z
}

func (o *Object) KineticEnergy() int {
	return o.Vel.Energy()
}

func (o *Object) PotentialEnergy() int {
	return o.Pos.Energy()
}

func (o *Object) TotalEnergy() int {
	return o.KineticEnergy() * o.PotentialEnergy()
}

func (o *Object) String() string {
	return fmt.Sprintf("pos=%s, vel=%s", o.Pos.String(), o.Vel.String())
}

type Universe []*Object

func (u Universe) Evolve() {
	for _, obj := range u {
		fmt.Println(obj.String())
	}
	fmt.Print("\n")
	// apply gravity
	for idx, obj := range u {
		for _, o := range u[idx+1:] {
			applyGravity(obj, o)
		}
	}
	// apply velocity
	for _, obj := range u {
		obj.Accelerate()
	}
}

func applyGravity(a, b *Object) {
	if a.Pos.X < b.Pos.X {
		a.Vel.X += 1
		b.Vel.X -= 1
	} else if a.Pos.X > b.Pos.X {
		a.Vel.X -= 1
		b.Vel.X += 1
	}
	if a.Pos.Y < b.Pos.Y {
		a.Vel.Y += 1
		b.Vel.Y -= 1
	} else if a.Pos.Y > b.Pos.Y {
		a.Vel.Y -= 1
		b.Vel.Y += 1
	}
	if a.Pos.Z < b.Pos.Z {
		a.Vel.Z += 1
		b.Vel.Z -= 1
	} else if a.Pos.Z > b.Pos.Z {
		a.Vel.Z -= 1
		b.Vel.Z += 1
	}
}

type Part1 struct{}

func (p *Part1) Solve(input string) (string, error) {
	lines := strings.Split(input, "\n")
	universe := make(Universe, len(lines))
	for idx, line := range lines {
		universe[idx] = NewObject(parseLine(line))
	}
	for i := 0; i < 1000; i++ {
		fmt.Printf("After %d steps:\n", i)
		universe.Evolve()
	}
	var energy int
	for _, obj := range universe {
		fmt.Printf("pot: %3d + %3d + %3d = %4d", obj.Pos.X, obj.Pos.Y, obj.Pos.Z, obj.PotentialEnergy())
		fmt.Printf("\tkin: %3d + %3d + %3d = %4d", obj.Vel.X, obj.Vel.Y, obj.Vel.Z, obj.KineticEnergy())
		fmt.Printf("\ttotal: %4d * %4d = %4d\n", obj.PotentialEnergy(), obj.KineticEnergy(), obj.TotalEnergy())
		energy += obj.TotalEnergy()
	}
	return strconv.Itoa(energy), nil
}

type Part2 struct{}

func (p *Part2) Solve(input string) (string, error) {
	return "n/a", nil
}
