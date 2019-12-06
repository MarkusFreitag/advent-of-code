package puzzles

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	p, err := Get(1, 1)
	require.Nil(t, p)
	require.Equal(t, "could not find puzzle for year 1 day 1", err.Error())

	p, err = Get(2019, 1)
	require.Nil(t, err)
	require.Equal(t, 2, len(p))
}
