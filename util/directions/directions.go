package directions

import (
	"image"
	"iter"
)

var directionPoints = [8]image.Point{
	image.Point{X: 0, Y: 1},
	image.Point{X: 1, Y: 1},
	image.Point{X: 1, Y: 0},
	image.Point{X: 1, Y: -1},
	image.Point{X: 0, Y: -1},
	image.Point{X: -1, Y: -1},
	image.Point{X: -1, Y: 0},
	image.Point{X: -1, Y: 1},
}

type dirConstraint interface {
	Move | Cardinal
}

type options struct {
	counterClockwise bool
	intermediate     bool
	infinite         bool
}

type Option func(*options)

func WithIntermediate() Option {
	return func(o *options) {
		o.intermediate = true
	}
}
func WithCounterClockwise() Option {
	return func(o *options) {
		o.counterClockwise = true
	}
}
func WithInfinite() Option {
	return func(o *options) {
		o.infinite = true
	}
}

func dirSeq[T dirConstraint](start T, opts ...Option) iter.Seq[T] {
	o := &options{
		counterClockwise: false,
		intermediate:     false,
	}
	for _, opt := range opts {
		opt(o)
	}

	var offset int
	if o.intermediate {
		if o.counterClockwise {
			offset = 7
		} else {
			offset = 1
		}
	} else {
		if o.counterClockwise {
			offset = 6
		} else {
			offset = 2
		}
	}

	return func(yield func(T) bool) {
		current := start
		for {
			if !yield(current) {
				return
			}
			current = T((int(current) + offset) % 8)
			if current == start && !o.infinite {
				return
			}
		}
	}
}
