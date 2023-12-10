package day10

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util/sliceutil"
)

var (
	neighs = [][2]int{
		[2]int{-1, 0},
		[2]int{0, 1},
		[2]int{1, 0},
		[2]int{0, -1},
	}
)

func getDir(a, b [2]int) string {
	if a[0] == b[0] {
		if a[1] > b[1] {
			return "west"
		}
		if a[1] < b[1] {
			return "east"
		}
	}
	if a[1] == b[1] {
		if a[0] > b[0] {
			return "north"
		}
		if a[0] < b[0] {
			return "south"
		}
	}
	panic(fmt.Sprintf("same points %v %v", a, b))
}

func searchLoopedPipe(grid [][]rune) [][2]int {
	pipe := make([][2]int, 1)
	for y, row := range grid {
		var found bool
		for x, cell := range row {
			if cell == 'S' {
				pipe[0] = [2]int{y, x}
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	for idx, neigh := range neighs {
		y := pipe[0][0] + neigh[0]
		if y < 0 || y >= len(grid) {
			continue
		}
		x := pipe[0][1] + neigh[1]
		if x < 0 || x >= len(grid[y]) {
			continue
		}
		if !strings.ContainsRune("|-LJ7F", grid[y][x]) {
			continue
		}
		switch idx {
		case 0:
			if strings.ContainsRune("|7F", grid[y][x]) {
				pipe = append(pipe, [2]int{y, x})
				break
			}
		case 1:
			if strings.ContainsRune("-7J", grid[y][x]) {
				pipe = append(pipe, [2]int{y, x})
				break
			}
		case 2:
			if strings.ContainsRune("|JL", grid[y][x]) {
				pipe = append(pipe, [2]int{y, x})
				break
			}
		case 3:
			if strings.ContainsRune("-FL", grid[y][x]) {
				pipe = append(pipe, [2]int{y, x})
				break
			}
		}
	}

	cur := pipe[len(pipe)-1]
	facing := getDir(pipe[0], cur)
	for {
		var next [2]int
		switch grid[cur[0]][cur[1]] {
		case '|':
			if facing == "south" {
				next = [2]int{cur[0] + 1, cur[1]}
			} else if facing == "north" {
				next = [2]int{cur[0] - 1, cur[1]}
			}
		case '-':
			if facing == "east" {
				next = [2]int{cur[0], cur[1] + 1}
			} else if facing == "west" {
				next = [2]int{cur[0], cur[1] - 1}
			}
		case 'L':
			if facing == "west" {
				next = [2]int{cur[0] - 1, cur[1]}
			} else if facing == "south" {
				next = [2]int{cur[0], cur[1] + 1}
			}
		case 'J':
			if facing == "south" {
				next = [2]int{cur[0], cur[1] - 1}
			} else if facing == "east" {
				next = [2]int{cur[0] - 1, cur[1]}
			}
		case '7':
			if facing == "east" {
				next = [2]int{cur[0] + 1, cur[1]}
			} else if facing == "north" {
				next = [2]int{cur[0], cur[1] - 1}
			}
		case 'F':
			if facing == "west" {
				next = [2]int{cur[0] + 1, cur[1]}
			} else if facing == "north" {
				next = [2]int{cur[0], cur[1] + 1}
			}
		}
		if grid[next[0]][next[1]] == 'S' {
			break
		}
		facing = getDir(cur, next)
		pipe = append(pipe, next)
		cur = next
	}
	return pipe
}

func parseInput(input string) [][]rune {
	lines := strings.Split(input, "\n")
	grid := make([][]rune, len(lines))
	for idx, line := range lines {
		grid[idx] = []rune(line)
	}
	return grid
}

func Part1(input string) (string, error) {
	pipe := searchLoopedPipe(parseInput(input))
	return strconv.Itoa(len(pipe) / 2), nil
}

func grow(grid [][]rune) [][]rune {
	largeGrid := make([][]rune, len(grid)*3)
	for idx := range largeGrid {
		row := make([]rune, len(grid[0])*3)
		sliceutil.Fill(row, ' ')
		largeGrid[idx] = row
	}
	for y, row := range grid {
		for x, cell := range row {
			largeGrid[3*y+1][3*x+1] = cell

			switch cell {
			case '|':
				largeGrid[3*y][3*x+1] = '|'
				largeGrid[3*y+2][3*x+1] = '|'
			case '-':
				largeGrid[3*y+1][3*x] = '-'
				largeGrid[3*y+1][3*x+2] = '-'
			case 'L':
				largeGrid[3*y][3*x+1] = '|'
				largeGrid[3*y+1][3*x+2] = '-'
			case 'J':
				largeGrid[3*y][3*x+1] = '|'
				largeGrid[3*y+1][3*x] = '-'
			case '7':
				largeGrid[3*y+1][3*x] = '-'
				largeGrid[3*y+2][3*x+1] = '|'
			case 'F':
				largeGrid[3*y+1][3*x+2] = '-'
				largeGrid[3*y+2][3*x+1] = '|'
			case '.':
			case 'S':
				if y-1 >= 0 {
					if r := grid[y-1][x]; r == '|' || r == '7' || r == 'F' {
						largeGrid[3*y][3*x+1] = '|'
					}
				}
				if y+1 < len(grid) {
					if r := grid[y+1][x]; r == '|' || r == 'L' || r == 'J' {
						largeGrid[3*y+2][3*x+1] = '|'
					}
				}
				if x-1 >= 0 {
					if r := grid[y][x-1]; r == '-' || r == 'L' || r == 'F' {
						largeGrid[3*y+1][3*x] = '-'
					}
				}
				if x+1 < len(row) {
					if r := grid[y][x+1]; r == '-' || r == 'J' || r == '7' {
						largeGrid[3*y+1][3*x+2] = '-'
					}
				}
			}
		}
	}
	return largeGrid
}

func fill(grid [][]rune, pts, seen map[[2]int]bool, pt [2]int, char rune) [][]rune {
	for _, dir := range neighs {
		p := [2]int{pt[0] + dir[0], pt[1] + dir[1]}
		if p[0] < 0 || p[1] < 0 || p[0] >= len(grid) || p[1] >= len(grid[0]) {
			continue
		}
		if _, ok := seen[p]; ok {
			continue
		}
		if _, ok := pts[p]; ok {
			continue
		}
		seen[p] = true
		grid[p[0]][p[1]] = char
		grid = fill(grid, pts, seen, p, char)
	}
	return grid
}

func Part2(input string) (string, error) {
	grid := grow(parseInput(input))

	pipe := searchLoopedPipe(grid)
	pts := make(map[[2]int]bool)
	for _, p := range pipe {
		pts[p] = true
	}

	grid = fill(grid, pts, make(map[[2]int]bool), [2]int{0, 0}, ' ')

	for pt := range pts {
		grid[pt[0]][pt[1]] = ' '
	}

	var count int
	for y := 1; y < len(grid); y += 3 {
		for x := 1; x < len(grid[0]); x += 3 {
			if grid[y][x] != ' ' {
				count++
			}
		}
	}

	return strconv.Itoa(count), nil
}
