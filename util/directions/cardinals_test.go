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

	assert.Equal(t, "EAST", directions.East.String())
	assert.Equal(t, "EAST", directions.East.Word())
	assert.Equal(t, "E", directions.East.Short())

	assert.Equal(t, "SOUTH", directions.South.String())
	assert.Equal(t, "SOUTH", directions.South.Word())
	assert.Equal(t, "S", directions.South.Short())

	assert.Equal(t, "WEST", directions.West.String())
	assert.Equal(t, "WEST", directions.West.Word())
	assert.Equal(t, "W", directions.West.Short())

	assert.Equal(
		t,
		[]directions.Cardinal{directions.North, directions.East, directions.South, directions.West},
		slices.Collect(directions.Cardinals()),
	)
	assert.Equal(
		t,
		[]directions.Cardinal{directions.North, directions.West, directions.South, directions.East},
		slices.Collect(directions.CardinalsCounterClockwise()),
	)
	assert.Equal(
		t,
		[]directions.Cardinal{directions.East, directions.South, directions.West, directions.North},
		slices.Collect(directions.CardinalsFrom(directions.East)),
	)
	assert.Equal(
		t,
		[]directions.Cardinal{directions.West, directions.South, directions.East, directions.North},
		slices.Collect(directions.CardinalsFromCounterClockwise(directions.West)),
	)
}
