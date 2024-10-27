package matrixutil_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/MarkusFreitag/advent-of-code/util/matrixutil"
)

func TestFromString(t *testing.T) {
	m := matrixutil.FromString(`123
456`)
	assert.Equal(t, 2, len(m))
	assert.Equal(t, 3, len(m[0]))
}

func TestTrimEdges(t *testing.T) {
	m := matrixutil.FromString(`123
456
789`)
	trimmed := matrixutil.TrimEdges(m)
	assert.Equal(t, 1, len(trimmed))
	assert.Equal(t, 1, len(trimmed[0]))
	assert.Equal(t, '5', trimmed[0][0])
}

func TestJoinHorizontal(t *testing.T) {
	s := `123
456
789`
	a := matrixutil.FromString(s)
	b := matrixutil.FromString(s)
	c := matrixutil.FromString(s)
	joined := matrixutil.JoinHorizontal(a, b, c)
	assert.Equal(t, 3, len(joined))
	assert.Equal(t, 9, len(joined[0]))
	rows := []string{
		"123123123",
		"456456456",
		"789789789",
	}
	for idx, row := range rows {
		assert.Equal(t, row, string(joined[idx]))
	}
}

func TestJoinVertical(t *testing.T) {
	s := `123
456
789`
	a := matrixutil.FromString(s)
	b := matrixutil.FromString(s)
	c := matrixutil.FromString(s)
	joined := matrixutil.JoinVertical(a, b, c)
	assert.Equal(t, 9, len(joined))
	assert.Equal(t, 3, len(joined[0]))
	rows := []string{
		"123",
		"456",
		"789",
		"123",
		"456",
		"789",
		"123",
		"456",
		"789",
	}
	for idx, row := range rows {
		assert.Equal(t, row, string(joined[idx]))
	}
}

func TestFlipVertical(t *testing.T) {
	a := matrixutil.FromString(`123
456
789`)
	flipped := matrixutil.FromString(`321
654
987`)
	matrixutil.FlipVertical(a)
	assert.Equal(t, flipped, a)
}

func TestFlipHorizontal(t *testing.T) {
	a := matrixutil.FromString(`123
456
789`)
	flipped := matrixutil.FromString(`789
456
123`)
	matrixutil.FlipHorizontal(a)
	assert.Equal(t, flipped, a)
}

func TestTranspose(t *testing.T) {
	a := matrixutil.FromString(`123
456
789`)
	transposed := matrixutil.FromString(`147
258
369`)
	matrixutil.Transpose(a)
	assert.Equal(t, transposed, a)
}

func TestRotateClockwise(t *testing.T) {
	a := matrixutil.FromString(`123
456
789`)
	rotated := matrixutil.FromString(`741
852
963`)
	matrixutil.RotateClockwise(a)
	assert.Equal(t, rotated, a)
}

func TestRotateCounterClockwise(t *testing.T) {
	a := matrixutil.FromString(`123
456
789`)
	rotated := matrixutil.FromString(`369
258
147`)
	matrixutil.RotateCounterClockwise(a)
	assert.Equal(t, rotated, a)
}

func TestCopy(t *testing.T) {
	a := matrixutil.FromString(`123
456
789`)
	b := matrixutil.Copy(a)
	a[1][1] = '0'
	assert.NotEqual(t, a, b)
}
