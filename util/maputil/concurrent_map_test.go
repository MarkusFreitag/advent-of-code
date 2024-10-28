package maputil_test

import (
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/MarkusFreitag/advent-of-code/util/maputil"
)

func TestConcurrentMap(t *testing.T) {
	cMap := maputil.NewConcurrentMap[string, int]()

	require.Equal(t, 0, cMap.Length())

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, cm *maputil.ConcurrentMap[string, int], i int) {
			defer wg.Done()
			cm.Set(strconv.Itoa(i+1), i+1)
		}(&wg, cMap, i)
	}

	wg.Wait()

	val, ok := cMap.Get("1")
	require.True(t, ok)
	require.Equal(t, 1, val)

	val, ok = cMap.Get("4")
	require.False(t, ok)
	require.Equal(t, 0, val)

	length := cMap.Length()
	require.Equal(t, 3, length)
	keys := make([]string, 0, length)
	values := make([]int, 0, length)
	for key, value := range cMap.Seq() {
		keys = append(keys, key)
		values = append(values, value)
	}

	require.ElementsMatch(t, []string{"1", "2", "3"}, keys)
	require.ElementsMatch(t, []string{"1", "2", "3"}, cMap.Keys())
	require.ElementsMatch(t, []int{1, 2, 3}, values)
	require.ElementsMatch(t, []int{1, 2, 3}, cMap.Values())

	cMap.Delete("1")
	require.Equal(t, length-1, cMap.Length())
	require.ElementsMatch(t, []string{"2", "3"}, cMap.Keys())
}
