package view

import (
	"so-p4_memory/src/view/lang"

	"github.com/gotk3/gotk3/gtk"
)

type CreateProcessPanelListeners interface {
}

type CreateProcessPanel struct {
	Box   *gtk.Box
	frame *gtk.Frame

	ProcessNameEntry            *gtk.Entry
	ProcessTimeEntry            *gtk.Entry
	IsProcessBlockedCheckButton *gtk.CheckButton
}

func CreateCreateProcessPanel() *CreateProcessPanel {
	panel := CreateProcessPanel{
		Box:   CreateBox(gtk.ORIENTATION_HORIZONTAL, SmallMargin),
		frame: CreateFrame(lang.CREATE_PROCESS),

        ProcessNameEntry: CreateLabel(lang.NAME),
        ProcessTimeEntry: CreateLabel(lang.TIME),
	}

	grid := CreateGrid()

	processNameLabel := CreateLabel("Nombre")
	processTimeLabel := CreateLabel("Tiempo")

	addProcessButton := CreateButton("Crear")
	addProcessButton.Connect("clicked", func() {
		processFrame.addProcess(listeners)
	})

	grid.Attach(processNameLabel, 0, 0, 1, 1)
	grid.Attach(processTimeLabel, 0, 1, 1, 1)
	grid.Attach(processFrame.ProcessNameEntry, 1, 0, 1, 1)
	grid.Attach(processFrame.ProcessTimeEntry, 1, 1, 1, 1)
	grid.Attach(processFrame.IsProcessBlockedCheckButton, 0, 2, 2, 1)
	grid.Attach(processFrame.IsProcessSuspendedAtRunningCheckButton, 0, 3, 2, 1)
	grid.Attach(processFrame.IsProcessSuspendedAtBlockedCheckButton, 0, 4, 2, 1)
	grid.Attach(processFrame.IsProcessSuspendedAtIOCompletionCheckButton, 0, 5, 2, 1)

	grid.Attach(addProcessButton, 0, 6, 2, 1)

	box.SetCenterWidget(grid)
	processFrame.Frame.Add(box)

	return &panel
}
