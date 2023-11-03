package day9

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCheckIfEscaped(t *testing.T) {
	assert.False(t, checkIfEscaped("{{<!>},{<!>},{<!>},{<a>}}", 0))
	assert.False(t, checkIfEscaped("{{<!>},{<!>},{<!>},{<a>}}", 3))
	assert.True(t, checkIfEscaped("{{<!>},{<!>},{<!>},{<a>}}", 4))
	assert.False(t, checkIfEscaped("{{<!>},{<!>},{<!>},{<a>}}", 5))
	assert.False(t, checkIfEscaped("{{<!>},{<!>},{<!>},{<a>}}", 9))
	assert.True(t, checkIfEscaped("{{<!>},{<!>},{<!>},{<a>}}", 10))
	assert.False(t, checkIfEscaped("{{<!>},{<!>},{<!>},{<a>}}", 11))
	assert.False(t, checkIfEscaped("{{<!>},{<!>},{<!>},{<a>}}", 15))
	assert.True(t, checkIfEscaped("{{<!>},{<!>},{<!>},{<a>}}", 16))
	assert.False(t, checkIfEscaped("{{<!>},{<!>},{<!>},{<a>}}", 17))

	assert.False(t, checkIfEscaped("{{<!!>},{<!!>},{<!!>},{<!!>}}", 3))
	assert.True(t, checkIfEscaped("{{<!!>},{<!!>},{<!!>},{<!!>}}", 4))
	assert.False(t, checkIfEscaped("{{<!!>},{<!!>},{<!!>},{<!!>}}", 5))
}

func TestCleanupGarbage(t *testing.T) {
	cases := map[string]string{
		"":                  "",
		"random characters": "random characters",
		"<<<":               "<<<",
		"{!>}":              "{}",
		"!!":                "",
		"!!!>":              "",
		`{o"i!a,<{i<a`:      `{o"i,<{i<a`,
	}
	for input, expected := range cases {
		require.Equal(t, expected, cleanupGarbage(input))
	}
}

func TestEval(t *testing.T) {
	garbageCases := map[string]int{
		"<>":                  0,
		"<random characters>": 17,
		"<<<<>":               3,
		"<{!>}>":              2,
		"<!!>":                0,
		"<!!!>>":              0,
		`<{o"i!a,<{i<a>`:      10,
	}
	for input, expected := range garbageCases {
		_, count := eval(input)
		require.Equal(t, expected, count)
	}

	groupCases := map[string]int{
		"{}":                        1,
		"{{{}}}":                    3,
		"{{},{}}":                   3,
		"{{{},{},{{}}}}":            6,
		"{<{},{},{{}}>}":            1,
		"{<a>,<a>,<a>,<a>}":         1,
		"{{<a>},{<a>},{<a>},{<a>}}": 5,
		"{{<!>},{<!>},{<!>},{<a>}}": 2,
	}
	for input, expected := range groupCases {
		lvls, _ := eval(input)
		require.Len(t, lvls, expected)
	}
}

func TestPart1(t *testing.T) {
	testcases := map[string]string{
		"{}":                            "1",
		"{{{}}}":                        "6",
		"{{},{}}":                       "5",
		"{{{},{},{{}}}}":                "16",
		"{<a>,<a>,<a>,<a>}":             "1",
		"{{<ab>},{<ab>},{<ab>},{<ab>}}": "9",
		"{{<!!>},{<!!>},{<!!>},{<!!>}}": "9",
		"{{<a!>},{<a!>},{<a!>},{<ab>}}": "3",
	}
	for input, expected := range testcases {
		solution, err := Part1(input)
		require.Nil(t, err)
		require.Equal(t, expected, solution)
	}
}
