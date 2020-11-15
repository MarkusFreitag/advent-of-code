package day2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	input := `abcdef
bababc
abbcde
abcccd
aabcdd
abcdee
ababab`
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "12", solution)
}

func TestPart2(t *testing.T) {
	input := `abcde
fghij
klmno
pqrst
fguij
axcye
wvxyz`
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "fgij", solution)
}
