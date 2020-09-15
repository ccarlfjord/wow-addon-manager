package gui

import (
	"github.com/gotk3/gotk3/gtk"
)

func BoxNew(name string) (*gtk.Box, error) {
	b, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 6)
	if err != nil {
		return b, err
	}

	label, err := gtk.LabelNew(name)
	if err != nil {
		return b, err
	}
	b.Add(label)
	return b, nil
}
