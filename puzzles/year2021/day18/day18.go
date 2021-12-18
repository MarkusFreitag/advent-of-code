package day18

import (
	"fmt"
	"strconv"
	"strings"
)

func intP(i int) *int { return &i }

type Tree struct {
	Value  *int
	Left   *Tree
	Right  *Tree
	Parent *Tree
}

func NewTree(str string) *Tree {
	tree := new(Tree)
	i, err := strconv.Atoi(str)
	if err == nil {
		tree.Value = intP(i)
		return tree
	}
	var depth int
	for idx, char := range str[1 : len(str)-1] {
		if char == '[' {
			depth++
		} else if char == ']' {
			depth--
		} else if char == ',' && depth == 0 {
			tree.Left = NewTree(str[1 : idx+1])
			tree.Right = NewTree(str[idx+2 : len(str)-1])
			tree.Left.Parent = tree
			tree.Right.Parent = tree
			break
		}
	}
	return tree
}

func (t *Tree) String() string {
	if t.Value == nil {
		return fmt.Sprintf("[%s,%s]", t.Left, t.Right)
	}
	return strconv.Itoa(*t.Value)
}

func (t *Tree) Add(tree *Tree) *Tree {
	newTree := new(Tree)
	newTree.Left = t
	newTree.Right = tree
	newTree.Left.Parent = newTree
	newTree.Right.Parent = newTree

	for newTree.Explode(0) || newTree.Split() {
		continue
	}

	return newTree
}

func (t *Tree) Split() bool {
	if t.Value == nil {
		return t.Left.Split() || t.Right.Split()
	}

	i := *t.Value
	if i > 9 {
		t.Value = nil

		t.Left = &Tree{
			Parent: t,
			Value:  intP(i / 2),
		}
		t.Right = &Tree{
			Parent: t,
			Value:  intP(i/2 + i%2),
		}

		return true
	}
	return false
}

func (t *Tree) Explode(depth int) bool {
	if t.Left != nil {
		if depth == 4 && t.Left.Left == nil && t.Right.Left == nil {
			left, old := t.Parent, t
			for left != nil && left.Left == old {
				old, left = left, left.Parent
			}
			if left != nil {
				left = left.Left
			}
			for left != nil && left.Right != nil {
				left = left.Right
			}

			right, old := t.Parent, t
			for right != nil && right.Right == old {
				old, right = right, right.Parent
			}
			if right != nil {
				right = right.Right
			}
			for right != nil && right.Left != nil {
				right = right.Left
			}

			if left != nil {
				left.Value = intP(*left.Value + *t.Left.Value)
			}
			if right != nil {
				right.Value = intP(*right.Value + *t.Right.Value)
			}

			t.Left = nil
			t.Right = nil
			t.Value = intP(0)
			return true
		}
		return t.Left.Explode(depth+1) || t.Right.Explode(depth+1)
	}
	return false
}

func (t *Tree) Magnitude() int {
	if t.Value != nil {
		return *t.Value
	}
	return 3*t.Left.Magnitude() + 2*t.Right.Magnitude()
}

func Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")

	tree := NewTree(lines[0])

	for _, line := range lines[1:] {
		tree = tree.Add(NewTree(line))
	}

	return strconv.Itoa(tree.Magnitude()), nil
}

func Part2(input string) (string, error) {
	lines := strings.Split(input, "\n")
	var max int
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			if i == j {
				continue
			}
			/*
				Parsing the same trees over and over again is not efficient,
				but reusing them will end in chaos due to the usage of pointers.
			*/
			if mag := NewTree(lines[i]).Add(NewTree(lines[j])).Magnitude(); mag > max {
				max = mag
			}
		}
	}
	return strconv.Itoa(max), nil
}
