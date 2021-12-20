package puzzles

import (
	"fmt"

	"github.com/MarkusFreitag/advent-of-code/puzzles/year2015"
	"github.com/MarkusFreitag/advent-of-code/puzzles/year2016"
	"github.com/MarkusFreitag/advent-of-code/puzzles/year2018"
	"github.com/MarkusFreitag/advent-of-code/puzzles/year2019"
	"github.com/MarkusFreitag/advent-of-code/puzzles/year2020"
	"github.com/MarkusFreitag/advent-of-code/puzzles/year2021"
	"github.com/MarkusFreitag/advent-of-code/util"
)

var Puzzles = map[string]map[string]util.Puzzle{
	"year2015": year2015.Puzzles,
	"year2016": year2016.Puzzles,
	"year2018": year2018.Puzzles,
	"year2019": year2019.Puzzles,
	"year2020": year2020.Puzzles,
	"year2021": year2021.Puzzles,
}

func Get(year, day int) (util.Puzzle, error) {
	if y, ok := Puzzles[fmt.Sprintf("year%d", year)]; ok {
		if d, ok := y[fmt.Sprintf("day%d", day)]; ok {
			return d, nil
		}
		return nil, fmt.Errorf("could not find puzzle for year %d day %d", year, day)
	}
	return nil, fmt.Errorf("could not find puzzles for year %d", year)
}
