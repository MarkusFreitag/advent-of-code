package day11

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
	"github.com/MarkusFreitag/advent-of-code/util/slice"
)

type Op struct {
	Left, Middle, Right string
}

type Test struct {
	Value int
	True  int
	False int
}

type Monkey struct {
	ID        int
	Items     []int
	Operation Op
	Test      Test
	Activity  int
}

func (m *Monkey) Add(item int) {
	m.Items = append(m.Items, item)
}

func (m *Monkey) Inspect(item int) int {
	nums := [2]int{}
	if m.Operation.Left == "old" {
		nums[0] = item
	} else {
		nums[0] = util.ParseInt(m.Operation.Left)
	}
	if m.Operation.Right == "old" {
		nums[1] = item
	} else {
		nums[1] = util.ParseInt(m.Operation.Right)
	}
	switch m.Operation.Middle {
	case "+":
		return nums[0] + nums[1]
	case "*":
		return nums[0] * nums[1]
	}
	panic("ERROR")
}

func (m *Monkey) DoTest(item int) bool {
	return item%m.Test.Value == 0
}

type Monkeys []*Monkey

func (m Monkeys) Round(relief func(int) int) {
	for _, monkey := range m {
		for len(monkey.Items) > 0 {
			var item int
			item, monkey.Items = slice.PopFront(monkey.Items)
			item = monkey.Inspect(item)
			item = relief(item)
			if monkey.DoTest(item) {
				m[monkey.Test.True].Add(item)
			} else {
				m[monkey.Test.False].Add(item)
			}
			monkey.Activity++
		}
	}
}

func (m Monkeys) MBL() int {
	nums := make([]int, len(m))
	for idx, monkey := range m {
		nums[idx] = monkey.Activity
	}
	slice.SortDesc(nums)
	return nums[0] * nums[1]
}

func parseInput(input string) Monkeys {
	blocks := strings.Split(input, "\n\n")
	monkeys := make(Monkeys, len(blocks))
	for idx, block := range blocks {
		monkey := &Monkey{ID: idx, Items: make([]int, 0)}
		monkeys[idx] = monkey
		lines := strings.Split(block, "\n")

		fields := strings.Fields(lines[1])
		for _, field := range fields[2:] {
			monkey.Items = append(monkey.Items, util.ParseInt(strings.TrimSuffix(field, ",")))
		}

		fields = strings.Fields(lines[2])
		monkey.Operation = Op{
			Left:   fields[3],
			Middle: fields[4],
			Right:  fields[5],
		}

		fields = strings.Fields(lines[3])
		monkey.Test = Test{
			Value: util.ParseInt(fields[3]),
		}

		fields = strings.Fields(lines[4])
		monkey.Test.True = util.ParseInt(fields[len(fields)-1])
		fields = strings.Fields(lines[5])
		monkey.Test.False = util.ParseInt(fields[len(fields)-1])
	}
	return monkeys
}

func Part1(input string) (string, error) {
	monkeys := parseInput(input)
	for i := 0; i < 20; i++ {
		monkeys.Round(func(item int) int {
			return item / 3
		})
	}
	return strconv.Itoa(monkeys.MBL()), nil
}

func Part2(input string) (string, error) {
	monkeys := parseInput(input)

	nums := make([]int, len(monkeys))
	for idx, monkey := range monkeys {
		nums[idx] = monkey.Test.Value
	}
	relief := numbers.Multiply(nums...)

	for i := 0; i < 10000; i++ {
		monkeys.Round(func(item int) int {
			return item % relief
		})
	}
	return strconv.Itoa(monkeys.MBL()), nil
}
