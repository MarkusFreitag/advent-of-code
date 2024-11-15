package util

import (
	"fmt"
	"sort"
	"strings"
)

type queueItem[E comparable] interface {
	CmpKey() E
	comparable
}

type PriorityQueue[Q queueItem[E], E comparable] struct {
	items      []Q
	priorities map[E]int
	maxPrio    bool
}

func NewMinPriorityQueue[Q queueItem[E], E comparable]() *PriorityQueue[Q, E] {
	return &PriorityQueue[Q, E]{
		items:      make([]Q, 0),
		priorities: make(map[E]int),
	}
}

func NewMaxPriorityQueue[Q queueItem[E], E comparable]() *PriorityQueue[Q, E] {
	return &PriorityQueue[Q, E]{
		items:      make([]Q, 0),
		priorities: make(map[E]int),
		maxPrio:    true,
	}
}

func (q *PriorityQueue[Q, E]) Get(item Q) (int, bool) {
	priority, ok := q.priorities[item.CmpKey()]
	return priority, ok
}

func (q *PriorityQueue[Q, E]) Set(item Q, priority int) {
	if _, ok := q.priorities[item.CmpKey()]; !ok {
		q.items = append(q.items, item)
	}
	q.priorities[item.CmpKey()] = priority
	sort.Sort(q)
}

func (q *PriorityQueue[Q, E]) Next() (Q, int) {
	var item Q
	item, q.items = q.items[0], q.items[1:]

	priority := q.priorities[item.CmpKey()]

	delete(q.priorities, item.CmpKey())

	return item, priority
}

func (q *PriorityQueue[Q, E]) Len() int {
	return len(q.items)
}

func (q *PriorityQueue[Q, E]) Swap(i, j int) {
	q.items[i], q.items[j] = q.items[j], q.items[i]
}

func (q *PriorityQueue[Q, E]) Less(i, j int) bool {
	if q.maxPrio {
		return q.priorities[q.items[i].CmpKey()] > q.priorities[q.items[j].CmpKey()]
	}
	return q.priorities[q.items[i].CmpKey()] < q.priorities[q.items[j].CmpKey()]
}

func (q *PriorityQueue[Q, E]) String() string {
	s := make([]string, len(q.items))
	for idx, item := range q.items {
		s[idx] = fmt.Sprintf("%v", item)
	}
	return strings.Join(s, ",")
}
