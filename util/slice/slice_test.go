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
}

func TestIndex(t *testing.T) {
	strSlice := []string{"a", "b", "c"}
	require.Equal(t, 1, slice.Index(strSlice, "b"))
	require.Equal(t, -1, slice.Index(strSlice, "d"))
	strSlice = append(strSlice, "b")
	require.Equal(t, 1, slice.Index(strSlice, "b"))
}

func TestLastIndex(t *testing.T) {
	strSlice := []string{"a", "b", "c"}
	require.Equal(t, 1, slice.LastIndex(strSlice, "b"))
	require.Equal(t, -1, slice.LastIndex(strSlice, "d"))
	strSlice = append(strSlice, "b")
	require.Equal(t, 3, slice.LastIndex(strSlice, "b"))
}

func TestAny(t *testing.T) {
	require.True(t, slice.Any([]string{"a", "a"}, "a"))
	require.True(t, slice.Any([]string{"a", "b"}, "a"))
	require.False(t, slice.Any([]string{"b", "b"}, "a"))

	require.True(t, slice.Any([]bool{true, true}, true))
	require.True(t, slice.Any([]bool{true, false}, true))
	require.False(t, slice.Any([]bool{false, false}, true))
}

func TestAll(t *testing.T) {
	require.True(t, slice.All([]string{"a", "a"}, "a"))
	require.False(t, slice.All([]string{"a", "b"}, "a"))
	require.False(t, slice.All([]string{"b", "b"}, "a"))

	require.True(t, slice.All([]bool{true, true}, true))
	require.False(t, slice.All([]bool{true, false}, true))
	require.False(t, slice.All([]bool{false, false}, true))
}

func TestReverse(t *testing.T) {
	strSlice := []string{"a", "b", "c"}
	slice.Reverse(strSlice)
	require.Equal(t, []string{"c", "b", "a"}, strSlice)
}

func TestDelete(t *testing.T) {
	strSlice := []string{"a", "b", "c", "d"}
	strSlice = slice.Delete(strSlice, 1)
	require.Equal(t, []string{"a", "c", "d"}, strSlice)
	strSlice = slice.Delete(strSlice, 2)
	require.Equal(t, []string{"a", "c"}, strSlice)
	strSlice = slice.Delete(strSlice, 0)
	require.Equal(t, []string{"c"}, strSlice)
}

func TestPop(t *testing.T) {
	strSlice := []string{"a", "b", "c"}
	item, strSlice := slice.Pop(strSlice)
	require.Equal(t, []string{"a", "b"}, strSlice)
	require.Equal(t, "c", item)
	item, strSlice = slice.Pop(strSlice)
	require.Equal(t, []string{"a"}, strSlice)
	require.Equal(t, "b", item)
	item, strSlice = slice.Pop(strSlice)
	require.Equal(t, []string{}, strSlice)
	require.Equal(t, "a", item)
}

func TestPopFront(t *testing.T) {
	strSlice := []string{"a", "b", "c"}
	item, strSlice := slice.PopFront(strSlice)
	require.Equal(t, []string{"b", "c"}, strSlice)
	require.Equal(t, "a", item)
	item, strSlice = slice.PopFront(strSlice)
	require.Equal(t, []string{"c"}, strSlice)
	require.Equal(t, "b", item)
	item, strSlice = slice.PopFront(strSlice)
	require.Equal(t, []string{}, strSlice)
	require.Equal(t, "c", item)
}

func TestPush(t *testing.T) {
	strSlice := []string{}
	strSlice = slice.Push(strSlice, "a")
	require.Equal(t, []string{"a"}, strSlice)
	strSlice = slice.Push(strSlice, "b")
	require.Equal(t, []string{"a", "b"}, strSlice)
	strSlice = slice.Push(strSlice, "c")
	require.Equal(t, []string{"a", "b", "c"}, strSlice)
}

func TestPushFront(t *testing.T) {
	strSlice := []string{}
	strSlice = slice.PushFront(strSlice, "c")
	require.Equal(t, []string{"c"}, strSlice)
	strSlice = slice.PushFront(strSlice, "b")
	require.Equal(t, []string{"b", "c"}, strSlice)
	strSlice = slice.PushFront(strSlice, "a")
	require.Equal(t, []string{"a", "b", "c"}, strSlice)
}

func TestInsert(t *testing.T) {
	strSlice := []string{}
	strSlice = slice.Insert(strSlice, "", 0)
	require.Equal(t, []string{""}, strSlice)
	strSlice = slice.Insert(strSlice, "a", 0)
	require.Equal(t, []string{"a", ""}, strSlice)
	strSlice = slice.Insert(strSlice, "b", 1)
	require.Equal(t, []string{"a", "b", ""}, strSlice)
}

func TestInsertVector(t *testing.T) {
	strSlice := []string{}
	strSlice = slice.InsertSlice(strSlice, []string{}, 0)
	require.Equal(t, []string{}, strSlice)
	strSlice = slice.InsertSlice(strSlice, []string{"a"}, 0)
	require.Equal(t, []string{"a"}, strSlice)
	strSlice = slice.InsertSlice(strSlice, []string{"b"}, 0)
	require.Equal(t, []string{"b", "a"}, strSlice)
	strSlice = slice.InsertSlice(strSlice, []string{"c"}, 1)
	require.Equal(t, []string{"b", "c", "a"}, strSlice)
	strSlice = slice.InsertSlice(strSlice, []string{"1", "2"}, 1)
	require.Equal(t, []string{"b", "1", "2", "c", "a"}, strSlice)
}

func TestCopy(t *testing.T) {
	strSlice := []string{"a", "b", "c"}
	strSliceCopy := slice.Copy(strSlice)
	strSliceCopy[1] = "d"
	require.NotEqual(t, strSlice, strSliceCopy)
}