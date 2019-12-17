package day12

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func abs(i int) int {
  return int(math.Abs(float64(i)))
}

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
	return abs(p.X) + abs(p.Y) + abs(p.Z)
}

func (p *Point) String() string {
	return fmt.Sprintf("<x=%3d,y=%3d,z=%3d>", p.X, p.Y, p.Z)
}

func (p *Point) State() string {
	return fmt.Sprintf("%d%d%d", p.X, p.Y, p.Z)
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

func (o *Object) State() string {
	return o.Pos.State()+o.Vel.State()
}

type Universe []*Object

func (u Universe) Evolve() {
  /*
  for _, obj := range u {
    fmt.Println(obj.String())
  }
  fmt.Print("\n")
  */
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

func (u Universe) State() string {
	//return fmt.Sprintf("%s|%s|%s|%s", u[0].String(), u[1].String(), u[2].String(), u[3].String())
	return u[0].State()+u[1].State()+u[2].State()+u[3].State()
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
		//fmt.Printf("After %d steps:\n", i)
		universe.Evolve()
	}
	var energy int
	for _, obj := range universe {
		//fmt.Printf("pot: %3d + %3d + %3d = %4d", obj.Pos.X, obj.Pos.Y, obj.Pos.Z, obj.PotentialEnergy())
		//fmt.Printf("\tkin: %3d + %3d + %3d = %4d", obj.Vel.X, obj.Vel.Y, obj.Vel.Z, obj.KineticEnergy())
		//fmt.Printf("\ttotal: %4d * %4d = %4d\n", obj.PotentialEnergy(), obj.KineticEnergy(), obj.TotalEnergy())
		energy += obj.TotalEnergy()
	}
	return strconv.Itoa(energy), nil
}

type Part2 struct{}

func (p *Part2) Solve(input string) (string, error) {
	//input = "<x=-1, y=0, z=2>\n<x=2, y=-10, z=-7>\n<x=4, y=-8, z=8>\n<x=3, y=5, z=-1>"
  //input = "<x=-8, y=-10, z=0>\n<x=5, y=5, z=10>\n<x=2, y=-7, z=3>\n<x=9, y=-8, z=-3>"
	lines := strings.Split(input, "\n")
	universe := make(Universe, len(lines))
  buf := make(Universe, len(lines))
	for idx, line := range lines {
		universe[idx] = NewObject(parseLine(line))
    buf[idx] = NewObject(parseLine(line))
	}

  var steps int
  var x,y,z int
  for steps=1;x==0||y==0||z==0;steps++{
    if steps%1000000 == 0 {
      fmt.Printf("steps: %d million\n", steps/1000000)
    }
		universe.Evolve()

    if x==0 {
      ok := true
      for idx, moon := range universe {
        if moon.Pos.X != buf[idx].Pos.X || moon.Vel.X != buf[idx].Vel.X {
          ok = false
          break
        }
      }
      if ok {
        x = steps
      }
    }
    if y==0 {
      ok := true
      for idx, moon := range universe {
        if moon.Pos.Y != buf[idx].Pos.Y || moon.Vel.Y != buf[idx].Vel.Y {
          ok = false
          break
        }
      }
      if ok {
        y = steps
      }
    }
    if z==0 {
      ok := true
      for idx, moon := range universe {
        if moon.Pos.Z != buf[idx].Pos.Z || moon.Vel.Z != buf[idx].Vel.Z {
          ok = false
          break
        }
      }
      if ok {
        z = steps
      }
    }
  }
  fmt.Printf("x: %d, y: %d, z: %d\n", x, y, z)
  return "n/a", nil

  /*
  var steps int
  for {
    if steps%1000000 == 0 {
      fmt.Printf("steps: %d million\n", steps/1000000)
    }
		universe.Evolve()
		steps++
    if "000000000000" == fmt.Sprintf(
      "%d%d%d%d%d%d%d%d%d%d%d%d",
      universe[0].Vel.X, universe[0].Vel.Y, universe[0].Vel.Z,
      universe[1].Vel.X, universe[1].Vel.Y, universe[1].Vel.Z,
      universe[2].Vel.X, universe[2].Vel.Y, universe[2].Vel.Z,
      universe[3].Vel.X, universe[3].Vel.Y, universe[3].Vel.Z,
    ) {
      return strconv.Itoa(steps*2), nil
		}
  }
  */

  /*
  // runs OOM at about 14 million steps
  states := make(map[string]int)
	var steps int
	for {
    if steps%1000000 == 0 {
      fmt.Printf("steps: %d\n", steps)
    }
    states[universe.State()] = 0
		universe.Evolve()
		steps++
    if _, ok := states[universe.State()]; ok {
      return strconv.Itoa(steps), nil
		}
	}

  // unbelievable slow
  states := make([]string, 0)
	var steps int
	for {
    if steps%1000 == 0 {
      fmt.Printf("steps: %d\n", steps)
    }
    states = append(states, universe.State())
		universe.Evolve()
		steps++
    for _, state := range states {
      if state == universe.State() {
        return strconv.Itoa(steps), nil
      }
		}
	}
  */
}
