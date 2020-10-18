package day13

import (
	"strconv"
	"strings"
)

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

func mapKeys(m map[string]map[string]int) []string {
	keys := make([]string, 0)
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Part1(input string) (string, error) {
	guests := make(map[string]map[string]int)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		person := parts[0]
		neighbor := strings.TrimSuffix(parts[len(parts)-1], ".")
		score, err := strconv.Atoi(parts[3])
		if err != nil {
			return "", err
		}
		if parts[2] == "lose" {
			score = score * -1
		}
		if _, ok := guests[person]; !ok {
			guests[person] = make(map[string]int)
		}
		guests[person][neighbor] = score
	}

	combinations := make([][]string, 0)
	Perm(mapKeys(guests), func(l []string) {
		tmp := make([]string, len(l))
		copy(tmp, l)
		combinations = append(combinations, tmp)
	})

	var highest int
	for _, comb := range combinations {
		var score int
		for idx, p := range comb {
			var leftNeigh, rightNeigh string
			if idx == 0 {
				leftNeigh = comb[len(comb)-1]
				rightNeigh = comb[1]
			} else if idx == len(comb)-1 {
				leftNeigh = comb[idx-1]
				rightNeigh = comb[0]
			} else {
				leftNeigh = comb[idx-1]
				rightNeigh = comb[idx+1]
			}
			score += guests[p][leftNeigh]
			score += guests[p][rightNeigh]
		}

		if score > highest {
			highest = score
		}
	}

	return strconv.Itoa(highest), nil
}

func Part2(input string) (string, error) {
	guests := make(map[string]map[string]int)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		person := parts[0]
		neighbor := strings.TrimSuffix(parts[len(parts)-1], ".")
		score, err := strconv.Atoi(parts[3])
		if err != nil {
			return "", err
		}
		if parts[2] == "lose" {
			score = score * -1
		}
		if _, ok := guests[person]; !ok {
			guests[person] = make(map[string]int)
		}
		guests[person][neighbor] = score
	}

	guests["santa"] = make(map[string]int)
	for _, name := range mapKeys(guests) {
		guests[name]["santa"] = 0
		guests["santa"][name] = 0
	}

	combinations := make([][]string, 0)
	Perm(mapKeys(guests), func(l []string) {
		tmp := make([]string, len(l))
		copy(tmp, l)
		combinations = append(combinations, tmp)
	})

	var highest int
	for _, comb := range combinations {
		var score int
		for idx, p := range comb {
			var leftNeigh, rightNeigh string
			if idx == 0 {
				leftNeigh = comb[len(comb)-1]
				rightNeigh = comb[1]
			} else if idx == len(comb)-1 {
				leftNeigh = comb[idx-1]
				rightNeigh = comb[0]
			} else {
				leftNeigh = comb[idx-1]
				rightNeigh = comb[idx+1]
			}
			score += guests[p][leftNeigh]
			score += guests[p][rightNeigh]
		}

		if score > highest {
			highest = score
		}
	}

	return strconv.Itoa(highest), nil
}
