package pane

import gc "github.com/rthornton128/goncurses"

type PacketPane struct {
	w *gc.Window
}

func NewPacketPane(w *gc.Window) *PacketPane {
	return &PacketPane{
		w: w,
	}
}
