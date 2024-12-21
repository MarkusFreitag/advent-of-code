package day21

import (
	"fmt"
	"image"
	"iter"
	"slices"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/directions"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

var (
	numericKeypad = [][]rune{
		[]rune{'7', '8', '9'},
		[]rune{'4', '5', '6'},
		[]rune{'1', '2', '3'},
		[]rune{'X', '0', 'A'},
	}
	directionalKeypad = [][]rune{
		[]rune{'X', '^', 'A'},
		[]rune{'<', 'v', '>'},
	}
	numericPaths     = make(map[rune]map[rune][]string)
	directionalPaths = make(map[rune]map[rune][]string)
)

func init() {
	numericPaths = shortestPaths(numericKeypad)
	directionalPaths = shortestPaths(directionalKeypad)
}

type GridItem struct {
	Pt  image.Point
	Val rune
}

func gridSeq(grid [][]rune) iter.Seq[GridItem] {
	return func(yield func(GridItem) bool) {
		for y, row := range grid {
			for x, col := range row {
				if !yield(GridItem{Pt: image.Pt(x, y), Val: col}) {
					return
				}
			}
		}
	}
}

func neighbours(grid [][]rune) util.NeighboursFunc[Move] {
	border := image.Rect(0, 0, len(grid[0]), len(grid))
	return func(m Move) iter.Seq[Move] {
		return func(yield func(Move) bool) {
			for neighbour := range directions.Moves() {
				np := m.Pos.Add(neighbour.Point())
				if np.In(border) && grid[np.Y][np.X] != 'X' {
					if !yield(Move{Pos: np, Dir: neighbour}) {
						return
					}
				}
			}
		}
	}
}

type Move struct {
	Pos image.Point
	Dir directions.Move
}

func shortestPaths(grid [][]rune) map[rune]map[rune][]string {
	shortest := make(map[rune]map[rune][]string)
	for from := range gridSeq(grid) {
		if from.Val == 'X' {
			continue
		}
		for to := range gridSeq(grid) {
			if to.Val == 'X' {
				continue
			}
			var seqs []string
			if from.Val == to.Val {
				// Even if we are already at the correct button, we need to confirm it by pressing A
				seqs = []string{"A"}
			} else {
				// The shortest path will nether be longer than the manhattan distance
				paths := util.AllPathsWithScore(
					Move{Pos: from.Pt},
					util.FakeCost(neighbours(grid)),
					func(m Move) bool { return grid[m.Pos.Y][m.Pos.X] == to.Val },
					numbers.Abs(from.Pt.X-to.Pt.X)+numbers.Abs(from.Pt.Y-to.Pt.Y),
				)
				seqs = make([]string, len(paths))
				for idx, val := range paths {
					path := slices.Collect(val.Seq())
					slices.Reverse(path)
					var str string
					for _, m := range path[1:] {
						switch m.Dir {
						case directions.Up:
							str += "v"
						case directions.Right:
							str += ">"
						case directions.Down:
							str += "^"
						case directions.Left:
							str += "<"
						}
					}
					seqs[idx] = str + "A"
				}
			}
			val, ok := shortest[from.Val]
			if !ok {
				val = make(map[rune][]string)
			}
			val[to.Val] = seqs
			shortest[from.Val] = val
		}
	}
	return shortest
}

func sequenceLength(seq string, depth, maxDepth int, cache map[string]int) int {
	if depth > maxDepth {
		return len(seq)
	}
	key := fmt.Sprintf("[%d]%s", depth, seq)
	if length, ok := cache[key]; ok && length > 0 {
		return length
	}
	var shortestPaths map[rune]map[rune][]string
	if depth == 0 {
		shortestPaths = numericPaths
	} else {
		shortestPaths = directionalPaths
	}
	last := 'A'
	var total int
	for _, char := range seq {
		length := numbers.MaxInteger
		for _, path := range shortestPaths[last][char] {
			length = min(length, sequenceLength(path, depth+1, maxDepth, cache))
		}
		total += length
		last = char
	}
	cache[key] = total
	return total
}

func Part1(input string) (string, error) {
	var complexity int
	for _, code := range strings.Fields(input) {
		complexity += sequenceLength(code, 0, 2, make(map[string]int)) * util.ParseInt(strings.TrimSuffix(code, "A"))
	}
	return strconv.Itoa(complexity), nil
}

func Part2(input string) (string, error) {
	var complexity int
	for _, code := range strings.Fields(input) {
		complexity += sequenceLength(code, 0, 25, make(map[string]int)) * util.ParseInt(strings.TrimSuffix(code, "A"))
	}
	return strconv.Itoa(complexity), nil
}
