package directions

import (
	"image"
	"iter"
)

const (
	North Cardinal = iota
	NorthEast
	East
	SouthEast
	South
	SouthWest
	West
	NorthWest
)

var (
	cardinalWordStrings  = [8]string{"NORTH", "NORTHEAST", "EAST", "SOUTHEAST", "SOUTH", "SOUTHWEST", "WEST", "NORTHWEST"}
	cardinalShortStrings = [8]string{"N", "NE", "E", "SE", "S", "SW", "W", "NW"}
)

type Cardinal int

func (c Cardinal) String() string {
	return c.Word()
}

func (c Cardinal) Short() string {
	return cardinalShortStrings[c]
}

func (c Cardinal) Word() string {
	return cardinalWordStrings[c]
}

func (c Cardinal) Point() image.Point {
	return directionPoints[c]
}

func Cardinals(opts ...Option) iter.Seq[Cardinal] {
	return dirSeq(North, opts...)
}

func CardinalsFrom(from Cardinal, opts ...Option) iter.Seq[Cardinal] {
	return dirSeq(from, opts...)
}
