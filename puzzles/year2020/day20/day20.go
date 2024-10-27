package day20

import (
	"iter"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util/directions"
	"github.com/MarkusFreitag/advent-of-code/util/matrixutil"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
	"github.com/MarkusFreitag/advent-of-code/util/setutil"
	"github.com/MarkusFreitag/advent-of-code/util/sliceutil"
)

var rgxHead = regexp.MustCompile(`^Tile\s(\d+):`)

type Piece struct {
	ID int
	M  matrixutil.Matrix[rune]
}

func solve(pieces []Piece, board matrixutil.Matrix[Piece], row, col int, visited setutil.Set[int]) bool {
	boardLength := len(board)
	if fullBoard(board) && validateBoard(board) {
		return true
	}
	for _, piece := range pieces {
		if _, ok := visited[piece.ID]; !ok {
			if row > 0 && !matrixutil.Equal(
				Border(board[row-1][col].M, directions.Down),
				Border(piece.M, directions.Up),
			) {
				continue
			}
			if col > 0 && !matrixutil.Equal(
				Border(board[row][col-1].M, directions.Right),
				Border(piece.M, directions.Left),
			) {
				continue
			}

			board[row][col] = piece
			visited[piece.ID] = struct{}{}
			if col == boardLength-1 {
				if solve(pieces, board, row+1, 0, visited) {
					return true
				}
			} else {
				if solve(pieces, board, row, col+1, visited) {
					return true
				}
			}
			delete(visited, piece.ID)
		}
	}
	return false
}

func validateBoard(board matrixutil.Matrix[Piece]) bool {
	rowLength := len(board)
	colLength := len(board[0])
	for r := 0; r < rowLength; r++ {
		if r+1 == rowLength {
			break
		}
		for c := 0; c < colLength; c++ {
			if c+1 == colLength {
				break
			}

			if !matrixutil.Equal(
				Border(board[r][c].M, directions.Down),
				Border(board[r+1][c].M, directions.Up),
			) {
				return false
			}
			if !matrixutil.Equal(
				Border(board[r][c].M, directions.Right),
				Border(board[r][c+1].M, directions.Left),
			) {
				return false
			}
		}
	}
	return true
}

func fullBoard(board matrixutil.Matrix[Piece]) bool {
	for _, row := range board {
		for _, piece := range row {
			if piece.ID == 0 {
				return false
			}
		}
	}
	return true
}

func monsterHunt(compacted matrixutil.Matrix[rune]) bool {
	monsterSeq := [][2]int{
		{0, 0}, {1, 1},
		{0, 3}, {-1, 1}, {0, 1}, {1, 1},
		{0, 3}, {-1, 1}, {0, 1}, {1, 1},
		{0, 3}, {-1, 1}, {0, 1}, {-1, 0}, {1, 1},
	}

	rowLen := len(compacted)
	colLen := len(compacted[0])
	var correctOrientation bool
	for r := 0; r < rowLen; r++ {
		for c := 0; c < colLen; c++ {
			current := [2]int{r, c}
			found := true
			for _, move := range monsterSeq {
				next := [2]int{current[0] + move[0], current[1] + move[1]}
				if next[0] < 0 || next[0] >= rowLen || next[1] < 0 || next[1] >= colLen {
					found = false
					break
				}
				if compacted[next[0]][next[1]] == '.' {
					found = false
					break
				}
				current = next
			}
			if found {
				correctOrientation = true
				current := [2]int{r, c}
				for _, move := range monsterSeq {
					next := [2]int{current[0] + move[0], current[1] + move[1]}
					compacted[next[0]][next[1]] = 'O'
					current = next
				}
			}
		}
	}

	return correctOrientation
}

func Border[T any](m matrixutil.Matrix[T], d directions.Move) iter.Seq[T] {
	return func(yield func(T) bool) {
		switch d {
		case directions.Up, directions.Down:
			var row int
			if d == directions.Down {
				row = len(m) - 1
			}
			for i := 0; i < len(m[row]); i++ {
				if !yield(m[row][i]) {
					return
				}
			}
		case directions.Right, directions.Left:
			var col int
			if d == directions.Right {
				col = len(m) - 1
			}
			for i := 0; i < len(m); i++ {
				if !yield(m[i][col]) {
					return
				}
			}
		}
	}
}

func Part1(input string) (string, error) {
	matrices := make(map[int]matrixutil.Matrix[rune])
	for _, block := range strings.Split(input, "\n\n") {
		head := rgxHead.FindAllStringSubmatch(block, -1)[0]
		id, _ := strconv.Atoi(head[1])
		matrices[id] = matrixutil.FromString(strings.Split(block, ":")[1])
	}

	allSides := make([]string, 0)
	for _, matrix := range matrices {
		for dir := range directions.Moves() {
			side := slices.Collect(Border(matrix, dir))
			allSides = append(allSides, string(side))
			slices.Reverse(side)
			allSides = append(allSides, string(side))
		}
	}

	edges := make([]int, 0)
	for id, matrix := range matrices {
		var count int
		for dir := range directions.Moves() {
			if sliceutil.Count(allSides, string(slices.Collect(Border(matrix, dir)))) == 1 {
				count++
			}
		}
		if count == 2 {
			edges = append(edges, id)
		}
	}

	return strconv.Itoa(numbers.Multiply(edges...)), nil
}

func Part2(input string) (string, error) {
	pieces := make([]Piece, 0)
	for _, block := range strings.Split(input, "\n\n") {
		head := rgxHead.FindAllStringSubmatch(block, -1)[0]
		id, _ := strconv.Atoi(head[1])
		matrix := matrixutil.FromString(strings.Split(block, ":")[1])
		for range directions.Moves() {
			pieces = append(pieces, Piece{ID: id, M: matrixutil.Copy(matrix)})
			matrixutil.RotateClockwise(matrix)
		}
		matrixutil.FlipHorizontal(matrix)
		for range directions.Moves() {
			pieces = append(pieces, Piece{ID: id, M: matrixutil.Copy(matrix)})
			matrixutil.RotateClockwise(matrix)
		}
	}

	boardSize := int(math.Sqrt(float64(len(pieces) / 8)))
	board := matrixutil.New[Piece](boardSize, boardSize)

	solve(pieces, board, 0, 0, make(setutil.Set[int]))

	borderless := matrixutil.New[matrixutil.Matrix[rune]](len(board), len(board[0]))
	for rIdx, row := range board {
		for cIdx, piece := range row {
			borderless[rIdx][cIdx] = matrixutil.TrimEdges(piece.M)
		}
	}

	rows := make([]matrixutil.Matrix[rune], 0)
	for _, row := range borderless {
		rows = append(rows, matrixutil.JoinHorizontal(row...))
	}
	compacted := matrixutil.JoinVertical(rows...)

	do := func() {
		for range directions.Moves() {
			if monsterHunt(compacted) {
				return
			}
			matrixutil.RotateClockwise(compacted)
		}
		matrixutil.FlipHorizontal(compacted)
		for range directions.Moves() {
			if monsterHunt(compacted) {
				return
			}
			matrixutil.RotateClockwise(compacted)
		}
	}
	do()
	var total int
	for _, row := range compacted {
		total += sliceutil.Count(row, '#')
	}
	return strconv.Itoa(total), nil
}
