package view

import (
	"fmt"
	"so-p4_memory/src/object"
	"so-p4_memory/src/view/lang"

	"github.com/gotk3/gotk3/gtk"
)

type SelectOptionPartitionListPanelListeners interface {
    ShowWithoutPartitionFilters()
    ShowPartitionsList()
}

type SelectOptionPartitionListPanel struct {
	ScrolledWindow *gtk.ScrolledWindow
	box            *gtk.Box
}

func CreateSelectOptionPartitionListPanel(listeners SelectOptionPartitionListPanelListeners) *SelectOptionPartitionListPanel {
	panel := SelectOptionPartitionListPanel{
		ScrolledWindow: CreateScrolledWindow(),
		box:            CreateBox(gtk.ORIENTATION_VERTICAL, MediumMargin),
	}

	panel.box.SetSpacing(int(MediumMargin))
    panel.box.Add(CreateButton(lang.PARTITION_LIST, listeners.ShowPartitionsList))
    panel.box.Add(CreateSeparator(gtk.ORIENTATION_HORIZONTAL))
    panel.box.Add(CreateButton(lang.WITHOUT_FILTERS, listeners.ShowWithoutPartitionFilters))

	panel.ScrolledWindow.Add(panel.box)
	return &panel
}

func (c *SelectOptionPartitionListPanel) AddPartitionSelectionButton(partition *object.Partition, onClick function) {
	c.box.Add(CreateButton(fmt.Sprintf("Filtrar por partición %s", partition.Name), onClick))
    c.box.ShowAll()
}
