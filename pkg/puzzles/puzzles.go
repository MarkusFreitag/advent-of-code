package puzzles

import "fmt"

type Puzzle interface {
	Solve(string) (string, error)
}

var puzzles = map[string][]Puzzle{
	"2019_1": {
		&y2019d1p1{},
		&y2019d1p2{},
	},
	"2019_2": {
		&y2019d2p1{},
		&y2019d2p2{},
	},
	"2019_6": {
		&y2019d6p1{},
		&y2019d6p2{},
	},
}

func Get(year, day int) ([]Puzzle, error) {
	key := fmt.Sprintf("%d_%d", year, day)
	if p, ok := puzzles[key]; ok {
		return p, nil
	}
	return nil, fmt.Errorf("could not find puzzle for year %d day %d", year, day)
}
