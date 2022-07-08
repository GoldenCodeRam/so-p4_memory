package view

import "github.com/gotk3/gotk3/gtk"

type SelectListElementsNotebookListeners interface {
}

type SelectListElementsNotebook struct {
	Notebook *gtk.Notebook
}

func CreateSelectListElementsNotebook(listeners SelectListElementsNotebookListeners) *SelectListElementsNotebook {
	notebook := SelectListElementsNotebook{
		Notebook: CreateNotebook(),
	}
	return &notebook
}
