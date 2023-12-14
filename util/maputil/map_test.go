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

func TestKeysFiltered(t *testing.T) {
	m := map[string]any{"aa": 1, "ab": true, "bc": "a", "ad": 1.2}
	fn := func(k string) bool {
		return string(k[0]) == "a"
	}
	require.ElementsMatch(t, []string{"aa", "ab", "ad"}, maputil.KeysFiltered(m, fn))
}

func TestValues(t *testing.T) {
	m := map[string]any{"a": 1, "b": true, "c": "a", "d": 1.2}
	require.ElementsMatch(t, []any{1, true, "a", 1.2}, maputil.Values(m))
}

func TestValuesFiltered(t *testing.T) {
	m := map[string]any{"a": 1, "b": true, "c": 2, "d": 1.2}
	fn := func(v any) bool {
		_, ok := v.(int)
		return ok
	}
	require.ElementsMatch(t, []any{1, 2}, maputil.ValuesFiltered(m, fn))
}

func TestGet(t *testing.T) {
	require.Equal(t, 1, maputil.Get(map[string]int{"a": 1, "c": 3}, "a", 2))
	require.Equal(t, 2, maputil.Get(map[string]int{"a": 1, "c": 3}, "b", 2))
}
