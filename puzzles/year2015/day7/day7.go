package day7

import (
	"fmt"
	"strconv"
	"strings"
)

type logicGate struct {
	left       uint16
	right      uint16
	leftReady  bool
	rightReady bool
	leftID     string
	rightID    string
	op         string
	dest       string
}

func newLogicGate() *logicGate {
	return new(logicGate)
}

func (l *logicGate) setLeft(i uint16) {
	l.left = i
	l.leftReady = true
}

func (l *logicGate) setRight(i uint16) {
	l.right = i
	l.rightReady = true
}

func (l *logicGate) ready() bool {
	switch l.op {
	case "SET":
		return l.leftReady
	case "NOT":
		return l.rightReady
	default:
		return l.leftReady && l.rightReady
	}
}

func (l *logicGate) solve() uint16 {
	var result uint16
	switch l.op {
	case "SET":
		result = l.left
	case "NOT":
		result = ^l.right
	case "AND":
		result = l.left & l.right
	case "OR":
		result = l.left | l.right
	case "LSHIFT":
		result = l.left << l.right
	case "RSHIFT":
		result = l.left >> l.right
	}
	return result
}

func parseLine(line string) *logicGate {
	parts := strings.Split(line, "->")
	gate := newLogicGate()
	var operation string
	operation, gate.dest = strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
	parts = strings.Split(operation, " ")
	switch len(parts) {
	case 1:
		gate.op = "SET"
		if num, err := strconv.Atoi(parts[0]); err == nil {
			gate.setLeft(uint16(num))
		} else {
			gate.leftID = parts[0]
		}
	case 2:
		gate.op = parts[0]
		if num, err := strconv.Atoi(parts[1]); err == nil {
			gate.setRight(uint16(num))
		} else {
			gate.rightID = parts[1]
		}
	case 3:
		gate.op = parts[1]
		if num, err := strconv.Atoi(parts[0]); err == nil {
			gate.setLeft(uint16(num))
		} else {
			gate.leftID = parts[0]
		}
		if num, err := strconv.Atoi(parts[2]); err == nil {
			gate.setRight(uint16(num))
		} else {
			gate.rightID = parts[2]
		}
	}
	return gate
}

func remove(s []*logicGate, i int) []*logicGate {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func solve(gates []*logicGate) map[string]uint16 {
	solved := make(map[string]uint16)
	for len(gates) > 0 {
		for idx, gate := range gates {
			if gate.ready() {
				solved[gate.dest] = gate.solve()
				gates = remove(gates, idx)
				break
			}
		}
		for key, value := range solved {
			for _, gate := range gates {
				if gate.leftID == key {
					gate.setLeft(value)
				}
				if gate.rightID == key {
					gate.setRight(value)
				}
			}
		}
	}
	return solved
}

func Part1(input string) (string, error) {
	unsolved := make([]*logicGate, 0)
	for _, line := range strings.Split(input, "\n") {
		unsolved = append(unsolved, parseLine(line))
	}
	solved := solve(unsolved)
	a = solved["a"]
	return fmt.Sprintf("%d", a), nil
}

var a uint16

func Part2(input string) (string, error) {
	unsolved := make([]*logicGate, 0)
	for _, line := range strings.Split(input, "\n") {
		gate := parseLine(line)
		if gate.dest == "b" {
			gate.setLeft(uint16(956))
		}
		unsolved = append(unsolved, gate)
	}
	solved := solve(unsolved)
	return fmt.Sprintf("%d", solved["a"]), nil
}
