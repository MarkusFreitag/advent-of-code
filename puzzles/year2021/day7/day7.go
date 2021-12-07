package day7

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

func travel(crabs []int, fuelCost func(int) int) int {
	minFuel := util.MaxInteger
	min, max := util.MinMaxInt(crabs...)
	for i := min; i <= max; i++ {
		var f int
		for _, crab := range crabs {
			f += fuelCost(util.Abs(i - crab))
		}
		minFuel = util.MinInt(minFuel, f)
	}
	return minFuel
}

func Part1(input string) (string, error) {
	parts := strings.Split(input, ",")
	fuel := travel(util.StrsToInts(parts), func(dist int) int {
		return dist
	})
	return strconv.Itoa(fuel), nil
}

func Part2(input string) (string, error) {
	parts := strings.Split(input, ",")
	fuel := travel(util.StrsToInts(parts), func(dist int) int {
		return (dist*dist + dist) / 2
	})
	return strconv.Itoa(fuel), nil
}
