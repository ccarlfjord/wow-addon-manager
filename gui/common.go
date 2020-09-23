package gui

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func BoxNew(name string) *gtk.Box {
	b, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 6)
	if err != nil {
		log.Fatal(err.Error())
	}

	label, err := gtk.LabelNew(name)
	if err != nil {
		log.Fatal(err.Error())
	}
	b.Add(label)
	return b
}
