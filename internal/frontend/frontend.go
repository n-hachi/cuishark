package frontend

import (
	"context"
	"fmt"

	"github.com/n-hachi/go-cuishark/internal/frontend/pane"
	"github.com/n-hachi/go-cuishark/internal/utils"
	gc "github.com/rthornton128/goncurses"
)

type Frontend struct {
	stdscr *gc.Window
	p0     *pane.PacketPane
}

func New() (f *Frontend, err error) {
	f = new(Frontend)

	f.stdscr, err = gc.Init()
	if err != nil {
		return nil, err
	}

	// Set as non-blocking read mode.
	f.stdscr.Timeout(-1)

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
		return nil, err
	}

	// Set sub window.
	height, width := f.stdscr.MaxYX()
	sub_height := height / 3

	sw0 := f.stdscr.Sub(sub_height-2, width, sub_height*0+2, 0)
	f.p0 = pane.NewPacketPane(sw0)

	return f, nil
}

func End() {
	gc.End()
}

func (f *Frontend) Height() (h int) {
	h, _ = f.stdscr.MaxYX()
	return h
}

func (f *Frontend) SubHeight() (sh int) {
	sh = f.Height()
	return sh
}

func (f *Frontend) Width() (w int) {
	_, w = f.stdscr.MaxYX()
	return w
}

func (f *Frontend) Draw() {
	f.stdscr.AttrOn(gc.A_REVERSE)
	s := fmt.Sprintf("%-5s %-13s %-20s %-20s %-6s %-5s %-10s",
		"No.", "Time", "Source", "Destination", "Proto", "Len", "Info")
	for i := len(s); i < f.Width(); i++ {
		s += " "
	}
	f.stdscr.MovePrint(1, 0, s)
	f.stdscr.AttrOff(gc.A_REVERSE)
	f.stdscr.Refresh()
}

func (f *Frontend) OpenChan(ctx context.Context) chan gc.Key {
	ch := make(chan gc.Key, 1)

	// Receive key input and relay to channe.
	go func() {
		defer close(ch)
		for {
			k := f.stdscr.GetChar()
			if k == 0 {
				break
			}
			ch <- k
		}
	}()

	// Watch context variable and close if context.Done is called.
	go func() {
		defer close(ch)
		select {
		case <-ctx.Done():
		}
	}()

	return ch
}

func (f *Frontend) Reflesh(s *utils.Status) {
	f.p0.Reflesh(s)
}
