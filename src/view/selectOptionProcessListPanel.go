package view

import (
	"so-p4_memory/src/view/lang"

	"github.com/gotk3/gotk3/gtk"
)

type SelectOptionProcessListPanelListeners interface {
	ShowStateInputList()
	ShowStateReadyList()
	ShowStateRunningList()
	ShowStateBlockedList()
	ShowStateFinishedList()
}

type SelectOptionProcessListPanel struct {
	ScrolledWindow *gtk.ScrolledWindow
}

func CreateSelectOptionProcessListPanel(listeners SelectOptionProcessListPanelListeners) *SelectOptionProcessListPanel {
	panel := SelectOptionProcessListPanel{
		ScrolledWindow: CreateScrolledWindow(),
	}
	box := CreateBox(gtk.ORIENTATION_VERTICAL, MediumMargin)

	statesFrame := CreateFrame("Estados")
	statesFrameBox := CreateBox(gtk.ORIENTATION_VERTICAL, MediumMargin)

	// States frame
	statesFrame.Add(statesFrameBox)
	statesFrameBox.Add(CreateButton(lang.READY_LIST, listeners.ShowStateReadyList))
	statesFrameBox.Add(CreateButton(lang.RUNNING_LIST, listeners.ShowStateRunningList))
	statesFrameBox.Add(CreateButton(lang.BLOCKED_LIST, listeners.ShowStateBlockedList))
	statesFrameBox.Add(CreateButton(lang.FINISHED_LIST, listeners.ShowStateFinishedList))

	panel.ScrolledWindow.Add(box)
	box.SetSpacing(int(SmallMargin))

    box.Add(CreateButton(lang.INPUT_LIST, listeners.ShowStateInputList))

	box.Add(statesFrame)
	return &panel
}
