package day9

import (
  "strconv"
  "strings"
  "sort"

  "github.com/MarkusFreitag/advent-of-code/util"
)

func check(nums []int, num int) bool {
  for i:=0;i<len(nums);i++ {
    for j:=i+1;j<len(nums);j++ {
      if nums[i] + nums[j] == num {
        return true
      }
    }
  }
  return false
}

func findInvalid(nums []int, bound int) int {
  for i:=bound;i<len(nums);i++ {
    preamble := nums[i-bound:i]
    if !check(preamble, nums[i]) {
      return nums[i]
    }
  }
  return -1
}

var preambleSize = 25

func Part1(input string) (string, error) {
  strs := strings.Split(input, "\n")
  nums := util.StrsToInts(strs)
  return strconv.Itoa(findInvalid(nums, preambleSize)), nil
}

func Part2(input string) (string, error) {
  strs := strings.Split(input, "\n")
  nums := util.StrsToInts(strs)
  invalid := findInvalid(nums, preambleSize)

  for i:=0;i<len(nums);i++ {
    count := 1
    for {
      if i+count == len(nums) {
        break
      }
      part := nums[i:i+count]
      if util.SumInts(part...) == invalid {
        sort.Ints(part)
        return strconv.Itoa(part[0] + part[len(part)-1]), nil
      }
      count++
    }
  }

  return "n/a", nil
}
