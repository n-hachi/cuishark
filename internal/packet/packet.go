package packet

import "github.com/google/gopacket"

type Layer struct {
	gl gopacket.Layer
}

func NewLayer(gl gopacket.Layer) (l *Layer) {
	return &Layer{
		gl: gl,
	}
}

type Packet struct {
	l []*Layer
}

func NewPacket(gp gopacket.Packet) (p *Packet) {
	p = new(Packet)
	for _, gl := range gp.Layers() {
		p.l = append(p.l, NewLayer(gl))
	}
	return p
}

func (p *Packet) Size() (s int) {
	return len(p.l)
}
