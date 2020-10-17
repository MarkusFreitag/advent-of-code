package day5

import (
  "strconv"
  "strings"
)

const (
  alphabet = "abcdefghijklmnopqrstuvwxyz"
  vowels   = "aeiou"
)

func threeVowels(str string) bool {
  var counter int
  for _, r := range str {
    if strings.ContainsRune(vowels, r) {
      counter++
    }
  }
  return counter >= 3
}

func doubleLetters(str string) bool {
  for _, l := range alphabet {
    if strings.Contains(str, string(l)+string(l)) {
      return true
    }
  }
  return false
}

func containsForbidden(str string) bool {
  for _, s := range []string{"ab", "cd", "pq", "xy"} {
    if strings.Contains(str, s) {
      return true
    }
  }
  return false
}

func containsPair(str string) bool {
  for i:=0;i+2<len(str);i++ {
    if strings.Count(str, str[i:i+2]) > 1 {
      return true
    }
  }
  return false
}

func containsRepeatPattern(str string) bool {
  for i:=0;i+3<=len(str);i++ {
    sub := str[i:i+3]
    if sub[0] == sub[2] {
      return true
    }
  }
  return false
}

func Part1(input string) (string, error) {
  var total int
  for _, str := range strings.Split(input, "\n") {
    if !threeVowels(str) || !doubleLetters(str) || containsForbidden(str) {
      continue
    }
    total++
  }
  return strconv.Itoa(total), nil
}

func Part2(input string) (string, error) {
  var total int
  for _, str := range strings.Split(input, "\n") {
    if !containsPair(str) || !containsRepeatPattern(str) {
      continue
    }
    total++
  }
  return strconv.Itoa(total), nil
}
