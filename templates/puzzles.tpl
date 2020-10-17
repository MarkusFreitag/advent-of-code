package puzzles

import (
  "fmt"
{{ range .Years}}
	"github.com/MarkusFreitag/advent-of-code/puzzles/{{.}}"
{{- end}}
	"github.com/MarkusFreitag/advent-of-code/util"
)

var Puzzles = map[string]map[string]util.Puzzle{
{{- range .Years}}
	"{{.}}": {{.}}.Puzzles,
{{- end}}
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
