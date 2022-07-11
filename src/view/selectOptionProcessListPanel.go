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

	ShowTransitionReadyToRunningList()
	ShowTransitionRunningToReadyList()
	ShowTransitionRunningToBlockedList()
	ShowTransitionBlockedToReadyList()
	ShowTransitionRunningToFinishedList()
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

	transitionsFrame := CreateFrame("Transiciones")
	transitionsFrameBox := CreateBox(gtk.ORIENTATION_VERTICAL, MediumMargin)

	// States frame
	statesFrame.Add(statesFrameBox)
	statesFrameBox.Add(CreateButton(lang.READY_LIST, listeners.ShowStateReadyList))
	statesFrameBox.Add(CreateButton(lang.RUNNING_LIST, listeners.ShowStateRunningList))
	statesFrameBox.Add(CreateButton(lang.BLOCKED_LIST, listeners.ShowStateBlockedList))
	statesFrameBox.Add(CreateButton(lang.FINISHED_LIST, listeners.ShowStateFinishedList))

	// Transitions frame
	transitionsFrame.Add(transitionsFrameBox)
	transitionsFrameBox.Add(CreateButton(lang.READY_TO_RUNNING, listeners.ShowTransitionReadyToRunningList))
	transitionsFrameBox.Add(CreateButton(lang.RUNNING_TO_READY, listeners.ShowTransitionRunningToReadyList))
	transitionsFrameBox.Add(CreateButton(lang.RUNNING_TO_BLOCKED, listeners.ShowTransitionRunningToBlockedList))
	transitionsFrameBox.Add(CreateButton(lang.BLOCKED_TO_READY, listeners.ShowTransitionBlockedToReadyList))
	transitionsFrameBox.Add(CreateButton(lang.RUNNING_TO_FINISHED, listeners.ShowTransitionRunningToFinishedList))

	panel.ScrolledWindow.Add(box)
	box.SetSpacing(int(SmallMargin))

    box.Add(CreateButton(lang.INPUT_LIST, listeners.ShowStateInputList))

	box.Add(statesFrame)
	box.Add(transitionsFrame)
	return &panel
}
