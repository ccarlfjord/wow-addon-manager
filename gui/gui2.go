package gui

// import (
// 	"errors"
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/ccarlfjord/wow-addon-manager/config"
// 	"github.com/gotk3/gotk3/glib"
// 	"github.com/gotk3/gotk3/gtk"
// )

// const appId = "com.github.ccarlfjord.wam.glade"

// func InitNew(cfg *config.Config) {
// 	app, err := gtk.ApplicationNew(appId, glib.APPLICATION_FLAGS_NONE)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	app.Connect("activate", func() {
// 		builder, err := gtk.BuilderNewFromFile("./data/ui/ui.glade")
// 		if err != nil {
// 			log.Fatal(err.Error())
// 		}
// 		obj, err := builder.GetObject("main")
// 		if err != nil {
// 			log.Fatal(err.Error())
// 		}
// 		win, err := isWindow(obj)
// 		if err != nil {
// 			log.Fatal(err.Error())
// 		}
// 		win.ShowAll()
// 		app.AddWindow(win)
// 	})
// 	app.Connect("shutdown", func() {
// 		log.Println("Exiting")
// 	})
// 	app.Run(os.Args)
// }

// func isWindow(obj glib.IObject) (*gtk.Window, error) {

// 	if win, ok := obj.(*gtk.Window); ok {
// 		return win, nil
// 	}
// 	return nil, errors.New("Not a *gtk.Window")
// }

// func setRetail() {
// 	fmt.Println("Setting Retail")
// }
