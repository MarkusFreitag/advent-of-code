package day23

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/MarkusFreitag/advent-of-code/pkg/intcode"
)

type Packet struct {
	Addr, X, Y int64
}

type Computer struct {
	sync.Mutex
	program intcode.IntCode
	Addr    int64
	Idle    bool
	Queue   []Packet
	InChan  chan int64
	OutChan chan intcode.Message
}

func NewComputer(addr int64, prog intcode.IntCode) *Computer {
	return &Computer{
		program: prog,
		Addr:    addr,
		Queue:   make([]Packet, 0),
		InChan:  make(chan int64),
		OutChan: make(chan intcode.Message, 10),
	}
}

func (c *Computer) Run(nw chan Packet, debug bool) {
	if debug {
		fmt.Printf("#[%d] run program\n", c.Addr)
	}
	go intcode.RunSync(c.program, c.InChan, c.OutChan)
	running := true
	first := true
	for running {
		msg := <-c.OutChan
		switch msg.Type {
		case intcode.MessageWantsInput:
			if first {
				c.InChan <- int64(c.Addr)
				first = false
			} else {
				p := c.PopQueue()
				if p == nil {
					c.Idle = true
					c.InChan <- -1
				} else {
					c.Idle = false
					if debug {
						fmt.Printf("#[%d] input %#v\n", c.Addr, p)
					}
					c.InChan <- p.X
					c.InChan <- p.Y
				}
			}
		case intcode.MessageOutput:
			var p Packet
			p.Addr = msg.Value

			msg = <-c.OutChan
			p.X = msg.Value

			msg = <-c.OutChan
			p.Y = msg.Value

			if debug {
				fmt.Printf("#[%d] output %#v\n", c.Addr, p)
			}
			nw <- p
		case intcode.MessageHalt:
			running = false
		}
	}
}

func (c *Computer) AddPacket(p Packet) {
	c.Lock()
	defer c.Unlock()
	c.Queue = append(c.Queue, p)
}

func (c *Computer) PopQueue() *Packet {
	c.Lock()
	defer c.Unlock()
	if len(c.Queue) == 0 {
		return nil
	}
	var p Packet
	p, c.Queue = c.Queue[0], c.Queue[1:]
	return &p
}

type Network []*Computer

func (n Network) Idle() bool {
	for _, c := range n {
		if !c.Idle {
			return false
		}
	}
	return true
}

func (n Network) ByAddr(addr int64) *Computer {
	for _, c := range n {
		if c.Addr == addr {
			return c
		}
	}
	return nil
}

func Part1(input string) (string, error) {
	icode, err := intcode.New(input)
	if err != nil {
		return "", err
	}
	net := make(chan Packet, 100)
	network := make(Network, 50)
	for idx := range network {
		com := NewComputer(int64(idx), icode)
		network[idx] = com
		go com.Run(net, false)
	}

	for packet := range net {
		if packet.Addr == int64(255) {
			return strconv.Itoa(int(packet.Y)), nil
		}
		network[int(packet.Addr)].AddPacket(packet)
	}

	return "n/a", nil
}

func Part2(input string) (string, error) {
	return "n/a", nil
}
