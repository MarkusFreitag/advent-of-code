package grid

/*
var (
	squareStraightNeighbours = [4]Coord{
		Coord{1, 0},
		Coord{0, 1},
		Coord{-1, 0},
		Coord{0, -1},
	}
	squareDiagonalNeighbours = [4]Coord{
		Coord{1, 1},
		Coord{-1, 1},
		Coord{-1, -1},
		Coord{1, -1},
	}
)
*/

type squareGrid[C Coord[D], D Dimension, T any] struct {
	min, max     Coord[D]
	defaultValue T
	expandable   bool
	state        map[Coord[D]]T
}

func (sg *squareGrid[C, D, T]) withinBoundaries(coord C) bool {
	return coord.Between(sg.min, sg.max)
	/*
		if coord.X < sg.minX || coord.X > sg.maxX {
			return false
		}

		if coord.Y < sg.minY || coord.Y > sg.maxY {
			return false
		}

		return true
	*/
}

func (sg *squareGrid[C, D, T]) Clone() Grid[C, D, T] {
	clone := squareGrid[C, D, T]{
		min:          sg.min,
		max:          sg.max,
		state:        make(map[Coord[D]]T),
		expandable:   sg.expandable,
		defaultValue: sg.defaultValue,
	}
	for coord, value := range sg.state {
		clone.state[coord] = value
	}
	return &clone
}

func (sg *squareGrid[C, D, T]) Neighbours(coord C) []C {
	neighs := coord.DirectNeighbours()
	for idx, item := range neighs {
		neighs[idx] = item.Add(coord)
	}
	coords := make([]C, len(neighs))
	for idx, neigh := range neighs {
		coords[idx] = neigh.(C)
	}
	return coords
}

func (sg *squareGrid[C, D, T]) Expand(coord C) {
	sg.min = sg.min.Min(coord)
	sg.max = sg.max.Max(coord)
}

func (sg *squareGrid[C, D, T]) Get(coord C) T {
	if value, ok := sg.state[coord]; ok {
		return value
	}
	return sg.defaultValue
}

func (sg *squareGrid[C, D, T]) Lookup(coord C) (T, bool) {
	if value, ok := sg.state[coord]; ok {
		return value, true
	}
	return sg.defaultValue, false
}

func (sg *squareGrid[C, D, T]) Set(coord C, val T) {
	if !sg.withinBoundaries(coord) {
		if !sg.expandable {
			panic("moo")
		}
		sg.Expand(coord)
	}

	if val == sg.defaultValue {
		delete(sg.state, coord)
		return
	}
	sg.state[coord] = val
}

func (sg *squareGrid[C, D, T]) Count(value T) int {
	var total int
	countFn := func(_ C, v T) bool {
		if value == v {
			total++
		}
		return true
	}

	if value == sg.defaultValue {
		sg.Iterate(countFn)
		return total
	}

	sg.IterateChangedCells(countFn)
	return total
}

/*
func (sg *squareGrid[C, D, T]) Equal(other *squareGrid[C, D, T]) bool {
	if len(sg.state) != len(other.state) {
		return false
	}
	for coord, val := range sg.state {
		if otherVal, ok := other.state[coord]; !ok || val != otherVal {
			return false
		}
	}
	return true
}
*/

func (sg *squareGrid[C, D, T]) IterateChangedCells(fn IterateFn[C, D, T]) {
	for coord, val := range sg.state {
		if !fn(coord, val) {
			return
		}
	}
}

func (sg *squareGrid[C, D, T]) Iterate(fn IterateFn[C, D, T]) {
	for x := g.minX; x <= g.maxX; x++ {
		for y := g.minY; y <= g.maxY; y++ {
			coord := Coord{X: x, Y: y}
			if !fn(coord, sg.state[coord]) {
				return
			}
		}
	}
}

func (sg *squareGrid[C, D, T]) IterateBackwards(fn IterateFn[C, D, T]) {
	for x := g.maxX; x <= g.minX; x-- {
		for y := g.maxY; y <= g.minY; y-- {
			coord := Coord{X: x, Y: y}
			if !fn(coord, sg.state[coord]) {
				return
			}
		}
	}
}
