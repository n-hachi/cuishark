package utils

import "github.com/n-hachi/go-cuishark/internal/packet"

type Direction int

const (
	Up = iota
	Down
)

type Status struct {
	pl        []*packet.Packet
	packetIdx int
	detailIdx int
	binaryIdx int
	paneIdx   int
}

func NewStatus() (s *Status, err error) {
	s = &Status{}
	s.packetIdx = 0
	s.detailIdx = 0
	s.binaryIdx = 0
	s.paneIdx = 0

	return s, nil
}

func (s *Status) AppendPacket(p *packet.Packet) {
	s.pl = append(s.pl, p)
}

func (s *Status) PacketList() (pl []*packet.Packet) {
	return s.pl
}

func (s *Status) PacketCount() (size int) {
	return len(s.PacketList())
}

func (s *Status) MovePacketIdx(d Direction) {
	switch d {
	case Up:
		if s.packetIdx > 0 {
			s.packetIdx--
		}
	case Down:
		if s.packetIdx < (s.PacketCount() - 1) {
			s.packetIdx++
		}
	default:
		panic("Code error, in MovePctIdx")
	}
}

func (s *Status) MoveDetailIdx(d Direction) {
	switch d {
	case Up:
		if s.detailIdx > 0 {
			s.detailIdx--
		}
	case Down:
		p := s.FocusedPacket()
		if s.detailIdx < (p.Size() - 1) {
			s.detailIdx++
		}
	default:
		panic("Code error, in MoveDetailIdx")
	}
}

func (s *Status) MoveBinaryIdx(d Direction) {
	switch d {
	case Up:
		if s.binaryIdx > 0 {
			s.binaryIdx--
		}
	case Down:
		p := s.FocusedPacket()
		if s.binaryIdx < (len(p.Binary()) - 1) {
			s.binaryIdx++
		}
	default:
		panic("Code error, in MoveBinaryIdx")
	}
}

func (s *Status) PaneIdx() (paneIdx int) {
	return s.paneIdx
}

func (s *Status) PacketIdx() (packetIdx int) {
	return s.packetIdx
}

func (s *Status) DetailIdx() (detailIdx int) {
	return s.detailIdx
}

func (s *Status) BinaryIdx() (binaryIdx int) {
	return s.binaryIdx
}

func (s *Status) RotatePaneIdx() (newIdx int) {
	s.paneIdx = (s.paneIdx + 1) % 3
	return s.paneIdx
}

func (s *Status) FocusedPacket() (p *packet.Packet) {
	return s.PacketList()[s.PacketIdx()]
}
