package sliceutil

import (
	"iter"
	"slices"
	"sync"
)

type ConcurrentSlice[S ~[]E, E any] struct {
	sync.RWMutex
	items S
}

func NewConcurrentSlice[S ~[]E, E any]() *ConcurrentSlice[S, E] {
	return &ConcurrentSlice[S, E]{items: make(S, 0)}
}

func (cs *ConcurrentSlice[S, E]) Length() int {
	cs.Lock()
	defer cs.Unlock()

	return len(cs.items)
}

func (cs *ConcurrentSlice[S, E]) Items() S {
	cs.Lock()
	defer cs.Unlock()

	return slices.Clone(cs.items)
}

func (cs *ConcurrentSlice[S, E]) Append(item E) {
	cs.Lock()
	defer cs.Unlock()

	cs.items = append(cs.items, item)
}

func (cs *ConcurrentSlice[S, E]) Delete(index int) {
	cs.Lock()
	defer cs.Unlock()

	var defaultValue E
	cs.items[index] = defaultValue
}

func (cs *ConcurrentSlice[S, E]) Set(index int, item E) {
	cs.Lock()
	defer cs.Unlock()

	cs.items[index] = item
}

func (cs *ConcurrentSlice[S, E]) Get(index int) E {
	cs.Lock()
	defer cs.Unlock()

	return cs.items[index]
}

func (cs *ConcurrentSlice[S, E]) Seq() iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		cs.Lock()
		defer cs.Unlock()
		for index, value := range cs.items {
			if !yield(index, value) {
				return
			}
		}
	}
}
