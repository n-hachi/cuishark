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
		bp.w.MovePrintf(i, 0, "%s", s)
	}

	bp.w.Refresh()

	return nil
}
