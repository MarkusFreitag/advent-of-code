package day23

import (
	"slices"
	"strconv"
	"strings"
)

func parseInput(input string) map[string]map[string]struct{} {
	computers := make(map[string]map[string]struct{})
	for _, line := range strings.Fields(input) {
		parts := strings.Split(line, "-")
		conns, ok := computers[parts[0]]
		if !ok {
			conns = make(map[string]struct{})
		}
		conns[parts[1]] = struct{}{}
		computers[parts[0]] = conns

		conns, ok = computers[parts[1]]
		if !ok {
			conns = make(map[string]struct{})
		}
		conns[parts[0]] = struct{}{}
		computers[parts[1]] = conns
	}
	return computers
}

func Part1(input string) (string, error) {
	computers := parseInput(input)
	triplets := make(map[string]struct{})
	for computer, connections := range computers {
		for conn1 := range connections {
			for conn2 := range connections {
				if conn1 == conn2 {
					continue
				}
				_, ok1 := computers[conn1][conn2]
				_, ok2 := computers[conn2][conn1]
				if ok1 && ok2 {
					triplet := []string{computer, conn1, conn2}
					slices.Sort(triplet)
					triplets[strings.Join(triplet, ",")] = struct{}{}
				}
			}
		}
	}
	var count int
	for triplet := range triplets {
		for _, pc := range strings.Split(triplet, ",") {
			if strings.HasPrefix(pc, "t") {
				count++
				break
			}
		}
	}
	return strconv.Itoa(count), nil
}

func Part2(input string) (string, error) {
	computers := parseInput(input)
	var largest []string
	for computer, connections := range computers {
		party := make([]string, 1)
		party[0] = computer

		for connectedComputer := range connections {
			connected := true
			for _, pc := range party {
				if _, ok := computers[connectedComputer][pc]; !ok {
					connected = false
					break
				}
			}
			if connected {
				party = append(party, connectedComputer)
			}
		}

		if len(party) > len(largest) {
			largest = party
		}
	}
	slices.Sort(largest)
	return strings.Join(largest, ","), nil
}
