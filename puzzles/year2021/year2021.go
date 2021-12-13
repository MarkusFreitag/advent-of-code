package year2021

import (
	"github.com/MarkusFreitag/advent-of-code/puzzles/year2021/day1"
	"github.com/MarkusFreitag/advent-of-code/puzzles/year2021/day10"
	"github.com/MarkusFreitag/advent-of-code/puzzles/year2021/day11"
	"github.com/MarkusFreitag/advent-of-code/puzzles/year2021/day12"
	"github.com/MarkusFreitag/advent-of-code/puzzles/year2021/day13"
	"github.com/MarkusFreitag/advent-of-code/puzzles/year2021/day2"
	"github.com/MarkusFreitag/advent-of-code/puzzles/year2021/day3"
	"github.com/MarkusFreitag/advent-of-code/puzzles/year2021/day4"
	"github.com/MarkusFreitag/advent-of-code/puzzles/year2021/day5"
	"github.com/MarkusFreitag/advent-of-code/puzzles/year2021/day6"
	"github.com/MarkusFreitag/advent-of-code/puzzles/year2021/day7"
	"github.com/MarkusFreitag/advent-of-code/puzzles/year2021/day9"
	"github.com/MarkusFreitag/advent-of-code/util"
)

var Puzzles = map[string]util.Puzzle{
	"day1":  {day1.Part1, day1.Part2},
	"day10": {day10.Part1, day10.Part2},
	"day11": {day11.Part1, day11.Part2},
	"day12": {day12.Part1, day12.Part2},
	"day13": {day13.Part1, day13.Part2},
	"day2":  {day2.Part1, day2.Part2},
	"day3":  {day3.Part1, day3.Part2},
	"day4":  {day4.Part1, day4.Part2},
	"day5":  {day5.Part1, day5.Part2},
	"day6":  {day6.Part1, day6.Part2},
	"day7":  {day7.Part1, day7.Part2},
	"day9":  {day9.Part1, day9.Part2},
}
