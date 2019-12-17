package day17

import (
	"fmt"
  "strconv"
  "strings"

	"github.com/MarkusFreitag/advent-of-code/pkg/intcode"
)

var grid = make(Grid, 0)

type Grid [][]string

type Part1 struct{}

func (p *Part1) Solve(input string) (string, error) {
	icode, err := intcode.New(input)
	if err != nil {
		return "", err
	}
	out := make(chan intcode.Message, 100)
	go intcode.RunAsync(icode, nil, out)

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
      case 60:
        row = append(row, "<")
      case 62:
        row = append(row, ">")
      case 94:
        row = append(row, "^")
      case 118:
        row = append(row, "v")
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

// L12L12 R4R10 R6 R4R4 L12L12 R4R6 L12L12 R10R6 R4R4 L12L12 R4R10R6R4R4R6 L12L12 R6 L12L12 R10R6 R4R4

type Part2 struct{}

func (p *Part2) Solve(input string) (string, error) {
  lines := make([]string, len(grid))
  for idx, row := range grid {
    lines[idx] = fmt.Sprintf("%2d %s", idx, strings.Join(row, ""))
  }
  fmt.Println(strings.Join(lines, "\n"))
  return "n/a", nil

	icode, err := intcode.New(input)
	if err != nil {
		return "", err
	}

  icode[0] = 2

  in := make(chan int64)
	out := make(chan intcode.Message, 100)
	go intcode.RunAsync(icode, in, out)

  var dust int
  running := true
  for running {
    msg := <-out
    switch msg.Type {
    case intcode.MessageOutput:
    case intcode.MessageHalt:
      running = false
    }
  }
  return strconv.Itoa(dust), nil
}
