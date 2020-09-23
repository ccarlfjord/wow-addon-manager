package gui

import (
	"github.com/ccarlfjord/wow-addon-manager/addon"
	"github.com/ccarlfjord/wow-addon-manager/config"
	"github.com/gotk3/gotk3/gtk"
)

func installedBoxNew(cfg config.Config) (*gtk.Box, error) {
	b := BoxNew("Installed Addons")
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
