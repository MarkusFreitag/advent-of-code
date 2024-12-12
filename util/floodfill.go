package util

import "github.com/MarkusFreitag/advent-of-code/util/sliceutil"

func Floodfill[E comparable](root E, neighboursFn NeighboursFunc[E]) []E {
	elements := make([]E, 0)
	queue := make([]E, 1)
	queue[0] = root
	seen := make(map[E]struct{})
	seen[root] = struct{}{}
	for len(queue) > 0 {
		var elem E
		elem, queue = sliceutil.PopFront(queue)
		elements = append(elements, elem)
		for neighbour := range neighboursFn(elem) {
			if _, ok := seen[neighbour]; !ok {
				seen[neighbour] = struct{}{}
				queue = append(queue, neighbour)
			}
		}
	}
	return elements
}
