package util

import (
	"iter"
)

const (
	SearchStateOpened SearchState = iota
	SearchStateClosed
)

type SearchState int

func NoopPainter[E comparable]() PaintersFunc[E] {
	return func(_ E, _ SearchState) { return }
}

type PaintersFunc[E comparable] func(E, SearchState)

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

func Dijkstra[E comparable](root E, paintFn PaintersFunc[E], neighboursFn NeighboursCostFunc[E], goalFn GoalFunc[E]) *SearchNode[E] {
	queue := NewMinPriorityQueue[*SearchNode[E]]()
	queue.Set(&SearchNode[E]{Value: root}, 0)
	seen := make(map[E]struct{})

	for queue.Len() > 0 {
		node, cost := queue.Next()
		if paintFn != nil {
			paintFn(node.Value, SearchStateClosed)
		}

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
				if paintFn != nil {
					paintFn(node.Value, SearchStateOpened)
				}
				queue.Set(neighbourNode, newCost)
				continue
			}
		}
	}

	return nil
}
