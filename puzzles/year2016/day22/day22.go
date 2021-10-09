package day22

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type Node struct {
	Name   string
	X, Y   int
	Size   int
	Used   int
	Avail  int
	Use    int
	Target bool
}

func rgxMatchGroups(rgx *regexp.Regexp, str string) map[string]string {
	match := rgx.FindStringSubmatch(str)
	result := make(map[string]string)
	for i, name := range rgx.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	return result
}

func parseName(name string) (int, int) {
	rgx := regexp.MustCompile(`/dev/grid/node-x(?P<x>\d+)-y(?P<y>\d+)`)

	matches := rgxMatchGroups(rgx, name)

	return num(matches["x"]), num(matches["y"])
}

func num(str string) int {
	n, _ := strconv.Atoi(str)
	return n
}

func trimUnit(str string) string {
	return strings.TrimRightFunc(str, func(r rune) bool {
		return !unicode.IsNumber(r)
	})
}

func ParseNode(str string) *Node {
	fields := strings.Fields(str)

	x, y := parseName(fields[0])

	return &Node{
		Name:  fields[0],
		X:     x,
		Y:     y,
		Size:  num(trimUnit(fields[1])),
		Used:  num(trimUnit(fields[2])),
		Avail: num(trimUnit(fields[3])),
		Use:   num(trimUnit(fields[4])),
	}
}

func NodeName(x, y int) string {
	return fmt.Sprintf("/dev/grid/node-x%d-y%d", x, y)
}

func Part1(input string) (string, error) {
	nodes := make(map[string]*Node)

	lines := strings.Split(input, "\n")[2:]
	for _, line := range lines {
		node := ParseNode(line)
		nodes[node.Name] = node
	}

	var pairs int
	for name, node := range nodes {
		for na, no := range nodes {
			if name == na {
				continue
			}
			if node.Used != 0 && node.Used < no.Avail {
				pairs++
			}
		}
	}

	return strconv.Itoa(pairs), nil
}

func Part2(input string) (string, error) {
	nodes := make(map[string]*Node)

	var highestY, highestX int
	lines := strings.Split(input, "\n")[2:]
	for _, line := range lines {
		node := ParseNode(line)
		nodes[node.Name] = node
		if node.X > highestX {
			highestX = node.X
		}
		if node.Y > highestY {
			highestY = node.Y
		}
	}

	nodes[NodeName(highestX, 0)].Target = true

	for y := 0; y <= highestY; y++ {
		for x := 0; x <= highestX; x++ {
			nn := NodeName(x, y)
			s := fmt.Sprintf("%2d/%2d", nodes[nn].Used, nodes[nn].Size)
			if y == 0 && x == 0 {
				s = "(" + s + ")"
			} else if nodes[nn].Target {
				s = "[" + s + "]"
			} else {
				s = " " + s + " "
			}
			fmt.Print(s)
			if x < highestX {
				fmt.Print(" -- ")
			}
		}
		fmt.Print("\n")
	}

	return "not solved yet", nil
}
