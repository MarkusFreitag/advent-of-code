package util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

const (
	URLSCHEMA          = "https://adventofcode.com/%d/day/%d/input"
	FILESCHEMA         = "%d_%d.txt"
	MaxUnsignedInteger = ^uint(0)
	MinUnsignedInteger = 0
	MaxInteger         = int(MaxUnsignedInteger >> 1)
	MinInteger         = -MaxInteger - 1
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

func StrInSlice(str string, strs []string) bool {
	for _, s := range strs {
		if s == str {
			return true
		}
	}
	return false
}

func IntInSlice(num int, nums []int) bool {
	for _, i := range nums {
		if i == num {
			return true
		}
	}
	return false
}

func SumInts(nums ...int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SubInts(nums ...int) int {
	var sum int
	for _, num := range nums {
		sum -= num
	}
	return sum
}

func MulInts(nums ...int) int {
	product := 1
	for _, num := range nums {
		product *= num
	}
	return product
}

func PowInt(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func StrsToInts(slice []string) []int {
	nums := make([]int, len(slice))
	for idx, str := range slice {
		num, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		nums[idx] = num
	}
	return nums
}

func IntsToStrs(slice []int) []string {
	nums := make([]string, len(slice))
	for idx, i := range slice {
		nums[idx] = strconv.Itoa(i)
	}
	return nums
}

func Abs(i int) int {
	if i < 0 {
		i *= -1
	}
	return i
}

func Between(i, min, max int) bool {
	return min <= i && i <= max
}

func ParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func ParseSignedInt(s string) int {
	if strings.HasPrefix(s, "-") {
		return ParseInt(strings.TrimPrefix(s, "-")) * -1
	}
	return ParseInt(strings.TrimPrefix(s, "+"))
}

func Reverse(str string) string {
	var s string
	for _, r := range str {
		s = string(r) + s
	}
	return s
}

func LeftPad(str, padStr string, length int) string {
	padCount := 1 + ((length - len(padStr)) / len(padStr))
	s := strings.Repeat(padStr, padCount) + str
	return s[(len(s) - length):]
}

func RightPad(str, padStr string, length int) string {
	padCount := 1 + ((length - len(padStr)) / len(padStr))
	s := str + strings.Repeat(padStr, padCount)
	return s[:length]
}

func BinStrToDecInt(bin string) int {
	dec, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(dec)
}

func DecIntToBinStr(dec int) string {
	return strconv.FormatInt(int64(dec), 2)
}

func InRange(i, min, max int) bool {
	return i >= min && i <= max
}

func MinInt(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	min := nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
	}
	return min
}

func MaxInt(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for _, n := range nums[1:] {
		if n > max {
			max = n
		}
	}
	return max
}

func MinMaxInt(nums ...int) (int, int) {
	if len(nums) == 0 {
		return 0, 0
	}
	min, max := nums[0], nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
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
	if Abs(crossproduct) != 0 {
		return false
	}

	dotproduct := (cX-aX)*(bX-aX) + (cY-aY)*(bY-aY)
	if dotproduct < 0 {
		return false
	}

	squaredlengthba := (bX-aX)*(bX-aX) + (bY-aY)*(bY-aY)

	return dotproduct <= squaredlengthba
}

type Bools []bool

func (b Bools) All(state bool) bool {
	for _, i := range b {
		if i != state {
			return false
		}
	}
	return true
}

func (b Bools) Any(state bool) bool {
	for _, i := range b {
		if i == state {
			return true
		}
	}
	return false
}

func StrToStrs(str string) []string {
	slice := make([]string, len(str))
	for idx, char := range str {
		slice[idx] = string(char)
	}
	return slice
}

func StringSorter(str string) string {
	slice := StrToStrs(str)
	sort.Strings(slice)
	return strings.Join(slice, "")
}
