package slice_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/MarkusFreitag/advent-of-code/util/slice"
)

func TestSetOperations(t *testing.T) {
	left := []int{1, 2, 3}
	middle := []int{4, 5}
	right := []int{6, 7, 8}

	a := append(left, middle...)
	b := append(middle, right...)

	require.ElementsMatch(t, a, slice.SetToSlice(slice.SliceToSet(a)))

	diffA, diffB := slice.Difference(a, b)
	require.ElementsMatch(t, left, diffA)
	require.ElementsMatch(t, right, diffB)

	require.ElementsMatch(t, middle, slice.Intersect(a, b))

	require.ElementsMatch(t, append(left, append(middle, right...)...), slice.Union(a, b))

	require.ElementsMatch(t, left, slice.LeftJoin(a, b))
	require.ElementsMatch(t, right, slice.RightJoin(a, b))
	require.ElementsMatch(t, middle, slice.InnerJoin(a, b))
	require.ElementsMatch(t, append(left, right...), slice.OuterJoin(a, b))
}
