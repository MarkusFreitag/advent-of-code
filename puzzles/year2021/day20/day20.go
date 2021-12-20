package day20

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

var light = struct{}{}

type Point struct {
	X, Y int
}

type Image map[Point]struct{}

func NewImage() Image { return make(map[Point]struct{}) }

func ParseImage(str string) Image {
	img := NewImage()
	for y, line := range strings.Split(str, "\n") {
		for x, char := range line {
			if char == '#' {
				img[Point{X: x, Y: y}] = light
			}
		}
	}
	return img
}

func (img Image) Size() (int, int, int, int) {
	minX, minY := util.MaxInteger, util.MaxInteger
	var maxX, maxY int
	for pixel := range img {
		minX = util.MinInt(minX, pixel.X)
		minY = util.MinInt(minY, pixel.Y)
		maxX = util.MaxInt(maxX, pixel.X)
		maxY = util.MaxInt(maxY, pixel.Y)
	}
	return minX, minY, maxX, maxY
}

func (img Image) Show() {
	lines := make([]string, 0)
	startX, startY, width, height := img.Size()
	for y := startX - 1; y <= height+1; y++ {
		var row string
		for x := startY; x <= width; x++ {
			if _, ok := img[Point{X: x, Y: y}]; ok {
				row += "#"
			} else {
				row += "."
			}
		}
		lines = append(lines, row)
	}
	fmt.Println(strings.Join(lines, "\n"))
}

func (img Image) Enhance(str string, voidLight bool) Image {
	blinking := str[0] == '#'
	newImage := NewImage()
	minX, minY, maxX, maxY := img.Size()

	for y := minY - 1; y <= maxY+1; y++ {
		for x := minX - 1; x <= maxX+1; x++ {
			var idx int

			for m := y - 1; m <= y+1; m++ {
				for n := x - 1; n <= x+1; n++ {
					idx = idx << 1
					if m < minY || m > maxY || n < minX || n > maxX {
						if blinking && voidLight {
							idx |= 1
						}
						continue
					}

					if _, ok := img[Point{X: n, Y: m}]; ok {
						idx |= 1
					}
				}
			}

			if str[idx] == '#' {
				newImage[Point{X: x, Y: y}] = light
			}
		}
	}
	return newImage
}

func Part1(input string) (string, error) {
	blocks := strings.Split(input, "\n\n")

	img := ParseImage(blocks[1])

	for i := 0; i < 2; i++ {
		img = img.Enhance(blocks[0], i%2 == 1)
	}

	return strconv.Itoa(len(img)), nil
}

func Part2(input string) (string, error) {
	blocks := strings.Split(input, "\n\n")

	img := ParseImage(blocks[1])

	for i := 0; i < 50; i++ {
		img = img.Enhance(blocks[0], i%2 == 1)
	}

	return strconv.Itoa(len(img)), nil
}
