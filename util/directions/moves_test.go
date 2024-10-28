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
	assert.Equal(t, "U", directions.Up.Short())
	assert.Equal(t, "(0,1)", directions.Up.Point().String())

	assert.Equal(t, "RIGHT", directions.Right.String())
	assert.Equal(t, "RIGHT", directions.Right.Word())
	assert.Equal(t, "R", directions.Right.Short())
	assert.Equal(t, "(1,0)", directions.Right.Point().String())

	assert.Equal(t, "DOWN", directions.Down.String())
	assert.Equal(t, "DOWN", directions.Down.Word())
	assert.Equal(t, "D", directions.Down.Short())
	assert.Equal(t, "(0,-1)", directions.Down.Point().String())

	assert.Equal(t, "LEFT", directions.Left.String())
	assert.Equal(t, "LEFT", directions.Left.Word())
	assert.Equal(t, "L", directions.Left.Short())
	assert.Equal(t, "(-1,0)", directions.Left.Point().String())

	assert.Equal(t, "UPRIGHT", directions.UpRight.String())
	assert.Equal(t, "UPRIGHT", directions.UpRight.Word())
	assert.Equal(t, "UR", directions.UpRight.Short())
	assert.Equal(t, "(1,1)", directions.UpRight.Point().String())

	assert.Equal(t, "DOWNRIGHT", directions.DownRight.String())
	assert.Equal(t, "DOWNRIGHT", directions.DownRight.Word())
	assert.Equal(t, "DR", directions.DownRight.Short())
	assert.Equal(t, "(1,-1)", directions.DownRight.Point().String())

	assert.Equal(t, "DOWNLEFT", directions.DownLeft.String())
	assert.Equal(t, "DOWNLEFT", directions.DownLeft.Word())
	assert.Equal(t, "DL", directions.DownLeft.Short())
	assert.Equal(t, "(-1,-1)", directions.DownLeft.Point().String())

	assert.Equal(t, "UPLEFT", directions.UpLeft.String())
	assert.Equal(t, "UPLEFT", directions.UpLeft.Word())
	assert.Equal(t, "UL", directions.UpLeft.Short())
	assert.Equal(t, "(-1,1)", directions.UpLeft.Point().String())

	assert.Equal(
		t,
		[]directions.Move{directions.Up, directions.Right, directions.Down, directions.Left},
		slices.Collect(directions.Moves()),
	)
	assert.Equal(
		t,
		[]directions.Move{directions.Up, directions.Left, directions.Down, directions.Right},
		slices.Collect(directions.Moves(directions.WithCounterClockwise())),
	)
	assert.Equal(
		t,
		[]directions.Move{directions.Right, directions.Down, directions.Left, directions.Up},
		slices.Collect(directions.MovesFrom(directions.Right)),
	)
	assert.Equal(
		t,
		[]directions.Move{directions.Left, directions.Down, directions.Right, directions.Up},
		slices.Collect(directions.MovesFrom(directions.Left, directions.WithCounterClockwise())),
	)

	assert.Equal(
		t,
		[]directions.Move{
			directions.Up,
			directions.UpRight,
			directions.Right,
			directions.DownRight,
			directions.Down,
			directions.DownLeft,
			directions.Left,
			directions.UpLeft,
		},
		slices.Collect(directions.Moves(directions.WithIntermediate())),
	)
	assert.Equal(
		t,
		[]directions.Move{
			directions.Up,
			directions.UpLeft,
			directions.Left,
			directions.DownLeft,
			directions.Down,
			directions.DownRight,
			directions.Right,
			directions.UpRight,
		},
		slices.Collect(directions.Moves(directions.WithIntermediate(), directions.WithCounterClockwise())),
	)
	assert.Equal(
		t,
		[]directions.Move{
			directions.Right,
			directions.DownRight,
			directions.Down,
			directions.DownLeft,
			directions.Left,
			directions.UpLeft,
			directions.Up,
			directions.UpRight,
		},
		slices.Collect(directions.MovesFrom(directions.Right, directions.WithIntermediate())),
	)
	assert.Equal(
		t,
		[]directions.Move{
			directions.Left,
			directions.DownLeft,
			directions.Down,
			directions.DownRight,
			directions.Right,
			directions.UpRight,
			directions.Up,
			directions.UpLeft,
		},
		slices.Collect(directions.MovesFrom(
			directions.Left,
			directions.WithIntermediate(),
			directions.WithCounterClockwise(),
		)),
	)
}
