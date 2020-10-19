package main

import (
	"github.com/ccarlfjord/wow-addon-manager/addon"
	"github.com/ccarlfjord/wow-addon-manager/config"
	"github.com/ccarlfjord/wow-addon-manager/curseforge"
	"go.uber.org/zap"
)

func main() {
	cfg := config.ReadFile("./config.yaml")
	cfg.GameType.Get()

	logger, err := zap.NewProduction()
	defer logger.Sync()
	if err != nil {
		panic(err.Error())
	}
	cfg.Logger = logger.Sugar()

	log := cfg.Logger

	cfg.GameType = "_retail_"
	addons, err := addon.ReadDir(cfg.GetAddonDir())
	if err != nil {
		log.Error(err)
	}

	for _, a := range addons {
		curseSearchResults, err := curseforge.Search(a.Name, addon.HumanReadableVersion(a.TOC.Interface))
		log.Info(a)
		if err != nil {
			log.Error(err)
		}
		if len(curseSearchResults) > 1 {
			log.Infof("Multiple search results found", "searchResults", curseSearchResults)
		} else {
			log.Infof("Found result for addon", "addon", curseSearchResults)
		}

	}
}
