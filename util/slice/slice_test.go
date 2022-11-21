package slice_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/MarkusFreitag/advent-of-code/util/slice"
)

func TestContains(t *testing.T) {
	strSlice := []string{"a", "b", "c"}
	require.True(t, slice.Contains(strSlice, "b"))
	require.False(t, slice.Contains(strSlice, "d"))

	intSlice := []int{1, 2, 3}
	require.True(t, slice.Contains(intSlice, 2))
	require.False(t, slice.Contains(intSlice, 4))

	floatSlice := []float64{1.1, 2.2, 3.3}
	require.True(t, slice.Contains(floatSlice, 2.2))
	require.False(t, slice.Contains(floatSlice, 4.4))
}

func TestIndexOf(t *testing.T) {
	strSlice := []string{"a", "b", "c"}
	require.Equal(t, 1, slice.IndexOf(strSlice, "b"))
	require.Equal(t, -1, slice.IndexOf(strSlice, "d"))
	strSlice = append(strSlice, "b")
	require.Equal(t, 1, slice.IndexOf(strSlice, "b"))

	intSlice := []int{1, 2, 3}
	require.Equal(t, 1, slice.IndexOf(intSlice, 2))
	require.Equal(t, -1, slice.IndexOf(intSlice, 4))
	intSlice = append(intSlice, 2)
	require.Equal(t, 1, slice.IndexOf(intSlice, 2))

	floatSlice := []float64{1.1, 2.2, 3.3}
	require.Equal(t, 1, slice.IndexOf(floatSlice, 2.2))
	require.Equal(t, -1, slice.IndexOf(floatSlice, 4.4))
	floatSlice = append(floatSlice, 2.2)
	require.Equal(t, 1, slice.IndexOf(floatSlice, 2.2))
}

func TestAny(t *testing.T) {
	require.True(t, slice.Any([]string{"a", "a"}, "a"))
	require.True(t, slice.Any([]string{"a", "b"}, "a"))
	require.False(t, slice.Any([]string{"b", "b"}, "a"))

	require.True(t, slice.Any([]int{1, 1}, 1))
	require.True(t, slice.Any([]int{1, 2}, 1))
	require.False(t, slice.Any([]int{2, 2}, 1))

	require.True(t, slice.Any([]float64{1.1, 1.1}, 1.1))
	require.True(t, slice.Any([]float64{1.1, 2.2}, 1.1))
	require.False(t, slice.Any([]float64{2.2, 2.2}, 1.1))

	require.True(t, slice.Any([]bool{true, true}, true))
	require.True(t, slice.Any([]bool{true, false}, true))
	require.False(t, slice.Any([]bool{false, false}, true))
}

func TestAll(t *testing.T) {
	require.True(t, slice.All([]string{"a", "a"}, "a"))
	require.False(t, slice.All([]string{"a", "b"}, "a"))
	require.False(t, slice.All([]string{"b", "b"}, "a"))

	require.True(t, slice.All([]int{1, 1}, 1))
	require.False(t, slice.All([]int{1, 2}, 1))
	require.False(t, slice.All([]int{2, 2}, 1))

	require.True(t, slice.All([]float64{1.1, 1.1}, 1.1))
	require.False(t, slice.All([]float64{1.1, 2.2}, 1.1))
	require.False(t, slice.All([]float64{2.2, 2.2}, 1.1))

	require.True(t, slice.All([]bool{true, true}, true))
	require.False(t, slice.All([]bool{true, false}, true))
	require.False(t, slice.All([]bool{false, false}, true))
}
