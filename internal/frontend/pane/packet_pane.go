package pane

import (
	"fmt"

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

func (pp *PacketPane) MaxYX() (y int, x int) {
	return pp.w.MaxYX()
}

func (pp *PacketPane) Reflesh(pl []*packet.Packet) {
	pp.w.Clear()
	_, x := pp.MaxYX()

	for i, p := range pl {
		s := fmt.Sprintf("%5d %s", i+1, p.Oneline())
		pp.w.MovePrint(i, 0, s[:x])
	}

	pp.w.Refresh()
}
