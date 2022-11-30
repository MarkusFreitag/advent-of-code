package util

import "github.com/MarkusFreitag/advent-of-code/util/constraints"

func Keys[K constraints.Comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))
	var idx int
	for key := range m {
		keys[idx] = key
		idx++
	}
	return keys
}

func Values[K constraints.Comparable, V any](m map[K]V) []V {
	values := make([]V, len(m))
	var idx int
	for _, value := range m {
		values[idx] = value
		idx++
	}
	return values
}
