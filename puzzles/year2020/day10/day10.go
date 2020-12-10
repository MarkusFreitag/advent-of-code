package day10

import (
  "strconv"
  "strings"
  "sort"

  "github.com/MarkusFreitag/advent-of-code/util"
)


func Part1(input string) (string, error) {
  adapters := util.StrsToInts(strings.Split(input, "\n"))
  sort.Ints(adapters)

  var one, three int
  var current int
  builtin := adapters[len(adapters)-1] +3

  for _, adap := range adapters {
    switch adap-current {
    case 1:
      one++
    case 3:
      three++
    default:
      continue
    }
    current = adap
  }

  if builtin-current != 3 {
    return "invalid chain", nil
  } else {
    three++
  }

  return strconv.Itoa(one*three), nil
}

func Part2(input string) (string, error) {
  return "n/a", nil
}
