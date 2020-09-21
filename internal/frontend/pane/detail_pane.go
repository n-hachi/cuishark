package pane

import (
	"github.com/n-hachi/cuishark/internal/utils"
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
		a := l.Abstract()

		flg := (status.DetailIdx() == i && status.PaneIdx() == dp.idx)
		if flg {
			if err = dp.w.AttrOn(gc.A_UNDERLINE); err != nil {
				return err
			}
		}

		dp.w.MovePrint(height, 0, a)
		height++

		if flg {
			if err = dp.w.AttrOff(gc.A_UNDERLINE); err != nil {
				return err
			}
		}

		if status.IsShowLayer(l.LayerType()) {
			for _, s := range l.Detail() {
				s = utils.CutStringTail(s, x)
				dp.w.MovePrintf(height, 0, "%s", s)
				height++
			}
		}
	}

	dp.w.Refresh()

	return nil
}
