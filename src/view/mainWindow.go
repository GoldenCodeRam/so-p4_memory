package view

import "github.com/gotk3/gotk3/gtk"

type MainWindow struct {
	Window *gtk.Window
}

type MainWindowListeners interface {
	CreateElementsNotebookListeners
	SelectListElementsNotebookListeners
}

func CreateMainWindow(listeners MainWindowListeners) *MainWindow {
	mainWindow := MainWindow{
		Window: CreateWindow(),
	}

	mainWindow.Window.Add(generatePanedContent(listeners))

	return &mainWindow
}

func generatePanedContent(listeners MainWindowListeners) *gtk.Paned {
	paned := CreatePaned(gtk.ORIENTATION_HORIZONTAL)

	paned.Pack1(generatePack1Content(listeners), true, false)
	paned.Pack2(generatePack2Content(listeners), true, false)

	return paned
}

func generatePack1Content(listeners MainWindowListeners) *gtk.Paned {
	paned := CreatePaned(gtk.ORIENTATION_VERTICAL)

	paned.Pack1(CreateCreateElementsNotebook(listeners).Notebook, true, false)
	paned.Pack2(CreateSelectListElementsNotebook(listeners).Notebook, true, false)

	return paned
}

func generatePack2Content(listeners MainWindowListeners) *gtk.ScrolledWindow {
	scrolledWindow := CreateScrolledWindow()
	return scrolledWindow
}
