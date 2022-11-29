package util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

const (
	URLSCHEMA  = "https://adventofcode.com/%d/day/%d/input"
	FILESCHEMA = "%d_%d.txt"
)

var ErrNotSolved = errors.New("this part is not solved yet")

type Solution func(string) (string, error)

type Puzzle []Solution

func inputFilename(year, day int) string {
	return filepath.Join("inputs", fmt.Sprintf("%d_%d.txt", year, day))
}

func InputFromURL(year, day int) (string, error) {
	if _, err := os.Stat("aoc.session"); os.IsNotExist(err) {
		return "", err
	}
	cookie, err := ioutil.ReadFile("aoc.session")
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("GET", fmt.Sprintf(URLSCHEMA, year, day), nil)
	if err != nil {
		return "", err
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: strings.TrimSpace(string(cookie))})
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("could not access input statuscode: %d", resp.StatusCode)
	}
	input, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(input)), nil
}

func InputFromFile(year, day int) (string, error) {
	content, err := ioutil.ReadFile(inputFilename(year, day))
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(content)), nil
}

func InputToFile(year, day int, input string) error {
	return ioutil.WriteFile(inputFilename(year, day), []byte(input), 0644)
}

func StringsToInts(slice []string) []int {
	nums := make([]int, len(slice))
	for idx, str := range slice {
		nums[idx] = ParseInt(str)
	}
	return nums
}

func IntsToStrings(slice []int) []string {
	nums := make([]string, len(slice))
	for idx, i := range slice {
		nums[idx] = strconv.Itoa(i)
	}
	return nums
}

func ParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func ParseFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return f
}

func BinStringToDecInt(bin string) int {
	dec, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(dec)
}

func DecIntToBinString(dec int) string {
	return strconv.FormatInt(int64(dec), 2)
}

func RangeInt(from, to, steps int) <-chan int {
	upwards := func(from, to, steps int, c chan int) {
		for n := from; n <= to; n += steps {
			c <- n
		}
		close(c)
	}
	downwards := func(from, to, steps int, c chan int) {
		for n := from; n >= to; n -= steps {
			c <- n
		}
		close(c)
	}

	c := make(chan int)
	if from > to {
		go downwards(from, to, steps, c)
		return c
	}
	go upwards(from, to, steps, c)
	return c
}

func OnLineInt(aX, aY, bX, bY, cX, cY int) bool {
	crossproduct := (cY-aY)*(bX-aX) - (cX-aX)*(bY-aY)
	if numbers.Abs(crossproduct) != 0 {
		return false
	}

	dotproduct := (cX-aX)*(bX-aX) + (cY-aY)*(bY-aY)
	if dotproduct < 0 {
		return false
	}

	squaredlengthba := (bX-aX)*(bX-aX) + (bY-aY)*(bY-aY)

	return dotproduct <= squaredlengthba
}

func Memoize(fn func(any) any) func(any) any {
	cache := make(map[any]any)

	return func(i any) any {
		if val, found := cache[i]; found {
			return val
		}

		result := fn(i)
		cache[i] = result
		return result
	}
}

func Flatten(slice []any) []any {
	flat := make([]any, 0)
	for _, item := range slice {
		switch v := item.(type) {
		case []any:
			flat = append(flat, Flatten(v)...)
		default:
			flat = append(flat, item)
		}
	}
	return flat
}
