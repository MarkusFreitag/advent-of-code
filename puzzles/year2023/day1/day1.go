package day1

import (
	"sort"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	Numbers "github.com/MarkusFreitag/advent-of-code/util/numbers"
)

func Part1(input string) (string, error) {
	nums := make([]int, 0)
	for _, line := range strings.Fields(input) {
		foundNumbers := searchNumbers(line, false)
		nums = append(nums, util.ParseInt(foundNumbers.String()))
	}
	return strconv.Itoa(Numbers.Sum(nums...)), nil
}

func Part2(input string) (string, error) {
	nums := make([]int, 0)
	for _, line := range strings.Fields(input) {
		foundNumbers := searchNumbers(line, true)
		nums = append(nums, util.ParseInt(foundNumbers.String()))
	}
	return strconv.Itoa(Numbers.Sum(nums...)), nil
}

type number struct {
	str string
	pos int
}

func (n number) String() string { return n.str }

type numbers []number

func (n numbers) String() string {
	if len(n) >= 1 {
		sort.Sort(n)
	}
	return n[0].String() + n[len(n)-1].String()
}
func (n numbers) Len() int           { return len(n) }
func (n numbers) Less(i, j int) bool { return n[i].pos < n[j].pos }
func (n numbers) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }

func searchNumbers(str string, includingWords bool) numbers {
	nums := make(numbers, 0)
	for idx, char := range str {
		if _, err := strconv.Atoi(string(char)); err == nil {
			nums = append(nums, number{str: string(char), pos: idx})
		}
	}

	if !includingWords {
		return nums
	}

	words := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	for word, num := range words {
		s := str
		for {
			if idx := strings.Index(s, word); idx != -1 {
				s = strings.Replace(s, word, strings.Repeat(num, len(word)), 1)
				nums = append(nums, number{str: num, pos: idx})
			} else {
				break
			}
		}
	}

	return nums
}
