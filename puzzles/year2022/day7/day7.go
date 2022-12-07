package day7

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
	"github.com/MarkusFreitag/advent-of-code/util/slice"
)

type Entry struct {
	name     string
	dir      bool
	parent   *Entry
	children []*Entry
	size     int
}

func NewDir(name string) *Entry {
	return &Entry{
		name:     name,
		dir:      true,
		children: make([]*Entry, 0),
	}
}

func NewFile(name string, size int) *Entry {
	return &Entry{
		name: name,
		size: size,
	}
}

func (e *Entry) Add(child *Entry) {
	child.parent = e
	e.children = append(e.children, child)
}

func (e *Entry) ChildByName(name string) *Entry {
	for _, child := range e.children {
		if child.name == name {
			return child
		}
	}
	return nil
}

func (e *Entry) Size() int {
	if e.dir {
		var sum int
		for _, child := range e.children {
			sum += child.Size()
		}
		return sum
	}
	return e.size
}

func parseInput(input string) [][]string {
	blocks := make([][]string, 0)
	block := make([]string, 0)
	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "$") {
			blocks = append(blocks, block)
			block = []string{line}
			continue
		}
		block = append(block, line)
	}
	blocks = append(blocks, block)
	return blocks
}

func execHistory(history [][]string) *Entry {
	root := NewDir("/")
	currentDir := root

	for _, block := range history[1:] {
		cmd, output := slice.PopFront(block)

		if strings.HasPrefix(cmd, "$ cd") {
			fields := strings.Fields(cmd)
			if fields[2] == ".." {
				currentDir = currentDir.parent
			} else {
				if child := currentDir.ChildByName(fields[2]); child != nil {
					currentDir = child
				}
			}
		}

		if strings.HasPrefix(cmd, "$ ls") {
			for _, line := range output {
				fields := strings.Fields(line)
				if fields[0] == "dir" {
					currentDir.Add(NewDir(fields[1]))
				} else {
					currentDir.Add(NewFile(fields[1], util.ParseInt(fields[0])))
				}
			}
		}
	}

	return root
}

func filter(e *Entry, fn func(*Entry) bool) []int {
	sizes := make([]int, 0)
	if e.dir {
		if fn(e) {
			sizes = append(sizes, e.Size())
		}
		for _, child := range e.children {
			sizes = append(sizes, filter(child, fn)...)
		}
	}
	return sizes
}

func Part1(input string) (string, error) {
	root := execHistory(parseInput(input))

	sizes := filter(root, func(e *Entry) bool {
		return e.Size() < 100000
	})

	return strconv.Itoa(numbers.Sum(sizes...)), nil
}

func Part2(input string) (string, error) {
	root := execHistory(parseInput(input))

	unused := 70000000 - root.Size()
	needed := 30000000 - unused

	sizes := filter(root, func(e *Entry) bool {
		return e.Size() >= needed
	})

	return strconv.Itoa(numbers.Min(sizes...)), nil
}
