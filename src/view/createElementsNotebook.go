package view

import "github.com/gotk3/gotk3/gtk"

type CreateElementsNotebookListeners interface {
}

type CreateElementsNotebook struct {
	Notebook              *gtk.Notebook
	CreateProcessTreeView *gtk.TreeView
}

func CreateCreateElementsNotebook(listeners CreateElementsNotebookListeners) *CreateElementsNotebook {
	notebook := CreateElementsNotebook{
		Notebook: CreateNotebook(),
        CreateProcessPanel: CreateCreateProcessPanel(),
	}

	return &notebook
}
