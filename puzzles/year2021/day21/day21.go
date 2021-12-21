package day21

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
)

var winScore int

type DetermisticDice struct {
	value  int
	rolled int
}

func NewDetermisticDice() *DetermisticDice { return new(DetermisticDice) }

func (d *DetermisticDice) Roll() int {
	d.rolled++
	d.value++
	if d.value == 101 {
		d.value = 1
	}
	return d.value
}

func (d *DetermisticDice) Rolled() int { return d.rolled }

type Player struct {
	ID         string
	Pos, Score int
}

func (p Player) String() string { return fmt.Sprintf("%s: %d=%d", p.ID, p.Pos, p.Score) }

func NewPlayer(str string) *Player {
	fields := strings.Fields(str)
	return &Player{ID: fields[1], Pos: util.ParseInt(fields[len(fields)-1])}
}

func (p *Player) Move(steps int) {
	p.Pos = (p.Pos+steps-1)%10 + 1
	p.Score += p.Pos
}

func (p *Player) Won() bool { return p.Score >= winScore }

func Part1(input string) (string, error) {
	winScore = 1000
	lines := strings.Split(input, "\n")
	players := make([]*Player, len(lines))
	for idx, line := range lines {
		players[idx] = NewPlayer(line)
	}

	dice := NewDetermisticDice()

	var winner *Player
	for {
		for _, player := range players {
			nums := []int{dice.Roll(), dice.Roll(), dice.Roll()}
			player.Move(util.SumInts(nums...))
			if player.Won() {
				winner = player
				break
			}
		}
		if winner != nil {
			break
		}
	}

	for _, player := range players {
		if winner.ID != player.ID {
			return strconv.Itoa(player.Score * dice.Rolled()), nil
		}
	}

	return "n/a", nil
}

type Wins [2]int

func quantumGame(p1, p2 Player, cache map[string]Wins) Wins {
	if p1.Won() {
		return Wins{1, 0}
	}
	if p2.Won() {
		return Wins{0, 1}
	}

	key := fmt.Sprintf("%s|%s", p1, p2)
	if wins, ok := cache[key]; ok {
		return wins
	}

	wins := Wins{}

	for roll1 := 1; roll1 <= 3; roll1++ {
		for roll2 := 1; roll2 <= 3; roll2++ {
			for roll3 := 1; roll3 <= 3; roll3++ {
				steps := roll1 + roll2 + roll3

				newP1 := p1
				newP1.Pos = (p1.Pos+steps-1)%10 + 1
				newP1.Score += newP1.Pos

				newWins := quantumGame(p2, newP1, cache)
				wins[0] += newWins[1]
				wins[1] += newWins[0]
			}
		}
	}

	cache[key] = wins
	return wins
}

func Part2(input string) (string, error) {
	winScore = 21
	lines := strings.Split(input, "\n")
	players := make([]Player, len(lines))
	for idx, line := range lines {
		players[idx] = *NewPlayer(line)
	}

	gameCache := make(map[string]Wins)
	wins := quantumGame(players[0], players[1], gameCache)
	return strconv.Itoa(util.MaxInt(wins[0], wins[1])), nil
}
