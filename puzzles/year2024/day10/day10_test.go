package day10

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	testcases := map[string]string{
		`...0...
...1...
...2...
6543456
7.....7
8.....8
9.....9`: "2",
		`..90..9
...1.98
...2..7
6543456
765.987
876....
987....`: "4",
		`10..9..
2...8..
3...7..
4567654
...8..3
...9..2
.....01`: "3",
		`89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`: "36",
	}
	for input, expected := range testcases {
		solution, err := Part1(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}

func TestPart2(t *testing.T) {
	testcases := map[string]string{
		`.....0.
..4321.
..5..2.
..6543.
..7..4.
..8765.
..9....`: "3",
		`..90..9
...1.98
...2..7
6543456
765.987
876....
987....`: "13",
		`012345
123456
234567
345678
4.6789
56789.`: "227",
		`89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`: "81",
	}
	for input, expected := range testcases {
		solution, err := Part2(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}
