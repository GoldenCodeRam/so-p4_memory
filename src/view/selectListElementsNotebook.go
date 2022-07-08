package view

import (
	"so-p4_memory/src/view/lang"

	"github.com/gotk3/gotk3/gtk"
)

type SelectListElementsNotebookListeners interface {
    SelectOptionProcessListPanelListeners
    SelectOptionPartitionListPanelListeners
}

type SelectListElementsNotebook struct {
	Notebook                       *gtk.Notebook
	SelectOptionProcessListPanel   *SelectOptionProcessListPanel
	SelectOptionPartitionListPanel *SelectOptionPartitionListPanel
}

func CreateSelectListElementsNotebook(listeners SelectListElementsNotebookListeners) *SelectListElementsNotebook {
	notebook := SelectListElementsNotebook{
		Notebook: CreateNotebook(),
        SelectOptionProcessListPanel: CreateSelectOptionProcessListPanel(listeners),
        SelectOptionPartitionListPanel: CreateSelectOptionPartitionListPanel(listeners),
	}

	notebook.Notebook.AppendPage(
		notebook.SelectOptionProcessListPanel.ScrolledWindow,
		CreateLabel(lang.PROCESSES),
	)
	notebook.Notebook.AppendPage(
		notebook.SelectOptionPartitionListPanel.Box,
		CreateLabel(lang.PARTITIONS),
	)

	return &notebook
}
