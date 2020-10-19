package gui

import "github.com/gotk3/gotk3/gtk"

func newUpdateButton() (*gtk.Button, error) {
	btn, err := gtk.ButtonNew()
	// btn.SetConne
	label, err := gtk.LabelNew("Update")
	if err != nil {
		return nil, err
	}
	btn.Add(label)

	return btn, nil
}
