package day10

import (
	"sort"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

func Part1(input string) (string, error) {
	adapters := util.StringsToInts(strings.Split(input, "\n"))
	sort.Ints(adapters)

	var one, three int
	var current int
	builtin := adapters[len(adapters)-1] + 3

	for _, adap := range adapters {
		switch adap - current {
		case 1:
			one++
		case 3:
			three++
		default:
			continue
		}
		current = adap
	}

	if builtin-current != 3 {
		return "invalid chain", nil
	} else {
		three++
	}

	return strconv.Itoa(one * three), nil
}

type Node struct {
	ID   int
	Next map[*Node]bool
}

func Part2(input string) (string, error) {
	adapters := util.StringsToInts(strings.Split(input, "\n"))
	sort.Ints(adapters)

	start := &Node{ID: 0, Next: make(map[*Node]bool)}
	end := &Node{ID: adapters[len(adapters)-1] + 3, Next: make(map[*Node]bool)}

	nodes := make([]*Node, 0)
	nodes = append(nodes, start)
	for _, adapter := range adapters {
		node := &Node{ID: adapter, Next: make(map[*Node]bool)}
		for _, n := range nodes {
			distance := node.ID - n.ID
			if distance >= 1 && distance <= 3 {
				n.Next[node] = true
			}
		}
		nodes = append(nodes, node)
	}

	for _, n := range nodes {
		distance := end.ID - n.ID
		if distance >= 1 && distance <= 3 {
			n.Next[end] = true
		}
	}

	var counter int
	for node := range start.Next {
		counter += path(node, end)
	}
	return strconv.Itoa(counter), nil
}

var state = make(map[int]int)

func path(node *Node, dest *Node) int {
	if node.ID == dest.ID {
		return 1
	}
	if v, ok := state[node.ID]; ok {
		return v
	}
	var counter int
	for n := range node.Next {
		counter += path(n, dest)
	}
	state[node.ID] = counter
	return counter
}
