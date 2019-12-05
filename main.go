package main // import "github.com/MarkusFreitag/advent-of-code"

import (
	"flag"
	"fmt"
	"os"

	"github.com/MarkusFreitag/advent-of-code/pkg/puzzles"
	"github.com/MarkusFreitag/advent-of-code/pkg/util"
)

var (
	yearFlag = flag.Int("year", -1, "specifying year")
	dayFlag  = flag.Int("day", -1, "specifying day")
)

func main() {
	flag.Parse()
	if *yearFlag == -1 || *dayFlag == -1 {
		fmt.Println("-year and -day flags are required")
		return
	}

	input, err := util.InputFromFile(*yearFlag, *dayFlag)
	if os.IsNotExist(err) {
		input, err = util.InputFromURL(*yearFlag, *dayFlag)
		if err != nil {
			fmt.Printf("error while loading input from url: %s\n", err.Error())
			return
		}
		err = util.InputToFile(*yearFlag, *dayFlag, input)
		if err != nil {
			fmt.Printf("error while loading input from url: %s\n", err.Error())
		}
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
		solution, err := part.Solve(input)
		if err != nil {
			fmt.Printf("error while solving puzzle %d_%d part %d: %s\n", *yearFlag, *dayFlag, idx+1, err.Error())
			return
		}
		fmt.Printf("solution for puzzle %d_%d part %d: %s\n", *yearFlag, *dayFlag, idx+1, solution)
	}
}
