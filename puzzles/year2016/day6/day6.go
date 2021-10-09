package day6

import (
	"sort"
	"strings"
)

type char struct {
	r rune
	c int
}

type chars []*char

func (c chars) Len() int      { return len(c) }
func (c chars) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c chars) Less(i, j int) bool {
	if c[i].c == c[j].c {
		return int(c[i].r) < int(c[j].r)
	}
	return c[i].c > c[j].c
}

func (c chars) Contains(r rune) int {
	for idx, char := range c {
		if char.r == r {
			return idx
		}
	}
	return -1
}

func calcCharDistributions(str string) []chars {
	dist := make([]chars, 0)
	for _, line := range strings.Split(str, "\n") {
		for idx, r := range line {
			if idx >= len(dist) {
				dist = append(dist, make(chars, 0))
			}
			if i := dist[idx].Contains(r); i > -1 {
				dist[idx][i].c++
			} else {
				dist[idx] = append(dist[idx], &char{r: r, c: 1})
			}
		}
	}
	return dist
}

func Part1(input string) (string, error) {
	distributions := calcCharDistributions(input)
	var msg string
	for _, dist := range distributions {
		sort.Sort(dist)
		msg += string(dist[0].r)
	}
	return msg, nil
}

func Part2(input string) (string, error) {
	distributions := calcCharDistributions(input)
	var msg string
	for _, dist := range distributions {
		sort.Sort(dist)
		msg += string(dist[len(dist)-1].r)
	}
	return msg, nil
}
