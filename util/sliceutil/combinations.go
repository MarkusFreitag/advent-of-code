package sliceutil

import "math/bits"

/*
	These functions are heavily inspired by github.com/mxschmitt/golang-combinations
	Returning all combinations does not always work in the AoC context, thats why
	AllCombinationsFunc and CombinationsFunc were added.
*/

func combinationsFn[S ~[]E, E any](set S, n int, fn func(subset S)) {
	length := uint(len(set))

	if n > len(set) {
		n = len(set)
	}

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		if n > 0 && bits.OnesCount(uint(subsetBits)) != n {
			continue
		}

		var subset S

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, set[object])
			}
		}

		// call provided function with subset
		fn(subset)
	}
}

// All returns all combinations for a given array.
// This is essentially a powerset of the given set except that the empty set is disregarded.
func AllCombinations[S ~[]E, E any](set S) (subsets []S) {
	combinationsFn(set, 0, func(subset S) {
		subsets = append(subsets, subset)
	})

	return subsets
}

// AllFunc calls a function with every combination for a given array.
func AllCombinationsFunc[S ~[]E, E any](set S, fn func(subset S)) {
	combinationsFn(set, 0, fn)
}

// Combinations returns combinations of n elements for a given string array.
// For n < 1, it equals to All and returns all combinations.
func Combinations[S ~[]E, E any](set S, n int) (subsets []S) {
	combinationsFn(set, n, func(subset S) {
		subsets = append(subsets, subset)
	})

	return subsets
}

// CombinationsFunc call a function for combinations of n elements for a given array.
// For n < 1, it equals to AllFunc.
func CombinationsFunc[S ~[]E, E any](set S, n int, fn func(subset S)) {
	combinationsFn(set, n, fn)
}
