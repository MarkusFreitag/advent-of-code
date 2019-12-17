package day17

import (
	"fmt"
  "strconv"
  "strings"

	"github.com/MarkusFreitag/advent-of-code/pkg/intcode"
)

type Pos struct {
	X, Y int
	Dir  string
}

func (p *Pos) Move(i int) {
	switch p.Dir {
	case "^":
		if i == 0 {
			p.Dir = "<"
			p.X -= 1
		} else {
			p.Dir = ">"
			p.X += 1
		}
	case "v":
		if i == 0 {
			p.Dir = ">"
			p.X += 1
		} else {
			p.Dir = "<"
			p.X -= 1
		}
	case ">":
		if i == 0 {
			p.Dir = "^"
			p.Y += 1
		} else {
			p.Dir = "v"
			p.Y -= 1
		}
	case "<":
		if i == 0 {
			p.Dir = "v"
			p.Y -= 1
		} else {
			p.Dir = "^"
			p.Y += 1
		}
	}
}

func (p *Pos) String() string {
	return fmt.Sprintf("%d|%d", p.X, p.Y)
}

type Grid [][]string

type Part1 struct{}

func (p *Part1) Solve(input string) (string, error) {
	icode, err := intcode.New(input)
	if err != nil {
		return "", err
	}
	out := make(chan intcode.Message, 100)
	go intcode.RunAsync(icode, nil, out)

  grid := make(Grid, 0)
  row := make([]string, 0)
  running := true
  for running {
    msg := <-out
    switch msg.Type {
    case intcode.MessageOutput:
      switch msg.Value {
      case 35:
        row = append(row, "#")
      case 46:
        row = append(row, ".")
      case 10:
        if len(row) > 0 {
          grid = append(grid, row)
          row = make([]string, 0)
        }
      }
    case intcode.MessageHalt:
      lines := make([]string, len(grid))
      for idx, row := range grid {
        lines[idx] = fmt.Sprintf("%2d %s", idx, strings.Join(row, ""))
      }
      fmt.Println(strings.Join(lines, "\n"))
      running = false
    }
  }

  fmt.Printf("len(grid): %d %d\n", len(grid), len(grid)-1)

  var sum int
  for y:=1;y<len(grid)-1;y++ {
    for x:=1;x<len(grid[y])-1;x++ {
      if grid[y][x] == "#" && grid[y-1][x] == "#" && grid[y][x+1] == "#" && grid[y+1][x] == "#" && grid[y][x-1] == "#" {
        fmt.Printf("intersection %d|%d\n", y, x)
        sum += y*x
      }
    }
  }

  return strconv.Itoa(sum), nil
}

type Part2 struct{}

func (p *Part2) Solve(input string) (string, error) {
  return "n/a", nil
}
