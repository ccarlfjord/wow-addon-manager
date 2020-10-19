// +build gui

package main

import (
	"github.com/ccarlfjord/wow-addon-manager/config"
	"github.com/ccarlfjord/wow-addon-manager/gui"
)

func init() {
	cfg := config.ReadFile("./config.yaml")
	cfg.GameType.Get()
	gui.Init(cfg)
}
