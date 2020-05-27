package pane

import (
	"fmt"

	"github.com/n-hachi/go-cuishark/internal/utils"
	gc "github.com/rthornton128/goncurses"
)

type PacketPane struct {
	w   *gc.Window
	idx int
}

func NewPacketPane(w *gc.Window, idx int) *PacketPane {
	return &PacketPane{
		w:   w,
		idx: idx,
	}
}

func (pp *PacketPane) MaxYX() (y int, x int) {
	return pp.w.MaxYX()
}

func (pp *PacketPane) Reflesh(status *utils.Status) (err error) {
	pp.w.Clear()
	_, x := pp.MaxYX()

	for i, p := range status.PacketList() {
		// Check whether current focused pane is myself, and current line.
		flg := (status.PacketIdx() == i && status.PaneIdx() == pp.idx)

		// Set underline on
		if flg {
			if err = pp.w.AttrOn(gc.A_UNDERLINE); err != nil {
				return err
			}
		}
		s := fmt.Sprintf("%5d %s", i+1, p.Oneline())
		pp.w.MovePrint(i, 0, s[:x])

		// Set underline off
		if flg {
			if err = pp.w.AttrOff(gc.A_UNDERLINE); err != nil {
				return err
			}
		}
	}

	pp.w.Refresh()
	return nil
}
