package utils

import "github.com/n-hachi/go-cuishark/internal/packet"

type Direction int

const (
	Up = iota
	Down
)

type Status struct {
	pl     []*packet.Packet
	pctIdx int
}

func NewStatus() (s *Status, err error) {
	s = &Status{}
	s.pctIdx = 0

	return s, nil
}

func (s *Status) AppendPacket(p *packet.Packet) {
	s.pl = append(s.pl, p)
}

func (s *Status) PacketList() (pl []*packet.Packet) {
	return s.pl
}

func (s *Status) MovePctIdx(d Direction) {
	switch d {
	case Up:
		if s.pctIdx > 0 {
			s.pctIdx--
		}
	case Down:
		if s.pctIdx < (len(s.pl) - 1) {
			s.pctIdx++
		}
	}
}

func (s *Status) PctIdx() (pctIdx int) {
	return s.pctIdx
}

func (s *Status) FocusedPacket() (p *packet.Packet) {
	return s.PacketList()[s.PctIdx()]
}
