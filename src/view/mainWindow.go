package view

import "github.com/gotk3/gotk3/gtk"

type MainWindow struct {
	Window                     *gtk.Window
	CreateElementsNotebook     *CreateElementsNotebook
	SelectListElementsNotebook *SelectListElementsNotebook
	ProcessesOutputPanel       *ProcessesOutputPanel
}

type MainWindowListeners interface {
	CreateElementsNotebookListeners
	SelectListElementsNotebookListeners
}

func CreateMainWindow(listeners MainWindowListeners) *MainWindow {
	mainWindow := MainWindow{
		Window:                     CreateWindow(),
		CreateElementsNotebook:     CreateCreateElementsNotebook(listeners),
		SelectListElementsNotebook: CreateSelectListElementsNotebook(listeners),
		ProcessesOutputPanel:       CreateProcessesOutputPanel(),
	}

	horizontalPanel := CreatePaned(gtk.ORIENTATION_HORIZONTAL)

	verticalPanel := CreatePaned(gtk.ORIENTATION_VERTICAL)

	verticalPanel.Pack1(mainWindow.CreateElementsNotebook.Notebook, true, false)
	verticalPanel.Pack2(mainWindow.SelectListElementsNotebook.Notebook, true, false)

	horizontalPanel.Pack1(verticalPanel, true, false)
	horizontalPanel.Pack2(mainWindow.ProcessesOutputPanel.Box, true, false)

	mainWindow.Window.Add(horizontalPanel)
	return &mainWindow
}
