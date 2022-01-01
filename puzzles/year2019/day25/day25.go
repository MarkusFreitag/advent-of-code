package day25

import (
	"bufio"
	"fmt"
	"os"

	"github.com/MarkusFreitag/advent-of-code/intcode"
)

func strToASCII(s string) []int64 {
	nums := make([]int64, len(s))
	for idx, r := range s {
		nums[idx] = int64(r)
	}
	return nums
}

func Part1(input string) (string, error) {
	icode, err := intcode.New(input)
	if err != nil {
		return "", err
	}
	in := make(chan int64)
	out := make(chan intcode.Message, 100)
	go intcode.RunSync(icode, in, out)

	running := true
	reader := bufio.NewReader(os.Stdin)
	var ascii []int64
	for running {
		msg := <-out
		switch msg.Type {
		case intcode.MessageWantsInput:
			if len(ascii) == 0 {
				input, _ := reader.ReadString('\n')
				ascii = strToASCII(string(input))
			}
			var num int64
			num, ascii = ascii[0], ascii[1:]
			in <- num
		case intcode.MessageOutput:
			if msg.Value < 127 {
				fmt.Printf("%d", msg.Value)
			} else {
				fmt.Println(msg.Value)
			}
		case intcode.MessageHalt:
			running = false
		}
	}

	return "n/a", nil
}

func Part2(input string) (string, error) {
	return "n/a", nil
}
