package day16

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
	"github.com/MarkusFreitag/advent-of-code/util/slice"
)

type Valve struct {
	Name    string
	Flow    int
	Tunnels map[string]*Valve
}

func parseInput(input string) map[string]*Valve {
	valves := make(map[string]*Valve)
	for _, line := range strings.Split(input, "\n") {
		fields := strings.Fields(line)
		r := strings.NewReplacer("rate=", "", ";", "")
		valve := &Valve{
			Name:    fields[1],
			Flow:    util.ParseInt(r.Replace(fields[4])),
			Tunnels: make(map[string]*Valve),
		}
		for _, field := range fields[9:] {
			valve.Tunnels[strings.TrimSuffix(field, ",")] = nil
		}
		valves[fields[1]] = valve
	}
	for _, valve := range valves {
		for name := range valve.Tunnels {
			valve.Tunnels[name] = valves[name]
		}
	}
	return valves
}

func search(valves map[string]*Valve, current *Valve, minutes int, opened []string) int {
	key := fmt.Sprintf("%s|%d|%s", current.Name, minutes, strings.Join(opened, ","))
	if v, ok := cache[key]; ok {
		return v
	}

	if minutes <= 0 {
		cache[key] = 0
		return 0
	}

	var max int
	for _, valve := range current.Tunnels {
		max = numbers.Max(max, search(valves, valve, minutes-1, opened))
	}

	if !slice.Contains(opened, current.Name) && current.Flow > 0 && minutes > 0 {
		newOpened := make([]string, 0)
		newOpened = append(newOpened, opened...)
		newOpened = append(newOpened, current.Name)
		minutes--
		sum := minutes * current.Flow

		for _, valve := range current.Tunnels {
			max = numbers.Max(max, sum+search(valves, valve, minutes-1, newOpened))
		}
	}

	cache[key] = max
	return max
}

var cache = make(map[string]int)

func Part1(input string) (string, error) {
	cache = make(map[string]int)
	valves := parseInput(input)
	return strconv.Itoa(search(valves, valves["AA"], 30, make([]string, 0))), nil
}

func Part2(input string) (string, error) {
	return "not solved yet", nil
}
