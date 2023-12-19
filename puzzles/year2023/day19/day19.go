package day19

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/maputil"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

type Rule struct {
	Var    string
	Op     string
	Val    int
	Result string
}

func (r Rule) Match(item map[string]int) bool {
	if r.Var == "" && r.Op == "" {
		return true
	}
	switch r.Op {
	case ">":
		return item[r.Var] > r.Val
	case "<":
		return item[r.Var] < r.Val
	}
	return false
}

func parseRule(str string) Rule {
	if strings.Contains(str, ":") {
		parts := strings.Split(str, ":")
		return Rule{
			Var:    parts[0][0:1],
			Op:     parts[0][1:2],
			Val:    util.ParseInt(parts[0][2:]),
			Result: parts[1],
		}
	}
	return Rule{Result: str}
}

type Workflow struct {
	Name  string
	Rules []Rule
}

func (w Workflow) Check(item map[string]int) string {
	for _, rule := range w.Rules {
		if rule.Match(item) {
			return rule.Result
		}
	}
	return "n/a"
}

type Workflows map[string]Workflow

func (w Workflows) Do(item map[string]int) bool {
	wf := w["in"]
	for {
		result := wf.Check(item)
		if result == "A" {
			return true
		}
		if result == "R" {
			return false
		}
		wf = w[result]
	}
	return false
}

func parseWorkflow(str string) Workflow {
	parts := strings.Split(str, "{")
	wf := Workflow{Name: parts[0], Rules: make([]Rule, 0)}
	parts[1] = strings.TrimSuffix(parts[1], "}")
	for _, sub := range strings.Split(parts[1], ",") {
		wf.Rules = append(wf.Rules, parseRule(sub))
	}
	return wf
}

func Part1(input string) (string, error) {
	blocks := strings.Split(input, "\n\n")

	workflows := make(Workflows)
	for _, line := range strings.Split(blocks[0], "\n") {
		wf := parseWorkflow(line)
		workflows[wf.Name] = wf
	}

	items := make([]map[string]int, 0)
	for _, line := range strings.Split(blocks[1], "\n") {
		line = strings.TrimPrefix(strings.TrimSuffix(line, "}"), "{")
		item := make(map[string]int)
		for _, pair := range strings.Split(line, ",") {
			parts := strings.Split(pair, "=")
			item[parts[0]] = util.ParseInt(parts[1])
		}
		items = append(items, item)
	}

	var total int
	for _, item := range items {
		if workflows.Do(item) {
			total += numbers.Sum(maputil.Values(item)...)
		}
	}

	return strconv.Itoa(total), nil
}

func solve(ranges map[string][2]int, wfs Workflows, n string) int {
	if n == "A" {
		return numbers.Multiply(
			(ranges["x"][1] - ranges["x"][0] + 1),
			(ranges["m"][1] - ranges["m"][0] + 1),
			(ranges["a"][1] - ranges["a"][0] + 1),
			(ranges["s"][1] - ranges["s"][0] + 1),
		)
	} else if n == "R" {
		return 0
	}

	var count int
	rules := wfs[n].Rules
	for idx, rule := range rules {
		if idx == len(rules)-1 {
			count += do(ranges, wfs, rule.Result)
		}
		newRanges := map[string][2]int{
			"x": [2]int{ranges["x"][0], ranges["x"][1]},
			"m": [2]int{ranges["m"][0], ranges["m"][1]},
			"a": [2]int{ranges["a"][0], ranges["a"][1]},
			"s": [2]int{ranges["s"][0], ranges["s"][1]},
		}
		if rule.Op == ">" && ranges[rule.Var][1] > rule.Val {
			if ranges[rule.Var][0] < rule.Val {
				r := newRanges[rule.Var]
				r[0] = rule.Val + 1
				newRanges[rule.Var] = r
				r = ranges[rule.Var]
				r[1] = rule.Val
				ranges[rule.Var] = r
			}
			count += do(newRanges, wfs, rule.Result)
		} else if rule.Op == "<" && ranges[rule.Var][0] < rule.Val {
			if ranges[rule.Var][1] > rule.Val {
				r := newRanges[rule.Var]
				r[1] = rule.Val - 1
				newRanges[rule.Var] = r
				r = ranges[rule.Var]
				r[0] = rule.Val
				ranges[rule.Var] = r
			}
			count += do(newRanges, wfs, rule.Result)
		}
	}
	return count
}

func Part2(input string) (string, error) {
	blocks := strings.Split(input, "\n\n")

	workflows := make(Workflows)
	for _, line := range strings.Split(blocks[0], "\n") {
		wf := parseWorkflow(line)
		workflows[wf.Name] = wf
	}

	ranges := map[string][2]int{
		"x": [2]int{1, 4000},
		"m": [2]int{1, 4000},
		"a": [2]int{1, 4000},
		"s": [2]int{1, 4000},
	}

	return strconv.Itoa(do(ranges, workflows, "in")), nil
}
