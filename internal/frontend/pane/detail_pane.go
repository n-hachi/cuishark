package pane

import (
	"github.com/n-hachi/go-cuishark/internal/utils"
	gc "github.com/rthornton128/goncurses"
)

type DetailPane struct {
	w *gc.Window
}

func NewDetailPane(w *gc.Window) *DetailPane {
	return &DetailPane{
		w: w,
	}
}

func (dp *DetailPane) Reflesh(status *utils.Status) (err error) {
	dp.w.Clear()

	p := status.FocusedPacket()
	for i, s := range p.Detail() {
		dp.w.MovePrintf(i, 0, "%s", s)
	}

	dp.w.Refresh()

	return nil
}
