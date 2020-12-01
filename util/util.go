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

func Sum(nums ...int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
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
