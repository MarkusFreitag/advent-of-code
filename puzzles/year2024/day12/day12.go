package day12

import (
	"image"
	"iter"
	"slices"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/directions"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
	"github.com/MarkusFreitag/advent-of-code/util/sliceutil"
)

func calcPerimeter(pts []image.Point) int {
	if len(pts) == 1 {
		return 4
	}
	touching := 4 * len(pts)
	for idx, pt := range pts {
		for i, p := range pts {
			if idx == i {
				continue
			}
			if ((pt.X == p.X) && ((pt.Y == p.Y+1) || (pt.Y == p.Y-1))) ||
				((pt.Y == p.Y) && ((pt.X == p.X+1) || (pt.X == p.X-1))) {
				touching--
			}
		}
	}
	return touching
}

func calcSides(pts []image.Point) int {
	if len(pts) == 1 {
		return 4
	}

	var maxX, maxY int
	minX, minY := numbers.MaxInteger, numbers.MaxInteger
	for _, pt := range pts {
		minX, minY = min(pt.X, minX), min(pt.Y, minY)
		maxX, maxY = max(pt.X, maxX), max(pt.Y, maxY)
	}

	type Range struct {
		From, To int
	}

	var sides int
	for y := minY; y <= maxY; y++ {
		row := sliceutil.Map(pts, func(p image.Point) (image.Point, bool) { return p, p.Y == y })
		slices.SortFunc(row, func(a, b image.Point) int { return a.X - b.X })
		var borders [2][]*Range // track upper and lower border separately
		for _, pt := range row {
			for idx, dir := range []image.Point{directions.Down.Point(), directions.Up.Point()} {
				if p := pt.Add(dir); !slices.Contains(pts, p) {
					if len(borders[idx]) == 0 {
						borders[idx] = []*Range{&Range{p.X, p.X}}
					} else {
						last := borders[idx][len(borders[idx])-1]
						if last.To+1 == p.X {
							last.To = p.X
						} else {
							borders[idx] = append(borders[idx], &Range{p.X, p.X})
						}
					}
				}
			}
		}
		sides += len(borders[0])
		sides += len(borders[1])
	}

	for x := minX; x <= maxX; x++ {
		col := sliceutil.Map(pts, func(p image.Point) (image.Point, bool) { return p, p.X == x })
		slices.SortFunc(col, func(a, b image.Point) int { return a.Y - b.Y })
		var borders [2][]*Range // track left and right border separately
		for _, pt := range col {
			for idx, dir := range []image.Point{directions.Left.Point(), directions.Right.Point()} {
				if p := pt.Add(dir); !slices.Contains(pts, p) {
					if len(borders[idx]) == 0 {
						borders[idx] = []*Range{&Range{p.Y, p.Y}}
					} else {
						last := borders[idx][len(borders[idx])-1]
						if last.To+1 == p.Y {
							last.To = p.Y
						} else {
							borders[idx] = append(borders[idx], &Range{p.Y, p.Y})
						}
					}
				}
			}
		}
		sides += len(borders[0])
		sides += len(borders[1])
	}

	return sides
}

func calcCosts(input string, fenceFn func([]image.Point) int) int {
	grid := make([][]rune, 0)
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, []rune(line))
	}

	border := image.Rect(0, 0, len(grid[0]), len(grid))
	neighboursFn := func(tile rune) util.NeighboursFunc[image.Point] {
		return func(p image.Point) iter.Seq[image.Point] {
			return func(yield func(image.Point) bool) {
				for neighbour := range directions.Moves() {
					np := p.Add(neighbour.Point())
					if np.In(border) && grid[np.Y][np.X] == tile {
						if !yield(np) {
							return
						}
					}
				}
			}
		}
	}

	type Area []image.Point
	regions := make(map[rune][]Area)
	for y, row := range grid {
		for x, col := range row {
			rgs, ok := regions[col]
			if !ok {
				rgs = append(rgs, Area(util.Floodfill(image.Pt(x, y), neighboursFn(col))))
			} else {
				pt := image.Pt(x, y)
				var found bool
				for _, region := range rgs {
					if slices.Contains(region, pt) {
						found = true
						break
					}
				}
				if !found {
					rgs = append(rgs, Area(util.Floodfill(pt, neighboursFn(col))))
				}
			}
			regions[col] = rgs
		}
	}

	var total int
	for _, rgs := range regions {
		for _, area := range rgs {
			total += len(area) * fenceFn(area)
		}
	}
	return total
}

func Part1(input string) (string, error) {
	return strconv.Itoa(calcCosts(input, calcPerimeter)), nil
}

func Part2(input string) (string, error) {
	return strconv.Itoa(calcCosts(input, calcSides)), nil
}
