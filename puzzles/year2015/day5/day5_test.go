package day5

import (
  "testing"

	"github.com/stretchr/testify/require"
)

func TestThreeVowels(t *testing.T) {
  require.True(t, threeVowels("aei"))
  require.True(t, threeVowels("xazegov"))
  require.True(t, threeVowels("aeiouaeiouaeiou"))
  require.False(t, threeVowels("abcdefg"))
  require.False(t, threeVowels("xzyzmop"))
}

func TestDoubleLetters(t *testing.T) {
  require.True(t, doubleLetters("xx"))
  require.True(t, doubleLetters("abcdde"))
  require.True(t, doubleLetters("aabbccdd"))
  require.False(t, doubleLetters("abcdef"))
}

func TestContainsForbidden(t *testing.T) {
  require.True(t, containsForbidden("abcdef"))
  require.False(t, containsForbidden("efkg"))
}

func TestContainsPair(t *testing.T) {
  require.True(t, containsPair("xyxy"))
  require.True(t, containsPair("aabcdefgaa"))
  require.False(t, containsPair("aaa"))
}

func TestContainsRepeatPattern(t *testing.T) {
  require.True(t, containsRepeatPattern("xyx"))
  require.True(t, containsRepeatPattern("abcdefeghi"))
  require.True(t, containsRepeatPattern("aaa"))
  require.False(t, containsRepeatPattern("abcddef"))
}

func TestPart1(t *testing.T) {
  testcases := map[string]string{
    "ugknbfddgicrmopn": "1",
    "aaa":              "1",
    "jchzalrnumimnmhp": "0",
    "haegwjzuvuyypxyu": "0",
    "dvszwmarrgswjxmb": "0",
  }
  for input, expected := range testcases {
    solution, err := Part1(input)
    require.Nil(t, err)
    require.Equal(t, expected, solution)
  }
}

func TestPart2(t *testing.T) {
  testcases := map[string]string{
    "qjhvhtzxzqqjkmpb": "1",
    "xxyxx":            "1",
    "uurcxstgmygtbstg": "0",
    "ieodomkazucvgmuy": "0",
  }
  for input, expected := range testcases {
    solution, err := Part2(input)
    require.Nil(t, err)
    require.Equal(t, expected, solution)
  }
}
