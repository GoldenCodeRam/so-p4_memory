package view

import (
	"so-p4_memory/src/object"

	"github.com/gotk3/gotk3/gtk"
)

type MainWindow struct {
	window                     *gtk.Window
	createElementsNotebook     *CreateElementsNotebook
	selectListElementsNotebook *SelectListElementsNotebook
	logsOutputPanel            *LogsOutputPanel
}

type MainWindowListeners interface {
	CreateElementsNotebookListeners
	SelectListElementsNotebookListeners
}

func CreateMainWindow(listeners MainWindowListeners) *MainWindow {
	mainWindow := MainWindow{
		window:                     CreateWindow(),
		createElementsNotebook:     CreateCreateElementsNotebook(listeners),
		selectListElementsNotebook: CreateSelectListElementsNotebook(listeners),
		logsOutputPanel:            CreateProcessesOutputPanel(),
	}

	horizontalPanel := CreatePaned(gtk.ORIENTATION_HORIZONTAL)

	verticalPanel := CreatePaned(gtk.ORIENTATION_VERTICAL)

	verticalPanel.Pack1(mainWindow.createElementsNotebook.Notebook, true, false)
	verticalPanel.Pack2(mainWindow.selectListElementsNotebook.Notebook, true, false)

	horizontalPanel.Pack1(verticalPanel, true, false)
	horizontalPanel.Pack2(mainWindow.logsOutputPanel.Box, true, false)

	mainWindow.window.Add(horizontalPanel)
	return &mainWindow
}

func (m *MainWindow) StartMainWindow() {
	m.window.SetDefaultSize(800, 600)
	m.window.ShowAll()
}

func (m *MainWindow) AddToPartitionList(partition *object.Partition, onButtonClick function) {
	m.selectListElementsNotebook.SelectOptionPartitionListPanel.AddPartitionSelectionButton(
		partition,
		onButtonClick,
	)
}

func (m *MainWindow) ShowProcessesLogTreeView(logTreeView LogTreeView) {
	m.logsOutputPanel.ShowProcessesLogTreeView(logTreeView)
}

func (m *MainWindow) ShowPartitionsLogTreeView() {
	m.logsOutputPanel.ShowPartitionsLogTreeView()
}

func (m *MainWindow) UpdateTreeView(logs []*object.ProcessLog, logTreeView LogTreeView) {
	m.logsOutputPanel.UpdateProcessesLogTreeView(logs, logTreeView)
}

func (m *MainWindow) UpdatePartitionLogsTreeView(logs []*object.Partition) {
	m.logsOutputPanel.UpdatePartitionsLogTreeView(logs)
}

func (m *MainWindow) UpdateProcessesLogTreeViewByPartition(
	logs []*object.ProcessLog,
	logTreeView LogTreeView,
	partition *object.Partition,
) {
	m.logsOutputPanel.UpdateProcessesLogTreeViewByPartition(logs, logTreeView, partition)
}
