package day10

import (
	"sort"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

type Component interface {
	ID() int
	Type() string
	AddValue(int)
	Values() []int
	SetLow(Component)
	SetHigh(Component)
}

type Output struct {
	id      int
	value   int
	written bool
}

func NewOutput(id int) *Output {
	return &Output{id: id}
}

func (o *Output) ID() int { return o.id }

func (o *Output) Type() string { return "output" }

func (o *Output) AddValue(v int) {
	o.value = v
	o.written = true
}

func (o *Output) Values() []int {
	values := make([]int, 1)
	values[0] = o.value
	return values
}

func (o *Output) SetLow(_ Component) {}

func (o *Output) SetHigh(_ Component) {}

type Bot struct {
	id     int
	values []int
	low    Component
	high   Component
}

func NewBot(id int) *Bot {
	return &Bot{
		id:     id,
		values: make([]int, 0),
	}
}

func (b *Bot) calc() {
	if b.low != nil && b.high != nil && len(b.values) == 2 {
		if b.values[0] > b.values[1] {
			b.high.AddValue(b.values[0])
			b.low.AddValue(b.values[1])
		} else {
			b.high.AddValue(b.values[1])
			b.low.AddValue(b.values[0])
		}
	}
}

func (b *Bot) ID() int { return b.id }

func (b *Bot) Type() string { return "bot" }

func (b *Bot) AddValue(v int) {
	b.values = append(b.values, v)
	b.calc()
}

func (b *Bot) Values() []int {
	return b.values
}

func (b *Bot) SetLow(c Component) {
	b.low = c
	b.calc()
}

func (b *Bot) SetHigh(c Component) {
	b.high = c
	b.calc()
}

type Components struct {
	items []Component
}

func (c *Components) Get(id int, t string) Component {
	for _, item := range c.items {
		if item.Type() == t && item.ID() == id {
			return item
		}
	}
	var com Component
	switch t {
	case "bot":
		com = NewBot(id)
	case "output":
		com = NewOutput(id)
	}
	c.items = append(c.items, com)
	return com
}

func parseInstructions(input string) *Components {
	components := &Components{items: make([]Component, 0)}
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "value":
			v, _ := strconv.Atoi(parts[1])
			id, _ := strconv.Atoi(parts[5])
			comp := components.Get(id, "bot")
			comp.AddValue(v)
		case "bot":
			botID, _ := strconv.Atoi(parts[1])
			lowID, _ := strconv.Atoi(parts[6])
			highID, _ := strconv.Atoi(parts[11])

			lowC := components.Get(lowID, parts[5])
			highC := components.Get(highID, parts[10])

			botC := components.Get(botID, "bot")
			botC.SetLow(lowC)
			botC.SetHigh(highC)
		}
	}
	return components
}

func Part1(input string) (string, error) {
	components := parseInstructions(input)
	for _, com := range components.items {
		if com.Type() != "bot" {
			continue
		}
		values := com.Values()
		sort.Ints(values)
		if len(values) == 2 && values[0] == 17 && values[1] == 61 {
			return strconv.Itoa(com.ID()), nil
		}
	}
	return "not solved yet", nil
}

func Part2(input string) (string, error) {
	components := parseInstructions(input)
	outs := make([]int, 0)
	for _, com := range components.items {
		if com.Type() != "output" {
			continue
		}
		if com.ID() == 0 || com.ID() == 1 || com.ID() == 2 {
			outs = append(outs, com.Values()[0])
		}
	}
	return strconv.Itoa(util.MulInts(outs...)), nil
}
