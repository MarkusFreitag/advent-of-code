package day10

import (
	"fmt"
	"strings"
	"sync"
)

func mutate(str string) string {
	var mStr string
	for pos := 0; pos < len(str); {
		var counter int
		for i := pos; i < len(str); i++ {
			if string(str[i]) != string(str[pos]) {
				break
			}
			counter++
		}
		mStr += fmt.Sprintf("%d%s", counter, string(str[pos]))
		pos += counter
	}
	return mStr
}

func split(str string, size int) []string {
	chunks := make([]string, 0)
	for str != "" {
		var sub string
		if len(str) > size {
			sub = str[:size]
			str = str[size:]
			for sub[len(sub)-1] == str[0] {
				sub += string(str[0])
				str = str[1:]
			}
		} else {
			sub = str
			str = ""
		}
		chunks = append(chunks, sub)
	}
	return chunks
}

func mutateFast(str string) string {
	chunkSize := 1000
	chunks := split(str, chunkSize)
	mutChunks := make([]string, len(chunks))
	var wg sync.WaitGroup
	for idx, chunk := range chunks {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int, s string) {
			defer wg.Done()
			mutChunks[i] = mutate(s)
		}(&wg, idx, chunk)
	}
	wg.Wait()
	return strings.Join(mutChunks, "")
}

func Part1(input string) (string, error) {
	str := mutateFast(input)
	for i := 0; i < 39; i++ {
		str = mutateFast(str)
	}
	return fmt.Sprintf("%d", len(str)), nil
}

func Part2(input string) (string, error) {
	str := mutateFast(input)
	for i := 0; i < 49; i++ {
		str = mutateFast(str)
	}
	return fmt.Sprintf("%d", len(str)), nil
}
