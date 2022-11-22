package day7

import (
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

type IPv7 struct {
	address string
}

func NewIPv7(addr string) IPv7 {
	return IPv7{address: addr}
}

func (i IPv7) SupernetSequences() []string {
	seqs := make([]string, 0)
	for _, seq := range strings.Split(i.address, "[") {
		if strings.Contains(seq, "]") {
			seqs = append(seqs, strings.Split(seq, "]")[1])
		} else {
			seqs = append(seqs, seq)
		}
	}
	return seqs
}

func (i IPv7) HypernetSequences() []string {
	seqs := make([]string, 0)
	for _, seq := range strings.Split(i.address, "[") {
		if strings.Contains(seq, "]") {
			seqs = append(seqs, strings.Split(seq, "]")[0])
		}
	}
	return seqs
}

func (i IPv7) TLSSupport() bool {
	for _, seq := range i.HypernetSequences() {
		if detectABBA(seq) {
			return false
		}
	}
	for _, seq := range i.SupernetSequences() {
		if detectABBA(seq) {
			return true
		}
	}
	return false
}

func (i IPv7) SSLSupport() bool {
	abas := make([]string, 0)
	for _, seq := range i.SupernetSequences() {
		abas = append(abas, getABAs(seq)...)
	}
	for _, aba := range abas {
		for _, seq := range i.HypernetSequences() {
			bab := string(aba[1]) + aba[:2]
			if strings.Contains(seq, bab) {
				return true
			}
		}
	}
	return false
}

func detectABBA(str string) bool {
	for i := 0; i < len(str)-3; i++ {
		abba := str[i : i+2]
		if abba[0] == abba[1] {
			continue
		}
		abba += util.StringReverse(abba)
		if str[i:i+4] == abba {
			return true
		}
	}
	return false
}

func getABAs(str string) []string {
	abas := make([]string, 0)
	for i := 0; i < len(str)-2; i++ {
		aba := str[i:i+2] + string(str[i])
		if aba[0] == aba[1] {
			continue
		}
		if str[i:i+3] == aba {
			abas = append(abas, aba)
		}
	}
	return abas
}

func Part1(input string) (string, error) {
	var counter int
	for _, line := range strings.Split(input, "\n") {
		if NewIPv7(line).TLSSupport() {
			counter++
		}
	}
	return strconv.Itoa(counter), nil
}

func Part2(input string) (string, error) {
	var counter int
	for _, line := range strings.Split(input, "\n") {
		if NewIPv7(line).SSLSupport() {
			counter++
		}
	}
	return strconv.Itoa(counter), nil
}
