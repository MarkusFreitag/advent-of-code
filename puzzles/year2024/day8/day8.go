package day8

import (
	"image"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util/sliceutil"
)

func parseInput(input string) (map[rune][]image.Point, func(image.Point) bool) {
	antennas := make(map[rune][]image.Point)
	lines := strings.Split(input, "\n")
	for y, line := range lines {
		for x, char := range line {
			if char == '.' {
				continue
			}
			if val, ok := antennas[char]; ok {
				antennas[char] = append(val, image.Pt(x, y))
			} else {
				antennas[char] = []image.Point{image.Pt(x, y)}
			}
		}
	}
	return antennas, func(p image.Point) bool {
		if p.Y < 0 || p.Y >= len(lines) {
			return false
		}
		if p.X < 0 || p.X >= len(lines[0]) {
			return false
		}
		return true
	}
}

func Part1(input string) (string, error) {
	antennas, bounds := parseInput(input)
	antinodes := make(map[image.Point]struct{})
	for _, antenna := range antennas {
		if len(antenna) < 2 {
			continue
		}
		sliceutil.CombinationsFunc(antenna, 2, func(pair []image.Point) {
			if len(pair) != 2 {
				return
			}
			if pt := image.Pt(2*pair[1].X-pair[0].X, 2*pair[1].Y-pair[0].Y); bounds(pt) {
				antinodes[pt] = struct{}{}
			}
			if pt := image.Pt(2*pair[0].X-pair[1].X, 2*pair[0].Y-pair[1].Y); bounds(pt) {
				antinodes[pt] = struct{}{}
			}
		})
	}
	return strconv.Itoa(len(antinodes)), nil
}

func Part2(input string) (string, error) {
	antennas, bounds := parseInput(input)
	antinodes := make(map[image.Point]struct{})
	for _, antenna := range antennas {
		if len(antenna) < 2 {
			continue
		}
		sliceutil.CombinationsFunc(antenna, 2, func(pair []image.Point) {
			if len(pair) != 2 {
				return
			}

			diff := pair[0].Sub(pair[1])
			for pt := pair[0]; bounds(pt); pt = pt.Sub(diff) {
				antinodes[pt] = struct{}{}
			}
			for pt := pair[0]; bounds(pt); pt = pt.Add(diff) {
				antinodes[pt] = struct{}{}
			}
		})
	}
	return strconv.Itoa(len(antinodes)), nil
}
