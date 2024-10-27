package directions_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/MarkusFreitag/advent-of-code/util/directions"
)

func TestMove(t *testing.T) {
	assert.Equal(t, "UP", directions.Up.String())
	assert.Equal(t, "UP", directions.Up.Word())
	assert.Equal(t, "^", directions.Up.Short())

	assert.Equal(t, "RIGHT", directions.Right.String())
	assert.Equal(t, "RIGHT", directions.Right.Word())
	assert.Equal(t, ">", directions.Right.Short())

	assert.Equal(t, "DOWN", directions.Down.String())
	assert.Equal(t, "DOWN", directions.Down.Word())
	assert.Equal(t, "v", directions.Down.Short())

	assert.Equal(t, "LEFT", directions.Left.String())
	assert.Equal(t, "LEFT", directions.Left.Word())
	assert.Equal(t, "<", directions.Left.Short())

	assert.Equal(
		t,
		[]directions.Move{directions.Up, directions.Right, directions.Down, directions.Left},
		slices.Collect(directions.Moves()),
	)
	assert.Equal(
		t,
		[]directions.Move{directions.Up, directions.Left, directions.Down, directions.Right},
		slices.Collect(directions.MovesCounterClockwise()),
	)
	assert.Equal(
		t,
		[]directions.Move{directions.Right, directions.Down, directions.Left, directions.Up},
		slices.Collect(directions.MovesFrom(directions.Right)),
	)
	assert.Equal(
		t,
		[]directions.Move{directions.Left, directions.Down, directions.Right, directions.Up},
		slices.Collect(directions.MovesFromCounterClockwise(directions.Left)),
	)
}
