package main

import (
	"github.com/ccarlfjord/wow-addon-manager/config"
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

	// log := cfg.Logger

	// state, err := state.Init("data.sqlite3")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// addons, err := addon.ReadDir(cfg.GetAddonDir())
	// if err != nil {
	// 	log.Error(err)
	// }

	// for _, a := range addons {
	// 	curseSearchResults, err := curseforge.Search(a.Name, addon.HumanReadableVersion(a.TOC.Interface))
	// 	log.Info(a)
	// 	if err != nil {
	// 		log.Error(err)
	// 	}
	// 	if len(curseSearchResults) > 1 {
	// 		log.Infof("Multiple search results found", "searchResults", curseSearchResults)
	// 	} else {
	// 		log.Infof("Found result for addon", "addon", curseSearchResults)
	// 	}

	// }
}
