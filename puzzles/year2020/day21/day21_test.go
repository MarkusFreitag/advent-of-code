package day21

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "5", solution)
}

func TestPart2(t *testing.T) {
	input := `mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)`
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "mxmxvkd,sqjhc,fvjkl", solution)
}
