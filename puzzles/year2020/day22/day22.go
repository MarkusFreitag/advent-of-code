package day22

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

func score(cards []int) int {
	var score int
	for idx, card := range cards {
		score += (len(cards) - idx) * card
	}
	return score
}

func Part1(input string) (string, error) {
	blocks := strings.Split(input, "\n\n")

	player1 := util.StrsToInts(strings.Split(blocks[0], "\n")[1:])
	player2 := util.StrsToInts(strings.Split(blocks[1], "\n")[1:])

	for {
		card1, card2 := player1[0], player2[0]
		player1, player2 = player1[1:], player2[1:]

		if card1 > card2 {
			player1 = append(player1, card1)
			player1 = append(player1, card2)
		} else {
			player2 = append(player2, card2)
			player2 = append(player2, card1)
		}

		if len(player1) == 0 {
			return strconv.Itoa(score(player2)), nil
		}
		if len(player2) == 0 {
			return strconv.Itoa(score(player1)), nil
		}
	}
}

func play(p1, p2 []int) (int, []int) {
	player1 := make([]int, len(p1))
	copy(player1, p1)
	player2 := make([]int, len(p2))
	copy(player2, p2)

	games := make(map[string]bool)
	for {
		hash := hashGame(player1, player2)
		if _, ok := games[hash]; ok {
			return 1, player1
		}
		games[hash] = true

		card1, card2 := player1[0], player2[0]
		player1, player2 = player1[1:], player2[1:]

		var winner int
		if len(player1) >= card1 && len(player2) >= card2 {
			winner, _ = play(player1[:card1], player2[:card2])
		} else {
			if card1 > card2 {
				winner = 1
			} else {
				winner = 2
			}
		}

		if winner == 1 {
			player1 = append(player1, card1)
			player1 = append(player1, card2)
		} else {
			player2 = append(player2, card2)
			player2 = append(player2, card1)
		}

		if len(player1) == 0 {
			return 2, player2
		}
		if len(player2) == 0 {
			return 1, player1
		}
	}
}

func hashGame(p1, p2 []int) string {
	return fmt.Sprintf("%v|%v", p1, p2)
}

func Part2(input string) (string, error) {
	blocks := strings.Split(input, "\n\n")

	player1 := util.StrsToInts(strings.Split(blocks[0], "\n")[1:])
	player2 := util.StrsToInts(strings.Split(blocks[1], "\n")[1:])

	_, deck := play(player1, player2)
	return strconv.Itoa(score(deck)), nil
}
