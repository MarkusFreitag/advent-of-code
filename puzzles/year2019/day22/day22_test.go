package day22

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewDeck(t *testing.T) {
	d := NewDeck(10)
	require.Equal(t, 10, len(d))
	require.Equal(t, "[0 1 2 3 4 5 6 7 8 9]", fmt.Sprintf("%v", d))
}

func TestDealIntoNewStack(t *testing.T) {
	d := NewDeck(10)
	require.Equal(t, 10, len(d))
	require.Equal(t, "[0 1 2 3 4 5 6 7 8 9]", fmt.Sprintf("%v", d))
	d = d.DealIntoNewStack()
	require.Equal(t, 10, len(d))
	require.Equal(t, "[9 8 7 6 5 4 3 2 1 0]", fmt.Sprintf("%v", d))
}

func TestCutN(t *testing.T) {
	d := NewDeck(10)
	require.Equal(t, 10, len(d))
	require.Equal(t, "[0 1 2 3 4 5 6 7 8 9]", fmt.Sprintf("%v", d))

	d = d.CutN(3)
	require.Equal(t, 10, len(d))
	require.Equal(t, "[3 4 5 6 7 8 9 0 1 2]", fmt.Sprintf("%v", d))

	d = NewDeck(10)
	require.Equal(t, 10, len(d))
	require.Equal(t, "[0 1 2 3 4 5 6 7 8 9]", fmt.Sprintf("%v", d))

	d = d.CutN(-4)
	require.Equal(t, 10, len(d))
	require.Equal(t, "[6 7 8 9 0 1 2 3 4 5]", fmt.Sprintf("%v", d))
}

func TestIncrement(t *testing.T) {
	d := NewDeck(10)
	require.Equal(t, 10, len(d))
	require.Equal(t, "[0 1 2 3 4 5 6 7 8 9]", fmt.Sprintf("%v", d))
	d = d.Increment(3)
	require.Equal(t, 10, len(d))
	require.Equal(t, "[0 7 4 1 8 5 2 9 6 3]", fmt.Sprintf("%v", d))
}
