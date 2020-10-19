package day14

import (
	"sort"
	"strconv"
	"strings"
)

type reindeer struct {
	name     string
	speed    int
	duration int
	rest     int
	traveled int
	score    int
}

func (r *reindeer) run(tick chan int) {
	var pause bool
	dur := r.duration
	for range tick {
		if !pause {
			r.traveled += r.speed
		}
		dur--
		if dur == 0 {
			if pause {
				dur = r.duration
				pause = false
			} else {
				dur = r.rest
				pause = true
			}
		}
	}
}

type byTraveled []*reindeer

func (s byTraveled) Len() int {
	return len(s)
}
func (s byTraveled) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byTraveled) Less(i, j int) bool {
	return s[i].traveled > s[j].traveled
}

type byScore []*reindeer

func (s byScore) Len() int {
	return len(s)
}
func (s byScore) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byScore) Less(i, j int) bool {
	return s[i].score > s[j].score
}

func Part1(input string) (string, error) {
	reindeers := make([]*reindeer, 0)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		speed, err := strconv.Atoi(parts[3])
		if err != nil {
			return "", err
		}
		duration, err := strconv.Atoi(parts[6])
		if err != nil {
			return "", err
		}
		rest, err := strconv.Atoi(parts[13])
		if err != nil {
			return "", err
		}

		reindeers = append(reindeers, &reindeer{
			name:     parts[0],
			speed:    speed,
			duration: duration,
			rest:     rest,
		})
	}

	tickers := make([]chan int, len(reindeers))
	for idx, r := range reindeers {
		tick := make(chan int)
		tickers[idx] = tick
		go r.run(tick)
	}
	for i := 0; i < 2503; i++ {
		for _, tick := range tickers {
			tick <- i
		}
	}
	for _, tick := range tickers {
		close(tick)
	}
	sort.Sort(byTraveled(reindeers))
	return strconv.Itoa(reindeers[0].traveled), nil
}

func Part2(input string) (string, error) {
	reindeers := make([]*reindeer, 0)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		speed, err := strconv.Atoi(parts[3])
		if err != nil {
			return "", err
		}
		duration, err := strconv.Atoi(parts[6])
		if err != nil {
			return "", err
		}
		rest, err := strconv.Atoi(parts[13])
		if err != nil {
			return "", err
		}

		reindeers = append(reindeers, &reindeer{
			name:     parts[0],
			speed:    speed,
			duration: duration,
			rest:     rest,
		})
	}

	tickers := make([]chan int, len(reindeers))
	for idx, r := range reindeers {
		tick := make(chan int)
		tickers[idx] = tick
		go r.run(tick)
	}
	for i := 0; i < 2503; i++ {
		for _, tick := range tickers {
			tick <- i
		}
		var highest int
		for _, r := range reindeers {
			if r.traveled > highest {
				highest = r.traveled
			}
		}
		for _, r := range reindeers {
			if r.traveled == highest {
				r.score++
			}
		}
	}
	for _, tick := range tickers {
		close(tick)
	}
	sort.Sort(byScore(reindeers))
	return strconv.Itoa(reindeers[0].score), nil
}
