package nineteen

import (
	"github.com/MarkusFreitag/advent-of-code/pkg/puzzles/nineteen/day1"
	"github.com/MarkusFreitag/advent-of-code/pkg/puzzles/nineteen/day11"
	"github.com/MarkusFreitag/advent-of-code/pkg/puzzles/nineteen/day13"
	"github.com/MarkusFreitag/advent-of-code/pkg/puzzles/nineteen/day2"
	"github.com/MarkusFreitag/advent-of-code/pkg/puzzles/nineteen/day3"
	"github.com/MarkusFreitag/advent-of-code/pkg/puzzles/nineteen/day4"
	"github.com/MarkusFreitag/advent-of-code/pkg/puzzles/nineteen/day5"
	"github.com/MarkusFreitag/advent-of-code/pkg/puzzles/nineteen/day6"
	"github.com/MarkusFreitag/advent-of-code/pkg/puzzles/nineteen/day7"
	"github.com/MarkusFreitag/advent-of-code/pkg/puzzles/nineteen/day8"
	"github.com/MarkusFreitag/advent-of-code/pkg/puzzles/nineteen/day9"
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
	"5": {
		&day5.Part1{},
		&day5.Part2{},
	},
	"6": {
		&day6.Part1{},
		&day6.Part2{},
	},
	"7": {
		&day7.Part1{},
		&day7.Part2{},
	},
	"8": {
		&day8.Part1{},
		&day8.Part2{},
	},
	"9": {
		&day9.Part1{},
		&day9.Part2{},
	},
	"11": {
		&day11.Part1{},
		&day11.Part2{},
	},
	"13": {
		&day13.Part1{},
		&day13.Part2{},
	},
}
