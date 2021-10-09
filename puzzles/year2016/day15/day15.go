package day15

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

type Disks []Disk

func parseInput(input string) Disks {
	lines := strings.Split(input, "\n")
	disks := make(Disks, len(lines))
	for idx, line := range lines {
		fields := strings.Fields(line)
		disks[idx] = NewDisk(
			util.ParseInt(fields[3]),
			util.ParseInt(strings.TrimSuffix(fields[len(fields)-1], ".")),
		)
	}
	return disks
}

func (d Disks) Passthrough(time int) bool {
	pos := d[0].CalculatePos(time + 1)
	for idx, disk := range d {
		if idx == 0 {
			continue
		}
		next := disk.CalculatePos(time + idx + 1)
		if next != pos {
			return false
		}
		pos = next
	}
	return true
}

func (d Disks) Timing() int {
	var time int
	for {
		if d.Passthrough(time) {
			break
		}
		time++
	}
	return time
}

type Disk struct {
	Positions int
	Current   int
}

func NewDisk(positions, startPos int) Disk {
	return Disk{
		Positions: positions,
		Current:   startPos,
	}
}

func (d Disk) CalculatePos(offset int) int {
	return (d.Current + offset) % d.Positions
}

func Part1(input string) (string, error) {
	disks := parseInput(input)
	return strconv.Itoa(disks.Timing()), nil
}

func Part2(input string) (string, error) {
	disks := parseInput(input)
	disks = append(disks, NewDisk(11, 0))
	return strconv.Itoa(disks.Timing()), nil
}
