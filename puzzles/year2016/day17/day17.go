package day17

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strconv"
)

type V struct {
	X, Y int
}

func (v V) Add(vec V) V {
	return V{
		X: v.X + vec.X,
		Y: v.Y + vec.Y,
	}
}

func (v V) String() string { return fmt.Sprintf("[%d|%d]", v.Y, v.X) }

func md5Hash(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func check(pos V, path string) {
	if pos.X == 3 && pos.Y == 3 {
		paths = append(paths, path)
		return
	}
	for i, r := range md5Hash(pass + path)[:4] {
		if int(r) >= 98 && int(r) <= 102 {
			p := pos.Add(directions[i])
			if p.Y < 0 || p.Y > 3 || p.X < 0 || p.X > 3 {
				continue
			}
			switch i {
			case 0:
				check(p, path+"U")
			case 1:
				check(p, path+"D")
			case 2:
				check(p, path+"L")
			case 3:
				check(p, path+"R")
			}
		}
	}
}

var (
	pass       string
	paths      = make([]string, 0)
	directions = []V{{Y: -1}, {Y: 1}, {X: -1}, {X: 1}}
)

func Part1(input string) (string, error) {
	pass = input
	var pos V
	check(pos, "")
	switch len(paths) {
	case 0:
		return "could not find path", nil
	case 1:
		return paths[0], nil
	default:
		sort.Slice(paths, func(i, j int) bool {
			return len(paths[i]) < len(paths[j])
		})
		return paths[0], nil
	}
}

func Part2(input string) (string, error) {
	pass = input
	var pos V
	check(pos, "")
	switch len(paths) {
	case 0:
		return "could not find path", nil
	case 1:
		return strconv.Itoa(len(paths[0])), nil
	default:
		sort.Slice(paths, func(i, j int) bool {
			return len(paths[i]) > len(paths[j])
		})
		return strconv.Itoa(len(paths[0])), nil
	}
}
