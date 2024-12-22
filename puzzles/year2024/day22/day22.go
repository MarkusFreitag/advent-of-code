package day22

import (
	"maps"
	"slices"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
	"github.com/MarkusFreitag/advent-of-code/util/sliceutil"
)

func mul(secret, n int) int {
	return prune(mix(secret, secret*n))
}

func div(secret, n int) int {
	return prune(mix(secret, secret/n))
}

func mix(a, b int) int {
	return a ^ b
}

func prune(secret int) int {
	return secret % 16777216
}

func evolve(secret int) int {
	secret = mul(secret, 64)
	secret = div(secret, 32)
	return mul(secret, 2048)
}

func secretPrize(secret int) int {
	return secret % 10
}

func Part1(input string) (string, error) {
	var sum int
	for _, line := range strings.Fields(input) {
		secret := util.ParseInt(line)
		for range 2000 {
			secret = evolve(secret)
		}
		sum += secret
	}
	return strconv.Itoa(sum), nil
}

func Part2(input string) (string, error) {
	total := make(map[[4]int]int)
	for _, line := range strings.Fields(input) {
		secret := util.ParseInt(line)
		lastPrize := secretPrize(secret)
		rounds := make([][2]int, 0)
		for range 2000 {
			secret = evolve(secret)
			prize := secretPrize(secret)
			rounds = append(rounds, [2]int{prize - lastPrize, prize})
			lastPrize = prize
		}
		seen := make(map[[4]int]struct{})
		for _, window := range sliceutil.SlidingWindow(rounds, 4) {
			if len(window) != 4 {
				break
			}
			var pattern [4]int
			for idx, item := range window {
				pattern[idx] = item[0]
			}
			if _, ok := seen[pattern]; !ok {
				seen[pattern] = struct{}{}
				total[pattern] = total[pattern] + window[3][1]
			}
		}
	}
	return strconv.Itoa(numbers.Max(slices.Collect(maps.Values(total))...)), nil
}
