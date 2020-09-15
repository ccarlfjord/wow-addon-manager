package gui

import (
	"github.com/ccarlfjord/wow-addon-manager/addon"
	"github.com/gotk3/gotk3/gtk"
)

func installedBoxNew() (*gtk.Box, error) {
	b, err := BoxNew("Installed Addons")
	if err != nil {
		return b, err
	}
	text, err := gtk.ComboBoxTextNew()
	if err != nil {
		return b, err
	}
	addons, err := addon.ReadDir("/home/charles/Games/battlenet/drive_c/Program Files (x86)/World of Warcraft/_classic_/Interface/AddOns")
	for _, addon := range addons {
		text.AppendText(addon.Name)
	}
	b.Add(text)
	return b, nil
}
