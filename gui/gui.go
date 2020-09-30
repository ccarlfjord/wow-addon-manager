package gui

import (
	"log"

	"github.com/ccarlfjord/wow-addon-manager/config"
	"github.com/gotk3/gotk3/gtk"
)

func Init(cfg config.Config) {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)

	// Create a new toplevel window, set its title, and connect it to the
	// "destroy" signal to exit the GTK main loop when it is destroyed.
	win := windowNew()

	mainBox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 6)
	if err != nil {
		log.Fatal(err)
	}

	installedBox, err := installedBoxNew(cfg)

	if err != nil {
		log.Fatal(err)
	}

	searchBox, err := BoxNew("Search Addons")
	if err != nil {
		log.Fatal(err)
	}
	searchResultBox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		log.Fatal(err)
	}
	gameVersionBox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		log.Fatal(err)
	}

	searchEntry, err := gtk.SearchEntryNew()
	if err != nil {
		log.Fatal(err)
	}
	searchBox.PackStart(searchEntry, false, false, 0)

	gameVersionListBox, err := gtk.ListBoxNew()
	if err != nil {
		log.Fatal(err)
	}
	gameVersionListBoxLabel, err := gtk.LabelNew("Select Game Version")
	gameVersionBox.Add(gameVersionListBoxLabel)
	classicListBoxRow, err := gtk.ListBoxRowNew()
	if err != nil {
		log.Fatal(err)
	}
	classicListBoxRowLabel, err := gtk.LabelNew("Classic")
	if err != nil {
		log.Fatal(err)
	}
	classicListBoxRow.Add(classicListBoxRowLabel)
	retailListBoxRow, err := gtk.ListBoxRowNew()
	if err != nil {
		log.Fatal(err)
	}
	retailListBoxRowLabel, err := gtk.LabelNew("Retail")
	retailListBoxRow.Add(retailListBoxRowLabel)

	gameVersionListBox.Add(classicListBoxRow)
	gameVersionListBox.Add(retailListBoxRow)
	gameVersionBox.PackStart(gameVersionListBox, false, false, 0)

	var boxes []*gtk.Box
	boxes = append(boxes, gameVersionBox, installedBox, searchBox, searchResultBox)

	for _, box := range boxes {
		mainBox.PackStart(box, true, false, 0)
	}

	win.Add(mainBox)
	// Set the default window size.
	win.SetDefaultSize(800, 600)

	// Recursively show all widgets contained in this window.
	win.ShowAll()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()

}
