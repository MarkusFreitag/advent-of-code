package day19

import (
	"strconv"
	"strings"
)

type replacement struct {
	From string
	To   string
}

func Part1(input string) (string, error) {
	blocks := strings.Split(input, "\n\n")

	molecule := strings.TrimSpace(blocks[1])

	replacements := make([]replacement, 0)
	for _, line := range strings.Split(blocks[0], "\n") {
		parts := strings.Split(line, "=>")
		replacements = append(
			replacements,
			replacement{
				From: strings.TrimSpace(parts[0]),
				To:   strings.TrimSpace(parts[1]),
			},
		)
	}

	newMolecules := make(map[string]bool)
	for _, repl := range replacements {
		if !strings.Contains(molecule, repl.From) {
			continue
		}
		var idx int
		for idx = strings.Index(molecule[idx:], repl.From); idx > -1; idx++ {
			if idx == len(molecule)-1 {
				break
			}

			newMol := molecule[:idx] + strings.Replace(molecule[idx:], repl.From, repl.To, 1)
			if newMol != molecule {
				newMolecules[newMol] = true
			}
		}
	}

	return strconv.Itoa(len(newMolecules)), nil
}

func Part2(input string) (string, error) {
	return "not solved yet", nil
}
