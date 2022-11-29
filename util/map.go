package util

import "github.com/MarkusFreitag/advent-of-code/util/constraints"

func Keys[T constraints.Comparable](m map[T]any) []T {
	keys := make([]T, len(m))
	var idx int
	for key := range m {
		keys[idx] = key
		idx++
	}
	return keys
}

func Values[T constraints.Comparable](m map[any]T) []T {
	values := make([]T, len(m))
	var idx int
	for _, value := range m {
		values[idx] = value
		idx++
	}
	return values
}
