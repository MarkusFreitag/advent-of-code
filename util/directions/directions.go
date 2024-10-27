package directions

import "iter"

func foursomeSeq[T ~int](start T, offset int) iter.Seq[T] {
	return func(yield func(T) bool) {
		current := start
		for {
			if !yield(current) {
				return
			}
			// clockwise (current + 1) % 4
			// counter-clockwise (current + 3) % 4
			current = T((int(current) + offset) % 4)
			if current == start {
				return
			}
		}
	}
}
