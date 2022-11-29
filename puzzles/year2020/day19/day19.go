package day19

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/slice"
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
		if slice.Contains(sub, r.ID) {
			looped = true
		}
	}

	subs := make([]string, 0)
	for _, sub := range r.Subs {
		if slice.Contains(sub, r.ID) {
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
			ids := util.StringsToInts(strings.Fields(s))
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
	blocks := strings.Split(input, "\n\n")
	parseRulesBlock(blocks[0])

	/*
		id, rule := parseRule("8: 42 | 42 8")
		rules[id] = rule
		id, rule = parseRule("11: 42 31 | 42 11 31")
		rules[id] = rule
	*/

	rule42 := rules[42].Resolve()
	rule31 := rules[31].Resolve()

	rgx42 := regexp.MustCompile(fmt.Sprintf("^(%s)", rule42))
	rgx31 := regexp.MustCompile(fmt.Sprintf("(%s)", rule31))

	var count int
	rgx := regexp.MustCompile(fmt.Sprintf("^(%s){2,}(%s)+$", rule42, rule31))
	for _, str := range strings.Split(blocks[1], "\n") {
		if rgx.MatchString(str) {
			// Make sure that rule42 has more matches than rule31
			count42, restStr := countMatches(str, rgx42)
			count31, _ := countMatches(restStr, rgx31)
			if count42 > count31 {
				count++
			}
		}
	}

	return strconv.Itoa(count), nil
}

func countMatches(str string, rgx *regexp.Regexp) (int, string) {
	var count int
	for {
		result := rgx.FindStringIndex(str)
		if result == nil {
			break
		}
		count++
		str = str[result[1]:]
	}
	return count, str
}
