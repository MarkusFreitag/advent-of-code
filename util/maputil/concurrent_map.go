package maputil

import (
	"iter"
	"sync"
)

type ConcurrentMap[K comparable, V any] struct {
	sync.RWMutex
	items map[K]V
}

func NewConcurrentMap[K comparable, V any]() *ConcurrentMap[K, V] {
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

func (cm *ConcurrentMap[K, V]) Delete(key K) {
	cm.Lock()
	defer cm.Unlock()

	delete(cm.items, key)
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

func (cm *ConcurrentMap[K, V]) Seq() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		cm.Lock()
		defer cm.Unlock()
		for key, value := range cm.items {
			if !yield(key, value) {
				return
			}
		}
	}
}
