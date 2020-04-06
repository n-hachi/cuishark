package cuishark

import "github.com/n-hachi/go-cuishark/internal/frontend"

type Cuishark struct {
	f *frontend.Frontend
}

func New() *Cuishark {
	n := new(Cuishark)
	return n
}
