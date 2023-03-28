package main

import (
	"strings"
	"github.com/pwiecz/go-fltk"
)



func main() {
	fltk.SetScheme("oxy")

	entries := []fltk.StyleTableEntry {
		fltk.StyleTableEntry{
			fltk.RED,
			fltk.HELVETICA,
			14,
		},
		fltk.StyleTableEntry{
			fltk.BLUE,
			fltk.HELVETICA,
			14,
		},
	}

	buf := fltk.NewTextBuffer()
	sbuf := fltk.NewTextBuffer()

	window := fltk.NewWindow(600, 400)
	disp := fltk.NewTextDisplay(0, 0, 600, 400)
	disp.SetBuffer(buf)
	disp.SetHighlightData(sbuf, entries)
	window.End()
	window.Show()

	buf.Append("Hello\n")
	sbuf.Append(strings.Repeat("A", len("Hello\n")))
	buf.Append("World\n")
	sbuf.Append(strings.Repeat("B", len("World\n")))

	fltk.Run()
}