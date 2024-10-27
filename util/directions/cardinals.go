package directions

import "iter"

const (
	North Cardinal = iota
	East
	South
	West
)

var (
	cardinalWordStrings  = [4]string{"NORTH", "EAST", "SOUTH", "WEST"}
	cardinalShortStrings = [4]string{"N", "E", "S", "W"}
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

func Cardinals() iter.Seq[Cardinal]                 { return foursomeSeq(North, 1) }
func CardinalsCounterClockwise() iter.Seq[Cardinal] { return foursomeSeq(North, 3) }

func CardinalsFrom(from Cardinal) iter.Seq[Cardinal]                 { return foursomeSeq(from, 1) }
func CardinalsFromCounterClockwise(from Cardinal) iter.Seq[Cardinal] { return foursomeSeq(from, 3) }
