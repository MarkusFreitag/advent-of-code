package sliceutil_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/MarkusFreitag/advent-of-code/util/sliceutil"
)

func TestAllCombinations(t *testing.T) {
	tt := []struct {
		name string
		in   []string
		out  [][]string
	}{
		{
			name: "Empty slice",
			in:   []string{},
			out:  nil,
		},
		{
			name: "Single item",
			in:   []string{"A"},
			out: [][]string{
				{"A"},
			},
		},
		{
			name: "Two items",
			in:   []string{"A", "B"},
			out: [][]string{
				{"A"},
				{"B"},
				{"A", "B"},
			},
		},
		{
			name: "Three items",
			in:   []string{"A", "B", "C"},
			out: [][]string{
				{"A"},
				{"B"},
				{"A", "B"},
				{"C"},
				{"A", "C"},
				{"B", "C"},
				{"A", "B", "C"},
			},
		},
		{
			name: "Four items",
			in:   []string{"A", "B", "C", "D"},
			out: [][]string{
				{"A"},
				{"B"},
				{"A", "B"},
				{"C"},
				{"A", "C"},
				{"B", "C"},
				{"A", "B", "C"},
				{"D"},
				{"A", "D"},
				{"B", "D"},
				{"A", "B", "D"},
				{"C", "D"},
				{"A", "C", "D"},
				{"B", "C", "D"},
				{"A", "B", "C", "D"},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.out, sliceutil.AllCombinations(tc.in))

			var all [][]string
			sliceutil.AllCombinationsFunc(tc.in, func(comb []string) { all = append(all, comb) })
			require.Equal(t, tc.out, all)
		})
	}
}

func TestCombinations(t *testing.T) {
	tt := []struct {
		name string
		in   []string
		n    int
		out  [][]string
	}{
		{
			name: "Empty slice",
			in:   []string{},
			n:    1,
			out:  nil,
		},
		{
			name: "Single item",
			in:   []string{"A"},
			n:    1,
			out: [][]string{
				{"A"},
			},
		},
		{
			name: "Two items, n = 0",
			in:   []string{"A", "B"},
			n:    0,
			out: [][]string{
				{"A"},
				{"B"},
				{"A", "B"},
			},
		},
		{
			name: "Two items, n = 1",
			in:   []string{"A", "B"},
			n:    1,
			out: [][]string{
				{"A"},
				{"B"},
			},
		}, {
			name: "Two items, n = 2",
			in:   []string{"A", "B"},
			n:    2,
			out: [][]string{
				{"A", "B"},
			},
		},
		{
			name: "Three items, n = 0",
			in:   []string{"A", "B", "C"},
			n:    0,
			out: [][]string{
				{"A"},
				{"B"},
				{"A", "B"},
				{"C"},
				{"A", "C"},
				{"B", "C"},
				{"A", "B", "C"},
			},
		},
		{
			name: "Three items, n = 1",
			in:   []string{"A", "B", "C"},
			n:    1,
			out: [][]string{
				{"A"},
				{"B"},
				{"C"},
			},
		},
		{
			name: "Three items, n = 2",
			in:   []string{"A", "B", "C"},
			n:    2,
			out: [][]string{
				{"A", "B"},
				{"A", "C"},
				{"B", "C"},
			},
		},
		{
			name: "Three items, n = 3",
			in:   []string{"A", "B", "C"},
			n:    3,
			out: [][]string{
				{"A", "B", "C"},
			},
		},
		{
			name: "Three items, n = 4",
			in:   []string{"A", "B", "C"},
			n:    4,
			out: [][]string{
				{"A", "B", "C"},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.out, sliceutil.Combinations(tc.in, tc.n))

			var all [][]string
			sliceutil.CombinationsFunc(tc.in, tc.n, func(comb []string) { all = append(all, comb) })
			require.Equal(t, tc.out, all)
		})
	}
}
