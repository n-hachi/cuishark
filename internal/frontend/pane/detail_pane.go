package pane

import (
	"github.com/n-hachi/go-cuishark/internal/utils"
	gc "github.com/rthornton128/goncurses"
)

type DetailPane struct {
	w   *gc.Window
	idx int
}

func NewDetailPane(w *gc.Window, idx int) *DetailPane {
	return &DetailPane{
		w:   w,
		idx: idx,
	}
}

func (dp *DetailPane) Reflesh(status *utils.Status) (err error) {
	dp.w.Clear()

	p := status.FocusedPacket()

	for i, s := range p.Detail() {
		// Check whether current focused pane is myself, and current line.
		flg := (status.DetailIdx() == i && status.PaneIdx() == dp.idx)

		// Set underline on
		if flg {
			if err = dp.w.AttrOn(gc.A_UNDERLINE); err != nil {
				return err
			}
		}

		dp.w.MovePrintf(i, 0, "%s", s)

		// Set underline off
		if flg {
			if err = dp.w.AttrOff(gc.A_UNDERLINE); err != nil {
				return err
			}
		}
	}

	dp.w.Refresh()

	return nil
}
