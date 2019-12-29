package day16

import (
	"strconv"
	"strings"
	"sync"
)

var basePattern = []int{0, 1, 0, -1}

func calcPVal(a, b int) int {
	idx := ((a + 1) / (b + 1)) % len(basePattern)
	return basePattern[idx]
}

type Response struct {
	Index int
	Value int
}

func IntSToStrS(nums []int) []string {
	s := make([]string, len(nums))
	for idx, num := range nums {
		s[idx] = strconv.Itoa(num)
	}
	return s
}

func StrToIntS(str string) []int {
	nums := make([]int, len(str))
	for idx, s := range str {
		num, _ := strconv.Atoi(string(s))
		nums[idx] = num
	}
	return nums
}

func Part1(input string) (string, error) {
	numbers := StrToIntS(input)

	calc := func(nums []int, index int, rChan chan Response, wg *sync.WaitGroup) {
		var sum int
		for idx, num := range nums {
			sum += num * calcPVal(idx, index)
		}
		if sum < 0 {
			sum *= -1
		}
		rChan <- Response{Index: index, Value: sum % 10}
		wg.Done()
	}

	for phase := 1; phase <= 100; phase++ {
		responses := make(chan Response, len(numbers))
		var wg sync.WaitGroup

		for i := 0; i < len(numbers); i++ {
			wg.Add(1)
			go calc(numbers, i, responses, &wg)
		}
		wg.Wait()
		close(responses)
		for resp := range responses {
			numbers[resp.Index] = resp.Value
		}
	}

	return strings.Join(IntSToStrS(numbers[:8]), ""), nil
}

func Part2(input string) (string, error) {
	return "n/a", nil
}
