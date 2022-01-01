package day18

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewTree(t *testing.T) {
	inputs := []string{
		"1",
		"[1,2]",
		"[[1,2],3]",
		"[1,[2,3]]",
		"[[1,2],[3,4]]",

		// examples from explode
		"[[[[[9,8],1],2],3],4]",
		"[7,[6,[5,[4,[3,2]]]]]",
		"[[6,[5,[4,[3,2]]]],1]",
		"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
		"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
	}
	for _, input := range inputs {
		require.Equal(t, input, NewTree(input).String())
	}
}

func TestTreeAdd(t *testing.T) {
	testcases := [][]string{
		{"[1,1]\n[2,2]\n[3,3]\n[4,4]", "[[[[1,1],[2,2]],[3,3]],[4,4]]"},
		{"[1,1]\n[2,2]\n[3,3]\n[4,4]\n[5,5]", "[[[[3,0],[5,3]],[4,4]],[5,5]]"},
		{"[1,1]\n[2,2]\n[3,3]\n[4,4]\n[5,5]\n[6,6]", "[[[[5,0],[7,4]],[5,5]],[6,6]]"},
	}
	for _, tc := range testcases {
		lines := strings.Split(tc[0], "\n")
		tree := NewTree(lines[0])
		for _, line := range lines[1:] {
			tree = tree.Add(NewTree(line))
		}
		require.Equal(t, tc[1], tree.String())
	}
}

func TestTreeExplose(t *testing.T) {
	testcases := [][]string{
		{"[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]"},
		{"[7,[6,[5,[4,[3,2]]]]]", "[7,[6,[5,[7,0]]]]"},
		{"[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]"},
		{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"},
		{"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"},
	}
	for _, tc := range testcases {
		tree := NewTree(tc[0])
		ok := tree.Explode(0)
		require.Equal(t, tc[1], tree.String())
		require.True(t, ok)
	}
}

func TestTreeSplit(t *testing.T) {
	testcases := [][]string{
		{"10", "[5,5]"},
		{"11", "[5,6]"},
	}
	for _, tc := range testcases {
		tree := NewTree(tc[0])
		ok := tree.Split()
		require.Equal(t, tc[1], tree.String())
		require.True(t, ok)
	}
}

var input = `[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`

func TestPart1(t *testing.T) {
	solution, err := Part1(input)
	require.Nil(t, err)
	require.Equal(t, "4140", solution)
}

func TestPart2(t *testing.T) {
	solution, err := Part2(input)
	require.Nil(t, err)
	require.Equal(t, "3993", solution)
}
