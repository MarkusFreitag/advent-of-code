package day19

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/MarkusFreitag/advent-of-code/pkg/intcode"
)

type Grid [][]string

func NewGrid(width, height int) Grid {
	grid := make(Grid, height)
	for idx := range grid {
		row := make([]string, width)
		for i := range row {
			row[i] = "x"
		}
		grid[idx] = row
	}
	return grid
}

func (g Grid) Show() {
	lines := make([]string, len(g))
	for idx, row := range g {
		lines[idx] = strings.Join(row, "")
	}
	fmt.Println(strings.Join(lines, "\n"))
}

type Point struct {
	Y, X, State int
}

func checkPoint(wg *sync.WaitGroup, icode intcode.IntCode, p *Point) {
	in := make(chan int64, 2)
	in <- int64(p.X)
	in <- int64(p.Y)
	out := make(chan intcode.Message, 100)
	go intcode.RunAsync(icode, in, out)

	msg := <-out
	p.State = int(msg.Value)
	wg.Done()
}

type Part1 struct{}

func (p *Part1) Solve(input string) (string, error) {
	icode, err := intcode.New(input)
	if err != nil {
		return "", err
	}

	width := 50
	height := 50

	points := make([]*Point, 0)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			points = append(points, &Point{Y: y, X: x})
		}
	}

	var wg sync.WaitGroup
	for _, pt := range points {
		wg.Add(1)
		go checkPoint(&wg, icode, pt)
	}
	wg.Wait()

	grid := NewGrid(width, height)
	for _, pt := range points {
		switch pt.State {
		case 0:
			grid[pt.Y][pt.X] = "."
		case 1:
			grid[pt.Y][pt.X] = "#"
		}
	}
	grid.Show()

	var total int
	for _, row := range grid {
		for _, c := range row {
			if c == "#" {
				total++
			}
		}
	}

	return strconv.Itoa(total), nil
}

type Part2 struct{}

func (p *Part2) Solve(input string) (string, error) {
	return "n/a", nil
}
