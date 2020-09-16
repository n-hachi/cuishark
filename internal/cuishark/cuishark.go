package cuishark

import (
	"context"
	"log"
	"runtime"
	"unsafe"

	"github.com/n-hachi/cuishark/internal/frontend"
	"github.com/n-hachi/cuishark/internal/handler"
	"github.com/n-hachi/cuishark/internal/packet"
	"github.com/n-hachi/cuishark/internal/utils"

	gc "github.com/rthornton128/goncurses"
)

const (
	Packet = iota
	Detail
	Binary
)

type Cuishark struct {
	f *frontend.Frontend
	p unsafe.Pointer
	h *handler.PcapHandler
	s *utils.Status
}

func New(path string) (c *Cuishark, err error) {
	c = &Cuishark{}
	c.f, err = frontend.New()
	if err != nil {
		return nil, err
	}

	c.h, err = handler.NewPcapHandler(path)
	if err != nil {
		return nil, err
	}

	c.s, err = utils.NewStatus()
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Cuishark) End() {
	frontend.End()
}

func (c *Cuishark) Run(ctx context.Context) (err error) {
	// logging
	pc, file, _, _ := runtime.Caller(0)
	fname := runtime.FuncForPC(pc).Name()
	log.Printf("[start] file=%s, func=%s\n", file, fname)
	defer func() {
		log.Printf("[end] file=%s, func=%s\n", file, fname)
	}()

	c.f.Draw()

	keyChan := c.f.OpenChan(ctx)
	pctChan := c.h.OpenChan()
	for {
		select {
		case k := <-keyChan:
			ch := string(rune(int(k)))
			if ch == "q" {
				goto L
			} else if ch == "k" || ch == "j" {
				var d utils.Direction
				if ch == "j" {
					d = utils.Down
				} else {
					d = utils.Up
				}

				switch c.s.PaneIdx() {
				case Packet:
					c.s.MovePacketIdx(d)
				case Detail:
					c.s.MoveDetailIdx(d)
				case Binary:
					c.s.MoveBinaryIdx(d)
				default:
					panic("Code error")
				}
			} else if k == gc.KEY_TAB {
				c.s.RotatePaneIdx()
			}

		case gp, ok := <-pctChan:
			if !ok {
				pctChan = nil
			} else {
				p := packet.NewPacket(gp)
				c.s.AppendPacket(p)
			}
		}

		// Update pane
		c.f.Reflesh(c.s)
	}
L:
	return nil
}
