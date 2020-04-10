package cuishark

import (
	"time"

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

func End() {
	frontend.End()
}

func (c *Cuishark) Run() (err error) {
	c.f.Draw()
	time.Sleep(time.Second * 1)
	return nil
}
