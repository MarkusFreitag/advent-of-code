package slice

import (
	"fmt"

	"github.com/MarkusFreitag/advent-of-code/util/constraints"
)

func Contains[T constraints.Comparable](slice []T, item T) bool {
	return Index(slice, item) >= 0
}

func Index[T constraints.Comparable](slice []T, item T) int {
	for idx, i := range slice {
		if i == item {
			return idx
		}
	}
	return -1
}

func LastIndex[T constraints.Comparable](slice []T, item T) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == item {
			return i
		}
	}
	return -1
}

func Any[T constraints.Comparable](slice []T, item T) bool {
	return Index(slice, item) >= 0
}

func All[T constraints.Comparable](slice []T, item T) bool {
	for _, i := range slice {
		if i != item {
			return false
		}
	}
	return true
}

func Reverse[T any](slice []T) {
	first := 0
	last := len(slice) - 1
	for first < last {
		slice[first], slice[last] = slice[last], slice[first]
		first++
		last--
	}
}

func Cut[T any](slice []T, start, end int) []T {
	return append(slice[:start], slice[end+1:]...)
}

/*
func Delete[T any](a []T, index int) []T {
	return append(a[:i], a[i+1:]...)
}

func FilterInPlace[T any](a []T, keep func(item T) bool) []T {
	var n int
	for _, x := range a {
		if keep(x) {
			a[n] = x
			n++
		}
	}
	return a[:n]
}

func Pop[T any](a []T) ([]T, T) {
	return a[len(a)-1], a[:len(a)-1]
}

func PopFront[T any](a []T) []T {
	return a[0], a[1:]
}

func Push[T any](a []T, x T) []T {
	return append(a, x)
}

func PushFront[T any](a []T, x T) []T {
	return append([]T{x}, a...)
}

func AppendVector[T any](a, b []T) []T {
	return append(a, b...)
}

func InsertVector[T any](a, b []T, i int) []T {
	return append(a[:i], append(b, a[i:]...)...)
}

func Expand[T any](a []T, i, j int) []T {
	return append(a[:i], append(make([]T, j), a[i:]...)...)
}

func Insert[T any](a []T, x T, i int) []T {
	return append(a[:i], append([]T{x}, a[i:]...)...)
}

func Copy[T any](a []T) []T {
	b = make([]T, len(a))
	copy(b, a)
	return b
}

func Extend[T any](a []T, j int) []T {
	return append(a, make([]T, j)...)
}
*/

var exists = struct{}{}

func SliceToSet[T constraints.Comparable](slice []T) map[T]struct{} {
	set := make(map[T]struct{})
	for _, item := range slice {
		set[item] = exists
	}
	return set
}

func SetToSlice[T constraints.Comparable](set map[T]struct{}) []T {
	slice := make([]T, len(set))
	var idx int
	for item := range set {
		slice[idx] = item
		idx++
	}
	return slice
}

func Head[T any](slice []T, length int) []T {
	if len(slice) <= length {
		return slice
	}
	return slice[:length]
}

func Tail[T any](slice []T, length int) []T {
	if len(slice) <= length {
		return slice
	}
	return slice[len(slice)-length:]
}

func Equal[T constraints.Comparable](sliceA, sliceB []T) bool {
	if len(sliceA) != len(sliceB) {
		return false
	}
	equal := true
	for idx, item := range sliceA {
		if sliceB[idx] != item {
			equal = false
		}
		if !equal {
			break
		}
	}
	return equal
}

func Chunks[T any](slice []T, size int) [][]T {
	length := len(slice)

	chunks := make([][]T, 0)
	for i := 0; i < length; i += size {
		end := i + size

		if end > length {
			end = length
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func GetChunk[T any](slice []T, size, index int) []T {
	length := len(slice)
	if length <= size && index == 0 {
		return slice
	}

	idx := size * index
	if idx >= length {
		panic(fmt.Errorf("chunk=%d does not exist for slice-length=%d and chunk-size=%d", index, length, size))
	}

	if idx+size >= length {
		return slice[idx:]
	}

	return slice[idx : idx+size]
}

func Compact[T constraints.Comparable](slice []T) []T {
	var null T
	result := make([]T, 0)
	for _, item := range slice {
		if item != null {
			result = append(result, item)
		}
	}
	return result
}

func Fill[T constraints.Comparable](slice []T, value T) []T {
	var null T
	for idx, item := range slice {
		if item == null {
			slice[idx] = value
		}
	}
	return slice
}

func Flatten[T constraints.Comparable](nested [][]T) []T {
	flat := make([]T, 0)
	for _, slice := range nested {
		flat = append(flat, slice...)
	}
	return flat
}

func Uniq[T constraints.Comparable](slice []T) []T {
	return SetToSlice(SliceToSet(slice))
}
