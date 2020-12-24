package day24

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var rgxMove = regexp.MustCompile("(e|se|sw|w|nw|ne)")

func Part1(input string) (string, error) {
	blackTiles := make(map[string]bool)
	for _, line := range strings.Fields(input) {
		moves := rgxMove.FindAllString(line, -1)
		if len(moves) == 0 {
			continue
		}
		var y, x int
		for _, move := range moves {
			switch move {
			case "e":
				x++
			case "se":
				y++
				x++
			case "sw":
				y++
			case "w":
				x--
			case "nw":
				y--
				x--
			case "ne":
				y--
			}
		}
		pos := fmt.Sprintf("%d|%d", y, x)
		if _, ok := blackTiles[pos]; ok {
			delete(blackTiles, pos)
		} else {
			blackTiles[pos] = true
		}
	}

	return strconv.Itoa(len(blackTiles)), nil
}

func Part2(input string) (string, error) {
	tiles := make(map[string]bool)
	for _, line := range strings.Fields(input) {
		moves := rgxMove.FindAllString(line, -1)
		if len(moves) == 0 {
			continue
		}
		var y, x int
		for _, move := range moves {
			switch move {
			case "e":
				x++
			case "se":
				y++
				x++
			case "sw":
				y++
			case "w":
				x--
			case "nw":
				y--
				x--
			case "ne":
				y--
			}
		}
		pos := fmt.Sprintf("%d|%d", y, x)
		if v, ok := tiles[pos]; ok {
			tiles[pos] = !v
		} else {
			tiles[pos] = true
		}
	}

	for i := 1; i <= 100; i++ {
		var minY, minX, maxY, maxX int
		for pos := range tiles {
			parts := strings.Split(pos, "|")
			y, _ := strconv.Atoi(parts[0])
			x, _ := strconv.Atoi(parts[1])
			if y < minY {
				minY = y
			} else if y > maxY {
				maxY = y
			}
			if x < minX {
				minX = x
			} else if x > maxX {
				maxX = x
			}
		}

		for y := minY - 2; y < maxY+2; y++ {
			for x := minX - 2; x < maxX+2; x++ {
				pos := fmt.Sprintf("%d|%d", y, x)
				if _, ok := tiles[pos]; !ok {
					tiles[pos] = false
				}
			}
		}

		tmp := make(map[string]bool)
		for pos, black := range tiles {
			tmp[pos] = black
		}

		for pos, black := range tiles {
			var count int
			for _, p := range neighbours(pos) {
				if black, ok := tiles[p]; ok && black {
					count++
				}
			}
			if black && (count == 0 || count > 2) {
				tmp[pos] = false
			}
			if !black && count == 2 {
				tmp[pos] = true
			}
		}

		tiles = tmp
		/*
		   if i<10 || i%10==0 {
		     fmt.Printf("Day %d: %d\n", i, countBlack(tiles))
		   }
		*/
	}

	return strconv.Itoa(countBlack(tiles)), nil
}

func countBlack(tiles map[string]bool) int {
	var count int
	for _, black := range tiles {
		if black {
			count++
		}
	}
	return count
}

func neighbours(pos string) []string {
	parts := strings.Split(pos, "|")
	y, _ := strconv.Atoi(parts[0])
	x, _ := strconv.Atoi(parts[1])
	return []string{
		fmt.Sprintf("%d|%d", y, x+1),
		fmt.Sprintf("%d|%d", y+1, x+1),
		fmt.Sprintf("%d|%d", y+1, x),
		fmt.Sprintf("%d|%d", y, x-1),
		fmt.Sprintf("%d|%d", y-1, x-1),
		fmt.Sprintf("%d|%d", y-1, x),
	}
}
