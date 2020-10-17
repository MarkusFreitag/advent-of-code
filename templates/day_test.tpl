package {{.Package}}

import (
  "testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	solution, err := Part1("dummy")
	require.Nil(t, err)
	require.Equal(t, "not solved yet", solution)
}

func TestPart2(t *testing.T) {
	solution, err := Part2("dummy")
	require.Nil(t, err)
	require.Equal(t, "not solved yet", solution)
}
