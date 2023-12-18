package day18

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

var (
	neighs = [][2]int{
		[2]int{-1, 0},
		[2]int{0, 1},
		[2]int{1, 0},
		[2]int{0, -1},
	}
)

func fill(grid [][]bool, pts, seen map[[2]int]bool, pt [2]int) [][]bool {
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
		grid[p[0]][p[1]] = true
		grid = fill(grid, pts, seen, p)
	}
	return grid
}

func Part1(input string) (string, error) {
	/*
		minY, minX := numbers.MaxInteger, numbers.MaxInteger
		maxY, maxX := 0, 0
		pos := [2]int{0, 0}
		for _, line := range strings.Split(input, "\n") {
			fields := strings.Fields(line)
			switch fields[0] {
			case "U":
				pos[0] = pos[0] - util.ParseInt(fields[1])
			case "R":
				pos[1] = pos[1] + util.ParseInt(fields[1])
			case "D":
				pos[0] = pos[0] + util.ParseInt(fields[1])
			case "L":
				pos[1] = pos[1] - util.ParseInt(fields[1])
			}

			minY = numbers.Min(minY, pos[0])
			minX = numbers.Min(minX, pos[1])
			maxY = numbers.Max(maxY, pos[0])
			maxX = numbers.Max(maxX, pos[1])

		}
		fmt.Printf("%d|%d => %d|%d\n", minY, minX, maxY, maxX)
	*/
	pos := [2]int{390, 90}
	grid := make([][]bool, 500)
	for idx := range grid {
		grid[idx] = make([]bool, 300)
	}
	grid[390][90] = true
	points := make(map[[2]int]bool)
	points[[2]int{390, 90}] = true
	for _, line := range strings.Split(input, "\n") {
		fields := strings.Fields(line)
		switch fields[0] {
		case "U":
			for off := 0; off < util.ParseInt(fields[1]); off++ {
				pos[0] = pos[0] - 1
				grid[pos[0]][pos[1]] = true
				points[pos] = true
			}
		case "R":
			for off := 0; off < util.ParseInt(fields[1]); off++ {
				pos[1] = pos[1] + 1
				grid[pos[0]][pos[1]] = true
				points[pos] = true
			}
		case "D":
			for off := 0; off < util.ParseInt(fields[1]); off++ {
				pos[0] = pos[0] + 1
				grid[pos[0]][pos[1]] = true
				points[pos] = true
			}
		case "L":
			for off := 0; off < util.ParseInt(fields[1]); off++ {
				pos[1] = pos[1] - 1
				grid[pos[0]][pos[1]] = true
				points[pos] = true
			}
		}
	}

	grid = fill(grid, points, make(map[[2]int]bool), [2]int{390, 90})

	for _, row := range grid {
		var line string
		for _, cell := range row {
			if cell {
				line += "#"
			} else {
				line += "."
			}
		}
		fmt.Println(line)
	}

	var count int
	for _, row := range grid {
		for _, cell := range row {
			if !cell {
				count++
			}
		}
	}
	return strconv.Itoa(count + len(points)), nil
}

func Part2(input string) (string, error) {
	/*
		minY, minX := numbers.MaxInteger, numbers.MaxInteger
		maxY, maxX := 0, 0
		pos := [2]int{0, 0}
		for _, line := range strings.Split(input, "\n") {
			fields := strings.Fields(line)

			hex := strings.TrimPrefix(strings.TrimSuffix(strings.TrimPrefix(fields[2], "("), ")"), "#")
			dir, _ := strconv.ParseInt(hex[5:], 16, 64)
			meters, _ := strconv.ParseInt(hex[:5], 16, 64)

			fmt.Printf("%s %d %d\n", hex, meters, dir)

			switch dir {
			case 3:
				pos[0] = pos[0] - int(meters)
			case 0:
				pos[1] = pos[1] + int(meters)
			case 1:
				pos[0] = pos[0] + int(meters)
			case 2:
				pos[1] = pos[1] - int(meters)
			}

			minY = numbers.Min(minY, pos[0])
			minX = numbers.Min(minX, pos[1])
			maxY = numbers.Max(maxY, pos[0])
			maxX = numbers.Max(maxX, pos[1])

		}
		fmt.Printf("%d|%d => %d|%d\n", minY, minX, maxY, maxX)
		return "na", nil
	*/
	pos := [2]int{19141420, 4698860}
	grid := make([][]bool, 19141420+623860)
	for idx := range grid {
		grid[idx] = make([]bool, 4698860+13464692)
	}
	grid[19141420][4698860] = true
	points := make(map[[2]int]bool)
	points[[2]int{19141420, 4698860}] = true
	for _, line := range strings.Split(input, "\n") {
		fields := strings.Fields(line)
		switch fields[0] {
		case "U":
			for off := 0; off < util.ParseInt(fields[1]); off++ {
				pos[0] = pos[0] - 1
				grid[pos[0]][pos[1]] = true
				points[pos] = true
			}
		case "R":
			for off := 0; off < util.ParseInt(fields[1]); off++ {
				pos[1] = pos[1] + 1
				grid[pos[0]][pos[1]] = true
				points[pos] = true
			}
		case "D":
			for off := 0; off < util.ParseInt(fields[1]); off++ {
				pos[0] = pos[0] + 1
				grid[pos[0]][pos[1]] = true
				points[pos] = true
			}
		case "L":
			for off := 0; off < util.ParseInt(fields[1]); off++ {
				pos[1] = pos[1] - 1
				grid[pos[0]][pos[1]] = true
				points[pos] = true
			}
		}
	}

	grid = fill(grid, points, make(map[[2]int]bool), [2]int{390, 90})

	var count int
	for _, row := range grid {
		for _, cell := range row {
			if !cell {
				count++
			}
		}
	}
	return strconv.Itoa(count + len(points)), nil
}
