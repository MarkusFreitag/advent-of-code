package day17

import (
	"container/heap"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

const (
	up int = iota
	right
	down
	left
)

var dirs = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

type Item struct {
	pos      [2]int
	dir      int
	dirCount int
	loss     int
	index    int
}

func (i *Item) Key() [4]int {
	return [4]int{i.pos[0], i.pos[1], i.dir, i.dirCount}
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].loss < pq[j].loss
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func search(g [][]int, start, end [2]int, part2 bool) int {
	queue := make(PriorityQueue, 1)
	queue[0] = &Item{
		pos:      [2]int{0, 0},
		dir:      -1,
		dirCount: -1,
		loss:     0,
	}
	heap.Init(&queue)
	seen := make(map[[4]int]int)
	minLoss := numbers.MaxInteger
	for queue.Len() > 0 {
		item := heap.Pop(&queue).(*Item)
		if item.pos == end {
			minLoss = numbers.Min(minLoss, item.loss)
			continue
		}
		if _, ok := seen[item.Key()]; ok {
			continue
		}
		seen[item.Key()] = item.loss
		for dirIdx, off := range dirs {
			newPos := [2]int{item.pos[0] + off[0], item.pos[1] + off[1]}
			if newPos[0] < 0 || newPos[0] >= len(g) || newPos[1] < 0 || newPos[1] >= len(g[0]) {
				continue
			}

			if item.dir == up && dirIdx == down {
				continue
			}
			if item.dir == right && dirIdx == left {
				continue
			}
			if item.dir == down && dirIdx == up {
				continue
			}
			if item.dir == left && dirIdx == right {
				continue
			}

			newDirCount := item.dirCount + 1
			if item.dir != dirIdx {
				newDirCount = 1
			}

			var validMove bool
			if part2 {
				validMove = newDirCount <= 10 && (dirIdx == item.dir || item.dirCount >= 4 || item.dirCount == -1)
			} else {
				validMove = newDirCount <= 3
			}

			if validMove {
				heap.Push(&queue, &Item{
					pos:      newPos,
					dir:      dirIdx,
					dirCount: newDirCount,
					loss:     item.loss + g[newPos[0]][newPos[1]],
				})
			}
		}

	}
	return minLoss
}

func parseInput(input string) [][]int {
	lines := strings.Fields(input)
	grid := make([][]int, len(lines))
	for idx, line := range lines {
		row := make([]int, len(line))
		for i, c := range line {
			row[i] = util.ParseInt(string(c))
		}
		grid[idx] = row
	}
	return grid
}

func Part1(input string) (string, error) {
	return strconv.Itoa(
		search(
			parseInput(input),
			[2]int{0, 0},
			[2]int{len(grid) - 1, len(grid[0]) - 1},
			false,
		),
	), nil
}

func Part2(input string) (string, error) {
	return strconv.Itoa(
		search(
			parseInput(input),
			[2]int{0, 0},
			[2]int{len(grid) - 1, len(grid[0]) - 1},
			true,
		),
	), nil
}
