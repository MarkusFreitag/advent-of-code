package day13

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

func Part1(input string) (string, error) {
	icode, err := intcode.New(input)
	if err != nil {
		return "", err
	}
	out := make(chan intcode.Message, 100)
	go intcode.RunAsync(icode, nil, out)

	grid := make(map[string]int)
	for {
		msg := <-out
		switch msg.Type {
		case intcode.MessageOutput:
			x := int(msg.Value)

			msg = <-out
			y := int(msg.Value)

			msg = <-out
			third := int(msg.Value)

			pos := fmt.Sprintf("%d|%d", y, x)
			grid[pos] = third

		case intcode.MessageHalt:
			var counter int
			for _, v := range grid {
				if v == 2 {
					counter++
				}
			}
			return strconv.Itoa(counter), nil
		}
	}
}

type Game struct {
	Grid  [][]int
	Score int
}

func NewGame(width, height int) *Game {
	grid := make([][]int, height+1)
	for i := 0; i <= height; i++ {
		row := make([]int, width+1)
		for i := 0; i < len(row); i++ {
			row[i] = -1
		}
		grid[i] = row
	}
	return &Game{
		Grid: grid,
	}
}

func (g *Game) find(tile int) (int, int) {
	for y, row := range g.Grid {
		for x, t := range row {
			if t == tile {
				return y, x
			}
		}
	}
	return -1, -1
}

func (g *Game) FindBall() (int, int) {
	return g.find(3)
}

func (g *Game) FindPaddle() (int, int) {
	return g.find(4)
}

func (g *Game) Initialized() bool {
	for _, row := range g.Grid {
		for _, n := range row {
			if n == -1 {
				return false
			}
		}
	}
	return true
}

func (g *Game) Show() {
	lines := make([]string, len(g.Grid))
	for idx, row := range g.Grid {
		var line string
		for _, t := range row {
			switch t {
			case 0:
				line += " "
			case 1:
				line += "W"
			case 2:
				line += "#"
			case 3:
				line += "-"
			case 4:
				line += "+"
			default:
				line += " "
			}
		}
		lines[idx] = line
	}
	fmt.Printf("Score: %d\n", g.Score)
	fmt.Println(strings.Join(lines, "\n"))
}

func Part2(input string) (string, error) {
	icode, err := intcode.New(input)
	if err != nil {
		return "", err
	}

	icode[0] = 2

	in := make(chan int64, 1)
	out := make(chan intcode.Message, 100)

	go intcode.RunSync(icode, in, out)

	game := NewGame(43, 20)
	for {
		msg := <-out
		switch msg.Type {
		case intcode.MessageWantsInput:
			_, ballX := game.FindBall()
			_, paddleX := game.FindPaddle()
			var move int
			if (ballX - paddleX) > 0 {
				move = -1
			} else if (ballX - paddleX) < 0 {
				move = 1
			}
			in <- int64(move)
		case intcode.MessageOutput:
			x := int(msg.Value)

			msg = <-out
			y := int(msg.Value)

			msg = <-out
			third := int(msg.Value)

			if x == -1 && y == 0 {
				game.Score = third
			} else {
				game.Grid[y][x] = third
			}
		case intcode.MessageHalt:
			return strconv.Itoa(game.Score), nil
		}
	}
}
