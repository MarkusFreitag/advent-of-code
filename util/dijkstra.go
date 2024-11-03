package util

import (
	"iter"
)

type NeighboursCostFunc[E comparable] func(E) iter.Seq2[E, int]

func FakeCost[E comparable](fn NeighboursFunc[E]) NeighboursCostFunc[E] {
	return func(p E) iter.Seq2[E, int] {
		return func(yield func(E, int) bool) {
			for neighbour := range fn(p) {
				if !yield(neighbour, 1) {
					return
				}
			}
		}
	}
}

func Dijkstra[E comparable](root E, neighboursFn NeighboursCostFunc[E], goalFn GoalFunc[E]) *SearchNode[E] {
	queue := NewMinPriorityQueue[*SearchNode[E]]()
	queue.Set(&SearchNode[E]{Value: root}, 0)
	seen := make(map[E]struct{})

	for queue.Len() > 0 {
		node, cost := queue.Next()

		if goalFn(node.Value) {
			return node
		}

		seen[node.Value] = struct{}{}

		for neighbour, neighbourCost := range neighboursFn(node.Value) {
			if _, ok := seen[neighbour]; ok {
				continue
			}

			neighbourNode := node.NewNode(neighbour)

			newCost := cost + neighbourCost
			queuedCost, ok := queue.Get(neighbourNode)
			if !ok || newCost < queuedCost {
				queue.Set(neighbourNode, newCost)
				continue
			}
		}
	}

	return nil
}
