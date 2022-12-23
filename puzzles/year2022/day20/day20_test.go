package day20

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testInput = `1
2
-3
10
-9
0
4`

func TestLookup(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	require.Equal(t, 1, lookup(slice, 0))
	require.Equal(t, 1, lookup(slice, 5))
	require.Equal(t, 1, lookup(slice, 10))
	require.Equal(t, 3, lookup(slice, 2))
	require.Equal(t, 3, lookup(slice, 7))
	require.Equal(t, 5, lookup(slice, 4))
}

func TestMove(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	nums = move(nums, 2, true)
	require.Equal(t, []int{1, 3, 2, 4, 5}, nums)
	nums = move(nums, 2, true)
	require.Equal(t, []int{1, 3, 4, 2, 5}, nums)
	nums = move(nums, 2, true)
	require.Equal(t, []int{1, 3, 4, 5, 2}, nums)
	nums = move(nums, 2, true)
	require.Equal(t, []int{1, 2, 3, 4, 5}, nums)

	nums = move(nums, 2, false)
	require.Equal(t, []int{2, 1, 3, 4, 5}, nums)
	nums = move(nums, 2, false)
	require.Equal(t, []int{1, 3, 4, 2, 5}, nums)
	nums = move(nums, 2, false)
	require.Equal(t, []int{1, 3, 2, 4, 5}, nums)
	nums = move(nums, 2, false)
	require.Equal(t, []int{1, 2, 3, 4, 5}, nums)
}

func TestPart1(t *testing.T) {
	solution, err := Part1(testInput)
	require.Nil(t, err)
	require.Equal(t, "3", solution)
}

func TestPart2(t *testing.T) {
	solution, err := Part2(testInput)
	require.Nil(t, err)
	require.Equal(t, "", solution)
}
