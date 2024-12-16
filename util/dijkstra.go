package util

import (
	"iter"
	"sort"
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
	queue := make([]*SearchNode[E], 1)
	queue[0] = &SearchNode[E]{Value: root}
	seen := make(map[E]struct{})

	for len(queue) > 0 {
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].Cost < queue[j].Cost
		})

		var node *SearchNode[E]
		node, queue = queue[0], queue[1:]

		if goalFn(node.Value) {
			return node
		}

		if _, ok := seen[node.Value]; ok {
			continue
		}

		seen[node.Value] = struct{}{}

		for neighbour, neighbourCost := range neighboursFn(node.Value) {
			if _, ok := seen[neighbour]; ok {
				continue
			}
			queue = append(queue, node.NewNodeWithCost(neighbour, neighbourCost))
		}
	}

	return nil
}

func AllPathsWithScore[E comparable](
	root E,
	neighboursFn NeighboursCostFunc[E],
	goalFn GoalFunc[E],
	expectedCost int,
) []*SearchNode[E] {
	queue := make([]*SearchNode[E], 1)
	queue[0] = &SearchNode[E]{Value: root}
	seen := make(map[E]int)
	paths := make([]*SearchNode[E], 0)

	for len(queue) > 0 {
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].Cost < queue[j].Cost
		})

		var node *SearchNode[E]
		node, queue = queue[0], queue[1:]

		if node.Cost > expectedCost {
			continue
		}

		if goalFn(node.Value) && node.Cost == expectedCost {
			paths = append(paths, node)
			continue
		}

		if cost, ok := seen[node.Value]; ok && cost < node.Cost {
			continue
		}

		seen[node.Value] = node.Cost

		for neighbour, neighbourCost := range neighboursFn(node.Value) {
			if _, ok := seen[neighbour]; ok {
				continue
			}
			queue = append(queue, node.NewNodeWithCost(neighbour, neighbourCost))
		}
	}

	return paths
}
