package day13

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Bus struct {
	ID     int
	Arrive int
}

type byID []*Bus

func (s byID) Len() int {
	return len(s)
}
func (s byID) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byID) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")
	t, _ := strconv.Atoi(lines[0])
	buses := make([]*Bus, 0)
	for _, b := range strings.Split(lines[1], ",") {
		if b == "x" {
			continue
		}
		num, _ := strconv.Atoi(b)
		buses = append(buses, &Bus{ID: num, Arrive: num})
	}
	sort.Sort(byID(buses))

	for {
		for _, bus := range buses {
			if bus.ID == 787 {
				continue
			}
			if bus.Arrive > t {
				fmt.Printf("%#v\n", *bus)
				return strconv.Itoa(bus.ID * (bus.Arrive - t)), nil
			}
			bus.Arrive += bus.ID
		}
	}

	return "n/a", nil
}

func Part2(input string) (string, error) {
	lines := strings.Split(input, "\n")
	buses := make([]*Bus, 0)
	for _, b := range strings.Split(lines[1], ",") {
		if b == "x" {
			buses = append(buses, nil)
		}
		num, _ := strconv.Atoi(b)
		buses = append(buses, &Bus{ID: num, Arrive: num})
	}

	for {
		found := true
		t := buses[0].Arrive
		for idx, bus := range buses {
			if bus == nil {
				continue
			}
			found = t+idx == bus.Arrive
		}
		if found {
			return strconv.Itoa(buses[0].Arrive), nil
		}

		for _, bus := range buses {
			if bus == nil {
				continue
			}
			bus.Arrive += bus.ID
		}
	}

	return "n/a", nil
}
