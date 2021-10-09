package day19

import (
	"fmt"
	"strconv"
)

type Elf struct {
	ID   int
	Next *Elf
}

func NewElf(id int) *Elf {
	return &Elf{
		ID: id,
	}
}

func (e *Elf) String() string {
	return fmt.Sprintf("#%d next:%d", e.ID, e.Next.ID)
}

func initCircle(count int) *Elf {
	first := NewElf(1)
	next := first
	for i := 1; i < count; i++ {
		next.Next = NewElf(i + 1)
		next = next.Next
	}
	next.Next = first
	return first
}

func Part1(input string) (string, error) {
	elfCount, err := strconv.Atoi(input)
	if err != nil {
		return "", err
	}
	elf := initCircle(elfCount)
	for elf.ID != elf.Next.ID {
		elf.Next = elf.Next.Next
		elf = elf.Next
	}
	return strconv.Itoa(elf.ID), nil
}

func jump(elf *Elf, steps int) {
	current := elf
	next := elf
	for i := 0; i < steps; i++ {
		next = next.Next
	}
	elf = current
	elf.Next = next.Next
	fmt.Printf("dropped %d\n", next.ID)
}

func Part2(input string) (string, error) {
	elfCount, err := strconv.Atoi(input)
	if err != nil {
		return "", err
	}
	elf := initCircle(elfCount)
	//for elf.ID != elf.Next.ID {
	for elfCount > 1 {
		jump(elf, elfCount/2)
		elfCount--
	}
	return strconv.Itoa(elf.ID), nil
}
