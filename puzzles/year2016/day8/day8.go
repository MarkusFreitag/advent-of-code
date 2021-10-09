package day8

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Display struct {
	pixels [][]bool
}

func NewDisplay(width, height int) *Display {
	d := Display{pixels: make([][]bool, height)}
	for idx := range d.pixels {
		d.pixels[idx] = make([]bool, width)
	}
	return &d
}

func (d *Display) Show() {
	for _, row := range d.pixels {
		for _, col := range row {
			if col {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func (d *Display) Count(state bool) int {
	var count int
	for _, row := range d.pixels {
		for _, col := range row {
			if col == state {
				count++
			}
		}
	}
	return count
}

func (d *Display) Do(action string) {
	if strings.HasPrefix(action, "rect") {
		matches := rectCmdRgx.FindAllStringSubmatch(action, -1)[0]
		w, _ := strconv.Atoi(matches[1])
		h, _ := strconv.Atoi(matches[2])
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				d.pixels[y][x] = true
			}
		}
	} else if strings.HasPrefix(action, "rotate") {
		matches := rotateCmdRgx.FindAllStringSubmatch(action, -1)[0]
		index, _ := strconv.Atoi(matches[2])
		times, _ := strconv.Atoi(matches[3])
		switch matches[1] {
		case "row":
			d.pixels[index] = rotate(d.pixels[index], times)
		case "column":
			col := make([]bool, len(d.pixels))
			for idx, row := range d.pixels {
				col[idx] = row[index]
			}
			col = rotate(col, times)
			for idx, value := range col {
				d.pixels[idx][index] = value
			}
		}
	}
}

var (
	displayWidth  = 50
	displayHeight = 6
	rectCmdRgx    = regexp.MustCompile(`^rect\s(\d+)x(\d+)$`)
	rotateCmdRgx  = regexp.MustCompile(`^rotate\s(row|column)\s[x-y]=(\d+)\sby\s(\d+)$`)
)

func rotate(slice []bool, times int) []bool {
	l := len(slice)
	pos := (l + times) % l
	buf := make([]bool, l)
	for _, v := range slice {
		buf[pos] = v
		pos++
		if pos == l {
			pos = 0
		}
	}
	return buf
}

func Part1(input string) (string, error) {
	display := NewDisplay(displayWidth, displayHeight)
	for _, line := range strings.Split(input, "\n") {
		display.Do(line)
	}
	return strconv.Itoa(display.Count(true)), nil
}

func Part2(input string) (string, error) {
	display := NewDisplay(displayWidth, displayHeight)
	for _, line := range strings.Split(input, "\n") {
		display.Do(line)
	}
	display.Show()
	return "see output", nil
}
