package year2017

import (
	"github.com/MarkusFreitag/advent-of-code/puzzles/year2017/day10"
	"github.com/MarkusFreitag/advent-of-code/puzzles/year2017/day9"
	"github.com/MarkusFreitag/advent-of-code/util"
)

var Puzzles = map[string]util.Puzzle{
	"day10": {day10.Part1, day10.Part2},
	"day9":  {day9.Part1, day9.Part2},
}
