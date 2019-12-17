package day12

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToNum(t *testing.T) {
	require.Equal(t, 0, toNum("x=0"))
	require.Equal(t, 10, toNum("y=10"))
	require.Equal(t, -2, toNum("z=-2"))

	require.Equal(t, 0, toNum(" x=0"))
	require.Equal(t, 10, toNum(" y=10"))
	require.Equal(t, -2, toNum(" z=-2"))
}

func TestParseLine(t *testing.T) {
	x, y, z := parseLine("<x=-1, y=0, z=2>")
	require.Equal(t, -1, x)
	require.Equal(t, 0, y)
	require.Equal(t, 2, z)

	x, y, z = parseLine("<x=2, y=-10, z=-7>")
	require.Equal(t, 2, x)
	require.Equal(t, -10, y)
	require.Equal(t, -7, z)

	x, y, z = parseLine("<x=4, y=-8, z=8>")
	require.Equal(t, 4, x)
	require.Equal(t, -8, y)
	require.Equal(t, 8, z)

	x, y, z = parseLine("<x=3, y=5, z=-1>")
	require.Equal(t, 3, x)
	require.Equal(t, 5, y)
	require.Equal(t, -1, z)
}

func TestPoint(t *testing.T) {
	pt := &Point{X: -1, Y: 0, Z: 2}
	require.Equal(t, 3, pt.Energy())
	require.Equal(t, "<x= -1,y=  0,z=  2>", pt.String())

	pt = &Point{X: 2, Y: -10, Z: -7}
	require.Equal(t, 19, pt.Energy())
	require.Equal(t, "<x=  2,y=-10,z= -7>", pt.String())

	pt = &Point{X: 4, Y: -8, Z: 8}
	require.Equal(t, 20, pt.Energy())
	require.Equal(t, "<x=  4,y= -8,z=  8>", pt.String())

	pt = &Point{X: 3, Y: 5, Z: -1}
	require.Equal(t, 9, pt.Energy())
	require.Equal(t, "<x=  3,y=  5,z= -1>", pt.String())
}

func TestObject(t *testing.T) {
	obj := NewObject(-1, 0, 2)
	require.Equal(t, -1, obj.Pos.X)
	require.Equal(t, 0, obj.Pos.Y)
	require.Equal(t, 2, obj.Pos.Z)
	require.Equal(t, 0, obj.Vel.X)
	require.Equal(t, 0, obj.Vel.Y)
	require.Equal(t, 0, obj.Vel.Z)

	obj.Accelerate()
	require.Equal(t, -1, obj.Pos.X)
	require.Equal(t, 0, obj.Pos.Y)
	require.Equal(t, 2, obj.Pos.Z)
	require.Equal(t, 0, obj.Vel.X)
	require.Equal(t, 0, obj.Vel.Y)
	require.Equal(t, 0, obj.Vel.Z)

	obj.Vel.X = -1
	obj.Vel.Y = 2

	obj.Accelerate()
	require.Equal(t, -2, obj.Pos.X)
	require.Equal(t, 2, obj.Pos.Y)
	require.Equal(t, 2, obj.Pos.Z)
	require.Equal(t, -1, obj.Vel.X)
	require.Equal(t, 2, obj.Vel.Y)
	require.Equal(t, 0, obj.Vel.Z)

	require.Equal(t, 6, obj.PotentialEnergy())
	require.Equal(t, 3, obj.KineticEnergy())
	require.Equal(t, 18, obj.TotalEnergy())

	require.Equal(t, "pos=<x= -2,y=  2,z=  2>, vel=<x= -1,y=  2,z=  0>", obj.String())
}

func TestPart2(t *testing.T) {
	p := &Part2{}
	solution, err := p.Solve("<x=-1, y=0, z=2>\n<x=2, y=-10, z=-7>\n<x=4, y=-8, z=8>\n<x=3, y=5, z=-1>")
	require.Nil(t, err)
	require.Equal(t, "2772", solution)
	solution, err = p.Solve("<x=-8, y=-10, z=0>\n<x=5, y=5, z=10>\n<x=2, y=-7, z=3>\n<x=9, y=-8, z=-3>")
	require.Nil(t, err)
	require.Equal(t, "4686774924", solution)
}
