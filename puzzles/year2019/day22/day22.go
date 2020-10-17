package day22

import (
	"strconv"
	"strings"
)

type Deck []int

func NewDeck(n int) Deck {
	d := make(Deck, n)
	for i := 0; i < n; i++ {
		d[i] = i
	}
	return d
}

func (d Deck) DealIntoNewStack() Deck {
	newDeck := make(Deck, 0)
	for i := len(d) - 1; i >= 0; i-- {
		newDeck = append(newDeck, d[i])
	}
	return newDeck
}

func (d Deck) cutTop(n int) Deck {
	newDeck := make(Deck, 0)
	newDeck = append(newDeck, d[n:]...)
	newDeck = append(newDeck, d[:n]...)
	return newDeck
}

func (d Deck) cutBottom(n int) Deck {
	newDeck := make(Deck, 0)
	newDeck = append(newDeck, d[len(d)-n:]...)
	newDeck = append(newDeck, d[:len(d)-n]...)
	return newDeck
}

func (d Deck) CutN(n int) Deck {
	if n < 0 {
		return d.cutBottom(-n)
	}
	return d.cutTop(n)
}

func (d Deck) Increment(n int) Deck {
	newDeck := make(Deck, len(d))
	var pos int
	for _, c := range d {
		newDeck[pos] = c
		pos += n
		if pos > len(d) {
			pos -= len(d)
		}
	}
	return newDeck
}

func Shuffle(instructions []string, d Deck) Deck {
	for _, i := range instructions {
		if i == "deal into new stack" {
			d = d.DealIntoNewStack()
		} else {
			parts := strings.Split(i, " ")
			num, _ := strconv.Atoi(parts[len(parts)-1])
			instr := strings.Join(parts[:len(parts)-1], " ")
			switch instr {
			case "cut":
				d = d.CutN(num)
			case "deal with increment":
				d = d.Increment(num)
			}
		}
	}
	return d
}

func Part1(input string) (string, error) {
	deck := NewDeck(10007)

	deck = Shuffle(strings.Split(input, "\n"), deck)

	for idx, card := range deck {
		if card == 2019 {
			return strconv.Itoa(idx), nil
		}
	}
	return "n/a", nil
}

func Part2(input string) (string, error) {
	return "n/a", nil
}
