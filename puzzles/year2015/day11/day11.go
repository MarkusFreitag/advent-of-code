package day11

import "strings"

func incrStraight(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		num := int(s[i])
		if int(s[i+1]) == num+1 && int(s[i+2]) == num+2 {
			return true
		}
	}
	return false
}

func notIOL(s string) bool {
	return !strings.Contains(s, "i") && !strings.Contains(s, "o") && !strings.Contains(s, "l")
}

func containsPairs(s string) bool {
	var counter int
	for i := 0; i < len(s)-1; i++ {
		if int(s[i]) == int(s[i+1]) {
			counter++
			i++
		}
	}
	return counter > 1
}

func valid(s string) bool {
	return len(s) == 8 && incrStraight(s) && notIOL(s) && containsPairs(s)
}

func incrString(s string) string {
	runes := []rune(s)
	for i := len(runes) - 1; i > 0; i-- {
		if runes[i] == 'z' {
			runes[i] = 'a'
		} else {
			runes[i] = runes[i] + 1
			break
		}
	}
	return string(runes)
}

func Part1(input string) (string, error) {
	for !valid(input) {
		input = incrString(input)
	}
	return input, nil
}

func Part2(input string) (string, error) {
	for !valid(input) || input == "hxbxxyzz" {
		input = incrString(input)
	}
	return input, nil
}
