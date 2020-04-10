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

func (f *Frontend) Init() (err error) {
	f.stdscr, err = gc.Init()
	if err != nil {
		return err
	}

	// Set as non-blocking read mode.
	f.stdscr.Timeout(0)

	// Turn off buffering to eliminate the enter key.
	gc.CBreak(true)

	// Set the cursor unvisible.
	gc.Cursor(0)

	// Turns off the printing of typed characters.
	gc.Echo(false)

	// Activate the cursor keys.
	f.stdscr.Keypad(true)

	// Enable to scroll
	f.stdscr.ScrollOk(true)

	// Discard all input.
	err = gc.FlushInput()
	if err != nil {
		return err
	}

	return nil
}

func End() {
	gc.End()
}
