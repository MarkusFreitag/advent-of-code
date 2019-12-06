package puzzles

import (
	"fmt"
	"strconv"

	"github.com/MarkusFreitag/advent-of-code/pkg/puzzles/nineteen"
	"github.com/MarkusFreitag/advent-of-code/pkg/util"
)

var puzzles = map[string]map[string][]util.Puzzle{
	"2019": nineteen.Puzzles,
}

func Get(year, day int) ([]util.Puzzle, error) {
	if y, ok := puzzles[strconv.Itoa(year)]; ok {
		if d, ok := y[strconv.Itoa(day)]; ok {
			return d, nil
		}
		return nil, fmt.Errorf("could not find puzzle for year %d day %d", year, day)
	}
	return nil, fmt.Errorf("could not find puzzles for year %d", year)
}
