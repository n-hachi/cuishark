package pane

import (
	"fmt"
	"log"

	"github.com/n-hachi/cuishark/internal/utils"
	gc "github.com/rthornton128/goncurses"
)

type PacketPane struct {
	w   *gc.Window
	idx int
	top int
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

	// Move window if needed
	pp.slideWindow(status)

	// Window line
	winLine := 0

	//for i, p := range status.PacketList() {
	for i := pp.top; i < pp.bottom(status); i++ {
		// Get print packet
		p := status.PacketList()[i]

		// Check whether current focused pane is myself, and current line.
		flg := (status.PacketIdx() == i && status.PaneIdx() == pp.idx)

		// Set underline on
		if flg {
			if err = pp.w.AttrOn(gc.A_UNDERLINE); err != nil {
				return err
			}
		}
		s := fmt.Sprintf("%5d %s", i+1, p.Oneline())
		s = utils.CutStringTail(s, x)
		pp.w.MovePrint(winLine, 0, s)
		winLine++

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

func (pp *PacketPane) bottom(status *utils.Status) (bottom int) {
	height, _ := pp.MaxYX()
	if height > len(status.PacketList()) {
		height = len(status.PacketList())
	}
	return pp.top + height
}

func (pp *PacketPane) slideWindow(status *utils.Status) {
	log.Printf("status.PacketIdx = %d, pp.top = %d", status.PacketIdx(), pp.top)
	if status.PacketIdx() < pp.top {
		pp.top = status.PacketIdx()
		log.Printf("slide upward")
	}
	if status.PacketIdx() >= pp.bottom(status) {
		slide := status.PacketIdx() - pp.bottom(status) + 1
		pp.top = pp.top + slide
		log.Printf("slide downward")
	}
}
