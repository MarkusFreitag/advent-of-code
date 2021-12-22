package main // import "github.com/MarkusFreitag/advent-of-code"

//go:generate go run gen.go $YEAR $DAY

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/MarkusFreitag/advent-of-code/puzzles"
	"github.com/MarkusFreitag/advent-of-code/util"
)

var (
	yearFlag = flag.Int("year", time.Now().Year(), "specifying year")
	dayFlag  = flag.Int("day", time.Now().Day(), "specifying day")
)

func main() {
	flag.Parse()

	input, err := util.InputFromFile(*yearFlag, *dayFlag)
	if os.IsNotExist(err) {
		fmt.Printf("input file for %d:%d does not exist\n", *yearFlag, *dayFlag)
		return
	} else if err != nil {
		fmt.Printf("error while loading input from file: %s\n", err.Error())
		return
	}

	puzzle, err := puzzles.Get(*yearFlag, *dayFlag)
	if err != nil {
		fmt.Printf("error while looking for puzzle %d_%d: %s\n", *yearFlag, *dayFlag, err.Error())
		return
	}
	for idx, part := range puzzle {
		start := time.Now()
		solution, err := part(input)
		if err != nil {
			fmt.Printf("error while solving puzzle %d_%d part %d: %s\n", *yearFlag, *dayFlag, idx+1, err.Error())
			return
		}
		fmt.Printf("[%s] solution for puzzle %d_%d part %d: %s\n", time.Since(start), *yearFlag, *dayFlag, idx+1, solution)
	}
}
