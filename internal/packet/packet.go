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
	l        []*Layer
	lt       gopacket.LayerType
	src, dst gopacket.Endpoint
}

func NewPacket(gp gopacket.Packet) (p *Packet) {
	p = new(Packet)
	for _, gl := range gp.Layers() {
		p.l = append(p.l, NewLayer(gl))
		p.lt = gl.LayerType()
	}

	// Get Source and Destination informatinos.
	// Note, do 'not' check transport layer.
	if net := gp.NetworkLayer(); net != nil {
		nf := net.NetworkFlow()
		p.src, p.dst = nf.Endpoints()
	} else if link := gp.LinkLayer(); link != nil {
		lf := link.LinkFlow()
		p.src, p.dst = lf.Endpoints()
	}

	return p
}

func (p *Packet) Size() (s int) {
	return len(p.l)
}

func (p *Packet) LayerType() (lt gopacket.LayerType) {
	return p.lt
}

func (p *Packet) Src() (src gopacket.Endpoint) {
	return p.src
}

func (p *Packet) Dst() (dst gopacket.Endpoint) {
	return p.dst
}
