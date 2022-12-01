package slice

import (
	"fmt"
	"sort"

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

func Delete[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

func Pop[T any](slice []T) (T, []T) {
	return slice[len(slice)-1], slice[:len(slice)-1]
}

func PopFront[T any](slice []T) (T, []T) {
	return slice[0], slice[1:]
}

func Push[T any](slice []T, item T) []T {
	return append(slice, item)
}

func PushFront[T any](slice []T, item T) []T {
	return append([]T{item}, slice...)
}

func InsertSlice[T any](sliceA, sliceB []T, index int) []T {
	return append(sliceA[:index], append(sliceB, sliceA[index:]...)...)
}

func Insert[T any](slice []T, item T, index int) []T {
	return append(slice[:index], append([]T{item}, slice[index:]...)...)
}

func Copy[T any](slice []T) []T {
	dup := make([]T, len(slice))
	copy(dup, slice)
	return dup
}

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

func Fill[T constraints.Comparable](slice []T, value T) {
	var null T
	for idx, item := range slice {
		if item == null {
			slice[idx] = value
		}
	}
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

func SortAsc[T constraints.Numbers](slice []T) {
	sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
}

func SortDesc[T constraints.Numbers](slice []T) {
	sort.Slice(slice, func(i, j int) bool { return slice[i] > slice[j] })
}
