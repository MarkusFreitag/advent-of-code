package day16

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MarkusFreitag/advent-of-code/util"
	"github.com/MarkusFreitag/advent-of-code/util/numbers"
)

func asBits(i int64) string {
	return util.StringPadLeft(strconv.FormatInt(i, 2), "0", 4)
}

type Packet struct {
	Version    int
	TypeID     int
	Data       string
	SubPackets []*Packet
}

func (p *Packet) VersionSum() int {
	sum := p.Version
	for _, sub := range p.SubPackets {
		sum += sub.VersionSum()
	}
	return sum
}

func (p *Packet) Evaluate() int {
	if p.TypeID == 4 {
		return p.LiteralValue()
	}

	nums := make([]int, len(p.SubPackets))
	for idx, packet := range p.SubPackets {
		nums[idx] = packet.Evaluate()
	}
	switch p.TypeID {
	case 0:
		return numbers.Sum(nums...)
	case 1:
		return numbers.Multiply(nums...)
	case 2:
		return numbers.Min(nums...)
	case 3:
		return numbers.Max(nums...)
	case 5:
		if len(p.SubPackets) != 2 {
			panic(fmt.Errorf("WARNING, TYPEID=5 ONLY ALLOWS 2 SUBPACKETS"))
		}
		if nums[0] > nums[1] {
			return 1
		}
		return 0
	case 6:
		if len(p.SubPackets) != 2 {
			panic(fmt.Errorf("WARNING, TYPEID=5 ONLY ALLOWS 2 SUBPACKETS"))
		}
		if nums[0] < nums[1] {
			return 1
		}
		return 0
	case 7:
		if len(p.SubPackets) != 2 {
			panic(fmt.Errorf("WARNING, TYPEID=5 ONLY ALLOWS 2 SUBPACKETS"))
		}
		if nums[0] == nums[1] {
			return 1
		}
		return 0
	}
	panic(fmt.Errorf("WARNING, UNKNOWN TYPEID!"))
}

func (p *Packet) LiteralValue() int {
	return util.BinStringToDecInt(p.Data)
}

func ParsePacket(data string) (*Packet, string) {
	var version, id string
	version, data = data[:3], data[3:]
	id, data = data[:3], data[3:]

	p := &Packet{
		Version: util.BinStringToDecInt(version),
		TypeID:  util.BinStringToDecInt(id),
	}

	if p.TypeID == 4 {
		p.Data, data = parseLiteralPacket(data)
	} else {
		p.SubPackets, data = parseOperatorPacket(data)
	}

	return p, data
}

func parseLiteralPacket(data string) (string, string) {
	var str string
	for {
		var group string
		group, data = data[:5], data[5:]
		str += group[1:]
		if strings.HasPrefix(group, "0") {
			break
		}
	}
	return str, data
}

func parseOperatorPacket(data string) ([]*Packet, string) {
	var lengthID int
	lengthID, data = util.BinStringToDecInt(data[:1]), data[1:]

	var length int
	if lengthID == 0 {
		length, data = util.BinStringToDecInt(data[:15]), data[15:]
		var subPacketStr string
		subPacketStr, data = data[:length], data[length:]
		subPackets := make([]*Packet, 0)
		for len(subPacketStr) > 0 {
			var packet *Packet
			packet, subPacketStr = ParsePacket(subPacketStr)
			subPackets = append(subPackets, packet)
		}
		return subPackets, data
	}

	if lengthID == 1 {
		length, data = util.BinStringToDecInt(data[:11]), data[11:]
		subPackets := make([]*Packet, 0)
		for i := 0; i < length; i++ {
			var packet *Packet
			packet, data = ParsePacket(data)
			subPackets = append(subPackets, packet)
		}
		return subPackets, data
	}
	panic(fmt.Errorf("WARNING, UNKNOWN LENGTHID!"))
}

func Part1(input string) (string, error) {
	var binary string
	for _, s := range input {
		i, _ := strconv.ParseInt(string(s), 16, 32)
		binary += asBits(i)
	}

	packet, _ := ParsePacket(binary)

	return strconv.Itoa(packet.VersionSum()), nil
}

func Part2(input string) (string, error) {
	var binary string
	for _, s := range input {
		i, _ := strconv.ParseInt(string(s), 16, 32)
		binary += asBits(i)
	}

	packet, _ := ParsePacket(binary)

	return strconv.Itoa(packet.Evaluate()), nil
}
