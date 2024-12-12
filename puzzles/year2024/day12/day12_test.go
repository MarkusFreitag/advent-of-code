package day12

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	testcases := map[string]string{
		`AAAA
BBCD
BBCC
EEEC`: "140",
		`OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`: "772",
		`RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`: "1930",
	}
	for input, expected := range testcases {
		solution, err := Part1(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}

func TestPart2(t *testing.T) {
	testcases := map[string]string{
		`AAAA
BBCD
BBCC
EEEC`: "80",
		`OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`: "436",
		`EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`: "236",
		`AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`: "368",
		`RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`: "1206",
		`AAAAAAAA
AACBBDDA
AACBBAAA
ABBAAAAA
ABBADDDA
AAAADADA
AAAAAAAA`: "946",
	}
	for input, expected := range testcases {
		solution, err := Part2(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}
