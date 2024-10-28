package util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"iter"
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

type RequestOption func(*http.Request)

func WithAoCUserAgent(repo, email string) RequestOption {
	return func(req *http.Request) {
		req.Header.Set("User-Agent", fmt.Sprintf("%s by %s", repo, email))
	}
}

func InputFromURL(year, day int, reqOptions ...RequestOption) (string, error) {
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
	for _, option := range reqOptions {
		option(req)
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
	return strings.TrimSuffix(string(input), "\n"), nil
}

func InputFromFile(year, day int) (string, error) {
	content, err := ioutil.ReadFile(inputFilename(year, day))
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(string(content), "\n"), nil
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

func RangeInt(from, to, steps int) iter.Seq[int] {
	upwards := func(yield func(int) bool) {
		for n := from; n <= to; n += steps {
			if !yield(n) {
				return
			}
		}
	}
	downwards := func(yield func(int) bool) {
		for n := from; n >= to; n -= steps {
			if !yield(n) {
				return
			}
		}
	}

	if from > to {
		return downwards
	}
	return upwards
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

func Repeat[E any](value E, count int) []E {
	slice := make([]E, count)
	for i := 0; i < count; i++ {
		slice[i] = value
	}
	return slice
}

// Inspired by slices.Collect, but when you also need to transform each item
func CollectT[E, R any](seq iter.Seq[E], transform func(item E) R) []R {
	slice := make([]R, 0)
	for value := range seq {
		slice = append(slice, transform(value))
	}
	return slice
}
