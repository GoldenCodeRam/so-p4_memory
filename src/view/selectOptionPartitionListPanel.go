package view

import "github.com/gotk3/gotk3/gtk"

type SelectOptionPartitionListPanelListeners interface {
}

type SelectOptionPartitionListPanel struct {
	Box *gtk.Box
}

func CreateSelectOptionPartitionListPanel(listeners SelectOptionPartitionListPanelListeners) *SelectOptionPartitionListPanel {
	panel := SelectOptionPartitionListPanel{
		Box: CreateBox(gtk.ORIENTATION_VERTICAL, SmallMargin),
	}
	return &panel
}
