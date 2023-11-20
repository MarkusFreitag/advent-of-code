package day9

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
	"github.com/MarkusFreitag/advent-of-code/util/sliceutil"
)

func Part1(input string) (string, error) {
	lines := strings.Fields(input)
	points := make([][]int, len(lines))
	for idx, line := range lines {
		points[idx] = util.StringsToInts(util.StringToStrings(line))
	}

	var sum int
	for y, row := range points {
		for x, col := range row {
			b := make([]bool, 0)
			if y-1 >= 0 {
				b = append(b, col < points[y-1][x])
			}
			if y+1 < len(points) {
				b = append(b, col < points[y+1][x])
			}

			if x-1 >= 0 {
				b = append(b, col < row[x-1])
			}
			if x+1 < len(row) {
				b = append(b, col < row[x+1])
			}

			if sliceutil.All(b, true) {
				sum += col + 1
			}
		}
	}

	return strconv.Itoa(sum), nil
}

func strPt(y, x int) string { return fmt.Sprintf("%d|%d", y, x) }

func checkNeighbours(points [][]int, y, x int, visited map[string]bool) int {
	if _, ok := visited[strPt(y, x)]; ok || points[y][x] == 9 {
		return 0
	}

	size := 1
	visited[strPt(y, x)] = true
	if y-1 >= 0 {
		size += checkNeighbours(points, y-1, x, visited)
	}
	if y+1 < len(points) {
		size += checkNeighbours(points, y+1, x, visited)
	}

	if x-1 >= 0 {
		size += checkNeighbours(points, y, x-1, visited)
	}
	if x+1 < len(points[y]) {
		size += checkNeighbours(points, y, x+1, visited)
	}
	return size
}

func Part2(input string) (string, error) {
	lines := strings.Fields(input)
	points := make([][]int, len(lines))
	for idx, line := range lines {
		points[idx] = util.StringsToInts(util.StringToStrings(line))
	}

	sizes := make([]int, 0)
	for y, row := range points {
		for x, col := range row {
			b := make([]bool, 0)
			if y-1 >= 0 {
				b = append(b, col < points[y-1][x])
			}
			if y+1 < len(points) {
				b = append(b, col < points[y+1][x])
			}

			if x-1 >= 0 {
				b = append(b, col < row[x-1])
			}
			if x+1 < len(row) {
				b = append(b, col < row[x+1])
			}

			if sliceutil.All(b, true) {
				visited := make(map[string]bool)
				sizes = append(sizes, checkNeighbours(points, y, x, visited))
			}
		}
	}

	sort.Ints(sizes)

	return strconv.Itoa(numbers.Multiply(sizes[len(sizes)-3:]...)), nil
}
