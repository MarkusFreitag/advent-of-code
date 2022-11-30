package util_test

import (
	"testing"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/stretchr/testify/require"
)

func TestKeys(t *testing.T) {
	m := map[string]any{"a": 1, "b": true, "c": "a", "d": 1.2}
	require.ElementsMatch(t, []string{"a", "b", "c", "d"}, util.Keys(m))
}

func TestValues(t *testing.T) {
	m := map[string]any{"a": 1, "b": true, "c": "a", "d": 1.2}
	require.ElementsMatch(t, []any{1, true, "a", 1.2}, util.Values(m))
}
