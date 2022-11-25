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

	require.Equal(t, a, slice.SetToSlice(slice.SliceToSet(a)))

	require.Equal(t, left, slice.LeftJoin(a, b))
	require.Equal(t, right, slice.RightJoin(a, b))
	require.Equal(t, middle, slice.InnerJoin(a, b))
	require.Equal(t, append(left, right...), slice.OuterJoin(a, b))
}
