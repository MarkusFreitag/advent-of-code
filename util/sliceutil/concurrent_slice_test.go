package sliceutil_test

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/MarkusFreitag/advent-of-code/util/sliceutil"
)

func TestConcurrentSlice(t *testing.T) {
	cSlice := sliceutil.NewConcurrentSlice[[]int, int]()

	require.Equal(t, 0, cSlice.Length())

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, cs *sliceutil.ConcurrentSlice[[]int, int], i int) {
			defer wg.Done()
			cs.Append(i + 1)
		}(&wg, cSlice, i)
	}

	wg.Wait()

	length := cSlice.Length()
	require.Equal(t, 3, length)
	items := make([]int, length)
	for index, value := range cSlice.Seq() {
		items[index] = value
	}

	require.ElementsMatch(t, []int{1, 2, 3}, items)
	require.ElementsMatch(t, []int{1, 2, 3}, cSlice.Items())
	require.Equal(t, 2, cSlice.Get(1))
}
