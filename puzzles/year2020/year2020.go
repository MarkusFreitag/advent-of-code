package year2020

import (
	"github.com/MarkusFreitag/advent-of-code/puzzles/year2020/day1"
	"github.com/MarkusFreitag/advent-of-code/puzzles/year2020/day2"
	"github.com/MarkusFreitag/advent-of-code/puzzles/year2020/day3"
	"github.com/MarkusFreitag/advent-of-code/puzzles/year2020/day4"
	"github.com/MarkusFreitag/advent-of-code/puzzles/year2020/day5"
	"github.com/MarkusFreitag/advent-of-code/util"
)

var Puzzles = map[string]util.Puzzle{
	"day1": {day1.Part1, day1.Part2},
	"day2": {day2.Part1, day2.Part2},
	"day3": {day3.Part1, day3.Part2},
	"day4": {day4.Part1, day4.Part2},
	"day5": {day5.Part1, day5.Part2},
}
