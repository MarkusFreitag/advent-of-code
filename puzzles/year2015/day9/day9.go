package day9

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

type route struct {
	start    string
	end      string
	distance int
}

// Perm calls f with each permutation of a.
func Perm(a []string, f func([]string)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []string, f func([]string), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func Part1(input string) (string, error) {
	locs := make([]string, 0)
	routes := make([]route, 0)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		dis, _ := strconv.Atoi(parts[4])
		r := route{
			start:    parts[0],
			end:      parts[2],
			distance: dis,
		}
		routes = append(routes, r)
		if !slices.Contains(locs, r.start) {
			locs = append(locs, r.start)
		}
		if !slices.Contains(locs, r.end) {
			locs = append(locs, r.end)
		}
	}

	combinations := make([][]string, 0)
	Perm(locs, func(l []string) {
		tmp := make([]string, len(l))
		copy(tmp, l)
		combinations = append(combinations, tmp)
	})

	shortest := math.MaxInt64
	for _, comb := range combinations {
		var distance int
		for i := 0; i < len(comb)-1; i++ {
			for _, r := range routes {
				if (r.start == comb[i] && r.end == comb[i+1]) || (r.start == comb[i+1] && r.end == comb[i]) {
					distance += r.distance
					break
				}
			}
		}
		if distance < shortest {
			shortest = distance
		}
	}
	return strconv.Itoa(shortest), nil
}

func Part2(input string) (string, error) {
	locs := make([]string, 0)
	routes := make([]route, 0)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		dis, _ := strconv.Atoi(parts[4])
		r := route{
			start:    parts[0],
			end:      parts[2],
			distance: dis,
		}
		routes = append(routes, r)
		if !slices.Contains(locs, r.start) {
			locs = append(locs, r.start)
		}
		if !slices.Contains(locs, r.end) {
			locs = append(locs, r.end)
		}
	}

	combinations := make([][]string, 0)
	Perm(locs, func(l []string) {
		tmp := make([]string, len(l))
		copy(tmp, l)
		combinations = append(combinations, tmp)
	})

	var longest int
	for _, comb := range combinations {
		var distance int
		for i := 0; i < len(comb)-1; i++ {
			for _, r := range routes {
				if (r.start == comb[i] && r.end == comb[i+1]) || (r.start == comb[i+1] && r.end == comb[i]) {
					distance += r.distance
					break
				}
			}
		}
		if distance > longest {
			longest = distance
		}
	}
	return strconv.Itoa(longest), nil
}
