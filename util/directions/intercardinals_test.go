package directions_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/MarkusFreitag/advent-of-code/util/directions"
)

func TestIntercardinal(t *testing.T) {
	assert.Equal(t, "NORTHEAST", directions.NorthEast.String())
	assert.Equal(t, "NORTHEAST", directions.NorthEast.Word())
	assert.Equal(t, "NE", directions.NorthEast.Short())

	assert.Equal(t, "SOUTHEAST", directions.SouthEast.String())
	assert.Equal(t, "SOUTHEAST", directions.SouthEast.Word())
	assert.Equal(t, "SE", directions.SouthEast.Short())

	assert.Equal(t, "SOUTHWEST", directions.SouthWest.String())
	assert.Equal(t, "SOUTHWEST", directions.SouthWest.Word())
	assert.Equal(t, "SW", directions.SouthWest.Short())

	assert.Equal(t, "NORTHWEST", directions.NorthWest.String())
	assert.Equal(t, "NORTHWEST", directions.NorthWest.Word())
	assert.Equal(t, "NW", directions.NorthWest.Short())

	assert.Equal(
		t,
		[]directions.Intercardinal{directions.NorthEast, directions.SouthEast, directions.SouthWest, directions.NorthWest},
		slices.Collect(directions.Intercardinals()),
	)
	assert.Equal(
		t,
		[]directions.Intercardinal{directions.NorthEast, directions.NorthWest, directions.SouthWest, directions.SouthEast},
		slices.Collect(directions.IntercardinalsCounterClockwise()),
	)
	assert.Equal(
		t,
		[]directions.Intercardinal{directions.NorthEast, directions.SouthEast, directions.SouthWest, directions.NorthWest},
		slices.Collect(directions.IntercardinalsFrom(directions.NorthEast)),
	)
	assert.Equal(
		t,
		[]directions.Intercardinal{directions.NorthWest, directions.SouthWest, directions.SouthEast, directions.NorthEast},
		slices.Collect(directions.IntercardinalsFromCounterClockwise(directions.NorthWest)),
	)
}
