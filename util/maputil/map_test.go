package maputil_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/MarkusFreitag/advent-of-code/util/maputil"
)

func TestKeys(t *testing.T) {
	m := map[string]any{"a": 1, "b": true, "c": "a", "d": 1.2}
	require.ElementsMatch(t, []string{"a", "b", "c", "d"}, maputil.Keys(m))
}

func TestValues(t *testing.T) {
	m := map[string]any{"a": 1, "b": true, "c": "a", "d": 1.2}
	require.ElementsMatch(t, []any{1, true, "a", 1.2}, maputil.Values(m))
}
