package day6

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

func Part1(input string) (string, error) {
	objects := buildMap(strings.Split(input, "\n"))
	var total int
	for _, obj := range objects {
		total += obj.orbits(0)
	}
	return strconv.Itoa(total), nil
}

func Part2(input string) (string, error) {
	objects := buildMap(strings.Split(input, "\n"))
	you := objects.search("YOU")
	if you == nil {
		return "", fmt.Errorf("could not find YOU")
	}
	san := objects.search("SAN")
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
