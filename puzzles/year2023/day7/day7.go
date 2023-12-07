package day7

import (
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/maputil"
	"github.com/MarkusFreitag/advent-of-code/util/sliceutil"
)

var score = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

type handType int

func (ht handType) String() string {
	switch ht {
	case HighCard:
		return "HighCard"
	case OnePair:
		return "OnePair"
	case TwoPair:
		return "TwoPair"
	case ThreeOfKind:
		return "ThreeOfKind"
	case FullHouse:
		return "FullHouse"
	case FourOfKind:
		return "FourOfKind"
	case FiveOfKind:
		return "FiveOfKind"
	}
	return ""
}

const (
	HighCard handType = iota
	OnePair
	TwoPair
	ThreeOfKind
	FullHouse
	FourOfKind
	FiveOfKind
)

func checkHand(hand string) handType {
	cardRates := util.StringTally(hand)
	rates := maputil.Values(cardRates)
	sliceutil.SortDesc(rates)

	switch len(rates) {
	case 1:
		return FiveOfKind
	case 2:
		if rates[0] == 4 {
			return FourOfKind
		}
		if rates[0] == 3 && rates[1] == 2 {
			return FullHouse
		}
	case 3:
		if rates[0] == 3 {
			return ThreeOfKind
		}
		if sliceutil.Count(rates, 2) == 2 {
			return TwoPair
		}
	case 4:
		return OnePair
	case 5:
		return HighCard
	}
	panic("unreachable within checkHand")
}

func compare(h1, h2 Hand) bool {
	if h1.Type > h2.Type {
		return true
	}
	if h1.Type < h2.Type {
		return false
	}
	for i := 0; i < len(h1.Cards); i++ {
		h1Score := score[rune(h1.Cards[i])]
		h2Score := score[rune(h2.Cards[i])]
		if h1Score > h2Score {
			return true
		}
		if h1Score < h2Score {
			return false
		}
	}
	panic("unreachable within compare")
}

type Hand struct {
	Cards string
	Bid   int
	Type  handType
}

type Hands []Hand

func (h Hands) Len() int           { return len(h) }
func (h Hands) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h Hands) Less(i, j int) bool { return compare(h[i], h[j]) }

func parseInput(input string) Hands {
	lines := strings.Split(input, "\n")
	hands := make(Hands, len(lines))
	for idx, line := range lines {
		parts := strings.Fields(line)
		hands[idx] = Hand{
			Cards: parts[0],
			Bid:   util.ParseInt(parts[1]),
			Type:  checkHand(parts[0]),
		}
	}
	return hands
}

func Part1(input string) (string, error) {
	hands := parseInput(input)

	sort.Sort(hands)
	slices.Reverse(hands)

	var total int
	for idx, hand := range hands {
		total += (idx + 1) * hand.Bid
	}
	return strconv.Itoa(total), nil
}

func Part2(input string) (string, error) {
	score['J'] = 1

	hands := parseInput(input)
	for idx, hand := range hands {
		if strings.Contains(hand.Cards, "J") && hand.Cards != "JJJJJ" {
			cardRates := util.StringTally(hand.Cards)
			delete(cardRates, "J")

			var mostCommonCard string
			var maxCount int
			for card, count := range cardRates {
				if count > maxCount {
					maxCount = count
					mostCommonCard = card
				}
			}
			hand.Type = checkHand(strings.ReplaceAll(hand.Cards, "J", mostCommonCard))

			hands[idx] = hand
		}
	}

	sort.Sort(hands)
	slices.Reverse(hands)

	var total int
	for idx, hand := range hands {
		total += (idx + 1) * hand.Bid
	}
	return strconv.Itoa(total), nil
}
