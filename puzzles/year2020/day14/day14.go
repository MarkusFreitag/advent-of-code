package day14

import (
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

const BITMASKSIZE = 36

var (
	rgxMask = regexp.MustCompile(`^mask\s=\s([X01]{36})$`)
	rgxVal  = regexp.MustCompile(`^mem\[(\d+)\]\s=\s(\d+)$`)
)

func getAddrAndValue(line string) (int, int) {
	matches := rgxVal.FindAllStringSubmatch(line, -1)[0]
	addr, err := strconv.Atoi(matches[1])
	if err != nil {
		panic(err)
	}
	val, err := strconv.Atoi(matches[2])
	if err != nil {
		panic(err)
	}
	return addr, val
}

func applyMaskV1(dec int, mask string) int {
	bin := strconv.FormatInt(int64(dec), 2)
	bin = util.LeftPad(bin, "0", BITMASKSIZE)
	var result string
	for idx, char := range bin {
		switch mask[idx] {
		case '0':
			result += "0"
		case '1':
			result += "1"
		default:
			result += string(char)
		}
	}

	v, err := strconv.ParseInt(result, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(v)
}

func Part1(input string) (string, error) {
	mem := make(map[int]int)
	var mask string
	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "mask") {
			mask = rgxMask.FindAllStringSubmatch(line, -1)[0][1]
		} else {
			addr, val := getAddrAndValue(line)
			mem[addr] = applyMaskV1(val, mask)
		}
	}
	var sum int
	for _, val := range mem {
		sum += val
	}
	return strconv.Itoa(sum), nil
}

func applyMaskV2(dec int, mask string) []int {
	bin := strconv.FormatInt(int64(dec), 2)
	bin = util.LeftPad(bin, "0", BITMASKSIZE)
	var result string
	for idx, char := range bin {
		switch mask[idx] {
		case '0':
			result += string(char)
		case '1':
			result += "1"
		default:
			result += "X"
		}
	}

	x := strings.Count(mask, "X")
	p := int(math.Pow(2, float64(x)))
	results := make([]int, 0)

	for i := 0; i < p; i++ {
		m := strconv.FormatInt(int64(i), 2)
		m = util.LeftPad(m, "0", x)

		r := result
		for _, b := range m {
			r = strings.Replace(r, "X", string(b), 1)
		}

		v, err := strconv.ParseInt(r, 2, 64)
		if err != nil {
			panic(err)
		}
		results = append(results, int(v))
	}
	return results
}

func Part2(input string) (string, error) {
	mem := make(map[int]int)
	var mask string
	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "mask") {
			mask = rgxMask.FindAllStringSubmatch(line, -1)[0][1]
		} else {
			addr, val := getAddrAndValue(line)
			for _, a := range applyMaskV2(addr, mask) {
				mem[a] = val
			}
		}
	}
	var sum int
	for _, val := range mem {
		sum += val
	}
	return strconv.Itoa(sum), nil
}
