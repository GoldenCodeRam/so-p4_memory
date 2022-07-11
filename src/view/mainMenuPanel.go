package view

import (
	"so-p4_memory/src/view/lang"

	"github.com/gotk3/gotk3/gtk"
)

type MainMenuPanelListeners interface {
    StartProcessor()
}

type MainMenuPanel struct {
	Box *gtk.Box
}

func CreateMainMenuPanel(listeners MainMenuPanelListeners) *MainMenuPanel {
	panel := MainMenuPanel{
		Box: CreateBox(gtk.ORIENTATION_VERTICAL, MediumMargin),
	}

    panel.Box.Add(CreateButton(lang.START_PROCESSOR, listeners.StartProcessor))

	return &panel
}
