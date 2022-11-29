package day14

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
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

func applyMask(dec int, mask string, ignore byte) string {
	var result string
	for idx, char := range util.StringPadLeft(util.DecIntToBinString(dec), "0", BITMASKSIZE) {
		if mask[idx] == ignore {
			result += string(char)
		} else {
			result += string(mask[idx])
		}
	}
	return result
}

func Part1(input string) (string, error) {
	mem := make(map[int]int)
	var mask string
	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "mask") {
			mask = rgxMask.FindAllStringSubmatch(line, -1)[0][1]
		} else {
			addr, val := getAddrAndValue(line)
			mem[addr] = util.BinStringToDecInt(applyMask(val, mask, 'X'))
		}
	}
	var sum int
	for _, val := range mem {
		sum += val
	}
	return strconv.Itoa(sum), nil
}

func Part2(input string) (string, error) {
	mem := make(map[int]int)
	var mask string
	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "mask") {
			mask = rgxMask.FindAllStringSubmatch(line, -1)[0][1]
		} else {
			addr, val := getAddrAndValue(line)
			modAddr := applyMask(addr, mask, '0')
			xCount := strings.Count(mask, "X")
			for i := 0; i < numbers.Pow(2, xCount); i++ {
				a := modAddr
				for _, b := range util.StringPadLeft(util.DecIntToBinString(i), "0", xCount) {
					a = strings.Replace(a, "X", string(b), 1)
				}
				mem[util.BinStringToDecInt(a)] = val
			}
		}
	}
	var sum int
	for _, val := range mem {
		sum += val
	}
	return strconv.Itoa(sum), nil
}
