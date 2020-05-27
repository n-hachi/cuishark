package pane

import (
	"github.com/n-hachi/go-cuishark/internal/utils"
	gc "github.com/rthornton128/goncurses"
)

type BinaryPane struct {
	w   *gc.Window
	idx int
}

func NewBinaryPane(w *gc.Window, idx int) *BinaryPane {
	return &BinaryPane{
		w:   w,
		idx: idx,
	}
}

func (bp *BinaryPane) Reflesh(status *utils.Status) (err error) {
	bp.w.Clear()

	p := status.FocusedPacket()

	for i, s := range p.Binary() {
		// Check whether current focused pane is myself, and current line.
		flg := (status.BinaryIdx() == i && status.PaneIdx() == bp.idx)

		// Set underline on
		if flg {
			if err = bp.w.AttrOn(gc.A_UNDERLINE); err != nil {
				return err
			}
		}

		bp.w.MovePrintf(i, 0, "%s", s)

		// Set underline off
		if flg {
			if err = bp.w.AttrOff(gc.A_UNDERLINE); err != nil {
				return err
			}
		}
	}

	bp.w.Refresh()

	return nil
}
