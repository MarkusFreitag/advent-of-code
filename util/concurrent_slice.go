package util

import (
	"sync"

	"github.com/MarkusFreitag/advent-of-code/util/slice"
)

type ConcurrentSlice[E any] struct {
	sync.RWMutex
	items []E
}

type ConcurrentSliceItem[E any] struct {
	Index int
	Value E
}

func NewConcurrentSlice[E any]() *ConcurrentSlice[E] {
	return &ConcurrentSlice[E]{items: make([]E, 0)}
}

func (cs *ConcurrentSlice[E]) Length() int {
	cs.Lock()
	defer cs.Unlock()

	return len(cs.items)
}

func (cs *ConcurrentSlice[E]) Items() []E {
	cs.Lock()
	defer cs.Unlock()

	return slice.Copy(cs.items)
}

func (cs *ConcurrentSlice[E]) Append(item E) {
	cs.Lock()
	defer cs.Unlock()

	cs.items = append(cs.items, item)
}

func (cs *ConcurrentSlice[E]) Iter() <-chan ConcurrentSliceItem[E] {
	c := make(chan ConcurrentSliceItem[E])

	f := func() {
		cs.Lock()
		defer cs.Unlock()
		for index, value := range cs.items {
			c <- ConcurrentSliceItem[E]{Index: index, Value: value}
		}
		close(c)
	}
	go f()

	return c
}
