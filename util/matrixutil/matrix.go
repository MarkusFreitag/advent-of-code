package matrixutil

import (
	"iter"
	"slices"
	"strings"
)

type Matrix[T any] [][]T

func New[T any](rows, cols int) Matrix[T] {
	return NewWithDefaultValue(rows, cols, *new(T))
}

func NewWithDefaultValue[T any](rows, cols int, val T) Matrix[T] {
	m := make(Matrix[T], rows)
	for r := range rows {
		m[r] = make([]T, cols)
		for c := range cols {
			m[r][c] = val
		}
	}
	return m
}

func FromString(s string) Matrix[rune] {
	lines := strings.Fields(s)
	m := New[rune](len(lines), 0)
	for idx, line := range lines {
		m[idx] = []rune(line)
	}
	return m
}

func TrimEdges[T any](m Matrix[T]) Matrix[T] {
	n := New[T](len(m)-2, 0)
	for idx, row := range m[1 : len(m)-1] {
		n[idx] = row[1 : len(row)-1]
	}
	return n
}

func JoinHorizontal[T any](matrices ...Matrix[T]) Matrix[T] {
	first := Copy(matrices[0])
	for _, matrix := range matrices[1:] {
		for idx := 0; idx < len(first); idx++ {
			first[idx] = append(first[idx], matrix[idx]...)
		}
	}
	return first
}

func JoinVertical[T any](matrices ...Matrix[T]) Matrix[T] {
	first := Copy(matrices[0])
	for _, matrix := range matrices[1:] {
		for _, row := range matrix {
			first = append(first, row)
		}
	}
	return first
}

func FlipVertical[T any](m Matrix[T]) {
	for _, row := range m {
		slices.Reverse(row)
	}
}

func FlipHorizontal[T any](m Matrix[T]) {
	slices.Reverse(m)
}

func Transpose[T any](m Matrix[T]) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < i; j++ {
			m[i][j], m[j][i] = m[j][i], m[i][j]
		}
	}
}

func RotateClockwise[T any](m Matrix[T]) {
	slices.Reverse(m)
	Transpose(m)
}

func RotateCounterClockwise[T any](m Matrix[T]) {
	Transpose(m)
	slices.Reverse(m)
}

func Copy[T any](m Matrix[T]) Matrix[T] {
	c := make(Matrix[T], len(m))
	for i := range m {
		c[i] = make([]T, len(m[i]))
		copy(c[i], m[i])
	}
	return c
}

func Equal[T comparable](s1, s2 iter.Seq[T]) bool {
	next1, stop1 := iter.Pull(s1)
	next2, stop2 := iter.Pull(s2)
	defer stop1()
	defer stop2()

	for {
		val1, ok1 := next1()
		val2, ok2 := next2()

		if !ok1 && !ok2 {
			return true
		}

		if (ok1 && !ok2) || (!ok1 && ok2) {
			return false
		}

		if val1 != val2 {
			return false
		}
	}
}
