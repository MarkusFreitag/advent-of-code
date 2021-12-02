package year2021

import (
	"github.com/MarkusFreitag/advent-of-code/puzzles/year2021/day1"
	"github.com/MarkusFreitag/advent-of-code/puzzles/year2021/day2"
	"github.com/MarkusFreitag/advent-of-code/util"
)

var Puzzles = map[string]util.Puzzle{
	"day1": {day1.Part1, day1.Part2},
	"day2": {day2.Part1, day2.Part2},
}
