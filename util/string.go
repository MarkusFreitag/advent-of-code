package util

import (
	"sort"
	"strings"
)

func StringReverse(str string) string {
	var s string
	for _, r := range str {
		s = string(r) + s
	}
	return s
}

func StringPadLeft(str, padStr string, length int) string {
	padCount := 1 + ((length - len(padStr)) / len(padStr))
	s := strings.Repeat(padStr, padCount) + str
	return s[(len(s) - length):]
}

func StringPadRight(str, padStr string, length int) string {
	padCount := 1 + ((length - len(padStr)) / len(padStr))
	s := str + strings.Repeat(padStr, padCount)
	return s[:length]
}

func StringToStrings(str string) []string {
	slice := make([]string, len(str))
	for idx, char := range str {
		slice[idx] = string(char)
	}
	return slice
}

func StringSorter(str string) string {
	slice := StringToStrings(str)
	sort.Strings(slice)
	return strings.Join(slice, "")
}

func StringContainsAny(s string, strs ...string) bool {
	for _, str := range strs {
		if strings.Contains(s, str) {
			return true
		}
	}
	return false
}