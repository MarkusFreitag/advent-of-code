package numbers

import (
	"math"
	"sort"

	"github.com/MarkusFreitag/advent-of-code/util/constraints"
)

const (
	MaxUnsignedInteger = ^uint(0)
	MinUnsignedInteger = 0
	MaxInteger         = int(MaxUnsignedInteger >> 1)
	MinInteger         = -MaxInteger - 1
)

func Sum[T constraints.Numbers](nums ...T) T {
	if len(nums) == 0 {
		return 0
	}
	total := nums[0]
	for _, num := range nums[1:] {
		total += num
	}
	return total
}

func Subtract[T constraints.Numbers](nums ...T) T {
	if len(nums) == 0 {
		return 0
	}
	total := nums[0]
	for _, num := range nums[1:] {
		total -= num
	}
	return total
}

func Multiply[T constraints.Numbers](nums ...T) T {
	if len(nums) == 0 {
		return 0
	}
	total := nums[0]
	for _, num := range nums[1:] {
		total *= num
	}
	return total
}

func Pow[T constraints.Numbers](a, b T) T {
	return T(math.Pow(float64(a), float64(b)))
}

func Abs[T constraints.Float | constraints.Signed](i T) T {
	if i < T(0) {
		i *= T(-1)
	}
	return i
}

func Between[T constraints.Numbers](i, min, max T) bool {
	return min <= i && i <= max
}

func Min[T constraints.Numbers](nums ...T) T {
	min, _ := MinMax(nums...)
	return min
}

func Max[T constraints.Numbers](nums ...T) T {
	_, max := MinMax(nums...)
	return max
}

func MinMax[T constraints.Numbers](nums ...T) (T, T) {
	if len(nums) == 0 {
		return T(0), T(0)
	}
	min, max := nums[0], nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

func MinN[T constraints.Numbers](n int, nums ...T) []T {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	if len(nums) <= n {
		return nums
	}
	return nums[:n]
}

func MaxN[T constraints.Numbers](n int, nums ...T) []T {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	if len(nums) <= n {
		return nums
	}
	return nums[:n]
}
