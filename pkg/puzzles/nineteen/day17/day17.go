package day17

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/pkg/intcode"
)

var grid = make(Grid, 0)

type Grid [][]string

func (g Grid) Show() {
	lines := make([]string, len(g))
	for idx, row := range g {
		lines[idx] = strings.Join(row, "")
	}
	fmt.Println(strings.Join(lines, "\n"))
}

func Part1(input string) (string, error) {
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
			running = false
		}
	}
	grid.Show()

	var sum int
	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[y])-1; x++ {
			if grid[y][x] == "#" && grid[y-1][x] == "#" && grid[y][x+1] == "#" && grid[y+1][x] == "#" && grid[y][x-1] == "#" {
				fmt.Printf("intersection %d|%d\n", y, x)
				sum += y * x
			}
		}
	}

	return strconv.Itoa(sum), nil
}

/*
  L12L12R4R10R6R4R4L12L12R4R6L12L12R10R6R4R4L12L12R4R10R6R4R4R6L12L12R6L12L12R10R6R4R4

  L12L12 R4R10R6R4R4 L12L12 R4R6 L12L12 R10R6R4R4 L12L12 R4R10R6R4R4R6 L12L12 R6 L12L12 R10R6R4R4

  A        B         A        C        B         A        B         C        C        B
  L12L12R4 R10R6R4R4 L12L12R4 R6L12L12 R10R6R4R4 L12L12R4 R10R6R4R4 R6L12L12 R6L12L12 R10R6R4R4

  => A,B,A,C,B,A,B,C,C,B
*/
func Part2(input string) (string, error) {
	icode, err := intcode.New(input)
	if err != nil {
		return "", err
	}

	icode[0] = 2

	inputs := make([]int64, 0)
	inputs = addInput(inputs, "A,B,A,C,B,A,B,C,C,B")
	inputs = addInput(inputs, "L,12,L,12,R,4")
	inputs = addInput(inputs, "R,10,R,6,R,4,R,4")
	inputs = addInput(inputs, "R,6,L,12,L,12")
	inputs = addInput(inputs, "n")

	in := make(chan int64, 100)
	out := make(chan intcode.Message, 100)
	go intcode.RunSync(icode, in, out)

	var resp string
	running := true
	for running {
		msg := <-out
		switch msg.Type {
		case intcode.MessageWantsInput:
			var i int64
			i, inputs = inputs[0], inputs[1:]
			in <- i
		case intcode.MessageOutput:
			if msg.Value < 127 {
				if msg.Value == 10 {
					fmt.Println(resp)
					resp = ""
				} else {
					resp += fmt.Sprintf("%c", rune(msg.Value))
				}
			} else {
				return strconv.Itoa(int(msg.Value)), nil
			}
		case intcode.MessageHalt:
			running = false
		}
	}

	return "n/a", nil
}

func addInput(inp []int64, str string) []int64 {
	items := strings.Split(str, ",")
	for idx, s := range items {
		for _, b := range []byte(s) {
			inp = append(inp, int64(b))
		}
		if idx < len(items)-1 {
			inp = append(inp, int64(44))
		}
	}
	return append(inp, int64(10))
}
