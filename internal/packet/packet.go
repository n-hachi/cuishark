package packet

import (
	"fmt"
	"strings"

	"github.com/google/gopacket"
)

type Layer struct {
	gl gopacket.Layer
}

func (l *Layer) Oneline() (s string) {
	s = Oneline(l.gl)
	s = strings.TrimLeft(s, "{")
	s = strings.TrimRight(s, "}")
	return s
}

func (l *Layer) Detail() (sl []string) {
	sl = append(sl, l.gl.LayerType().String())
	sl = append(sl, Detail(l.gl)...)
	return sl
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
	gp       gopacket.Packet
}

func NewPacket(gp gopacket.Packet) (p *Packet) {
	p = new(Packet)
	for _, gl := range gp.Layers() {
		// Do not regard 'Payload' layertype as main target layer.
		if gl.LayerType() == gopacket.LayerTypePayload {
			break
		}
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

	// Save original gopacket.Packet
	p.gp = gp

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

func (p *Packet) Layer(index int) (l *Layer) {
	return p.l[index]
}

func (p *Packet) LastLayer() (l *Layer) {
	i := p.Size() - 1
	return p.l[i]
}

func (p *Packet) UnixTime() (t int64) {
	return p.gp.Metadata().Timestamp.Unix()
}

func (p *Packet) Length() (l int) {
	return p.gp.Metadata().CaptureLength
}

func (p *Packet) Oneline() (s string) {
	return fmt.Sprintf("%-13v %-20v %-20v %-6v %-5v %-10v",
		p.UnixTime(),
		p.Src(),
		p.Dst(),
		p.LayerType(),
		p.Length(),
		p.LastLayer().Oneline())
}

func (p *Packet) Detail() (sl []string) {
	for _, l := range p.l {
		sl = append(sl, l.Detail()...)
	}
	return sl
}
