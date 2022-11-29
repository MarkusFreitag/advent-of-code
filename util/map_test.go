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
	m := map[any]string{1: "a", true: "b", "a": "c", 1.2: "d"}
	require.ElementsMatch(t, []string{"a", "b", "c", "d"}, util.Values(m))
}
