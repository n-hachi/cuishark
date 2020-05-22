package pane

import (
	"github.com/n-hachi/go-cuishark/internal/packet"
	gc "github.com/rthornton128/goncurses"
)

type PacketPane struct {
	w *gc.Window
}

func NewPacketPane(w *gc.Window) *PacketPane {
	return &PacketPane{
		w: w,
	}
}

func (pp *PacketPane) Reflesh(pl []*packet.Packet) {
	pp.w.Clear()

	for i, p := range pl {
		pp.w.MovePrint(i, 0, p.Oneline())
	}

	pp.w.Refresh()
}
