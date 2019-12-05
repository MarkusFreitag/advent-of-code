package puzzles

import "fmt"

type Puzzle interface {
	Solve(string) (string, error)
}

var puzzles = map[string][]Puzzle{}

func Get(year, day int) ([]Puzzle, error) {
	key := fmt.Sprintf("%d_%d", year, day)
	if p, ok := puzzles[key]; ok {
		return p, nil
	}
	return nil, fmt.Errorf("could not find puzzle for year %d day %d", year, day)
}
