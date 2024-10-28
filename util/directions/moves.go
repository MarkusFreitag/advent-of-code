package directions

import (
	"image"
	"iter"
)

const (
	Up Move = iota
	UpRight
	Right
	DownRight
	Down
	DownLeft
	Left
	UpLeft
)

var (
	moveWordStrings  = [8]string{"UP", "UPRIGHT", "RIGHT", "DOWNRIGHT", "DOWN", "DOWNLEFT", "LEFT", "UPLEFT"}
	moveShortStrings = [8]string{"U", "UR", "R", "DR", "D", "DL", "L", "UL"}
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

func (m Move) Point() image.Point {
	return directionPoints[m]
}

func Moves(opts ...Option) iter.Seq[Move] {
	return dirSeq(Up, opts...)
}

func MovesFrom(from Move, opts ...Option) iter.Seq[Move] {
	return dirSeq(from, opts...)
}
