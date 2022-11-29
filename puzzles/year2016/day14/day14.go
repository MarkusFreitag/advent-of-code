package day14

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

func searchTriplet(hash string) string {
	for i := 0; i < len(hash)-2; i++ {
		if hash[i] == hash[i+1] && hash[i] == hash[i+2] {
			return string(hash[i])
		}
	}
	return ""
}

func md5Hash(str string, times int) string {
	hashString := str
	for i := 0; i < times; i++ {
		hash := md5.Sum([]byte(hashString))
		hashString = hex.EncodeToString(hash[:])
	}
	return hashString
}

type Triplet struct {
	Index int
	Char  string
}

func getLastKeyIndex(salt string, hashingRounds int) int {
	var keyCounter int
	triplets := make([]Triplet, 0)
	for i := 0; true; i++ {
		hash := md5Hash(salt+strconv.Itoa(i), hashingRounds)

		for _, triplet := range triplets {
			if numbers.Between(i, triplet.Index+1, triplet.Index+1000) {
				if strings.Contains(hash, strings.Repeat(triplet.Char, 5)) {
					keyCounter++
					if keyCounter == 64 {
						return triplet.Index
					}
				}
			}
		}

		if triplet := searchTriplet(hash); triplet != "" {
			triplets = append(triplets, Triplet{Index: i, Char: triplet})
		}
	}
	return -1
}

func Part1(input string) (string, error) {
	return strconv.Itoa(getLastKeyIndex(input, 1)), nil
}

func Part2(input string) (string, error) {
	return strconv.Itoa(getLastKeyIndex(input, 2017)), nil
}
