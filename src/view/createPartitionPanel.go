package view

import (
	"log"
	"so-p4_memory/src/object"
	"so-p4_memory/src/view/lang"
	"so-p4_memory/src/view/utils"

	"github.com/gotk3/gotk3/gtk"
)

type CreatePartitionPanelListeners interface {
	CreatePartition(partition *object.Partition)
}

type CreatePartitionPanel struct {
	Box *gtk.Box

	PartitionNameEntry *gtk.Entry
	PartitionSizeEntry *gtk.Entry
}

func CreateCreatePartitionPanel(listeners CreatePartitionPanelListeners) *CreatePartitionPanel {
	panel := CreatePartitionPanel{
		Box: CreateBox(gtk.ORIENTATION_HORIZONTAL, SmallMargin),

		PartitionNameEntry: CreateEntry(),
		PartitionSizeEntry: CreateEntry(),
	}

	grid := CreateGrid()

	partitionNameLabel := CreateLabel(lang.PARTITION_NAME)
	partitionSizeLabel := CreateLabel(lang.PARTITION_SIZE)
	addPartitionButton := CreateButton(lang.CREATE, func() {
		panel.createPartition(listeners)
	})

	grid.Attach(partitionNameLabel, 0, 0, 1, 1)
	grid.Attach(partitionSizeLabel, 0, 1, 1, 1)
	grid.Attach(panel.PartitionNameEntry, 1, 0, 1, 1)
	grid.Attach(panel.PartitionSizeEntry, 1, 1, 1, 1)
	grid.Attach(addPartitionButton, 0, 3, 2, 1)

	panel.Box.SetCenterWidget(grid)
	return &panel
}

func (c *CreatePartitionPanel) createPartition(listeners CreatePartitionPanelListeners) {
	log.Default().Println("Creating new partition...")
	name, err := utils.ExtractTextFromEntry(c.PartitionNameEntry)
	if err != nil {
		utils.ShowErrorDialog(err)
		c.resetFields()
		return
	}
	size, err := utils.ExtractIntFromEntry(c.PartitionSizeEntry)
	if err != nil {
		utils.ShowErrorDialog(err)
		c.resetFields()
		return
	}

	listeners.CreatePartition(&object.Partition{
		Name: name,
		Size: size,
	})
	c.resetFields()
}

func (c *CreatePartitionPanel) resetFields() {
	c.PartitionNameEntry.SetText("")
	c.PartitionSizeEntry.SetText("")
}
