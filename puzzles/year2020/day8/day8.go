package day8

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

type Cmd struct {
	Action string
	Value  int
	Runned bool
}

type Cmds []*Cmd

func LoadCmds(table string) Cmds {
	cmds := make(Cmds, 0)
	for _, line := range strings.Split(table, "\n") {
		parts := strings.Split(line, " ")
		cmds = append(cmds, &Cmd{
			Action: parts[0],
			Value:  util.ParseInt(parts[1]),
		})
	}
	return cmds
}

func (c Cmds) Run() (int, bool) {
	var global int
	var pos int
	var fullCycle bool
	for {
		if pos == len(c) {
			fullCycle = true
			break
		}
		cmd := c[pos]
		if cmd.Runned {
			break
		}
		switch cmd.Action {
		case "acc":
			global += cmd.Value
			pos++
		case "jmp":
			pos += cmd.Value
		case "nop":
			pos++
		}
		cmd.Runned = true
	}

	return global, fullCycle
}

func Part1(input string) (string, error) {
	cmds := LoadCmds(input)
	global, _ := cmds.Run()
	return strconv.Itoa(global), nil
}

func Part2(input string) (string, error) {
	cmds := LoadCmds(input)

	var global int
	for idx, cmd := range cmds {
		if cmd.Action == "acc" {
			continue
		}

		program := newCmds(cmds)
		if cmd.Action == "jmp" {
			program[idx].Action = "nop"
		} else if cmd.Action == "nop" {
			program[idx].Action = "jmp"
		}

		var fullCycle bool
		global, fullCycle = program.Run()

		if fullCycle {
			break
		}
	}
	return strconv.Itoa(global), nil
}

func newCmds(cmds Cmds) Cmds {
	nCmds := make(Cmds, len(cmds))
	for i, c := range cmds {
		if c == nil {
			continue
		}
		v := *c
		nCmds[i] = &v
	}
	return nCmds
}
