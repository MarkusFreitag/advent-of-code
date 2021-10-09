package day4

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

var rgx = regexp.MustCompile(`(.*?)-(\d+)\[([a-z]+)\]`)

type char struct {
	r rune
	c int
}

type chars []*char

func (c chars) Len() int      { return len(c) }
func (c chars) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c chars) Less(i, j int) bool {
	if c[i].c == c[j].c {
		return int(c[i].r) < int(c[j].r)
	}
	return c[i].c > c[j].c
}

func (c chars) Contains(r rune) int {
	for idx, char := range c {
		if char.r == r {
			return idx
		}
	}
	return -1
}

type room struct {
	ID       int
	Name     string
	Checksum string
}

func parseRoom(str string) *room {
	matches := rgx.FindAllStringSubmatch(str, -1)
	id, _ := strconv.Atoi(matches[0][2])
	return &room{
		ID:       id,
		Name:     matches[0][1],
		Checksum: matches[0][3],
	}
}

func (r *room) Valid() bool {
	distribution := make(chars, 0)
	for _, r := range r.Name {
		if !unicode.IsLower(r) {
			continue
		}
		idx := distribution.Contains(r)
		if idx == -1 {
			distribution = append(distribution, &char{r: r, c: 1})
		} else {
			distribution[idx].c++
		}
	}
	sort.Sort(distribution)

	var order string
	for _, char := range distribution {
		order += string(char.r)
		if len(order) == 5 {
			break
		}
	}

	return order == r.Checksum
}

func (r *room) Decrypt() string {
	realOffset := r.ID % 26
	var realName string
	for _, r := range r.Name {
		if r == '-' {
			realName += " "
		} else {
			n := int(r) + realOffset
			if n > int('z') {
				realName += string(rune(n - int('z') + int('a') - 1))
			} else {
				realName += string(rune(n))
			}
		}
	}
	return realName
}

func Part1(input string) (string, error) {
	var counter int
	for _, line := range strings.Split(input, "\n") {
		if room := parseRoom(line); room.Valid() {
			counter += room.ID
		}
	}
	return strconv.Itoa(counter), nil
}

func Part2(input string) (string, error) {
	for _, line := range strings.Split(input, "\n") {
		if room := parseRoom(line); room.Valid() {
			if name := room.Decrypt(); strings.Contains(name, "north") && strings.Contains(name, "pole") && strings.Contains(name, "object") {
				return strconv.Itoa(room.ID), nil
			}
		}
	}
	return "n/a", nil
}
