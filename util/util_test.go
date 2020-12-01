package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInputFilename(t *testing.T) {
	require.Equal(t, "inputs/2019_1.txt", inputFilename(2019, 1))
}

func TestStrInSlice(t *testing.T) {
	require.True(t, StrInSlice("3", []string{"1", "2", "3"}))
	require.False(t, StrInSlice("4", []string{"1", "2", "3"}))
}

func TestIntInSlice(t *testing.T) {
	require.True(t, IntInSlice(3, []int{1, 2, 3}))
	require.False(t, IntInSlice(4, []int{1, 2, 3}))
}

func TestSum(t *testing.T) {
	require.Equal(t, 5, Sum(5))
	require.Equal(t, 5, Sum(2, 3))
}

func TestStrsToInts(t *testing.T) {
	nums := StrsToInts([]string{"1", "2", "3"})
	require.Equal(t, []int{1, 2, 3}, nums)

	require.Panics(t, func() { StrsToInts([]string{"1", "a", "3"}) })
}
