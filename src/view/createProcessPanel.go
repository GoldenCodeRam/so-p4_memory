package view

import (
	"log"
	"so-p4_memory/src/object"
	"so-p4_memory/src/view/lang"
	"so-p4_memory/src/view/utils"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type CreateProcessPanelListeners interface {
	AddProcess(process *object.Process)
}

type CreateProcessPanel struct {
	Box *gtk.Box

	ProcessNameEntry            *gtk.Entry
	ProcessTimeEntry            *gtk.Entry
	IsProcessBlockedCheckButton *gtk.CheckButton
}

func CreateCreateProcessPanel(listeners CreateProcessPanelListeners) *CreateProcessPanel {
	panel := CreateProcessPanel{
		Box: CreateBox(gtk.ORIENTATION_HORIZONTAL, SmallMargin),

		ProcessNameEntry:            CreateEntry(),
		ProcessTimeEntry:            CreateEntry(),
		IsProcessBlockedCheckButton: CreateCheckButton(lang.IS_BLOCKED),
	}

	grid := CreateGrid()

	processNameLabel := CreateLabel(lang.NAME)
	processTimeLabel := CreateLabel(lang.TIME)

	addProcessButton := CreateButton(lang.CREATE, func() {
        panel.createProcess(listeners)
    })

    // TODO
    test, _ := gtk.ListStoreNew(glib.TYPE_BOOLEAN, glib.TYPE_STRING)
    test.Set(test.Append(), []int{0, 1}, []interface{}{true, "test"})
    combo, _ := gtk.ComboBoxNewWithModelAndEntry(test)
    test.Clear()
    combo.SetEntryTextColumn(1)
    combo.SetActive(1)
    // TODO

	grid.Attach(processNameLabel, 0, 0, 1, 1)
	grid.Attach(processTimeLabel, 0, 1, 1, 1)
	grid.Attach(panel.ProcessNameEntry, 1, 0, 1, 1)
	grid.Attach(panel.ProcessTimeEntry, 1, 1, 1, 1)
	grid.Attach(panel.IsProcessBlockedCheckButton, 0, 2, 2, 1)
	grid.Attach(addProcessButton, 0, 3, 2, 1)
	grid.Attach(combo, 0, 4, 2, 1)

	panel.Box.SetCenterWidget(grid)

	return &panel
}

func (c *CreateProcessPanel) createProcess(listeners CreateProcessPanelListeners) {
	log.Default().Println("Creating new process...")

	name, err := utils.ExtractTextFromEntry(c.ProcessTimeEntry)
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

	listeners.AddProcess(&object.Process{
		Name:          name,
		Time:          time,
		TimeRemaining: time,
		IsBlocked:     c.IsProcessBlockedCheckButton.GetActive(),
		State:         object.READY,
	})

	c.resetFields()
}

func (c *CreateProcessPanel) resetFields() {
	c.ProcessNameEntry.SetText("")
	c.ProcessTimeEntry.SetText("")
    c.IsProcessBlockedCheckButton.SetActive(false)
}
