package cuishark

import (
	"context"

	"github.com/n-hachi/go-cuishark/internal/frontend"
	"github.com/n-hachi/go-cuishark/internal/handler"
	"github.com/n-hachi/go-cuishark/internal/packet"
	"github.com/n-hachi/go-cuishark/internal/utils"
)

type Cuishark struct {
	f *frontend.Frontend
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
	c.f.Draw()

	keyChan := c.f.OpenChan(ctx)
	pctChan := c.h.OpenChan()
	for {
		select {
		case k := <-keyChan:
			ch := string(int(k))
			if ch == "q" {
				goto L
			} else if ch == "k" {
				c.s.MovePctIdx(utils.Up)
			} else if ch == "j" {
				c.s.MovePctIdx(utils.Down)
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
