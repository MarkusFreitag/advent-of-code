package day24

import (
	"slices"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

func binaryFromWires(states map[string]int, prefix string) string {
	type Wire struct {
		Name  string
		Value int
	}
	wires := make([]Wire, 0)
	for name, val := range states {
		if strings.HasPrefix(name, prefix) {
			wires = append(wires, Wire{Name: name, Value: val})
		}
	}
	slices.SortFunc(wires, func(wa, wb Wire) int { return strings.Compare(wa.Name, wb.Name) })
	slices.Reverse(wires)
	var binary string
	for _, wire := range wires {
		binary += strconv.Itoa(wire.Value)
	}
	return binary
}

func simulate(states map[string]int, todo []string) map[string]int {
	for len(todo) > 0 {
		dup := make([]string, 0)
		for _, line := range todo {
			parts := strings.Fields(line)
			leftVal, leftOk := states[parts[0]]
			rightVal, rightOk := states[parts[2]]
			if !leftOk || !rightOk {
				dup = append(dup, line)
				continue
			}
			var result int
			switch parts[1] {
			case "AND":
				if leftVal == 1 && rightVal == 1 {
					result = 1
				}
			case "OR":
				if leftVal == 1 || rightVal == 1 {
					result = 1
				}
			case "XOR":
				if leftVal != rightVal {
					result = 1
				}
			}
			states[parts[4]] = result
		}
		todo = dup
	}
	return states
}

func Part1(input string) (string, error) {
	blocks := strings.Split(input, "\n\n")
	states := make(map[string]int)
	for _, line := range strings.Split(blocks[0], "\n") {
		parts := strings.Split(line, ":")
		states[parts[0]] = util.ParseInt(strings.TrimSpace(parts[1]))
	}
	states = simulate(states, strings.Split(blocks[1], "\n"))
	return strconv.Itoa(util.BinStringToDecInt(binaryFromWires(states, "z"))), nil
}

func Part2(input string) (string, error) {
	blocks := strings.Split(input, "\n\n")
	states := make(map[string]int)
	for _, line := range strings.Split(blocks[0], "\n") {
		parts := strings.Split(line, ":")
		states[parts[0]] = util.ParseInt(strings.TrimSpace(parts[1]))
	}
	ops := make(map[string][3]string)
	for _, line := range strings.Split(blocks[1], "\n") {
		parts := strings.Fields(line)
		ops[parts[4]] = [3]string{parts[0], parts[1], parts[2]}
	}
	broken := make([]string, 0)
	for out, op := range ops {
		inpIsX := strings.HasPrefix(op[0], "x") || strings.HasPrefix(op[2], "x")
		inpIsY := strings.HasPrefix(op[0], "y") || strings.HasPrefix(op[2], "y")

		if strings.HasPrefix(out, "z") && out != "z45" && op[1] != "XOR" {
			broken = append(broken, out)
			continue
		}

		if !strings.HasPrefix(out, "z") && !inpIsX && !inpIsY && op[1] != "AND" && op[1] != "OR" {
			broken = append(broken, out)
			continue
		}

		if op[1] == "XOR" && inpIsX && inpIsY {
			var found bool
			for _, nextOp := range ops {
				if nextOp[1] == "XOR" && (nextOp[0] == out || nextOp[2] == out) {
					found = true
					break
				}
			}
			if !found && op[0] != "x00" && op[0] != "y00" {
				broken = append(broken, out)
				continue
			}
		}

		if op[1] == "AND" {
			var found bool
			for _, nextOp := range ops {
				if nextOp[1] == "OR" && (nextOp[0] == out || nextOp[2] == out) {
					found = true
					break
				}
			}
			if !found && op[0] != "x00" && op[0] != "y00" {
				broken = append(broken, out)
				continue
			}
		}
	}

	slices.Sort(broken)
	return strings.Join(broken, ","), nil
}
