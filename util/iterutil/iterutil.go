package iterutil

import (
	"iter"

	"github.com/MarkusFreitag/advent-of-code/util/setutil"
)

type MapFn[V1, V2 any] func(V1) (V2, bool)
type Map2Fn[K1, K2, V1, V2 any] func(K1, V1) (K2, V2, bool)

func SeqFromSet[E comparable](set setutil.Set[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		for value := range set {
			if !yield(value) {
				return
			}
		}
	}
}

func SeqFromSlice[S ~[]E, E any](slice S) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, item := range Seq2FromSlice(slice) {
			if !yield(item) {
				return
			}
		}
	}
}

func Seq2FromSlice[S ~[]E, E any](slice S) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		for index, item := range slice {
			if !yield(index, item) {
				return
			}
		}
	}
}

func SeqFromSeq2[K, V, R any](seq iter.Seq2[K, V], fn func(K, V) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		for key, value := range seq {
			if !yield(fn(key, value)) {
				return
			}
		}
	}
}

func Count[S iter.Seq[E], E any](seq S) int {
	var counter int
	for range seq {
		counter++
	}
	return counter
}

func Count2[S iter.Seq2[K, V], K, V any](seq S) int {
	var counter int
	for range seq {
		counter++
	}
	return counter
}

func Map[V1, V2 any](seq iter.Seq[V1], fn MapFn[V1, V2]) iter.Seq[V2] {
	return func(yield func(V2) bool) {
		for value := range seq {
			if val, ok := fn(value); ok {
				if !yield(val) {
					return
				}
			}
		}
	}
}
func Map2[K1, K2, V1, V2 any](seq iter.Seq2[K1, V1], fn Map2Fn[K1, K2, V1, V2]) iter.Seq2[K2, V2] {
	return func(yield func(K2, V2) bool) {
		for key, value := range seq {
			if k, v, ok := fn(key, value); ok {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}
