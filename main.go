package main

import (
	"github.com/ccarlfjord/wow-addon-manager/config"
	"github.com/ccarlfjord/wow-addon-manager/gui"
)

func main() {
	cfg := config.ReadFile("./config.yaml")
	gui.Init(cfg)
}
