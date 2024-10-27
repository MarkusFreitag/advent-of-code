package directions

import "iter"

const (
	NorthEast Intercardinal = iota
	SouthEast
	SouthWest
	NorthWest
)

var (
	intercardinalWordStrings  = [4]string{"NORTHEAST", "SOUTHEAST", "SOUTHWEST", "NORTHWEST"}
	intercardinalShortStrings = [4]string{"NE", "SE", "SW", "NW"}
)

type Intercardinal int

func (c Intercardinal) String() string {
	return c.Word()
}

func (c Intercardinal) Short() string {
	return intercardinalShortStrings[c]
}

func (c Intercardinal) Word() string {
	return intercardinalWordStrings[c]
}

func Intercardinals() iter.Seq[Intercardinal]                 { return foursomeSeq(NorthEast, 1) }
func IntercardinalsCounterClockwise() iter.Seq[Intercardinal] { return foursomeSeq(NorthEast, 3) }

func IntercardinalsFrom(from Intercardinal) iter.Seq[Intercardinal] { return foursomeSeq(from, 1) }
func IntercardinalsFromCounterClockwise(from Intercardinal) iter.Seq[Intercardinal] {
	return foursomeSeq(from, 3)
}
