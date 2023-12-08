package year2023

import (
  "github.com/MarkusFreitag/advent-of-code/puzzles/year2023/day1"
  "github.com/MarkusFreitag/advent-of-code/puzzles/year2023/day2"
  "github.com/MarkusFreitag/advent-of-code/puzzles/year2023/day3"
  "github.com/MarkusFreitag/advent-of-code/puzzles/year2023/day4"
  "github.com/MarkusFreitag/advent-of-code/puzzles/year2023/day5"
  "github.com/MarkusFreitag/advent-of-code/puzzles/year2023/day6"
  "github.com/MarkusFreitag/advent-of-code/puzzles/year2023/day8"
  "github.com/MarkusFreitag/advent-of-code/util"
)

var Puzzles = map[string]util.Puzzle{
	"day1":  { day1.Part1, day1.Part2 },
	"day2":  { day2.Part1, day2.Part2 },
	"day3":  { day3.Part1, day3.Part2 },
	"day4":  { day4.Part1, day4.Part2 },
	"day5":  { day5.Part1, day5.Part2 },
	"day6":  { day6.Part1, day6.Part2 },
	"day8":  { day8.Part1, day8.Part2 },
}