package year2018

import (
	"github.com/MarkusFreitag/advent-of-code/puzzles/year2018/day1"
	"github.com/MarkusFreitag/advent-of-code/puzzles/year2018/day2"
	"github.com/MarkusFreitag/advent-of-code/puzzles/year2018/day3"
	"github.com/MarkusFreitag/advent-of-code/util"
)

var Puzzles = map[string]util.Puzzle{
	"day1": {day1.Part1, day1.Part2},
	"day2": {day2.Part1, day2.Part2},
	"day3": {day3.Part1, day3.Part2},
}
