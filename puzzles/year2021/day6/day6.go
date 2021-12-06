package day6

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

type Pool map[int]int

func NewPool(fish []int) Pool {
	p := make(Pool)
	for _, f := range fish {
		c, _ := p[f]
		c++
		p[f] = c
	}
	return p
}

func (p Pool) add(fish, count int) {
	v, _ := p[fish]
	v += count
	p[fish] = v
}

func (p Pool) Evolution() Pool {
	newPool := make(Pool)
	for fish, count := range p {
		if fish == 0 {
			newPool.add(6, count)

			newPool.add(8, count)
		} else {
			newPool.add(fish-1, count)
		}
	}
	return newPool
}

func (p Pool) Population() int {
	var count int
	for _, c := range p {
		count += c
	}
	return count
}

func Part1(input string) (string, error) {
	nums := util.StrsToInts(strings.Split(input, ","))
	pool := NewPool(nums)
	for d := 0; d < 80; d++ {
		pool = pool.Evolution()
	}
	return strconv.Itoa(pool.Population()), nil
}

func Part2(input string) (string, error) {
	nums := util.StrsToInts(strings.Split(input, ","))
	pool := NewPool(nums)
	for d := 0; d < 256; d++ {
		pool = pool.Evolution()
	}
	return strconv.Itoa(pool.Population()), nil
}
