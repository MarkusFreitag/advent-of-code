package day4

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func Part1(input string) (string, error) {
	var key int
	for {
		data := []byte(fmt.Sprintf("%s%d", input, key))
		sum := fmt.Sprintf("%x", md5.Sum(data))
		if strings.HasPrefix(sum, "00000") {
			return strconv.Itoa(key), nil
		}
		key++
	}
}

func Part2(input string) (string, error) {
	var key int
	for {
		data := []byte(fmt.Sprintf("%s%d", input, key))
		sum := fmt.Sprintf("%x", md5.Sum(data))
		if strings.HasPrefix(sum, "000000") {
			return strconv.Itoa(key), nil
		}
		key++
	}
}
