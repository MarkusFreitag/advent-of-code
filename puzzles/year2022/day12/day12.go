package day12

import (
	"image"
	"image/color"
	"image/gif"
	"iter"
	"strconv"
	"strings"

	"github.com/shomali11/gridder"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/directions"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

func parseInput(input string) ([][]int, image.Rectangle, image.Point, image.Point) {
	var start, dest image.Point
	lines := strings.Split(input, "\n")
	grid := make([][]int, len(lines))
	for idx, line := range lines {
		if v := strings.Index(line, "S"); v != -1 {
			start = image.Pt(v, idx)
			line = strings.ReplaceAll(line, "S", "a")
		}
		if v := strings.Index(line, "E"); v != -1 {
			dest = image.Pt(v, idx)
			line = strings.ReplaceAll(line, "E", "z")
		}
		bs := []byte(line)
		grid[idx] = make([]int, len(bs))
		for i, b := range bs {
			grid[idx][i] = int(b)
		}
	}
	return grid, image.Rect(0, 0, len(grid[0]), len(grid)), start, dest
}

func neighbours(grid [][]int, border image.Rectangle) util.NeighboursFunc[image.Point] {
	return func(p image.Point) iter.Seq[image.Point] {
		return func(yield func(image.Point) bool) {
			for neighbour := range directions.Moves() {
				np := p.Add(neighbour.Point())
				if np.In(border) && grid[np.Y][np.X]-grid[p.Y][p.X] <= 1 {
					if !yield(np) {
						return
					}
				}
			}
		}
	}
}

func Part1(input string) (string, error) {
	grid, border, pos, dest := parseInput(input)

	imageConfig := gridder.ImageConfig{
		Width:  len(grid[0]) * 50,
		Height: len(grid) * 50,
		Name:   "2022-12-part1.png",
	}
	gridConfig := gridder.GridConfig{
		Rows:              len(grid),
		Columns:           len(grid[0]),
		LineStrokeWidth:   2,
		BorderStrokeWidth: 4,
	}

	gridd, err := gridder.New(imageConfig, gridConfig, gridder.WithGIF(&gif.GIF{LoopCount: -1}))
	if err != nil {
		return "", err
	}
	defer gridd.SaveGIF()

	gridd.PaintCell(pos.Y, pos.X, color.RGBA{G: 255, A: 255})
	gridd.PaintCell(dest.Y, dest.X, color.RGBA{R: 255, A: 255})

	goalFn := func(p image.Point) bool { return p.Eq(dest) }
	closedColor := color.RGBA{B: 255, A: 255}
	openedColor := color.RGBA{B: 255 / 2, A: 255}
	pathColor := color.RGBA{R: 255 / 2, G: 255 / 2, B: 255 / 2, A: 255}
	painterFn := func(p image.Point, s util.SearchState) {
		if p.Eq(pos) || p.Eq(dest) {
			return
		}
		switch s {
		case util.SearchStateOpened:
			gridd.PaintCell(p.Y, p.X, openedColor)
		case util.SearchStateClosed:
			gridd.PaintCell(p.Y, p.X, closedColor)
		}
	}

	result := util.Dijkstra(pos, painterFn, util.FakeCost(neighbours(grid, border)), goalFn)
	//result := util.BFS(pos, painterFn, neighbours(grid, border), goalFn)
	last := result.Value
	for node := range result.Seq() {
		if node.Eq(pos) || node.Eq(dest) {
			continue
		}
		gridd.PaintCell(node.Y, node.X, pathColor)
		gridd.DrawPath(last.Y, last.X, node.Y, node.X, gridder.PathConfig{Dashes: 0, StrokeWidth: 10, Color: color.Black})
		last = node
	}
	gridd.DrawPath(last.Y, last.X, pos.Y, pos.X, gridder.PathConfig{Dashes: 0, StrokeWidth: 10, Color: color.Black})
	return strconv.Itoa(result.Dist()), nil
}

func Part2(input string) (string, error) {
	grid, border, _, dest := parseInput(input)
	dist := numbers.MaxInteger
	goalFn := func(p image.Point) bool { return p.Eq(dest) }
	for y, row := range grid {
		for x, col := range row {
			if col == 97 {
				if sn := util.BFS(image.Pt(x, y), nil, neighbours(grid, border), goalFn); sn != nil {
					dist = min(dist, sn.Dist())
				}
			}
		}
	}
	return strconv.Itoa(dist), nil
}
