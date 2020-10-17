package day15

import (
	"fmt"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/intcode"
)

const (
	NORTH = (iota + 1)
	SOUTH
	WEST
	EAST
)

type Vector2D struct {
	Y, X int
}

func (v Vector2D) String() string {
	return fmt.Sprintf("%d|%d", v.Y, v.X)
}

func (v Vector2D) Add(vec Vector2D) Vector2D {
	return Vector2D{
		Y: v.Y + vec.Y,
		X: v.X + vec.X,
	}
}

func (v Vector2D) Sub(vec Vector2D) Vector2D {
	return Vector2D{
		Y: v.Y - vec.Y,
		X: v.X - vec.X,
	}
}

func (v Vector2D) Next(dir int) Vector2D {
	var next Vector2D
	switch dir {
	case NORTH:
		next = Vector2D{
			Y: v.Y + 1,
			X: v.X,
		}
	case SOUTH:
		next = Vector2D{
			Y: v.Y - 1,
			X: v.X,
		}
	case WEST:
		next = Vector2D{
			Y: v.Y,
			X: v.X - 1,
		}
	case EAST:
		next = Vector2D{
			Y: v.Y,
			X: v.X + 1,
		}
	}
	return next
}

func show(world map[Vector2D]string) {
	lines := make([][]string, 60)
	for y := range lines {
		row := make([]string, 60)
		for x := range row {
			row[x] = " "
		}
		lines[y] = row
	}
	for pos, s := range world {
		lines[pos.Y+30][pos.X+30] = s
	}
	out := make([]string, 0)
	for _, line := range lines {
		out = append(out, strings.Join(line, ""))
	}
	fmt.Println(strings.Join(out, "\n"))
}

func Part1(input string) (string, error) {
	icode, err := intcode.New(input)
	if err != nil {
		return "", err
	}
	in := make(chan int64)
	out := make(chan intcode.Message, 100)
	go intcode.RunAsync(icode, in, out)

	world := make(map[Vector2D]string)
	pos := Vector2D{Y: 0, X: 0}
	var target *Vector2D
	var nextMove int
	var halted bool
	var moveCounter int
	for !halted {
		show(world)

		var inp int
		fmt.Scanf("%d", &inp)
		switch inp {
		case 2:
			nextMove = NORTH
		case 6:
			nextMove = EAST
		case 8:
			nextMove = SOUTH
		case 4:
			nextMove = WEST
		}

		in <- int64(nextMove)

		nextPos := pos.Next(nextMove)
		fmt.Printf("try to move from %s to %s...", pos.String(), nextPos.String())

		msg := <-out
		switch msg.Type {
		case intcode.MessageOutput:
			switch msg.Value {
			case 0:
				world[nextPos] = "#"
				fmt.Printf("could not move, try %d\n", nextMove)
			case 1:
				moveCounter++
				fmt.Print("ok\n")
				world[nextPos] = "."
				pos = nextPos
			case 2:
				fmt.Print("found target\n")
				target = &nextPos
				halted = true
			}
		case intcode.MessageHalt:
			halted = true
		}
	}
	fmt.Printf("%d %d\n", target.Y, target.X)
	fmt.Println(moveCounter + 1)
	return "n/a", nil
}

func Part2(input string) (string, error) {
	return "n/a", nil
}
