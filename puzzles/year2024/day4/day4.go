package day4

import (
	"image"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/directions"
	"github.com/MarkusFreitag/advent-of-code/util/sliceutil"
)

func extractWord(grid [][]rune, pos, dir image.Point, dist int) string {
	bounds := image.Rect(0, 0, len(grid[0]), len(grid))

	var word string
	for range dist {
		if pos.In(bounds) {
			word += string(grid[pos.Y][pos.X])
		}
		pos = pos.Add(dir)
		if !pos.In(bounds) {
			break
		}
	}
	return word
}

func Part1(input string) (string, error) {
	const xmas string = "XMAS"
	var xmasReverse = util.StringReverse(xmas)

	grid := make([][]rune, 0)
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, []rune(line))
	}

	var counter int
	for y, row := range grid {
		for x, char := range row {
			if char != 'X' && char != 'S' {
				continue
			}
			words := []string{
				extractWord(grid, image.Pt(x, y), directions.Right.Point(), 4),
				extractWord(grid, image.Pt(x, y), directions.Down.Point(), 4),
				extractWord(grid, image.Pt(x, y), directions.DownLeft.Point(), 4),
				extractWord(grid, image.Pt(x, y), directions.DownRight.Point(), 4),
			}
			counter += sliceutil.Count(words, xmas)
			counter += sliceutil.Count(words, xmasReverse)
		}
	}
	return strconv.Itoa(counter), nil
}

func Part2(input string) (string, error) {
	grid := make([][]rune, 0)
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, []rune(line))
	}

	var counter int
	for y, row := range grid {
		for x, char := range row {
			if char != 'A' {
				continue
			}
			uldr := extractWord(grid, image.Pt(x, y).Add(directions.UpLeft.Point()), directions.DownRight.Point(), 3)
			urdl := extractWord(grid, image.Pt(x, y).Add(directions.UpRight.Point()), directions.DownLeft.Point(), 3)
			if (uldr == "MAS" || uldr == "SAM") && (urdl == "MAS" || urdl == "SAM") {
				counter++
			}
		}
	}
	return strconv.Itoa(counter), nil
}
