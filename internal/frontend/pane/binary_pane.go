package pane

import (
	"log"

	"github.com/n-hachi/cuishark/internal/utils"
	gc "github.com/rthornton128/goncurses"
)

type BinaryPane struct {
	w   *gc.Window
	idx int
	top int
}

func NewBinaryPane(w *gc.Window, idx int) *BinaryPane {
	return &BinaryPane{
		w:   w,
		idx: idx,
	}
}

func (bp *BinaryPane) MaxYX() (y int, x int) {
	return bp.w.MaxYX()
}

func (bp *BinaryPane) Reflesh(status *utils.Status) (err error) {
	bp.w.Clear()
	_, x := bp.MaxYX()

	// Move window if needed
	bp.slideWindow(status)

	p := status.FocusedPacket()

	// Window line
	winLine := 0

	for i := bp.top; i < bp.bottom(status); i++ {
		s := p.Binary()[i]

		// Check whether current focused pane is myself, and current line.
		flg := (status.BinaryIdx() == i && status.PaneIdx() == bp.idx)

		// Set underline on
		if flg {
			if err = bp.w.AttrOn(gc.A_UNDERLINE); err != nil {
				return err
			}
		}

		s = utils.CutStringTail(s, x)
		bp.w.MovePrintf(winLine, 0, "%s", s)
		winLine++

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

func (bp *BinaryPane) bottom(status *utils.Status) (bottom int) {
	height, _ := bp.MaxYX()
	p := status.FocusedPacket()
	if height > len(p.Binary()) {
		height = len(p.Binary())
	}
	return bp.top + height
}

func (bp *BinaryPane) slideWindow(status *utils.Status) {
	log.Printf("status.BinaryIdx = %d, bp.top = %d", status.BinaryIdx(), bp.top)
	if status.BinaryIdx() < bp.top {
		bp.top = status.BinaryIdx()
		log.Printf("slide upward")
	}
	if status.BinaryIdx() >= bp.bottom(status) {
		slide := status.BinaryIdx() - bp.bottom(status) + 1
		bp.top = bp.top + slide
		log.Printf("slide downward")
	}
}
