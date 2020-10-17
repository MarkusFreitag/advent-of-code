package day4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIncludeDouble(t *testing.T) {
	require.True(t, includeDouble("122"))
	require.True(t, includeDouble("111"))
	require.False(t, includeDouble("101"))
}

func TestOnlyDoubles(t *testing.T) {
	require.True(t, onlyDoubles("122"))
	require.False(t, onlyDoubles("111"))
	require.False(t, onlyDoubles("101"))
}

func TestOnlyIncreasing(t *testing.T) {
	require.True(t, onlyIncreasing("149"))
	require.False(t, onlyIncreasing("1428"))
}
