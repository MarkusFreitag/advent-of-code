package day9

import (
	"slices"
	"strconv"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/sliceutil"
)

func Part1(input string) (string, error) {
	diskmap := make([]int, 0)
	for idx, char := range input {
		id := -1
		if idx%2 == 0 {
			id = idx / 2
		}
		diskmap = append(diskmap, util.Repeat(id, util.ParseInt(string(char)))...)
	}
	for idx := len(diskmap) - 1; idx >= 0; idx-- {
		id := diskmap[idx]
		if id == -1 {
			continue
		}
		spaceIdx := slices.IndexFunc(diskmap, func(id int) bool { return id == -1 })
		if spaceIdx == -1 || spaceIdx >= idx {
			// no more free space blocks left from the current file block
			break
		}
		diskmap[spaceIdx], diskmap[idx] = diskmap[idx], diskmap[spaceIdx]
	}
	var checksum int
	for idx, id := range diskmap {
		if id == -1 {
			continue
		}
		checksum += idx * id
	}
	return strconv.Itoa(checksum), nil

}

func Part2(input string) (string, error) {
	type File struct {
		ID    int
		Size  int
		Space bool
	}

	diskmap := make([]File, 0)
	for idx, char := range input {
		diskmap = append(diskmap, File{
			ID:    idx / 2,
			Size:  util.ParseInt(string(char)),
			Space: idx%2 != 0,
		})
	}

	files := sliceutil.Map(diskmap, func(f File) (File, bool) { return f, !f.Space })
	slices.Reverse(files)

	for _, file := range files {
		idx := slices.IndexFunc(diskmap, func(f File) bool { return f.ID == file.ID })
		spaceIdx := slices.IndexFunc(diskmap, func(f File) bool { return f.Space && f.Size >= file.Size })
		if spaceIdx == -1 || spaceIdx > idx {
			continue
		}
		sizeLeft := diskmap[spaceIdx].Size - file.Size
		if sizeLeft == 0 {
			diskmap[idx], diskmap[spaceIdx] = diskmap[spaceIdx], diskmap[idx]
		} else {
			diskmap[spaceIdx].Size = sizeLeft
			diskmap[idx].Space = true
			diskmap = slices.Insert(diskmap, spaceIdx, File{ID: file.ID, Size: file.Size})
		}
	}
	var checksum int
	var index int
	for _, file := range diskmap {
		if file.Space {
			index += file.Size
			continue
		}
		for range file.Size {
			checksum += index * file.ID
			index++
		}
	}
	return strconv.Itoa(checksum), nil
}
