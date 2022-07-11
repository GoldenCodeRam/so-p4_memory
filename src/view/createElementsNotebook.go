package view

import (
	"so-p4_memory/src/view/lang"

	"github.com/gotk3/gotk3/gtk"
)

type CreateElementsNotebookListeners interface {
	CreateProcessPanelListeners
	CreatePartitionPanelListeners
	MainMenuPanelListeners
}

type CreateElementsNotebook struct {
	Notebook             *gtk.Notebook
	CreateProcessPanel   *CreateProcessPanel
	CreatePartitionPanel *CreatePartitionPanel
	MainMenuPanel        *MainMenuPanel
}

func CreateCreateElementsNotebook(listeners CreateElementsNotebookListeners) *CreateElementsNotebook {
	notebook := CreateElementsNotebook{
		Notebook:             CreateNotebook(),
		CreateProcessPanel:   CreateCreateProcessPanel(listeners),
		CreatePartitionPanel: CreateCreatePartitionPanel(listeners),
		MainMenuPanel:        CreateMainMenuPanel(listeners),
	}

	notebook.Notebook.AppendPage(
		notebook.MainMenuPanel.Box,
		CreateLabel(lang.MAIN_MENU),
	)
	notebook.Notebook.AppendPage(
		notebook.CreateProcessPanel.Box,
		CreateLabel(lang.CREATE_PROCESS),
	)
	notebook.Notebook.AppendPage(
		notebook.CreatePartitionPanel.Box,
		CreateLabel(lang.CREATE_PARTITION),
	)

	return &notebook
}
