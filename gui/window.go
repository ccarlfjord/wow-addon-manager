package gui

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func windowNew() *gtk.Window {
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Wow Addon Manager")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	return win
}
