package view

import (
	"so-p4_memory/src/view/lang"

	"github.com/gotk3/gotk3/gtk"
)

type CreateElementsNotebookListeners interface {
	CreateProcessPanelListeners
    CreatePartitionPanelListeners
}

type CreateElementsNotebook struct {
	Notebook             *gtk.Notebook
	CreateProcessPanel   *CreateProcessPanel
	CreatePartitionPanel *CreatePartitionPanel
}

func CreateCreateElementsNotebook(listeners CreateElementsNotebookListeners) *CreateElementsNotebook {
	notebook := CreateElementsNotebook{
		Notebook:             CreateNotebook(),
		CreateProcessPanel:   CreateCreateProcessPanel(listeners),
		CreatePartitionPanel: CreateCreatePartitionPanel(listeners),
	}

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
