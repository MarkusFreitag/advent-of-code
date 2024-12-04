package day4

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(input string) (string, error) {
	grid := make([][]rune, 0)
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, []rune(line))
	}

	inBounds := func(y, x int) bool {
		if y < 0 || y >= len(grid) {
			return false
		}
		if x < 0 || x >= len(grid[0]) {
			return false
		}
		return true
	}

	words := make(map[string]struct{})
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			char := grid[y][x]
			if char != 'X' && char != 'S' {
				continue
			}
			// horizontal
			if inBounds(y, x+1) && inBounds(y, x+2) && inBounds(y, x+3) {
				if char == 'X' && grid[y][x+1] == 'M' && grid[y][x+2] == 'A' && grid[y][x+3] == 'S' {
					words[fmt.Sprintf("%d|%d-%d|%d-%d|%d-%d|%d", y, x, y, x+1, y, x+2, y, x+3)] = struct{}{}
				}
				if char == 'S' && grid[y][x+1] == 'A' && grid[y][x+2] == 'M' && grid[y][x+3] == 'X' {
					words[fmt.Sprintf("%d|%d-%d|%d-%d|%d-%d|%d", y, x, y, x+1, y, x+2, y, x+3)] = struct{}{}
				}
			}
			//vertically
			if inBounds(y+1, x) && inBounds(y+2, x) && inBounds(y+3, x) {
				if char == 'X' && grid[y+1][x] == 'M' && grid[y+2][x] == 'A' && grid[y+3][x] == 'S' {
					words[fmt.Sprintf("%d|%d-%d|%d-%d|%d-%d|%d", y, x, y+1, x, y+2, x, y+3, x)] = struct{}{}
				}
				if char == 'S' && grid[y+1][x] == 'A' && grid[y+2][x] == 'M' && grid[y+3][x] == 'X' {
					words[fmt.Sprintf("%d|%d-%d|%d-%d|%d-%d|%d", y, x, y+1, x, y+2, x, y+3, x)] = struct{}{}
				}
			}
			//diagonal
			//bottom left
			if inBounds(y+1, x-1) && inBounds(y+2, x-2) && inBounds(y+3, x-3) {
				if char == 'X' && grid[y+1][x-1] == 'M' && grid[y+2][x-2] == 'A' && grid[y+3][x-3] == 'S' {
					words[fmt.Sprintf("%d|%d-%d|%d-%d|%d-%d|%d", y, x, y+1, x-1, y+2, x-2, y+3, x-3)] = struct{}{}
				}
				if char == 'S' && grid[y+1][x-1] == 'A' && grid[y+2][x-2] == 'M' && grid[y+3][x-3] == 'X' {
					words[fmt.Sprintf("%d|%d-%d|%d-%d|%d-%d|%d", y, x, y+1, x-1, y+2, x-2, y+3, x-3)] = struct{}{}
				}
			}
			//bottom right
			if inBounds(y+1, x+1) && inBounds(y+2, x+2) && inBounds(y+3, x+3) {
				if char == 'X' && grid[y+1][x+1] == 'M' && grid[y+2][x+2] == 'A' && grid[y+3][x+3] == 'S' {
					words[fmt.Sprintf("%d|%d-%d|%d-%d|%d-%d|%d", y, x, y+1, x+1, y+2, x+2, y+3, x+3)] = struct{}{}
				}
				if char == 'S' && grid[y+1][x+1] == 'A' && grid[y+2][x+2] == 'M' && grid[y+3][x+3] == 'X' {
					words[fmt.Sprintf("%d|%d-%d|%d-%d|%d-%d|%d", y, x, y+1, x+1, y+2, x+2, y+3, x+3)] = struct{}{}
				}
			}

		}
	}

	return strconv.Itoa(len(words)), nil
}

func Part2(input string) (string, error) {
	grid := make([][]rune, 0)
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, []rune(line))
	}

	inBounds := func(y, x int) bool {
		if y < 0 || y >= len(grid) {
			return false
		}
		if x < 0 || x >= len(grid[0]) {
			return false
		}
		return true
	}

	words := make(map[string]struct{})
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if char := grid[y][x]; char != 'A' {
				continue
			}
			if inBounds(y-1, x-1) && inBounds(y-1, x+1) && inBounds(y+1, x+1) && inBounds(y+1, x-1) {
				ul := grid[y-1][x-1]
				ur := grid[y-1][x+1]
				br := grid[y+1][x+1]
				bl := grid[y+1][x-1]
				if (ul == 'M' && ur == 'M' && bl == 'S' && br == 'S') ||
					(ul == 'S' && ur == 'S' && bl == 'M' && br == 'M') ||
					(ul == 'M' && ur == 'S' && bl == 'M' && br == 'S') ||
					(ul == 'S' && ur == 'M' && bl == 'S' && br == 'M') {

					words[fmt.Sprintf("%d|%d-%d|%d-%d|%d-%d|%d-%d|%d", y, x, y-1, x-1, y-1, x+1, y+1, x+1, y+1, x-1)] = struct{}{}
				}
			}
		}
	}

	return strconv.Itoa(len(words)), nil
}
