package day11

import (
	"strconv"
	"strings"
)

const (
	FLOOR    = '.'
	EMPTY    = 'L'
	OCCUPIED = '#'
)

type Grid [][]rune

func ParseGrid(layout string) Grid {
	grid := make(Grid, 0)
	for _, line := range strings.Split(layout, "\n") {
		cols := make([]rune, len(line))
		for idx, char := range line {
			cols[idx] = char
		}
		grid = append(grid, cols)
	}
	return grid
}

func (g Grid) Copy() Grid {
	dupl := make(Grid, len(g))
	for idx := range g {
		dupl[idx] = make([]rune, len(g[idx]))
		copy(dupl[idx], g[idx])
	}
	return dupl
}

func (g Grid) Count(char rune) int {
	var count int
	for _, row := range g {
		for _, col := range row {
			if col == char {
				count++
			}
		}
	}
	return count
}

func (g Grid) ValidPos(rIdx, cIdx int) bool {
	if rIdx < 0 || rIdx >= len(g) {
		return false
	}
	if cIdx < 0 || cIdx >= len(g[rIdx]) {
		return false
	}
	return true
}

func (g Grid) Neighbours(rIdx, cIdx int, inSight bool) (int, int) {
	var empty, occupied int
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			a, b := rIdx+i, cIdx+j
			if !g.ValidPos(a, b) {
				continue
			}
			switch g[a][b] {
			case EMPTY:
				empty++
			case OCCUPIED:
				occupied++
			case FLOOR:
				if inSight {
					for {
						a, b = a+i, b+j
						if !g.ValidPos(a, b) {
							break
						}
						switch g[a][b] {
						case FLOOR:
							continue
						case EMPTY:
							empty++
						case OCCUPIED:
							occupied++
						}
						break
					}
				}
			}
		}
	}
	return empty, occupied
}

func Part1(input string) (string, error) {
	grid := ParseGrid(input)

	for {
		dupl := grid.Copy()

		var changes int
		for r := 0; r < len(grid); r++ {
			for c := 0; c < len(grid[r]); c++ {
				_, occupied := grid.Neighbours(r, c, false)
				switch grid[r][c] {
				case EMPTY:
					if occupied == 0 {
						dupl[r][c] = OCCUPIED
						changes++
					}
				case OCCUPIED:
					if occupied >= 4 {
						dupl[r][c] = EMPTY
						changes++
					}
				}
			}
		}

		if changes == 0 {
			break
		}
		grid = dupl
	}

	return strconv.Itoa(grid.Count(OCCUPIED)), nil
}

func Part2(input string) (string, error) {
	grid := ParseGrid(input)

	for {
		dupl := grid.Copy()

		var changes int
		for r := 0; r < len(grid); r++ {
			for c := 0; c < len(grid[r]); c++ {
				_, occupied := grid.Neighbours(r, c, true)
				switch grid[r][c] {
				case EMPTY:
					if occupied == 0 {
						dupl[r][c] = OCCUPIED
						changes++
					}
				case OCCUPIED:
					if occupied >= 5 {
						dupl[r][c] = EMPTY
						changes++
					}
				}
			}
		}

		if changes == 0 {
			break
		}
		grid = dupl
	}

	return strconv.Itoa(grid.Count(OCCUPIED)), nil
}
