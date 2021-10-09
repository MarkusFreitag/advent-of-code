package day20

import (
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func ParseRange(s string) Range {
	parts := strings.Split(s, "-")
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])
	return Range{
		start: start,
		end:   end,
	}
}

func Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")
	ranges := make([]Range, len(lines))
	for index, line := range lines {
		ranges[index] = ParseRange(line)
	}
	var counter int
	for {
		var blocked bool
		for idx := len(ranges) - 1; idx >= 0; idx-- {
			if counter > ranges[idx].end {
				ranges = append(ranges[:idx], ranges[idx+1:]...)
				continue
			}
			if counter == ranges[idx].start {
				counter = ranges[idx].end
				blocked = true
				break
			} else if counter > ranges[idx].start && counter <= ranges[idx].end {
				blocked = true
				break
			}
		}
		if !blocked {
			break
		}
		counter++
	}
	return strconv.Itoa(counter), nil
}

var max = 4294967295

func Part2(input string) (string, error) {
	lines := strings.Split(input, "\n")
	ranges := make([]Range, len(lines))
	for index, line := range lines {
		ranges[index] = ParseRange(line)
	}
	allowed := make([]int, 0)
	for counter := 0; counter <= max; counter++ {
		var blocked bool
		for idx := len(ranges) - 1; idx >= 0; idx-- {
			if counter > ranges[idx].end {
				ranges = append(ranges[:idx], ranges[idx+1:]...)
				continue
			}
			if counter == ranges[idx].start {
				counter = ranges[idx].end
				ranges = append(ranges[:idx], ranges[idx+1:]...)
				blocked = true
				break
			} else if counter > ranges[idx].start && counter <= ranges[idx].end {
				blocked = true
				break
			}
		}
		if !blocked {
			allowed = append(allowed, counter)
		}
	}
	return strconv.Itoa(len(allowed)), nil
}
