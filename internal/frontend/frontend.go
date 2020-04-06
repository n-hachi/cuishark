package frontend

import (
	gc "github.com/rthornton128/goncurses"
)

type Frontend struct {
	stdscr *gc.Window
}

func New() *Frontend {
	f := new(Frontend)
	return f
}
