package cuishark

import (
	"context"

	"github.com/n-hachi/go-cuishark/internal/frontend"
)

type Cuishark struct {
	f *frontend.Frontend
}

func New() *Cuishark {
	return &Cuishark{
		f: frontend.New(),
	}
}

func (c *Cuishark) Init() (err error) {
	err = c.f.Init()
	if err != nil {
		return err
	}
	return nil
}

func (c *Cuishark) End() {
	frontend.End()
}

func (c *Cuishark) Run(ctx context.Context) (err error) {
	c.f.Draw()

	ch := c.f.OpenChan(ctx)
	for {
		select {
		case k := <-ch:
			c := string(int(k))
			if c == "q" {
				goto L
			}
		}
	}
L:
	return nil
}
