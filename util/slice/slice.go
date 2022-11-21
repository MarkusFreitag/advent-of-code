package slice

import "github.com/MarkusFreitag/advent-of-code/util/constraints"

func Contains[T constraints.Ordered](slice []T, item T) bool {
	return IndexOf(slice, item) >= 0
}

func IndexOf[T constraints.Ordered](slice []T, item T) int {
	for idx, i := range slice {
		if i == item {
			return idx
		}
	}
	return -1
}

func Any[T constraints.Comparable](slice []T, item T) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}
	return false
}

func All[T constraints.Comparable](slice []T, item T) bool {
	for _, i := range slice {
		if i != item {
			return false
		}
	}
	return true
}
