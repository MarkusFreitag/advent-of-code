package day1

import (
	"container/ring"
	"strconv"
	"strings"
)

func Part1(input string) (string, error) {
	var count int
	r := ring.New(100)
	n := r.Len()
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}
	for r.Value != 50 {
		r = r.Next()
	}
	for _, line := range strings.Fields(input) {
		dir, rest := line[0], line[1:]
		num, _ := strconv.Atoi(rest)
		switch dir {
		case 'R':
			r = r.Move(num)
		case 'L':
			r = r.Move(-num)
		}
		if r.Value == 0 {
			count++
		}
	}
	return strconv.Itoa(count), nil
}

func Part2(input string) (string, error) {
	type info struct {
		val   int
		count int
	}

	r := ring.New(100)
	n := r.Len()
	for i := 0; i < n; i++ {
		r.Value = &info{val: i}
		r = r.Next()
	}
	for r.Value.(*info).val != 50 {
		r = r.Next()
	}
	for _, line := range strings.Fields(input) {
		dir, rest := line[0], line[1:]
		num, _ := strconv.Atoi(rest)
		switch dir {
		case 'R':
			for range num {
				r = r.Next()
				r.Value.(*info).count++
			}
		case 'L':
			for range num {
				r = r.Prev()
				r.Value.(*info).count++
			}
		}
	}
	for r.Value.(*info).val != 0 {
		r = r.Next()
	}
	return strconv.Itoa(r.Value.(*info).count), nil
}
