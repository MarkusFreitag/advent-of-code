package util

import (
	"fmt"
	"iter"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util/iterutil"
	"github.com/MarkusFreitag/advent-of-code/util/sliceutil"
)

type SearchNode[E comparable] struct {
	Parent *SearchNode[E]
	Value  E
}

func (sn *SearchNode[E]) NewNode(value E) *SearchNode[E] {
	return &SearchNode[E]{Parent: sn, Value: value}
}

func (sn *SearchNode[E]) Dist() int {
	return iterutil.Count(sn.Seq()) - 1
}

func (sn *SearchNode[E]) Seq() iter.Seq[E] {
	return func(yield func(E) bool) {
		current := sn
		for {
			if current == nil {
				return
			}
			if !yield(current.Value) {
				return
			}
			if current.Parent == nil {
				return
			}
			current = current.Parent
		}
	}
}

type NeighboursFunc[E comparable] func(E) iter.Seq[E]
type GoalFunc[E comparable] func(E) bool

func BFS[E comparable](root E, neighboursFn NeighboursFunc[E], goalFn GoalFunc[E]) *SearchNode[E] {
	queue := make([]*SearchNode[E], 1)
	queue[0] = &SearchNode[E]{Value: root}
	seen := make(map[E]struct{})
	for len(queue) > 0 {
		var sn *SearchNode[E]
		sn, queue = sliceutil.PopFront(queue)
		if goalFn(sn.Value) {
			return sn
		}
		for neighbour := range neighboursFn(sn.Value) {
			if _, ok := seen[neighbour]; !ok {
				seen[neighbour] = struct{}{}
				queue = append(queue, sn.NewNode(neighbour))
			}
		}
	}
	return nil
}

func AllPathsBFS[E comparable](root E, neighboursFn NeighboursFunc[E], goalFn GoalFunc[E]) map[string]*SearchNode[E] {
	paths := make(map[string]*SearchNode[E])
	queue := make([]*SearchNode[E], 1)
	queue[0] = &SearchNode[E]{Value: root}
	for len(queue) > 0 {
		var sn *SearchNode[E]
		sn, queue = sliceutil.PopFront(queue)
		if goalFn(sn.Value) {
			paths[pathKey(sn)] = sn
			continue
		}
		for neighbour := range neighboursFn(sn.Value) {
			queue = append(queue, sn.NewNode(neighbour))
		}
	}
	return paths
}

func pathKey[E comparable](sn *SearchNode[E]) string {
	strs := make([]string, 0)
	for p := range sn.Seq() {
		strs = append(strs, fmt.Sprintf("%v", p))
	}
	return strings.Join(strs, "|")
}

func DFS[E comparable](root E, neighboursFn NeighboursFunc[E], goalFn GoalFunc[E]) *SearchNode[E] {
	seen := make(map[E]struct{})
	return dfs(&SearchNode[E]{Value: root}, seen, neighboursFn, goalFn)
}

func dfs[E comparable](root *SearchNode[E], seen map[E]struct{}, neighboursFn NeighboursFunc[E], goalFn GoalFunc[E]) *SearchNode[E] {
	seen[root.Value] = struct{}{}
	for neighbour := range neighboursFn(root.Value) {
		if _, ok := seen[neighbour]; !ok {
			sn := root.NewNode(neighbour)
			if goalFn(neighbour) {
				return sn
			}
			return dfs(sn, seen, neighboursFn, goalFn)
		}
	}
	return nil
}
