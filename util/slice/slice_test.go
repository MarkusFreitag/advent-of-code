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

func TestIndex(t *testing.T) {
	strSlice := []string{"a", "b", "c"}
	require.Equal(t, 1, slice.Index(strSlice, "b"))
	require.Equal(t, -1, slice.Index(strSlice, "d"))
	strSlice = append(strSlice, "b")
	require.Equal(t, 1, slice.Index(strSlice, "b"))

	intSlice := []int{1, 2, 3}
	require.Equal(t, 1, slice.Index(intSlice, 2))
	require.Equal(t, -1, slice.Index(intSlice, 4))
	intSlice = append(intSlice, 2)
	require.Equal(t, 1, slice.Index(intSlice, 2))

	floatSlice := []float64{1.1, 2.2, 3.3}
	require.Equal(t, 1, slice.Index(floatSlice, 2.2))
	require.Equal(t, -1, slice.Index(floatSlice, 4.4))
	floatSlice = append(floatSlice, 2.2)
	require.Equal(t, 1, slice.Index(floatSlice, 2.2))
}

func TestLastIndex(t *testing.T) {
	strSlice := []string{"a", "b", "c"}
	require.Equal(t, 1, slice.LastIndex(strSlice, "b"))
	require.Equal(t, -1, slice.LastIndex(strSlice, "d"))
	strSlice = append(strSlice, "b")
	require.Equal(t, 3, slice.LastIndex(strSlice, "b"))

	intSlice := []int{1, 2, 3}
	require.Equal(t, 1, slice.LastIndex(intSlice, 2))
	require.Equal(t, -1, slice.LastIndex(intSlice, 4))
	intSlice = append(intSlice, 2)
	require.Equal(t, 3, slice.LastIndex(intSlice, 2))

	floatSlice := []float64{1.1, 2.2, 3.3}
	require.Equal(t, 1, slice.LastIndex(floatSlice, 2.2))
	require.Equal(t, -1, slice.LastIndex(floatSlice, 4.4))
	floatSlice = append(floatSlice, 2.2)
	require.Equal(t, 3, slice.LastIndex(floatSlice, 2.2))
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

func TestReverse(t *testing.T) {
	strSlice := []string{"a", "b", "c"}
	slice.Reverse(strSlice)
	require.Equal(t, []string{"c", "b", "a"}, strSlice)

	intSlice := []int{1, 2, 3}
	slice.Reverse(intSlice)
	require.Equal(t, []int{3, 2, 1}, intSlice)

	floatSlice := []float64{1.1, 2.2, 3.3}
	slice.Reverse(floatSlice)
	require.Equal(t, []float64{3.3, 2.2, 1.1}, floatSlice)
}

func TestCut(t *testing.T) {
	strSlice := []string{"a", "b", "c", "d", "e"}
	require.Equal(t, strSlice, slice.Cut(strSlice, 0, 0))
	require.Equal(t, []string{"c", "d", "e"}, slice.Cut(strSlice, 0, 2))
	require.Equal(t, []string{"a", "b", "c", "d"}, slice.Cut(strSlice, 4, 4))
	require.Equal(t, []string{"a", "e"}, slice.Cut(strSlice, 1, 3))
	require.Equal(t, []string{"a", "b", "d", "e"}, slice.Cut(strSlice, 2, 2))

	/*
		intSlice := []int{1, 2, 3, 4, 5}
		require.Equal(t, intSlice, slice.Cut(intSlice, 0, 0))
		require.Equal(t, []int{2, 3, 4}, slice.Cut(intSlice, 0, 0))
		require.Equal(t, []int{1, 2, 3, 4}, slice.Cut(intSlice, 0, 0))
		require.Equal(t, []int{1, 5}, slice.Cut(intSlice, 0, 0))
		require.Equal(t, []int{1, 2, 4, 5}, slice.Cut(intSlice, 0, 0))

		floatSlice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
		require.Equal(t, floatSlice, slice.Cut(floatSlice, 0, 0))
		require.Equal(t, []float64{3.3, 4.4, 5.5}, slice.Cut(floatSlice, 0, 0))
		require.Equal(t, []float64{1.1, 2.2, 3.3, 4.4}, slice.Cut(floatSlice, 0, 0))
		require.Equal(t, []float64{1.1, 5.5}, slice.Cut(floatSlice, 0, 0))
		require.Equal(t, []float64{1.1, 2.2, 4.4, 5.5}, slice.Cut(floatSlice, 0, 0))
	*/
}
