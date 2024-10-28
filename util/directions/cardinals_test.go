package directions_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/MarkusFreitag/advent-of-code/util/directions"
)

func TestCardinal(t *testing.T) {
	assert.Equal(t, "NORTH", directions.North.String())
	assert.Equal(t, "NORTH", directions.North.Word())
	assert.Equal(t, "N", directions.North.Short())
	assert.Equal(t, "(0,1)", directions.North.Point().String())

	assert.Equal(t, "EAST", directions.East.String())
	assert.Equal(t, "EAST", directions.East.Word())
	assert.Equal(t, "E", directions.East.Short())
	assert.Equal(t, "(1,0)", directions.East.Point().String())

	assert.Equal(t, "SOUTH", directions.South.String())
	assert.Equal(t, "SOUTH", directions.South.Word())
	assert.Equal(t, "S", directions.South.Short())
	assert.Equal(t, "(0,-1)", directions.South.Point().String())

	assert.Equal(t, "WEST", directions.West.String())
	assert.Equal(t, "WEST", directions.West.Word())
	assert.Equal(t, "W", directions.West.Short())
	assert.Equal(t, "(-1,0)", directions.West.Point().String())

	assert.Equal(t, "NORTHEAST", directions.NorthEast.String())
	assert.Equal(t, "NORTHEAST", directions.NorthEast.Word())
	assert.Equal(t, "NE", directions.NorthEast.Short())
	assert.Equal(t, "(1,1)", directions.NorthEast.Point().String())

	assert.Equal(t, "SOUTHEAST", directions.SouthEast.String())
	assert.Equal(t, "SOUTHEAST", directions.SouthEast.Word())
	assert.Equal(t, "SE", directions.SouthEast.Short())
	assert.Equal(t, "(1,-1)", directions.SouthEast.Point().String())

	assert.Equal(t, "SOUTHWEST", directions.SouthWest.String())
	assert.Equal(t, "SOUTHWEST", directions.SouthWest.Word())
	assert.Equal(t, "SW", directions.SouthWest.Short())
	assert.Equal(t, "(-1,-1)", directions.SouthWest.Point().String())

	assert.Equal(t, "NORTHWEST", directions.NorthWest.String())
	assert.Equal(t, "NORTHWEST", directions.NorthWest.Word())
	assert.Equal(t, "NW", directions.NorthWest.Short())
	assert.Equal(t, "(-1,1)", directions.NorthWest.Point().String())

	assert.Equal(
		t,
		[]directions.Cardinal{directions.North, directions.East, directions.South, directions.West},
		slices.Collect(directions.Cardinals()),
	)
	assert.Equal(
		t,
		[]directions.Cardinal{directions.North, directions.West, directions.South, directions.East},
		slices.Collect(directions.Cardinals(directions.WithCounterClockwise())),
	)
	assert.Equal(
		t,
		[]directions.Cardinal{directions.East, directions.South, directions.West, directions.North},
		slices.Collect(directions.CardinalsFrom(directions.East)),
	)
	assert.Equal(
		t,
		[]directions.Cardinal{directions.West, directions.South, directions.East, directions.North},
		slices.Collect(directions.CardinalsFrom(directions.West, directions.WithCounterClockwise())),
	)

	assert.Equal(
		t,
		[]directions.Cardinal{
			directions.North,
			directions.NorthEast,
			directions.East,
			directions.SouthEast,
			directions.South,
			directions.SouthWest,
			directions.West,
			directions.NorthWest,
		},
		slices.Collect(directions.Cardinals(directions.WithIntermediate())),
	)
	assert.Equal(
		t,
		[]directions.Cardinal{
			directions.North,
			directions.NorthWest,
			directions.West,
			directions.SouthWest,
			directions.South,
			directions.SouthEast,
			directions.East,
			directions.NorthEast,
		},
		slices.Collect(directions.Cardinals(directions.WithIntermediate(), directions.WithCounterClockwise())),
	)
	assert.Equal(
		t,
		[]directions.Cardinal{
			directions.East,
			directions.SouthEast,
			directions.South,
			directions.SouthWest,
			directions.West,
			directions.NorthWest,
			directions.North,
			directions.NorthEast,
		},
		slices.Collect(directions.CardinalsFrom(directions.East, directions.WithIntermediate())),
	)
	assert.Equal(
		t,
		[]directions.Cardinal{
			directions.West,
			directions.SouthWest,
			directions.South,
			directions.SouthEast,
			directions.East,
			directions.NorthEast,
			directions.North,
			directions.NorthWest,
		},
		slices.Collect(directions.CardinalsFrom(
			directions.West,
			directions.WithIntermediate(),
			directions.WithCounterClockwise(),
		)),
	)
}
