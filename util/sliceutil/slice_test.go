package sliceutil_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/MarkusFreitag/advent-of-code/util/sliceutil"
)

func TestLastIndex(t *testing.T) {
	strSlice := []string{"a", "b", "c"}
	require.Equal(t, 1, sliceutil.LastIndex(strSlice, "b"))
	require.Equal(t, -1, sliceutil.LastIndex(strSlice, "d"))
	strSlice = append(strSlice, "b")
	require.Equal(t, 3, sliceutil.LastIndex(strSlice, "b"))
}

func TestAny(t *testing.T) {
	require.True(t, sliceutil.Any([]string{"a", "a"}, "a"))
	require.True(t, sliceutil.Any([]string{"a", "b"}, "a"))
	require.False(t, sliceutil.Any([]string{"b", "b"}, "a"))

	require.True(t, sliceutil.Any([]bool{true, true}, true))
	require.True(t, sliceutil.Any([]bool{true, false}, true))
	require.False(t, sliceutil.Any([]bool{false, false}, true))
}

func TestAnyFunc(t *testing.T) {
	fn := func(i string) bool { return string(i[1]) == "a" }
	require.True(t, sliceutil.AnyFunc([]string{"aa", "aa"}, fn))
	require.True(t, sliceutil.AnyFunc([]string{"aa", "ab"}, fn))
	require.False(t, sliceutil.AnyFunc([]string{"ab", "ab"}, fn))
}

func TestAll(t *testing.T) {
	require.True(t, sliceutil.All([]string{"a", "a"}, "a"))
	require.False(t, sliceutil.All([]string{"a", "b"}, "a"))
	require.False(t, sliceutil.All([]string{"b", "b"}, "a"))

	require.True(t, sliceutil.All([]bool{true, true}, true))
	require.False(t, sliceutil.All([]bool{true, false}, true))
	require.False(t, sliceutil.All([]bool{false, false}, true))
}

func TestAllFunc(t *testing.T) {
	fn := func(i string) bool { return string(i[1]) == "a" }
	require.True(t, sliceutil.AllFunc([]string{"aa", "aa"}, fn))
	require.False(t, sliceutil.AllFunc([]string{"aa", "ab"}, fn))
	require.False(t, sliceutil.AllFunc([]string{"ab", "ab"}, fn))
}

func TestPop(t *testing.T) {
	strSlice := []string{"a", "b", "c"}
	item, strSlice := sliceutil.Pop(strSlice)
	require.Equal(t, []string{"a", "b"}, strSlice)
	require.Equal(t, "c", item)
	item, strSlice = sliceutil.Pop(strSlice)
	require.Equal(t, []string{"a"}, strSlice)
	require.Equal(t, "b", item)
	item, strSlice = sliceutil.Pop(strSlice)
	require.Equal(t, []string{}, strSlice)
	require.Equal(t, "a", item)
}

func TestPopN(t *testing.T) {
	strSlice := []string{"a", "b", "c", "d", "e"}
	popped, strSlice := sliceutil.PopN(strSlice, 1)
	require.Equal(t, []string{"a", "b", "c", "d"}, strSlice)
	require.Equal(t, []string{"e"}, popped)
	popped, strSlice = sliceutil.PopN(strSlice, 3)
	require.Equal(t, []string{"a"}, strSlice)
	require.Equal(t, []string{"b", "c", "d"}, popped)
	popped, strSlice = sliceutil.PopN(strSlice, 0)
	require.Equal(t, []string{"a"}, strSlice)
	require.Empty(t, popped)
}

func TestPopFront(t *testing.T) {
	strSlice := []string{"a", "b", "c"}
	item, strSlice := sliceutil.PopFront(strSlice)
	require.Equal(t, []string{"b", "c"}, strSlice)
	require.Equal(t, "a", item)
	item, strSlice = sliceutil.PopFront(strSlice)
	require.Equal(t, []string{"c"}, strSlice)
	require.Equal(t, "b", item)
	item, strSlice = sliceutil.PopFront(strSlice)
	require.Equal(t, []string{}, strSlice)
	require.Equal(t, "c", item)
}

func TestPopFrontN(t *testing.T) {
	strSlice := []string{"a", "b", "c", "d", "e"}
	popped, strSlice := sliceutil.PopFrontN(strSlice, 1)
	require.Equal(t, []string{"b", "c", "d", "e"}, strSlice)
	require.Equal(t, []string{"a"}, popped)
	popped, strSlice = sliceutil.PopFrontN(strSlice, 3)
	require.Equal(t, []string{"e"}, strSlice)
	require.Equal(t, []string{"b", "c", "d"}, popped)
	popped, strSlice = sliceutil.PopFrontN(strSlice, 0)
	require.Equal(t, []string{"e"}, strSlice)
	require.Empty(t, popped)
}

func TestPopIndex(t *testing.T) {
	strSlice := []string{"a", "b", "c", "d", "e"}
	popped, strSlice := sliceutil.PopIndex(strSlice, 0)
	require.Equal(t, []string{"b", "c", "d", "e"}, strSlice)
	require.Equal(t, "a", popped)
	popped, strSlice = sliceutil.PopIndex(strSlice, len(strSlice)-1)
	require.Equal(t, []string{"b", "c", "d"}, strSlice)
	require.Equal(t, "e", popped)
	popped, strSlice = sliceutil.PopIndex(strSlice, 1)
	require.Equal(t, []string{"b", "d"}, strSlice)
	require.Equal(t, "c", popped)
}

func TestPush(t *testing.T) {
	strSlice := []string{}
	strSlice = sliceutil.Push(strSlice, "a")
	require.Equal(t, []string{"a"}, strSlice)
	strSlice = sliceutil.Push(strSlice, "b")
	require.Equal(t, []string{"a", "b"}, strSlice)
	strSlice = sliceutil.Push(strSlice, "c")
	require.Equal(t, []string{"a", "b", "c"}, strSlice)
}

func TestPushFront(t *testing.T) {
	strSlice := []string{}
	strSlice = sliceutil.PushFront(strSlice, "c")
	require.Equal(t, []string{"c"}, strSlice)
	strSlice = sliceutil.PushFront(strSlice, "b")
	require.Equal(t, []string{"b", "c"}, strSlice)
	strSlice = sliceutil.PushFront(strSlice, "a")
	require.Equal(t, []string{"a", "b", "c"}, strSlice)
}

func TestSort(t *testing.T) {
	num := []int{2, 1, 4, 5, 8, 6, 3, 9, 7, 8, 4}
	sliceutil.SortAsc(num)
	require.Equal(t, []int{1, 2, 3, 4, 4, 5, 6, 7, 8, 8, 9}, num)

	num = []int{2, 1, 4, 5, 8, 6, 3, 9, 7, 8, 4}
	sliceutil.SortDesc(num)
	require.Equal(t, []int{9, 8, 8, 7, 6, 5, 4, 4, 3, 2, 1}, num)
}

func TestSlidingWindow(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	windows := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5},
	}
	for index, values := range sliceutil.SlidingWindow(nums, 2) {
		require.Equal(t, windows[index], values)
	}
}

func TestCount(t *testing.T) {
	require.Equal(t, 1, sliceutil.Count([]int{1, 2, 4, 2}, 1))
	require.Equal(t, 2, sliceutil.Count([]int{1, 2, 4, 2}, 2))
	require.Equal(t, 0, sliceutil.Count([]int{1, 2, 4, 2}, 3))
}

func TestTally(t *testing.T) {
	require.Equal(t, map[rune]int{'a': 2, 'b': 2, 'c': 1, 'd': 2, 'E': 1, 'A': 1}, sliceutil.Tally([]rune("abcddEaAb")))
}

func TestMap(t *testing.T) {
	assert.Equal(t,
		[]int{1, 2, 3, 4, 5},
		sliceutil.Map(
			[]string{"a", "1", "2", "b", "3", "c", "4", "d", "e", "5"},
			func(str string) (int, bool) {
				n, err := strconv.Atoi(str)
				return n, err == nil
			},
		),
	)
}
