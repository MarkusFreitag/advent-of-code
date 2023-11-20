package sliceutil

import (
	"cmp"
	"fmt"
	"slices"
)

func LastIndex[S ~[]E, E comparable](slice S, item E) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == item {
			return i
		}
	}
	return -1
}

func Any[S ~[]E, E comparable](slice S, item E) bool {
	return slices.Index(slice, item) >= 0
}

func All[S ~[]E, E comparable](slice S, item E) bool {
	for _, i := range slice {
		if i != item {
			return false
		}
	}
	return true
}

func Pop[S ~[]E, E any](slice S) (E, S) {
	return slice[len(slice)-1], slice[:len(slice)-1]
}

func PopN[S ~[]E, E any](slice S, n int) (S, S) {
	if n == 0 {
		return nil, slice
	}
	return slice[len(slice)-n:], slice[:len(slice)-n]
}

func PopFront[S ~[]E, E comparable](slice S) (E, S) {
	return slice[0], slice[1:]
}

func PopFrontN[S ~[]E, E comparable](slice S, n int) (S, S) {
	if n == 0 {
		return nil, slice
	}
	return slice[:n], slice[n:]
}

func PopIndex[S ~[]E, E comparable](slice S, index int) (E, S) {
	if index == 0 {
		return slice[0], slice[1:]
	}

	length := len(slice)
	if index == length-1 {
		return slice[length-1], slice[:length-1]
	}

	val := slice[index]
	return val, append(slice[:index], slice[index+1:]...)
}

func Push[S ~[]E, E comparable](slice S, item E) S {
	return append(slice, item)
}

func PushFront[S ~[]E, E comparable](slice S, item E) S {
	return append(S{item}, slice...)
}

func Head[S ~[]E, E any](slice S, length int) S {
	if len(slice) <= length {
		return slice
	}
	return slice[:length]
}

func Tail[S ~[]E, E any](slice S, length int) S {
	if len(slice) <= length {
		return slice
	}
	return slice[len(slice)-length:]
}

func Chunks[S ~[]E, E any](slice S, size int) []S {
	length := len(slice)

	chunks := make([]S, 0)
	for i := 0; i < length; i += size {
		end := i + size

		if end > length {
			end = length
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func GetChunk[S ~[]E, E any](slice S, size, index int) S {
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

func Fill[S ~[]E, E comparable](slice S, value E) {
	var null E
	for idx, item := range slice {
		if item == null {
			slice[idx] = value
		}
	}
}

func Flatten[S ~[]E, E any](nested []S) S {
	flat := make(S, 0)
	for _, slice := range nested {
		flat = append(flat, slice...)
	}
	return flat
}

func Uniq[S ~[]E, E comparable](slice S) S {
	return slices.Compact(slice)
}

func SortAsc[S ~[]E, E cmp.Ordered](slice S) {
	slices.Sort(slice)
}

func SortDesc[S ~[]E, E cmp.Ordered](slice S) {
	slices.Sort(slice)
	slices.Reverse(slice)
}

type Slide[S ~[]E, E any] struct {
	Index  int
	Values S
}

func SlidingWindow[S ~[]E, E any](slice S, size int) chan Slide[S, E] {
	slider := make(chan Slide[S, E], 0)
	go func() {
		length := len(slice)
		for idx := range slice {
			end := idx + size
			if end > length {
				end = length
			}
			slider <- Slide[S, E]{
				Index:  idx,
				Values: slice[idx:end],
			}
		}
		close(slider)
	}()
	return slider
}
