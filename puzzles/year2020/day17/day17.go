package day17

import (
	"strconv"
	"strings"
)

type Cube struct {
	Z, Y, X int
}

type Grid map[Cube]bool

func (g Grid) AddNeighbours() Grid {
	grid := make(Grid)
	for cube, state := range g {
		grid[cube] = state
	}
	for cube := range g {
		for _, c := range neighbourCubes(cube) {
			if _, ok := g[c]; !ok {
				g[c] = false
			}
		}
	}
	return grid
}

func (g Grid) CheckPos(z, y, x int) bool {
	if v, ok := g[Cube{Z: z, Y: y, X: x}]; ok {
		return v
	}
	return false
}

func (g Grid) CountActive() int {
	var count int
	for _, state := range g {
		if state {
			count++
		}
	}
	return count
}

func neighbourCubes(cube Cube) []Cube {
	neighbours := make([]Cube, 0)
	for z := -1; z <= 1; z++ {
		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {
				if z == 0 && y == 0 && x == 0 {
					continue
				}
				neighbours = append(neighbours, Cube{Z: cube.Z + z, Y: cube.Y + y, X: cube.X + x})
			}
		}
	}
	return neighbours
}

func Part1(input string) (string, error) {
	grid := make(Grid)
	for y, row := range strings.Split(input, "\n") {
		for x, col := range row {
			if rune(col) == '#' {
				grid[Cube{Y: y, X: x}] = true
			}
		}
	}

	for i := 1; i <= 6; i++ {
		dupl := grid.AddNeighbours()

		for cube, state := range grid {
			var active int
			for _, c := range neighbourCubes(cube) {
				if v, ok := grid[c]; ok {
					if v {
						active++
					}
				}
			}
			if state {
				dupl[cube] = (active == 2 || active == 3)
			} else {
				dupl[cube] = (active == 3)
			}
		}

		grid = dupl
	}
	return strconv.Itoa(grid.CountActive()), nil
}

type HyperCube struct {
	Z, Y, X, W int
}

type HyperGrid map[HyperCube]bool

func (g HyperGrid) AddNeighbours() HyperGrid {
	grid := make(HyperGrid)
	for cube, state := range g {
		grid[cube] = state
	}
	for cube := range g {
		for _, c := range neighbourHyperCubes(cube) {
			if _, ok := g[c]; !ok {
				g[c] = false
			}
		}
	}
	return grid
}

func (g HyperGrid) CheckPos(z, y, x, w int) bool {
	if v, ok := g[HyperCube{Z: z, Y: y, X: x, W: w}]; ok {
		return v
	}
	return false
}

func (g HyperGrid) CountActive() int {
	var count int
	for _, state := range g {
		if state {
			count++
		}
	}
	return count
}

func neighbourHyperCubes(hc HyperCube) []HyperCube {
	neighbours := make([]HyperCube, 0)
	for z := -1; z <= 1; z++ {
		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {
				for w := -1; w <= 1; w++ {
					if z == 0 && y == 0 && x == 0 && w == 0 {
						continue
					}
					neighbours = append(neighbours, HyperCube{Z: hc.Z + z, Y: hc.Y + y, X: hc.X + x, W: hc.W + w})
				}
			}
		}
	}
	return neighbours
}

func Part2(input string) (string, error) {
	grid := make(HyperGrid)
	for y, row := range strings.Split(input, "\n") {
		for x, col := range row {
			if rune(col) == '#' {
				grid[HyperCube{Y: y, X: x}] = true
			}
		}
	}

	for i := 1; i <= 6; i++ {
		dupl := grid.AddNeighbours()

		for cube, state := range grid {
			var active int
			for _, c := range neighbourHyperCubes(cube) {
				if v, ok := grid[c]; ok {
					if v {
						active++
					}
				}
			}
			if state {
				dupl[cube] = (active == 2 || active == 3)
			} else {
				dupl[cube] = (active == 3)
			}
		}

		grid = dupl
	}
	return strconv.Itoa(grid.CountActive()), nil
}
