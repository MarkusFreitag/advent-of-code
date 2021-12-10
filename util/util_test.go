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

func TestSumInts(t *testing.T) {
	require.Equal(t, 5, SumInts(5))
	require.Equal(t, 5, SumInts(2, 3))
	require.Equal(t, -1, SumInts(2, -3))
}

func TestSubInts(t *testing.T) {
	require.Equal(t, -5, SubInts(5))
	require.Equal(t, -5, SubInts(2, 3))
	require.Equal(t, 0, SubInts(-3, 3))
	require.Equal(t, 5, SubInts(-2, -3))
}

func TestMulInts(t *testing.T) {
	require.Equal(t, 5, MulInts(5))
	require.Equal(t, 1, MulInts(1))
	require.Equal(t, 6, MulInts(2, 3))
	require.Equal(t, 0, MulInts(2, 0, 3))
}

func TestPowInt(t *testing.T) {
	require.Equal(t, 8, PowInt(2, 3))
	require.Equal(t, 9, PowInt(3, 2))
}

func TestStrsToInts(t *testing.T) {
	require.Equal(t, []int{1, 2, 3}, StrsToInts([]string{"1", "2", "3"}))

	require.Panics(t, func() { StrsToInts([]string{"1", "a", "3"}) })
}

func TestIntsToStrs(t *testing.T) {
	require.Equal(t, []string{"1", "2", "3"}, IntsToStrs([]int{1, 2, 3}))
}

func TestAbs(t *testing.T) {
	testcases := map[int]int{
		1:  1,
		-2: 2,
	}
	for input, expected := range testcases {
		require.Equal(t, expected, Abs(input))
	}
}

func TestBetween(t *testing.T) {
	require.True(t, Between(1, 1, 3))
	require.True(t, Between(2, 1, 3))
	require.True(t, Between(3, 1, 3))
	require.False(t, Between(4, 1, 3))
}

func TestParseSignedInt(t *testing.T) {
	testcases := map[string]int{
		"1":  1,
		"+2": 2,
		"-3": -3,
	}
	for input, expected := range testcases {
		require.Equal(t, expected, ParseSignedInt(input))
	}
}

func TestReverse(t *testing.T) {
	testcases := map[string]string{
		"abab": "baba",
		"abba": "abba",
	}
	for input, expected := range testcases {
		require.Equal(t, expected, Reverse(input))
	}
}

func TestLeftPad(t *testing.T) {
	testcases := map[string]string{
		"1":     "0001",
		"1010":  "1010",
		"11111": "1111",
	}
	for input, expected := range testcases {
		require.Equal(t, expected, LeftPad(input, "0", 4))
	}
}

func TestRightPad(t *testing.T) {
	testcases := map[string]string{
		"1":     "1000",
		"1010":  "1010",
		"11111": "1111",
	}
	for input, expected := range testcases {
		require.Equal(t, expected, RightPad(input, "0", 4))
	}
}

func TestBinStrToDecInt(t *testing.T) {
	testcases := map[string]int{
		"1":     1,
		"1010":  10,
		"11111": 31,
	}
	for input, expected := range testcases {
		require.Equal(t, expected, BinStrToDecInt(input))
	}
}

func TestDecIntToBinStr(t *testing.T) {
	testcases := map[int]string{
		1:  "1",
		10: "1010",
		31: "11111",
	}
	for input, expected := range testcases {
		require.Equal(t, expected, DecIntToBinStr(input))
	}
}

func TestInRange(t *testing.T) {
	require.False(t, InRange(0, 1, 3))
	require.True(t, InRange(1, 1, 3))
	require.True(t, InRange(2, 1, 3))
	require.True(t, InRange(3, 1, 3))
}

func TestMinInt(t *testing.T) {
	require.Equal(t, 0, MinInt())
	require.Equal(t, 1, MinInt(1))
	require.Equal(t, 1, MinInt(1, 3))
	require.Equal(t, 1, MinInt(3, 1))
}

func TestMaxInt(t *testing.T) {
	require.Equal(t, 0, MaxInt())
	require.Equal(t, 1, MaxInt(1))
	require.Equal(t, 3, MaxInt(1, 3))
	require.Equal(t, 3, MaxInt(3, 1))
}

func TestMinMaxInt(t *testing.T) {
	min, max := MinMaxInt()
	require.Equal(t, 0, min)
	require.Equal(t, 0, max)

	min, max = MinMaxInt(1)
	require.Equal(t, 1, min)
	require.Equal(t, 1, max)

	min, max = MinMaxInt(1, 3)
	require.Equal(t, 1, min)
	require.Equal(t, 3, max)

	min, max = MinMaxInt(3, 1)
	require.Equal(t, 1, min)
	require.Equal(t, 3, max)
}

func TestRangeInt(t *testing.T) {
	nums := make([]int, 0)
	for i := range RangeInt(1, 1, 1) {
		nums = append(nums, i)
	}
	require.Equal(t, []int{1}, nums)

	nums = make([]int, 0)
	for i := range RangeInt(1, 3, 1) {
		nums = append(nums, i)
	}
	require.Equal(t, []int{1, 2, 3}, nums)

	nums = make([]int, 0)
	for i := range RangeInt(3, 1, 1) {
		nums = append(nums, i)
	}
	require.Equal(t, []int{3, 2, 1}, nums)

	nums = make([]int, 0)
	for i := range RangeInt(1, 6, 2) {
		nums = append(nums, i)
	}
	require.Equal(t, []int{1, 3, 5}, nums)

	nums = make([]int, 0)
	for i := range RangeInt(6, 1, 2) {
		nums = append(nums, i)
	}
	require.Equal(t, []int{6, 4, 2}, nums)
}

func TestOnLineInt(t *testing.T) {
	require.True(t, OnLineInt(0, 0, 5, 5, 1, 1))
	require.True(t, OnLineInt(5, 5, 0, 0, 1, 1))
	require.True(t, OnLineInt(0, 0, 0, 5, 0, 1))
	require.True(t, OnLineInt(0, 5, 0, 0, 0, 1))
	require.True(t, OnLineInt(0, 0, 5, 0, 1, 0))
	require.True(t, OnLineInt(5, 0, 0, 0, 1, 0))
	require.False(t, OnLineInt(0, 0, 5, 5, 2, 3))
}

func TestBoolsAll(t *testing.T) {
	b := make(Bools, 3)
	require.True(t, b.All(false))
	b[0] = true
	require.False(t, b.All(false))
}

func TestBoolsAny(t *testing.T) {
	b := make(Bools, 3)
	require.True(t, b.Any(false))
	require.False(t, b.Any(true))
	b[0] = true
	require.True(t, b.Any(false))
	require.True(t, b.Any(true))
	b[1] = true
	b[2] = true
	require.False(t, b.Any(false))
	require.True(t, b.Any(true))
}

func TestStrToStrs(t *testing.T) {
	slice := StrToStrs("a")
	require.Equal(t, 1, len(slice))
	require.Equal(t, "a", slice[0])

	slice = StrToStrs("abc")
	require.Equal(t, 3, len(slice))
	require.Equal(t, "a", slice[0])
	require.Equal(t, "b", slice[1])
	require.Equal(t, "c", slice[2])
}

func TestStringSorter(t *testing.T) {
	testcases := map[string]string{
		"a":   "a",
		"abc": "abc",
		"zoa": "aoz",
	}
	for input, expected := range testcases {
		require.Equal(t, expected, StringSorter(input))
	}
}

func TestStrContainsAny(t *testing.T) {
	require.True(t, StrContainsAny("abc", "a"))
	require.True(t, StrContainsAny("abc", "a", "b"))
	require.True(t, StrContainsAny("abc", "d", "b"))
	require.False(t, StrContainsAny("abc", "d"))
}
