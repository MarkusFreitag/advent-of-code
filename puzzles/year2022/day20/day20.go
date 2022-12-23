package day20

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/slice"
)

/*
1 2 -3 10 -9 0 4

2 1 -3 10 -9 0 4
1 -3 2 10 -9 0 4
1 2 10 -9 -3 0 4
1 2 -9 -3 0 10 4
1 2 -3 0 10 4 -9
1 2 -3 0 10 4 -9
1 2 -3 4 0 10 -9
*/

func move(nums []int, num int, forward bool) []int {
	length := len(nums)
	index := slice.Index(nums, num)
	if index == -1 {
		panic("num not found")
	}
	if forward {
		if index == length-1 {
			nums = append(nums[:index], nums[index+1:]...)
			nums = append(nums[:1], append([]int{num}, nums[1:]...)...)
		} else {
			nums[index+1], nums[index] = nums[index], nums[index+1]
		}
		return nums
	}
	if index == 0 {
		_, nums = slice.PopFront(nums)
		var last int
		last, nums = slice.Pop(nums)
		nums = append(nums, num)
		nums = append(nums, last)
	} else {
		nums[index-1], nums[index] = nums[index], nums[index-1]
	}
	return nums
}

func mix(nums []int) {
	for _, num := range slice.Copy(nums) {
		index := slice.Index(nums, num)

		// remove num from slice
		//nums = append(nums[:index], nums[index+1:]...)
		_, nums = slice.PopIndex(nums, index)

		// calc new index
		newIndex := (((index + num) % len(nums)) + len(nums)) % (len(nums))
		/*
			newIndex := (index + num) % len(nums)
			if newIndex < 0 {
				newIndex += len(nums)
			}
		*/

		// insert num into slice
		//nums = append(nums[:newIndex], append([]int{num}, nums[newIndex:]...)...)
		nums = slice.Insert(nums, num, newIndex)
		fmt.Println(nums)
	}
}

func lookup(nums []int, index int) int {
	length := len(nums)
	if index < length {
		return nums[index]
	}
	return nums[index%length]
}

func Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")
	nums := make([]int, len(lines))
	for idx, line := range lines {
		nums[idx] = util.ParseInt(line)
	}
	mix(nums)
	for idx, num := range nums {
		if num == 0 {
			var sum int
			sum += lookup(nums, 1000+idx)
			sum += lookup(nums, 2000+idx)
			sum += lookup(nums, 3000+idx)
			return strconv.Itoa(sum), nil
		}
	}

	return "n/a", nil
}

func Part2(input string) (string, error) {
	return "not solved yet", nil
}
