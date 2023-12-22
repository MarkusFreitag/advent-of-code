package day22

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
	"github.com/MarkusFreitag/advent-of-code/util/sliceutil"
)

type Point struct {
	X, Y, Z int
}

func (p Point) String() string {
	return fmt.Sprintf("%d,%d,%d", p.X, p.Y, p.Z)
}

type Points []Point

type Brick struct {
	ID    int
	Start Point
	End   Point
}

func parseBrick(id int, str string) Brick {
	parts := strings.Split(str, "~")
	start := util.StringsToInts(strings.Split(parts[0], ","))
	end := util.StringsToInts(strings.Split(parts[1], ","))
	return Brick{
		ID:    id,
		Start: Point{X: start[0], Y: start[1], Z: start[2]},
		End:   Point{X: end[0], Y: end[1], Z: end[2]},
	}
}

func (b Brick) String() string {
	return fmt.Sprintf("%s~%s", b.Start, b.End)
}

func (b Brick) GroundPoints() Points {
	ground := numbers.Min(b.Start.Z, b.End.Z)
	points := make(Points, 0)
	for x := b.Start.X; x <= b.End.X; x++ {
		for y := b.Start.Y; y <= b.End.Y; y++ {
			points = append(points, Point{X: x, Y: y, Z: ground})
		}
	}
	return points
}

func (b Brick) Interfere(pt Point) bool {
	return numbers.Between(pt.X, b.Start.X, b.End.X) && numbers.Between(pt.Y, b.Start.Y, b.End.Y) && numbers.Between(pt.Z, b.Start.Z, b.End.Z)
}

type Bricks []Brick

func (b Bricks) Len() int           { return len(b) }
func (b Bricks) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b Bricks) Less(i, j int) bool { return b[i].Start.Z < b[j].Start.Z }

func (b Bricks) Peak() int {
	var peak int
	for _, brick := range b {
		peak = numbers.Max(peak, brick.End.Z)
	}
	return peak
}

func (b Bricks) Interfere(pt Point) (int, bool) {
	for _, brick := range b {
		if brick.Interfere(pt) {
			return brick.ID, true
		}
	}
	return -1, false
}

func letFall(bricks Bricks) (Bricks, int, map[int]map[int]bool, map[int]map[int]bool) {
	settled := make(Bricks, 0, len(bricks))
	supports := make(map[int]map[int]bool)
	supportFrom := make(map[int]map[int]bool)
	var fallCounter int
	for len(bricks) > 0 {
		var brick Brick
		brick, bricks = sliceutil.PopFront(bricks)
		supports[brick.ID] = make(map[int]bool)
		supportFrom[brick.ID] = make(map[int]bool)

		orig := Brick{
			ID:    brick.ID,
			Start: brick.Start,
			End:   brick.End,
		}

		if len(settled) == 0 {
			dropped := brick.Start.Z - 1
			brick.Start.Z = 1
			brick.End.Z -= dropped
		} else {
			peak := settled.Peak() + 1
			dropped := brick.Start.Z - peak
			brick.Start.Z = peak
			brick.End.Z -= dropped

			for {
				if brick.Start.Z <= 1 {
					break
				}
				newBrick := Brick{
					ID: brick.ID,
					Start: Point{
						X: brick.Start.X,
						Y: brick.Start.Y,
						Z: brick.Start.Z - 1,
					},
					End: Point{
						X: brick.End.X,
						Y: brick.End.Y,
						Z: brick.End.Z - 1,
					},
				}
				var interfere bool
				for _, pt := range newBrick.GroundPoints() {
					for _, b := range settled {
						if b.Interfere(pt) {
							supports[b.ID][newBrick.ID] = true
							supportFrom[newBrick.ID][b.ID] = true
							interfere = true
							break
						}
					}
				}
				if interfere {
					break
				}
				brick = newBrick
			}
		}

		if orig.String() != brick.String() {
			fallCounter++
		}

		settled = append(settled, brick)
	}
	return settled, fallCounter, supports, supportFrom
}

func simulateFalling(input string) (Bricks, map[int]bool) {
	bricks := make(Bricks, 0)
	for idx, line := range strings.Fields(input) {
		bricks = append(bricks, parseBrick(idx+1, line))
	}
	sort.Sort(bricks)

	var supports, supportFrom map[int]map[int]bool
	bricks, _, supports, supportFrom = letFall(bricks)

	safeBricks := make(map[int]bool)
	for _, brick := range bricks {
		if len(supports[brick.ID]) == 0 {
			safeBricks[brick.ID] = true
			continue
		}

		safe := true
		for id := range supports[brick.ID] {
			if len(supportFrom[id]) < 2 {
				safe = false
				break
			}
		}
		if safe {
			safeBricks[brick.ID] = true
		}
	}

	return bricks, safeBricks
}

func Part1(input string) (string, error) {
	_, safe := simulateFalling(input)

	return strconv.Itoa(len(safe)), nil
}

func Part2(input string) (string, error) {
	bricks, safe := simulateFalling(input)

	var total int
	for idx, brick := range bricks {
		if _, ok := safe[brick.ID]; ok {
			continue
		}
		newBricks := make(Bricks, len(bricks))
		copy(newBricks, bricks)
		_, fallCounter, _, _ := letFall(slices.Delete(newBricks, idx, idx+1))
		total += fallCounter
	}

	return strconv.Itoa(total), nil
}
