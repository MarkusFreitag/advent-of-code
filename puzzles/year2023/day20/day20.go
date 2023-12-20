package day20

import (
	"errors"
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/MarkusFreitag/advent-of-code/util/maputil"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
	"github.com/MarkusFreitag/advent-of-code/util/sliceutil"
)

type Module interface {
	Name() string
	AddDestination(string)
	Process(item) []item
}

type Broadcast struct {
	name         string
	destinations []string
}

func (bc *Broadcast) Name() string { return bc.name }
func (bc *Broadcast) AddDestination(name string) {
	bc.destinations = append(bc.destinations, name)
}
func (bc *Broadcast) Process(i item) []item {
	items := make([]item, len(bc.destinations))
	for idx, dest := range bc.destinations {
		items[idx] = item{from: bc.name, to: dest, state: i.state}
	}
	return items
}

type FlipFlop struct {
	name         string
	state        bool
	destinations []string
}

func (ff *FlipFlop) Name() string { return ff.name }
func (ff *FlipFlop) AddDestination(name string) {
	ff.destinations = append(ff.destinations, name)
}
func (ff *FlipFlop) Process(i item) []item {
	if i.state {
		return nil
	}
	ff.state = !ff.state

	items := make([]item, len(ff.destinations))
	for idx, dest := range ff.destinations {
		items[idx] = item{from: ff.name, to: dest, state: ff.state}
	}
	return items
}

type Conjunction struct {
	name         string
	inputs       map[string]bool
	destinations []string
}

func (cj *Conjunction) Name() string { return cj.name }
func (cj *Conjunction) AddDestination(name string) {
	cj.destinations = append(cj.destinations, name)
}
func (cj *Conjunction) Process(i item) []item {
	cj.inputs[i.from] = i.state
	pulse := !sliceutil.All(maputil.Values(cj.inputs), true)

	items := make([]item, len(cj.destinations))
	for idx, dest := range cj.destinations {
		items[idx] = item{from: cj.name, to: dest, state: pulse}
	}
	return items
}

type item struct {
	from  string
	to    string
	state bool
}

func (i item) Count(low, high *atomic.Int64) {
	if i.state {
		high.Add(1)
	} else {
		low.Add(1)
	}
}

func parseInput(input string) map[string]Module {
	modules := make(map[string]Module)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Fields(line)
		var mod Module
		if strings.HasPrefix(parts[0], "%") {
			mod = &FlipFlop{
				name: strings.TrimPrefix(parts[0], "%"),
			}
		} else if strings.HasPrefix(parts[0], "&") {
			mod = &Conjunction{
				name:   strings.TrimPrefix(parts[0], "&"),
				inputs: make(map[string]bool),
			}
		} else {
			mod = &Broadcast{
				name: parts[0],
			}
		}
		modules[mod.Name()] = mod
	}

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Fields(line)

		name := strings.TrimPrefix(strings.TrimPrefix(parts[0], "%"), "&")

		mod := modules[name]
		for _, dest := range parts[2:] {
			dest = strings.TrimSuffix(dest, ",")
			mod.AddDestination(dest)

			if m, ok := modules[dest]; ok {
				if conjunction, ok := m.(*Conjunction); ok {
					conjunction.inputs[mod.Name()] = false
				}
			}
		}
	}

	return modules
}

func Part1(input string) (string, error) {
	modules := parseInput(input)

	var low, high atomic.Int64
	for idx := 0; idx < 1000; idx++ {
		queue := make([]item, 0)
		queue = append(queue, item{from: "button", to: "broadcaster", state: false})

		for len(queue) > 0 {
			var i item
			i, queue = sliceutil.PopFront(queue)
			i.Count(&low, &high)

			m, ok := modules[i.to]
			if !ok {
				continue
			}
			items := m.Process(i)
			if items != nil {
				queue = append(queue, items...)
			}
		}
	}

	return strconv.Itoa(int(low.Load()) * int(high.Load())), nil
}

func searchSources(modules map[string]Module, name string) []Module {
	sources := make([]Module, 0)
	for _, mod := range modules {
		switch v := mod.(type) {
		case *Broadcast:
			for _, dest := range v.destinations {
				if dest == name {
					sources = append(sources, v)
					break
				}
			}
		case *FlipFlop:
			for _, dest := range v.destinations {
				if dest == name {
					sources = append(sources, v)
					break
				}
			}
		case *Conjunction:
			for _, dest := range v.destinations {
				if dest == name {
					sources = append(sources, v)
					break
				}
			}
		}
	}
	return sources
}

func Part2(input string) (string, error) {
	modules := parseInput(input)

	var rxSource *Conjunction
	if sources := searchSources(modules, "rx"); len(sources) == 1 {
		rxSource = sources[0].(*Conjunction)
	} else {
		return "", errors.New("expected only one source for module named 'rx'")
	}

	sources := searchSources(modules, rxSource.name)

	targets := make(map[string]int)

	for round := 1; len(targets) != len(sources); round++ {
		queue := make([]item, 0)
		queue = append(queue, item{from: "button", to: "broadcaster", state: false})

		for len(queue) > 0 {
			var i item
			i, queue = sliceutil.PopFront(queue)

			m, ok := modules[i.to]
			if !ok {
				continue
			}
			items := m.Process(i)
			if items != nil {
				queue = append(queue, items...)
			}

			for _, mod := range sources {
				_, ok := targets[mod.Name()]
				if !ok && rxSource.inputs[mod.Name()] {
					targets[mod.Name()] = round
				}
			}
		}
	}

	values := maputil.Values(targets)
	return strconv.Itoa(numbers.LCM(values[0], values[1], values[2:]...)), nil
}
