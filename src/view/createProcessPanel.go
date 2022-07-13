package view

import (
	"log"
	"so-p4_memory/src/object"
	"so-p4_memory/src/view/lang"
	"so-p4_memory/src/view/utils"

	"github.com/gotk3/gotk3/gtk"
)

type CreateProcessPanelListeners interface {
	AddProcess(process *object.Process)
	GetPartitionWithName(partitionName string) *object.Partition
}

type CreateProcessPanel struct {
	Box *gtk.Box

	ProcessNameEntry            *gtk.Entry
	ProcessTimeEntry            *gtk.Entry
	ProcessSizeEntry            *gtk.Entry
	IsProcessBlockedCheckButton *gtk.CheckButton
}

func CreateCreateProcessPanel(listeners CreateProcessPanelListeners) *CreateProcessPanel {
	panel := CreateProcessPanel{
		Box: CreateBox(gtk.ORIENTATION_HORIZONTAL, SmallMargin),

		ProcessNameEntry:            CreateEntry(),
		ProcessTimeEntry:            CreateEntry(),
		ProcessSizeEntry:            CreateEntry(),
		IsProcessBlockedCheckButton: CreateCheckButton(lang.IS_BLOCKED),
	}

	grid := CreateGrid()

	processNameLabel := CreateLabel(lang.NAME)
	processTimeLabel := CreateLabel(lang.TIME)
	processSizeLabel := CreateLabel(lang.SIZE)

	addProcessButton := CreateButton(lang.CREATE, func() {
		panel.createProcess(listeners)
	})

	grid.Attach(processNameLabel, 0, 0, 1, 1)
	grid.Attach(processTimeLabel, 0, 1, 1, 1)
	grid.Attach(processSizeLabel, 0, 2, 1, 1)
	grid.Attach(panel.ProcessNameEntry, 1, 0, 1, 1)
	grid.Attach(panel.ProcessTimeEntry, 1, 1, 1, 1)
	grid.Attach(panel.ProcessSizeEntry, 1, 2, 1, 1)
	grid.Attach(panel.IsProcessBlockedCheckButton, 0, 3, 2, 1)
	grid.Attach(addProcessButton, 0, 4, 2, 1)

	panel.Box.SetCenterWidget(grid)

	return &panel
}

func (c *CreateProcessPanel) createProcess(listeners CreateProcessPanelListeners) {
	log.Default().Println("Creating new process...")

	name, err := utils.ExtractTextFromEntry(c.ProcessNameEntry)
	if err != nil {
		utils.ShowErrorDialog(err)
		c.resetFields()
		return
	}

	time, err := utils.ExtractIntFromEntry(c.ProcessTimeEntry)
	if err != nil {
		utils.ShowErrorDialog(err)
		c.resetFields()
		return
	}

	size, err := utils.ExtractIntFromEntry(c.ProcessSizeEntry)
	if err != nil {
		utils.ShowErrorDialog(err)
		c.resetFields()
		return
	}

	listeners.AddProcess(&object.Process{
		Name:          name,
		Time:          time,
		TimeRemaining: time,
		Size:          size,
		IsBlocked:     object.IsBlocked(c.IsProcessBlockedCheckButton.GetActive()),
		State:         object.READY,
	})

	c.resetFields()
}

func (c *CreateProcessPanel) resetFields() {
	c.ProcessNameEntry.SetText("")
	c.ProcessTimeEntry.SetText("")
	c.ProcessSizeEntry.SetText("")
	c.IsProcessBlockedCheckButton.SetActive(false)
}
