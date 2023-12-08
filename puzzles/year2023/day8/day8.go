package day8

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util/maputil"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
	"github.com/MarkusFreitag/advent-of-code/util/sliceutil"
)

type Node struct {
	Name  string
	Left  *Node
	Right *Node
}

func parseInput(input string) ([]rune, map[string]*Node) {
	nodes := make(map[string]*Node)
	blocks := strings.Split(input, "\n\n")
	for _, line := range strings.Split(blocks[1], "\n") {
		parts := strings.Split(line, "=")

		name := strings.TrimSpace(parts[0])
		node, ok := nodes[name]
		if !ok {
			node = &Node{Name: name}
			nodes[name] = node
		}

		nameLeft := strings.TrimPrefix(strings.Split(parts[1], ",")[0], " (")
		nodeLeft, ok := nodes[nameLeft]
		if !ok {
			nodeLeft = &Node{Name: nameLeft}
			nodes[nameLeft] = nodeLeft
		}
		node.Left = nodeLeft

		nameRight := strings.TrimSpace(strings.TrimSuffix(strings.Split(parts[1], ",")[1], ")"))
		nodeRight, ok := nodes[nameRight]
		if !ok {
			nodeRight = &Node{Name: nameRight}
			nodes[nameRight] = nodeRight
		}
		node.Right = nodeRight
	}
	return []rune(strings.TrimSpace(blocks[0])), nodes
}

func Part1(input string) (string, error) {
	moves, nodes := parseInput(input)

	var steps int
	node := nodes["AAA"]
	var idx int
	for node.Name != "ZZZ" {
		switch moves[idx] {
		case 'L':
			node = node.Left
		case 'R':
			node = node.Right
		}
		idx = (idx + 1) % len(moves)
		steps++
	}
	return strconv.Itoa(steps), nil
}

func destinationReached(n *Node) bool {
	return strings.HasSuffix(n.Name, "Z")
}

func Part2(input string) (string, error) {
	moves, nodes := parseInput(input)

	cNodes := maputil.ValuesFiltered(nodes, func(n *Node) bool {
		return strings.HasSuffix(n.Name, "A")
	})

	var steps int
	nSteps := make([]int, len(cNodes))
	var idx int
	for !sliceutil.AllFunc(cNodes, destinationReached) {
		for i, node := range cNodes {
			if strings.HasSuffix(node.Name, "Z") {
				continue
			}
			switch moves[idx] {
			case 'L':
				cNodes[i] = node.Left
			case 'R':
				cNodes[i] = node.Right
			}
			if strings.HasSuffix(cNodes[i].Name, "Z") {
				nSteps[i] = steps + 1
			}
		}
		idx = (idx + 1) % len(moves)
		steps++
	}
	return strconv.Itoa(numbers.LCM(nSteps[0], nSteps[1], nSteps[2:]...)), nil
}
