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
	blocks := strings.Split(input, "\n\n")
	parseRulesBlock(blocks[0])

	/*
		id, rule := parseRule("8: 42 | 42 8")
		rules[id] = rule
		id, rule = parseRule("11: 42 31 | 42 11 31")
		rules[id] = rule
	*/

	strs := strings.Split(blocks[1], "\n")
	count := try1(strs)
	//count := try2(strs)
	//count := try3(strs)

	return strconv.Itoa(count), nil
}

func try1(strs []string) int {
	var count int
	rgx := regexp.MustCompile(fmt.Sprintf("^(%s){2,}(%s)+$", rules[42].Resolve(), rules[31].Resolve()))
	for _, str := range strs {
		if rgx.MatchString(str) {
			count++
		}
	}
	return count
}

func try2(strs []string) int {
	var count int
	for _, str := range strs {
		count42, restStr := countMatches(str, regexp.MustCompile(fmt.Sprintf("^(%s)", rules[42].Resolve())))
		count31, _ := countMatches(restStr, regexp.MustCompile(fmt.Sprintf("(%s)", rules[31].Resolve())))
		if count42 > count31 {
			count++
		}
	}
	return count
}

func try3(strs []string) int {
	var count int
	for _, str := range strs {
		rgx42 := regexp.MustCompile(fmt.Sprintf("(%s)", rules[42].Resolve()))
		rgx31 := regexp.MustCompile(fmt.Sprintf("(%s)", rules[31].Resolve()))

		m42 := rgx42.FindAllStringIndex(str, -1)
		if m42 == nil {
			continue
		}
		m31 := rgx31.FindAllStringIndex(str[m42[len(m42)-1][1]:], -1)
		if m31 == nil {
			continue
		}

		if len(m42) > len(m31) {
			count++
		}
	}
	return count
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
