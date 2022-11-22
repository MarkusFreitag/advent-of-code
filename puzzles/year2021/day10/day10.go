package day10

import (
	"sort"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/slice"
)

var (
	brackets = []*Bracket{
		{Open: '(', Close: ')', CorruptScore: 3, IncompleteScore: 1},
		{Open: '[', Close: ']', CorruptScore: 57, IncompleteScore: 2},
		{Open: '{', Close: '}', CorruptScore: 1197, IncompleteScore: 3},
		{Open: '<', Close: '>', CorruptScore: 25137, IncompleteScore: 4},
	}
	bracketByOpen  = make(map[rune]*Bracket)
	bracketByClose = make(map[rune]*Bracket)
)

type Bracket struct {
	Open            rune
	Close           rune
	CorruptScore    int
	IncompleteScore int
}

func (b Bracket) Pair() string {
	return string(b.Open) + string(b.Close)
}

func init() {
	for _, bracket := range brackets {
		bracketByOpen[bracket.Open] = bracket
		bracketByClose[bracket.Close] = bracket
	}
}

func invalidPairs(bracket rune) []string {
	if _, ok := bracketByOpen[bracket]; ok {
		pairs := make([]string, 0)
		for _, b := range brackets {
			if bracket == b.Open {
				continue
			}
			pairs = append(pairs, string(bracket)+string(b.Close))
		}
		return pairs
	}

	if _, ok := bracketByClose[bracket]; ok {
		pairs := make([]string, 0)
		for _, b := range brackets {
			if bracket == b.Close {
				continue
			}
			pairs = append(pairs, string(b.Open)+string(bracket))
		}
		return pairs
	}

	return nil
}

func Part1(input string) (string, error) {
	var sum int
	for _, line := range strings.Fields(input) {
		for len(line) > 0 {
			newLine := line
			for _, bracket := range brackets {
				newLine = strings.ReplaceAll(newLine, bracket.Pair(), "")
			}

			if len(newLine) == len(line) {
				line = newLine
				break
			}
			line = newLine
		}

		for char := range bracketByClose {
			if strings.ContainsRune(line, char) {
				if util.StringContainsAny(line, invalidPairs(char)...) {
					sum += bracketByClose[char].CorruptScore
				}
			}
		}
	}

	return strconv.Itoa(sum), nil
}

func Part2(input string) (string, error) {
	scores := make([]int, 0)
	for _, line := range strings.Fields(input) {
		for len(line) > 0 {
			newLine := line
			for _, bracket := range brackets {
				newLine = strings.ReplaceAll(newLine, bracket.Pair(), "")
			}

			if len(newLine) == len(line) {
				bools := make([]bool, 0)
				for char := range bracketByClose {
					bools = append(bools, strings.ContainsRune(newLine, char))
				}

				if slice.All(bools, false) {
					var score int
					for _, char := range complete(newLine) {
						score *= 5
						score += bracketByClose[char].IncompleteScore
					}
					scores = append(scores, score)
				}
				break
			}
			line = newLine
		}
	}

	sort.Ints(scores)
	return strconv.Itoa(scores[len(scores)/2]), nil
}

func complete(str string) string {
	var s string
	for _, char := range str {
		s += string(bracketByOpen[char].Close)
	}
	return util.StringReverse(s)
}
