package util

import "sort"

type PriorityQueue[E comparable] struct {
	items      []E
	priorities map[E]int
	maxPrio    bool
}

func NewMinPriorityQueue[E comparable]() *PriorityQueue[E] {
	return &PriorityQueue[E]{
		items:      make([]E, 0),
		priorities: make(map[E]int),
	}
}

func NewMaxPriorityQueue[E comparable]() *PriorityQueue[E] {
	return &PriorityQueue[E]{
		items:      make([]E, 0),
		priorities: make(map[E]int),
		maxPrio:    true,
	}
}

func (q *PriorityQueue[E]) Get(item E) (int, bool) {
	priority, ok := q.priorities[item]
	return priority, ok
}

func (q *PriorityQueue[E]) Set(item E, priority int) {
	if _, ok := q.priorities[item]; !ok {
		q.items = append(q.items, item)
	}
	q.priorities[item] = priority
	sort.Sort(q)
}

func (q *PriorityQueue[E]) Next() (E, int) {
	var item E
	item, q.items = q.items[0], q.items[1:]

	priority := q.priorities[item]

	delete(q.priorities, item)

	return item, priority
}

func (q *PriorityQueue[E]) Len() int {
	return len(q.items)
}

func (q *PriorityQueue[E]) Swap(i, j int) {
	q.items[i], q.items[j] = q.items[j], q.items[i]
}

func (q *PriorityQueue[E]) Less(i, j int) bool {
	if q.maxPrio {
		return q.priorities[q.items[i]] > q.priorities[q.items[j]]
	}
	return q.priorities[q.items[i]] < q.priorities[q.items[j]]
}
