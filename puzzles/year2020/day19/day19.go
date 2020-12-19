package day19

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

type Rule struct {
	ID    int
	Match string
	Subs  [][]int
}

func (r Rule) Resolve() string {
	if r.Match != "" {
		return r.Match
	}

	var looped bool
	for _, sub := range r.Subs {
		if util.IntInSlice(r.ID, sub) {
			looped = true
		}
	}

	subs := make([]string, 0)
	for _, sub := range r.Subs {
		if util.IntInSlice(r.ID, sub) {
			continue
		}
		var m string
		for _, s := range sub {
			if v, ok := rules[s]; ok {
				m += v.Resolve()
			}
			if looped {
				m += "+"
			}
		}
		subs = append(subs, m)
	}
	return fmt.Sprintf("(%s)", strings.Join(subs, "|"))
}

func parseRule(str string) (int, Rule) {
	parts := strings.Split(str, ":")
	id, _ := strconv.Atoi(parts[0])
	var rule Rule
	rule.ID = id
	value := strings.TrimSpace(parts[1])
	if strings.Contains(value, `"`) {
		rule.Match, _ = strconv.Unquote(value)
	} else {
		rule.Subs = make([][]int, 0)
		for _, s := range strings.Split(value, "|") {
			ids := util.StrsToInts(strings.Fields(s))
			rule.Subs = append(rule.Subs, ids)
		}
	}
	return id, rule
}

func parseRulesBlock(block string) {
	// first only basic rules
	for _, line := range strings.Split(block, "\n") {
		if !strings.Contains(line, `"`) {
			continue
		}
		id, rule := parseRule(line)
		rules[id] = rule
	}

	// now only nested rules
	for _, line := range strings.Split(block, "\n") {
		if strings.Contains(line, `"`) {
			continue
		}
		id, rule := parseRule(line)
		rules[id] = rule
	}
}

var rules = make(map[int]Rule)

func Part1(input string) (string, error) {
	blocks := strings.Split(input, "\n\n")
	parseRulesBlock(blocks[0])

	rgx := regexp.MustCompile("^" + rules[0].Resolve() + "$")
	var count int
	for _, line := range strings.Split(blocks[1], "\n") {
		if rgx.MatchString(line) {
			count++
		}
	}
	return strconv.Itoa(count), nil
}

func Part2(input string) (string, error) {
	input = strings.ReplaceAll(input, "8: 42", "8: 42 | 42 8")
	input = strings.ReplaceAll(input, "11: 42 31", "11: 42 31 | 42 11 31")
	//input = strings.ReplaceAll(input, "42: 9 14 | 10 1", `42: "MOO"`)
	//input = strings.ReplaceAll(input, "31: 14 17 | 1 13", `31: "FOO"`)
	blocks := strings.Split(input, "\n\n")
	parseRulesBlock(blocks[0])

	//fmt.Println(rules[8].Resolve())
	//fmt.Println(rules[11].Resolve())

	rgx := regexp.MustCompile("^" + rules[0].Resolve() + "$")
	var count int
	for _, line := range strings.Split(blocks[1], "\n") {
		if rgx.MatchString(line) {
			count++
		}
	}
	return strconv.Itoa(count), nil
}
