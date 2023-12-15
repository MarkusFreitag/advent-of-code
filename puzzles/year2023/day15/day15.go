package day15

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
	"github.com/MarkusFreitag/advent-of-code/util/sliceutil"
)

func hash(s string) int {
	var current int
	for _, char := range s {
		ascii := int(rune(char))
		current += ascii
		current = current * 17
		current = current % 256
	}
	return current
}

func Part1(input string) (string, error) {
	var sum int
	for _, word := range strings.Split(input, ",") {
		sum += hash(word)
	}
	return strconv.Itoa(sum), nil
}

type Lens struct {
	Name  string
	Focal int
}

type Lenses []Lens

func (l Lenses) Len() int           { return len(l) }
func (l Lenses) Less(i, j int) bool { return l[i].Focal < l[j].Focal }
func (l Lenses) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

type Box struct {
	Slots Lenses
}

func (b Box) Index(l string) int {
	for idx, lens := range b.Slots {
		if lens.Name == l {
			return idx
		}
	}
	return -1
}

func Part2(input string) (string, error) {
	boxes := make([]Box, 256)
	for _, word := range strings.Split(input, ",") {
		if strings.Contains(word, "=") {
			parts := strings.Split(word, "=")
			lens := Lens{Name: parts[0], Focal: util.ParseInt(parts[1])}

			boxIdx := hash(lens.Name)
			box := boxes[boxIdx]
			if lensIdx := box.Index(lens.Name); lensIdx != -1 {
				box.Slots[lensIdx] = lens
			} else {
				box.Slots = append(box.Slots, lens)
			}
			boxes[boxIdx] = box
		} else {
			label := strings.TrimSuffix(word, "-")
			boxIdx := hash(label)
			box := boxes[boxIdx]
			if lensIdx := box.Index(label); lensIdx != -1 {
				_, box.Slots = sliceutil.PopIndex(box.Slots, lensIdx)
			}
			boxes[boxIdx] = box
		}
	}

	var sum int
	for bIdx, box := range boxes {
		for lIdx, lens := range box.Slots {
			sum += numbers.Multiply((bIdx + 1) * (lIdx + 1) * lens.Focal)
		}
	}
	return strconv.Itoa(sum), nil
}
