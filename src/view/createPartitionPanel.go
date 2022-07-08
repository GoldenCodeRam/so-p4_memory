package view

import (
	"log"
	"so-p4_memory/src/object"
	"so-p4_memory/src/view/lang"

	"github.com/gotk3/gotk3/gtk"
)

type CreatePartitionPanelListeners interface {
	CreatePartition(partition *object.Partition)
}

type CreatePartitionPanel struct {
	Box *gtk.Box

	PartitionNumberEntry *gtk.Entry
	PartitionSizeEntry   *gtk.Entry
}

func CreateCreatePartitionPanel(listeners CreatePartitionPanelListeners) *CreatePartitionPanel {
	panel := CreatePartitionPanel{
		Box: CreateBox(gtk.ORIENTATION_HORIZONTAL, SmallMargin),

		PartitionNumberEntry: CreateEntry(),
		PartitionSizeEntry:   CreateEntry(),
	}

	grid := CreateGrid()

	partitionNumberLabel := CreateLabel(lang.PARTITION_NUMBER)
	partitionSizeLabel := CreateLabel(lang.PARTITION_SIZE)
	addPartitionButton := CreateButton(lang.CREATE, func() {
		panel.createPartition(listeners)
	})

	grid.Attach(partitionNumberLabel, 0, 0, 1, 1)
	grid.Attach(partitionSizeLabel, 0, 1, 1, 1)
	grid.Attach(panel.PartitionNumberEntry, 1, 0, 1, 1)
	grid.Attach(panel.PartitionSizeEntry, 1, 1, 1, 1)
	grid.Attach(addPartitionButton, 0, 3, 2, 1)

	panel.Box.SetCenterWidget(grid)
	return &panel
}

func (c *CreatePartitionPanel) createPartition(listeners CreatePartitionPanelListeners) {
	log.Default().Println("Creating new partition...")

}
