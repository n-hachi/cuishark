package cuishark

import (
	"context"

	"github.com/n-hachi/go-cuishark/internal/frontend"
	"github.com/n-hachi/go-cuishark/internal/handler"
	"github.com/n-hachi/go-cuishark/internal/packet"
)

type Cuishark struct {
	f  *frontend.Frontend
	h  *handler.PcapHandler
	pl []*packet.Packet
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

	return c, nil
}

func (c *Cuishark) End() {
	frontend.End()
}

func (c *Cuishark) PacketList() (pl []*packet.Packet) {
	return c.pl
}

func (c *Cuishark) Run(ctx context.Context) (err error) {
	c.f.Draw()

	keyChan := c.f.OpenChan(ctx)
	pctChan := c.h.OpenChan()
	for {
		select {
		case k := <-keyChan:
			c := string(int(k))
			if c == "q" {
				goto L
			}
		case gp, ok := <-pctChan:
			if !ok {
				pctChan = nil
			} else {
				p := packet.NewPacket(gp)
				c.pl = append(c.pl, p)
			}
		}

		// Update pane
		c.f.Reflesh(c.pl)
	}
L:
	return nil
}
