package {{.Package}}

import (
{{- range .Imports}}
  "{{.}}"
{{- end}}
  "github.com/MarkusFreitag/advent-of-code/util"
)

var Puzzles = map[string]util.Puzzle{
{{- range .Days}}
	"{{.}}":  { {{.}}.Part1, {{.}}.Part2 },
{{- end}}
}
