package util

import (
	"sync"

	"github.com/MarkusFreitag/advent-of-code/util/constraints"
)

type ConcurrentMap[K constraints.Comparable, V any] struct {
	sync.RWMutex
	items map[K]V
}

type ConcurrentMapItem[K constraints.Comparable, V any] struct {
	Key   K
	Value V
}

func NewConcurrentMap[K constraints.Comparable, V any]() *ConcurrentMap[K, V] {
	return &ConcurrentMap[K, V]{items: make(map[K]V)}
}

func (cm *ConcurrentMap[K, V]) Length() int {
	cm.Lock()
	defer cm.Unlock()

	return len(cm.items)
}

func (cm *ConcurrentMap[K, V]) Keys() []K {
	cm.Lock()
	defer cm.Unlock()

	return Keys(cm.items)
}

func (cm *ConcurrentMap[K, V]) Values() []V {
	cm.Lock()
	defer cm.Unlock()

	return Values(cm.items)
}

func (cm *ConcurrentMap[K, V]) Set(key K, value V) {
	cm.Lock()
	defer cm.Unlock()

	cm.items[key] = value
}

func (cm *ConcurrentMap[K, V]) Get(key K) (V, bool) {
	cm.Lock()
	defer cm.Unlock()

	value, ok := cm.items[key]

	return value, ok
}

func (cm *ConcurrentMap[K, V]) Iter() <-chan ConcurrentMapItem[K, V] {
	c := make(chan ConcurrentMapItem[K, V])

	f := func() {
		cm.Lock()
		defer cm.Unlock()

		for k, v := range cm.items {
			c <- ConcurrentMapItem[K, V]{Key: k, Value: v}
		}
		close(c)
	}
	go f()

	return c
}
