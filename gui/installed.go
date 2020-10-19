package gui

import (
	"fmt"

	"github.com/ccarlfjord/wow-addon-manager/addon"
	"github.com/gotk3/gotk3/gtk"
)

func (g *gui) installedBoxNew() (*gtk.Box, error) {
	b, err := BoxNew("Update Addon")
	if err != nil {
		return b, err
	}

	c := make(chan addon.Addon)
	go func() error {
		addons, err := addon.ReadDir(g.cfg.GetAddonDir())
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
		text.Insert(-1, addon.Name, addon.Name)
	}
	text.Insert(0, "all", "All")
	b.Add(text)
	btn, err := newUpdateButton()
	if err != nil {
		return b, err
	}
	btn.Connect("clicked", func() {
		fmt.Println(text.GetActiveID())
	})
	b.PackStart(btn, true, false, 3)
	return b, nil
}
