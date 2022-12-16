package day16

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
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

func search(valves map[string]*Valve, current *Valve, opened map[string]*Valve, timeLeft, released int) int {
	fmt.Printf("%s: %v\n", current.Name, util.Keys(opened))
	tick := func(opened map[string]*Valve, timeLeft, released int) (int, int) {
		timeLeft--
		var sum int
		for _, valve := range opened {
			sum += valve.Flow
		}
		released = sum * timeLeft
		return timeLeft, released
	}

	// open current valve
	timeLeft, released = tick(opened, timeLeft, released)
	if timeLeft <= 0 {
		return released
	}

	opened[current.Name] = current

	options := make([]*Valve, 0)
	for _, valve := range current.Tunnels {
		if _, ok := opened[valve.Name]; ok {
			// skip if already opened before
			continue
		}
		options = append(options, valve)
	}
	sort.Slice(options, func(i, j int) bool {
		return options[i].Flow > options[j].Flow
	})

	results := make([]int, len(options)+1)
	results = append(results, released)
	for _, valve := range options {

		// move to valve
		timeLeft, released = tick(opened, timeLeft, released)
		if timeLeft <= 0 {
			return released
		}

		results = append(results, search(valves, valve, opened, timeLeft, released))
	}
	return numbers.Max(results...)
}

func Part1(input string) (string, error) {
	valves := parseInput(input)
	var max int
	for _, valve := range valves["AA"].Tunnels {
		if released := search(valves, valve, make(map[string]*Valve, 0), 30, 0); released > max {
			max = released
		}
	}
	return strconv.Itoa(max), nil
	//return strconv.Itoa(search(valves, valves["AA"], make(map[string]*Valve, 0), 30, 0)), nil
}

func Part2(input string) (string, error) {
	return "not solved yet", nil
}
