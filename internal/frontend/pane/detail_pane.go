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

func (dp *DetailPane) MaxYX() (y int, x int) {
	return dp.w.MaxYX()
}

func (dp *DetailPane) Reflesh(status *utils.Status) (err error) {
	dp.w.Clear()
	_, x := dp.MaxYX()

	p := status.FocusedPacket()

	height := 0
	for i, l := range p.LayerList() {
		for j, s := range l.Detail() {
			// Check whether current focused pane is myself, and current line.
			flg := (status.DetailIdx() == i && status.PaneIdx() == dp.idx && j == 0)

			// Set underline on
			if flg {
				if err = dp.w.AttrOn(gc.A_UNDERLINE); err != nil {
					return err
				}
			}

			s = utils.CutStringTail(s, x)
			dp.w.MovePrintf(height, 0, "%s", s)

			// Set underline off
			if flg {
				if err = dp.w.AttrOff(gc.A_UNDERLINE); err != nil {
					return err
				}
			}

			height++
		}
	}

	dp.w.Refresh()

	return nil
}
