package view

import (
	"so-p4_memory/src/object"
	"so-p4_memory/src/view/lang"

	"github.com/gotk3/gotk3/gtk"
)

type ProcessesOutputPanel struct {
	Box            *gtk.Box
	headerBar      *gtk.HeaderBar
	scrolledWindow *gtk.ScrolledWindow

	InputProcessesTreeView    *ProcessLogTreeView
	ReadyProcessesTreeView    *ProcessLogTreeView
	RunningProcessesTreeView  *ProcessLogTreeView
	BlockedProcessesTreeView  *ProcessLogTreeView
	FinishedProcessesTreeView *ProcessLogTreeView

	ReadyToRunningProcessesTreeView    *ProcessLogTreeView
	RunningToReadyProcessesTreeView    *ProcessLogTreeView
	RunningToBlockedProcessesTreeView  *ProcessLogTreeView
	BlockedToReadyProcessesTreeView    *ProcessLogTreeView
	RunningToFinishedProcessesTreeView *ProcessLogTreeView
}

func CreateProcessesOutputPanel() *ProcessesOutputPanel {
	panel := ProcessesOutputPanel{
		Box:            CreateBox(gtk.ORIENTATION_VERTICAL, ZeroMargin),
		headerBar:      CreateHeaderBar(),
		scrolledWindow: CreateScrolledWindow(),

		InputProcessesTreeView:    CreateProcessLogTreeView(),
		ReadyProcessesTreeView:    CreateProcessLogTreeView(),
		RunningProcessesTreeView:  CreateProcessLogTreeView(),
		BlockedProcessesTreeView:  CreateProcessLogTreeView(),
		FinishedProcessesTreeView: CreateProcessLogTreeView(),

		ReadyToRunningProcessesTreeView:    CreateProcessLogTreeView(),
		RunningToReadyProcessesTreeView:    CreateProcessLogTreeView(),
		RunningToBlockedProcessesTreeView:  CreateProcessLogTreeView(),
		BlockedToReadyProcessesTreeView:    CreateProcessLogTreeView(),
		RunningToFinishedProcessesTreeView: CreateProcessLogTreeView(),
	}

	panel.headerBar.SetHasSubtitle(false)
	panel.headerBar.SetTitle(lang.INPUT_LIST)

	panel.Box.PackStart(panel.headerBar, false, true, uint(ZeroMargin))
	panel.Box.PackEnd(panel.scrolledWindow, true, true, uint(ZeroMargin))

	panel.scrolledWindow.Add(panel.InputProcessesTreeView.TreeView)

	return &panel
}

func (c *ProcessesOutputPanel) ShowInputProcessesTreeView() {
	c.headerBar.SetTitle(lang.INPUT_LIST)
	c.scrolledWindow.GetChildren().Foreach(func(item interface{}) {
		c.scrolledWindow.Remove(item.(gtk.IWidget))
	})
	c.scrolledWindow.Add(c.InputProcessesTreeView.TreeView)
	c.Box.ShowAll()
}

func (c *ProcessesOutputPanel) ShowReadyProcessesTreeView() {
	c.headerBar.SetTitle(lang.READY_LIST)
	c.scrolledWindow.GetChildren().Foreach(func(item interface{}) {
		c.scrolledWindow.Remove(item.(gtk.IWidget))
	})
	c.scrolledWindow.Add(c.ReadyProcessesTreeView.TreeView)
	c.Box.ShowAll()
}

func (c *ProcessesOutputPanel) ShowRunningProcessesTreeView() {
	c.headerBar.SetTitle(lang.RUNNING_LIST)
	c.scrolledWindow.GetChildren().Foreach(func(item interface{}) {
		c.scrolledWindow.Remove(item.(gtk.IWidget))
	})
	c.scrolledWindow.Add(c.RunningProcessesTreeView.TreeView)
	c.Box.ShowAll()
}

func (c *ProcessesOutputPanel) ShowBlockedProcessesTreeView() {
	c.headerBar.SetTitle(lang.BLOCKED_LIST)
	c.scrolledWindow.GetChildren().Foreach(func(item interface{}) {
		c.scrolledWindow.Remove(item.(gtk.IWidget))
	})
	c.scrolledWindow.Add(c.BlockedProcessesTreeView.TreeView)
	c.Box.ShowAll()
}

func (c *ProcessesOutputPanel) ShowFinishedProcessesTreeView() {
	c.headerBar.SetTitle(lang.FINISHED_LIST)
	c.scrolledWindow.GetChildren().Foreach(func(item interface{}) {
		c.scrolledWindow.Remove(item.(gtk.IWidget))
	})
	c.scrolledWindow.Add(c.FinishedProcessesTreeView.TreeView)
	c.Box.ShowAll()
}

func (c *ProcessesOutputPanel) ShowReadyToRunningProcessesTreeView() {
	c.headerBar.SetTitle(lang.READY_TO_RUNNING)
	c.scrolledWindow.GetChildren().Foreach(func(item interface{}) {
		c.scrolledWindow.Remove(item.(gtk.IWidget))
	})
	c.scrolledWindow.Add(c.ReadyToRunningProcessesTreeView.TreeView)
	c.Box.ShowAll()
}

func (c *ProcessesOutputPanel) ShowRunningToReadyProcessesTreeView() {
	c.headerBar.SetTitle(lang.RUNNING_TO_READY)
	c.scrolledWindow.GetChildren().Foreach(func(item interface{}) {
		c.scrolledWindow.Remove(item.(gtk.IWidget))
	})
	c.scrolledWindow.Add(c.RunningToReadyProcessesTreeView.TreeView)
	c.Box.ShowAll()
}

func (c *ProcessesOutputPanel) ShowRunningToBlockedProcessesTreeView() {
	c.headerBar.SetTitle(lang.RUNNING_TO_BLOCKED)
	c.scrolledWindow.GetChildren().Foreach(func(item interface{}) {
		c.scrolledWindow.Remove(item.(gtk.IWidget))
	})
	c.scrolledWindow.Add(c.RunningToBlockedProcessesTreeView.TreeView)
	c.Box.ShowAll()
}

func (c *ProcessesOutputPanel) ShowBlockedToReadyProcessesTreeView() {
	c.headerBar.SetTitle(lang.BLOCKED_TO_READY)
	c.scrolledWindow.GetChildren().Foreach(func(item interface{}) {
		c.scrolledWindow.Remove(item.(gtk.IWidget))
	})
	c.scrolledWindow.Add(c.BlockedToReadyProcessesTreeView.TreeView)
	c.Box.ShowAll()
}

func (c *ProcessesOutputPanel) ShowRunningToFinishedProcessesTreeView() {
	c.headerBar.SetTitle(lang.RUNNING_TO_FINISHED)
	c.scrolledWindow.GetChildren().Foreach(func(item interface{}) {
		c.scrolledWindow.Remove(item.(gtk.IWidget))
	})
	c.scrolledWindow.Add(c.RunningToFinishedProcessesTreeView.TreeView)
	c.Box.ShowAll()
}

func (c *ProcessesOutputPanel) UpdateInputProcessesTreeView(partition *object.Partition, processes []*object.Process) {
	c.InputProcessesTreeView.Clear()
	if partition != nil {
		for _, element := range processes {
			if partition.Number == element.Partition.Number {
				c.InputProcessesTreeView.AddRow(element)
			}
		}
	} else {
		for _, element := range processes {
			c.InputProcessesTreeView.AddRow(element)
		}
	}
}

func (c *ProcessesOutputPanel) UpdateReadyProcessesTreeView(partition *object.Partition, processes []*object.Process) {
	c.ReadyProcessesTreeView.Clear()
	if partition != nil {
		for _, element := range processes {
			if partition.Number == element.Partition.Number {
				c.ReadyProcessesTreeView.AddRow(element)
			}
		}
	} else {
		for _, element := range processes {
			c.ReadyProcessesTreeView.AddRow(element)
		}
	}
}

func (c *ProcessesOutputPanel) UpdateRunningProcessesTreeView(partition *object.Partition, processes []*object.Process) {
	c.RunningProcessesTreeView.Clear()
	if partition != nil {
		for _, element := range processes {
			if partition.Number == element.Partition.Number {
				c.RunningProcessesTreeView.AddRow(element)
			}
		}
	} else {
		for _, element := range processes {
			c.RunningProcessesTreeView.AddRow(element)
		}
	}
}

func (c *ProcessesOutputPanel) UpdateBlockedProcessesTreeView(partition *object.Partition, processes []*object.Process) {
	c.BlockedProcessesTreeView.Clear()
	if partition != nil {
		for _, element := range processes {
			if partition.Number == element.Partition.Number {
				c.BlockedProcessesTreeView.AddRow(element)
			}
		}
	} else {
		for _, element := range processes {
			c.BlockedProcessesTreeView.AddRow(element)
		}
	}
}

func (c *ProcessesOutputPanel) UpdateFinishedProcessesTreeView(partition *object.Partition, processes []*object.Process) {
	c.FinishedProcessesTreeView.Clear()
	if partition != nil {
		for _, element := range processes {
			if partition.Number == element.Partition.Number {
				c.FinishedProcessesTreeView.AddRow(element)
			}
		}
	} else {
		for _, element := range processes {
			c.FinishedProcessesTreeView.AddRow(element)
		}
	}
}

func (c *ProcessesOutputPanel) UpdateReadyToRunningProcessesTreeView(partition *object.Partition, processes []*object.Process) {
	c.ReadyToRunningProcessesTreeView.Clear()
	if partition != nil {
		for _, element := range processes {
			if partition.Number == element.Partition.Number {
				c.ReadyToRunningProcessesTreeView.AddRow(element)
			}
		}
	} else {
		for _, element := range processes {
			c.ReadyToRunningProcessesTreeView.AddRow(element)
		}
	}
}

func (c *ProcessesOutputPanel) UpdateRunningToReadyProcessesTreeView(partition *object.Partition, processes []*object.Process) {
	c.RunningToReadyProcessesTreeView.Clear()
	if partition != nil {
		for _, element := range processes {
			if partition.Number == element.Partition.Number {
				c.RunningToReadyProcessesTreeView.AddRow(element)
			}
		}
	} else {
		for _, element := range processes {
			c.RunningToReadyProcessesTreeView.AddRow(element)
		}
	}
}

func (c *ProcessesOutputPanel) UpdateRunningToBlockedProcessesTreeView(partition *object.Partition, processes []*object.Process) {
	c.RunningToBlockedProcessesTreeView.Clear()
	if partition != nil {
		for _, element := range processes {
			if partition.Number == element.Partition.Number {
				c.RunningToBlockedProcessesTreeView.AddRow(element)
			}
		}
	} else {
		for _, element := range processes {
			c.RunningToBlockedProcessesTreeView.AddRow(element)
		}
	}
}

func (c *ProcessesOutputPanel) UpdateBlockedToReadyProcessesTreeView(partition *object.Partition, processes []*object.Process) {
	c.BlockedToReadyProcessesTreeView.Clear()
	if partition != nil {
		for _, element := range processes {
			if partition.Number == element.Partition.Number {
				c.BlockedToReadyProcessesTreeView.AddRow(element)
			}
		}
	} else {
		for _, element := range processes {
			c.BlockedToReadyProcessesTreeView.AddRow(element)
		}
	}
}

func (c *ProcessesOutputPanel) UpdateRunningToFinishedProcessesTreeView(partition *object.Partition, processes []*object.Process) {
	c.RunningToFinishedProcessesTreeView.Clear()
	if partition != nil {
		for _, element := range processes {
			if partition.Number == element.Partition.Number {
				c.RunningToFinishedProcessesTreeView.AddRow(element)
			}
		}
	} else {
		for _, element := range processes {
			c.RunningToFinishedProcessesTreeView.AddRow(element)
		}
	}
}
