package day7

import (
	"strconv"
	"strings"
)

type Bag struct {
	Color   string
	Content []*Bag
}

func (b *Bag) Count() int {
	count := len(b.Content)
	for _, bag := range b.Content {
		count += bag.Count()
	}
	return count
}

func (b *Bag) CanCarry(color string, bags map[string]bool) map[string]bool {
	if b.Color == color {
		return bags
	}
	for _, bag := range b.Content {
		bags = bag.CanCarry(color, bags)
		if bag.Color == color {
			bags[b.Color] = true
		}
		if _, ok := bags[bag.Color]; ok {
			bags[b.Color] = true
		}
	}
	return bags
}

func Part1(input string) (string, error) {
	bags := make(map[string]*Bag)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, "contain")
		color := strings.TrimSuffix(parts[0], " bags ")

		bag, ok := bags[color]
		if !ok {
			bag = &Bag{Color: color}
			bags[color] = bag
		}
		if parts[1] != " no other bags." {
			for _, item := range strings.Split(parts[1], ",") {
				item := strings.TrimSpace(item)
				strs := strings.Split(item, " ")
				color := strings.Join(strs[1:3], " ")
				b, ok := bags[color]
				if !ok {
					b = &Bag{Color: color}
					bags[color] = b
				}
				bag.Content = append(bag.Content, b)
			}
		}
	}

	carrier := make(map[string]bool)
	for _, bag := range bags {
		carrier = bag.CanCarry("shiny gold", carrier)
	}
	return strconv.Itoa(len(carrier)), nil
}

func Part2(input string) (string, error) {
	bags := make(map[string]*Bag)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, "contain")
		color := strings.TrimSuffix(parts[0], " bags ")

		bag, ok := bags[color]
		if !ok {
			bag = &Bag{Color: color}
			bags[color] = bag
		}
		if parts[1] != " no other bags." {
			for _, item := range strings.Split(parts[1], ",") {
				item := strings.TrimSpace(item)
				strs := strings.Split(item, " ")
				color := strings.Join(strs[1:3], " ")
				b, ok := bags[color]
				if !ok {
					b = &Bag{Color: color}
					bags[color] = b
				}
				count, _ := strconv.Atoi(strs[0])
				for i := 0; i < count; i++ {
					bag.Content = append(bag.Content, b)
				}
			}
		}
	}

	b := bags["shiny gold"]
	return strconv.Itoa(b.Count()), nil
}
