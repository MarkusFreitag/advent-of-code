package day5

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func Part1(input string) (string, error) {
	var pass string
	for i := 0; true; i++ {
		md5Sum := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", input, i))))
		if strings.HasPrefix(md5Sum, "00000") {
			pass += string(md5Sum[5])
		}
		if len(pass) == 8 {
			break
		}
	}
	return pass, nil
}

func Part2(input string) (string, error) {
	chars := make([]string, 8)
	for i := 0; true; i++ {
		md5Sum := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", input, i))))
		if strings.HasPrefix(md5Sum, "00000") {
			pos, err := strconv.Atoi(string(md5Sum[5]))
			if err != nil {
				continue
			}
			if pos < len(chars) && chars[pos] == "" {
				chars[pos] = string(md5Sum[6])
			}
		}
		if len(strings.Join(chars, "")) == 8 {
			break
		}
	}
	return strings.Join(chars, ""), nil
}
