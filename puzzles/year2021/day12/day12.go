package day12

import (
	"strconv"
	"strings"
)

type Cave struct {
	ID          string
	Small       bool
	Connections []*Cave
}

func NewCave(id string) *Cave {
	c := Cave{
		ID:          id,
		Connections: make([]*Cave, 0),
	}
	c.Small = id == strings.ToLower(id)
	return &c
}

func parseInput(input string) map[string]*Cave {
	caveMap := make(map[string]*Cave)
	for _, line := range strings.Fields(input) {
		parts := strings.Split(line, "-")
		a, ok := caveMap[parts[0]]
		if !ok {
			a = NewCave(parts[0])
			caveMap[a.ID] = a
		}
		b, ok := caveMap[parts[1]]
		if !ok {
			b = NewCave(parts[1])
			caveMap[b.ID] = b
		}
		a.Connections = append(a.Connections, b)
		b.Connections = append(b.Connections, a)
	}
	return caveMap
}

func findPaths(c *Cave, visited map[string]bool) int {
	if c.ID == "end" {
		return 1
	}
	var paths int
	for _, conn := range c.Connections {
		if conn.ID == "start" {
			continue
		}
		if _, ok := visited[conn.ID]; ok && conn.Small {
			continue
		}
		visited[conn.ID] = true
		paths += findPaths(conn, visited)
		delete(visited, conn.ID)
	}
	return paths
}

func Part1(input string) (string, error) {
	caveMap := parseInput(input)
	return strconv.Itoa(
		findPaths(
			caveMap["start"],
			make(map[string]bool),
		),
	), nil
}

func findPaths2(c *Cave, visited map[string]bool, searched bool) int {
	if c.ID == "end" {
		return 1
	}

	var paths int
	for _, conn := range c.Connections {
		if conn.ID == "start" {
			continue
		}
		if _, ok := visited[conn.ID]; ok && conn.Small {
			if !searched {
				paths += findPaths2(conn, visited, true)
			}
		} else {
			visited[conn.ID] = true
			paths += findPaths2(conn, visited, searched)
			delete(visited, conn.ID)
		}
	}
	return paths
}

func Part2(input string) (string, error) {
	caveMap := parseInput(input)
	return strconv.Itoa(
		findPaths2(
			caveMap["start"],
			make(map[string]bool),
			false,
		),
	), nil
}
