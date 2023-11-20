package setutil_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/MarkusFreitag/advent-of-code/util/setutil"
)

func TestSetOperations(t *testing.T) {
	left := []int{1, 2, 3}
	middle := []int{4, 5}
	right := []int{6, 7, 8}

	a := append(left, middle...)
	b := append(middle, right...)

	require.ElementsMatch(t, a, setutil.SetToSlice(setutil.SliceToSet(a)))

	diffA, diffB := setutil.Difference(a, b)
	require.ElementsMatch(t, left, diffA)
	require.ElementsMatch(t, right, diffB)

	require.ElementsMatch(t, middle, setutil.Intersect(a, b))

	require.ElementsMatch(t, append(left, append(middle, right...)...), setutil.Union(a, b))

	require.ElementsMatch(t, left, setutil.LeftJoin(a, b))
	require.ElementsMatch(t, right, setutil.RightJoin(a, b))
	require.ElementsMatch(t, middle, setutil.InnerJoin(a, b))
	require.ElementsMatch(t, append(left, right...), setutil.OuterJoin(a, b))
}
