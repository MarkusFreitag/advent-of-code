package nineteen

import (
	"github.com/MarkusFreitag/advent-of-code/pkg/puzzles/nineteen/day1"
	"github.com/MarkusFreitag/advent-of-code/pkg/puzzles/nineteen/day2"
	"github.com/MarkusFreitag/advent-of-code/pkg/puzzles/nineteen/day3"
	"github.com/MarkusFreitag/advent-of-code/pkg/puzzles/nineteen/day4"
	"github.com/MarkusFreitag/advent-of-code/pkg/puzzles/nineteen/day6"
	"github.com/MarkusFreitag/advent-of-code/pkg/util"
)

var Puzzles = map[string][]util.Puzzle{
	"1": {
		&day1.Part1{},
		&day1.Part2{},
	},
	"2": {
		&day2.Part1{},
		&day2.Part2{},
	},
	"3": {
		&day3.Part1{},
		&day3.Part2{},
	},
	"4": {
		&day4.Part1{},
		&day4.Part2{},
	},
	"6": {
		&day6.Part1{},
		&day6.Part2{},
	},
}
