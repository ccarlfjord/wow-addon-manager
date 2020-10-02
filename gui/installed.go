package gui

import (
	"github.com/ccarlfjord/wow-addon-manager/addon"
	"github.com/ccarlfjord/wow-addon-manager/config"
	"github.com/gotk3/gotk3/gtk"
)

func installedBoxNew(cfg config.Config) (*gtk.Box, error) {
	b, err := BoxNew("Installed Addons")
	if err != nil {
		return b, err
	}

	c := make(chan addon.Addon)
	go func() error {
		addons, err := addon.ReadDir(cfg.GetAddonDir())
		if err != nil {
			return err
		}
		for _, addon := range addons {
			c <- addon
		}
		close(c)
		return nil
	}()

	text, err := gtk.ComboBoxTextNew()
	if err != nil {
		return b, err
	}
	for addon := range c {
		text.AppendText(addon.Name)
	}
	b.Add(text)
	btn, err := newUpdateButton()
	if err != nil {
		return b, err
	}
	b.PackEnd(btn, true, false, 0)
	return b, nil
}
