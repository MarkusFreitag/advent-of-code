package directions

import "iter"

const (
	Up Move = iota
	Right
	Down
	Left
)

var (
	moveWordStrings  = [4]string{"UP", "RIGHT", "DOWN", "LEFT"}
	moveShortStrings = [4]string{"^", ">", "v", "<"}
)

type Move int

func (m Move) String() string {
	return m.Word()
}

func (m Move) Short() string {
	return moveShortStrings[m]
}

func (m Move) Word() string {
	return moveWordStrings[m]
}

func Moves() iter.Seq[Move]                 { return foursomeSeq(Up, 1) }
func MovesCounterClockwise() iter.Seq[Move] { return foursomeSeq(Up, 3) }

func MovesFrom(from Move) iter.Seq[Move]                 { return foursomeSeq(from, 1) }
func MovesFromCounterClockwise(from Move) iter.Seq[Move] { return foursomeSeq(from, 3) }
