package numbers_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

func TestSum(t *testing.T) {
	require.Equal(t, 0, numbers.Sum(0))
	require.Equal(t, 1, numbers.Sum(1))

	require.Equal(t, 5, numbers.Sum(2, 3))
	require.Equal(t, -1, numbers.Sum(2, -3))

	require.Equal(t, 5.300000000000001, numbers.Sum(2.2, 3.1))
	require.Equal(t, -0.8999999999999999, numbers.Sum(2.1, -3.0))
}

func TestSubtract(t *testing.T) {
	require.Equal(t, 0, numbers.Subtract(0))
	require.Equal(t, 1, numbers.Subtract(1))

	require.Equal(t, -1, numbers.Subtract(2, 3))
	require.Equal(t, -6, numbers.Subtract(-3, 3))
	require.Equal(t, 1, numbers.Subtract(-2, -3))

	require.Equal(t, -0.5, numbers.Subtract(2.5, 3.0))
	require.Equal(t, -6.2, numbers.Subtract(-3.1, 3.1))
	require.Equal(t, 1.1, numbers.Subtract(-2.5, -3.6))
}

func TestMultiply(t *testing.T) {
	require.Equal(t, 0, numbers.Multiply(0))
	require.Equal(t, 1, numbers.Multiply(1))

	require.Equal(t, 6, numbers.Multiply(2, 3))
	require.Equal(t, 0, numbers.Multiply(2, 0, 3))

	require.Equal(t, 7.040000000000001, numbers.Multiply(2.2, 3.2))
	require.Equal(t, 0.0, numbers.Multiply(2.1, 0.0, 3.1))
}

func TestPow(t *testing.T) {
	require.Equal(t, 8, numbers.Pow(2, 3))
	require.Equal(t, 9, numbers.Pow(3, 2))

	require.Equal(t, 8.0, numbers.Pow(2.0, 3.0))
	require.Equal(t, 15.588457268119894, numbers.Pow(3.0, 2.5))
}

func TestAbs(t *testing.T) {
	require.Equal(t, 1, numbers.Abs(1))
	require.Equal(t, 2, numbers.Abs(-2))

	require.Equal(t, 1.1, numbers.Abs(1.1))
	require.Equal(t, 2.1, numbers.Abs(-2.1))
}

func TestBetween(t *testing.T) {
	require.True(t, numbers.Between(1, 1, 3))
	require.True(t, numbers.Between(2, 1, 3))
	require.True(t, numbers.Between(3, 1, 3))
	require.False(t, numbers.Between(4, 1, 3))
}

func TestMin(t *testing.T) {
	require.Equal(t, 1, numbers.Min(1))
	require.Equal(t, 1, numbers.Min(1, 3))
	require.Equal(t, 1, numbers.Min(3, 1))

	require.Equal(t, 1.1, numbers.Min(1.1))
	require.Equal(t, 1.1, numbers.Min(1.1, 1.2))
	require.Equal(t, 1.1, numbers.Min(1.2, 1.1))
}

func TestMax(t *testing.T) {
	require.Equal(t, 1, numbers.Max(1))
	require.Equal(t, 3, numbers.Max(1, 3))
	require.Equal(t, 3, numbers.Max(3, 1))

	require.Equal(t, 1.1, numbers.Max(1.1))
	require.Equal(t, 1.2, numbers.Max(1.1, 1.2))
	require.Equal(t, 1.2, numbers.Max(1.2, 1.1))
}

func TestMinMax(t *testing.T) {
	minInt, maxInt := numbers.MinMax(1)
	require.Equal(t, 1, minInt)
	require.Equal(t, 1, maxInt)

	minInt, maxInt = numbers.MinMax(1, 3)
	require.Equal(t, 1, minInt)
	require.Equal(t, 3, maxInt)

	minInt, maxInt = numbers.MinMax(3, 1)
	require.Equal(t, 1, minInt)
	require.Equal(t, 3, maxInt)

	minFloat, maxFloat := numbers.MinMax(1.1)
	require.Equal(t, 1.1, minFloat)
	require.Equal(t, 1.1, maxFloat)

	minFloat, maxFloat = numbers.MinMax(1.1, 1.2)
	require.Equal(t, 1.1, minFloat)
	require.Equal(t, 1.2, maxFloat)

	minFloat, maxFloat = numbers.MinMax(1.2, 1.1)
	require.Equal(t, 1.1, minFloat)
	require.Equal(t, 1.2, maxFloat)
}
