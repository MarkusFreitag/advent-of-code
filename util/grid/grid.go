package grid

import (
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

type Dimension interface {
	Coord2 | Coord3
}

type Coord[D Dimension] interface {
	Add(Coord[D]) Coord[D]
	Min(Coord[D]) Coord[D]
	Max(Coord[D]) Coord[D]
	Between(Coord[D], Coord[D]) bool
	DirectNeighbours() []Coord[D]
}

type Coord2 struct {
	X, Y int
}

func (c Coord2) Add(other Coord2) Coord2 {
	return Coord2{
		X: c.X + other.X,
		Y: c.Y + other.Y,
	}
}

func (c Coord2) Min(other Coord2) Coord2 {
	return Coord2{
		X: numbers.Min(c.X, other.X),
		Y: numbers.Min(c.Y, other.Y),
	}
}

func (c Coord2) Max(other Coord2) Coord2 {
	return Coord2{
		X: numbers.Max(c.X, other.X),
		Y: numbers.Max(c.Y, other.Y),
	}
}

func (c Coord2) Between(a, b Coord2) bool {
	if c.X < a.X || c.X > b.X {
		return false
	}

	if c.Y < a.Y || c.Y > b.Y {
		return false
	}

	return true
}

func (c Coord2) DirectNeighbours() []Coord2 {
	return []Coord2{
		Coord2{X: 1, Y: 0},
		Coord2{X: 0, Y: 1},
		Coord2{X: -1, Y: 0},
		Coord2{X: 0, Y: -1},
	}
}

type Coord3 struct {
	X, Y, Z int
}

func (c Coord3) Add(other Coord3) Coord3 {
	return Coord3{
		X: c.X + other.X,
		Y: c.Y + other.Y,
		Z: c.Z + other.Z,
	}
}

func (c Coord3) Min(other Coord3) Coord3 {
	return Coord3{
		X: numbers.Min(c.X, other.X),
		Y: numbers.Min(c.Y, other.Y),
		Z: numbers.Min(c.Z, other.Z),
	}
}

func (c Coord3) Max(other Coord3) Coord3 {
	return Coord3{
		X: numbers.Max(c.X, other.X),
		Y: numbers.Max(c.Y, other.Y),
		Z: numbers.Max(c.Z, other.Z),
	}
}

func (c Coord3) Between(a, b Coord3) bool {
	return true
}

func (c Coord3) DirectNeighbours() []Coord3 {
	return []Coord3{}
}

type Cell[C Coord[D], D Dimension, T any] struct {
	Coord C
	Value T
}

type IterateFn[C Coord[D], D Dimension, T any] func(C, T) bool

type Grid[C Coord[D], D Dimension, T any] interface {
	//Distance(C, C) int
	Clone() Grid[C, D, T]
	Neighbours(C) []C
	Expand(C)
	Set(C, T)
	Get(C) T
	Lookup(C) (T, bool)
	Count(T) int
	//Equal(Grid[C, D, T]) bool
	IterateChangedCells(IterateFn[C, D, T])
	Iterate(IterateFn[C, D, T])
	IterateBackwards(IterateFn[C, D, T])
}
