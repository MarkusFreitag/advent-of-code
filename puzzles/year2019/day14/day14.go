package day14

import (
	"strconv"
	"strings"
)

func parseVariable(s string) (int, string) {
	s = strings.TrimSpace(s)
	parts := strings.Split(s, " ")
	v, _ := strconv.Atoi(parts[0])
	return v, parts[1]
}

type Chemical struct {
	Name   string
	Amount int
}

func parseChemical(s string) Chemical {
	i, n := parseVariable(s)
	return Chemical{
		Name:   n,
		Amount: i,
	}
}

type Formula struct {
	Inputs []Chemical
	Output Chemical
}

func NewFormula() Formula {
	return Formula{
		Inputs: make([]Chemical, 0),
	}
}

func (f Formula) Calc(inv Inventar) Inventar {
	if v, ok := inv[f.Output.Name]; ok {
		inv[f.Output.Name] = v + f.Output.Amount
	} else {
		inv[f.Output.Name] = f.Output.Amount
	}
	for _, inp := range f.Inputs {
		if v, ok := inv[inp.Name]; ok {
			inv[inp.Name] = v - inp.Amount
		} else {
			inv[inp.Name] = -inp.Amount
		}
	}
	return inv
}

type Formulas map[string]Formula

func loadFormulas(str string) Formulas {
	items := make(map[string]Formula)
	for _, line := range strings.Split(str, "\n") {
		parts := strings.Split(line, "=>")
		in := strings.TrimSpace(parts[0])
		out := strings.TrimSpace(parts[1])
		form := NewFormula()
		form.Output = parseChemical(out)
		for _, i := range strings.Split(in, ",") {
			form.Inputs = append(form.Inputs, parseChemical(i))
		}
		items[form.Output.Name] = form
	}
	return items
}

func (f Formulas) OrePerFuel() int {
	inv := make(Inventar)
	inv = f["FUEL"].Calc(inv)
	for {
		solved := true
		for k, v := range inv {
			if k == "ORE" {
				continue
			}
			if v < 0 {
				inv = f[k].Calc(inv)
				solved = false
			}
		}
		if solved {
			break
		}
	}
	return -inv["ORE"]
}

type Inventar map[string]int

func Part1(input string) (string, error) {
	formulas := loadFormulas(input)
	return strconv.Itoa(formulas.OrePerFuel()), nil
}

func Part2(input string) (string, error) {
	return "n/a", nil
}
