package puzzles

import (
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	Name   string
	Parent *node
}

func newNode(name string, parent *node) *node {
	return &node{
		Name:   name,
		Parent: parent,
	}
}

func (n *node) orbits(d int) int {
	if n.Parent == nil {
		return d
	}
	return n.Parent.orbits(d + 1)
}

func (n *node) transfersTo(target string) int {
	var count int
	parent := n.Parent
	for parent.Name != target {
		count++
		parent = parent.Parent
	}
	return count
}

func parents(node *node) nodes {
	parents := make(nodes, 0)
	n := node
	for n.Parent != nil {
		parents = append(parents, n.Parent)
		n = n.Parent
	}
	return parents
}

type nodes []*node

func (ns nodes) search(name string) *node {
	for _, node := range ns {
		if node.Name == name {
			return node
		}
	}
	return nil
}

func buildMap(pairs []string) nodes {
	objects := make(nodes, 0)
	for _, pair := range pairs {
		parts := strings.Split(pair, ")")
		left := objects.search(parts[0])
		if left == nil {
			left = newNode(parts[0], nil)
			objects = append(objects, left)
		}
		right := objects.search(parts[1])
		if right == nil {
			right = newNode(parts[1], nil)
			objects = append(objects, right)
		}
		right.Parent = left
	}
	return objects
}

type y2019d6p1 struct {
	objects nodes
}

func (p *y2019d6p1) Solve(input string) (string, error) {
	p.objects = buildMap(strings.Split(input, "\n"))
	var total int
	for _, obj := range p.objects {
		total += obj.orbits(0)
	}
	return strconv.Itoa(total), nil
}

type y2019d6p2 struct {
	objects nodes
}

func (p *y2019d6p2) Solve(input string) (string, error) {
	p.objects = buildMap(strings.Split(input, "\n"))
	you := p.objects.search("YOU")
	if you == nil {
		return "", fmt.Errorf("could not find YOU")
	}
	san := p.objects.search("SAN")
	if san == nil {
		return "", fmt.Errorf("could not find SAN")
	}
	youParents := parents(you)
	sanParents := parents(san)
	var crossing string
	for _, y := range youParents {
		for _, s := range sanParents {
			if crossing == "" && y.Name == s.Name {
				crossing = y.Name
			}
		}
	}
	if crossing == "" {
		return "", fmt.Errorf("you do not cross with santa")
	}
	return strconv.Itoa(you.transfersTo(crossing) + san.transfersTo(crossing)), nil
}
