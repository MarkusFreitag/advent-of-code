package iterutil_test

import (
	"slices"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/MarkusFreitag/advent-of-code/util/iterutil"
)

func TestMap(t *testing.T) {
	slice := []string{"a", "1", "2", "b", "3", "c", "4", "d", "e", "5"}

	assert.Equal(
		t,
		[]string{"1", "2", "3", "4", "5"},
		slices.Collect(
			iterutil.Map(
				iterutil.SeqFromSlice(slice), func(str string) (string, bool) {
					_, err := strconv.Atoi(str)
					return str, err == nil
				},
			),
		),
	)
	assert.Equal(
		t,
		[]int{1, 2, 3, 4, 5},
		slices.Collect(
			iterutil.Map(
				iterutil.SeqFromSlice(slice), func(str string) (int, bool) {
					n, err := strconv.Atoi(str)
					return n, err == nil
				},
			),
		),
	)
	assert.Equal(
		t,
		[]string{"2", "3", "4"},
		slices.Collect(
			iterutil.SeqFromSeq2(
				iterutil.Map2(
					iterutil.Seq2FromSlice(slice), func(index int, item string) (int, string, bool) {
						_, err := strconv.Atoi(item)
						if err != nil {
							return 0, "", false
						}
						return index, item, index%2 == 0
					},
				),
				func(_ int, v string) string { return v },
			),
		),
	)
	assert.Equal(
		t,
		[]int{2, 3, 4},
		slices.Collect(
			iterutil.SeqFromSeq2(
				iterutil.Map2(
					iterutil.Seq2FromSlice(slice), func(index int, item string) (int, int, bool) {
						n, err := strconv.Atoi(item)
						if err != nil {
							return 0, 0, false
						}
						return index, n, index%2 == 0
					},
				),
				func(_, v int) int { return v },
			),
		),
	)
}
