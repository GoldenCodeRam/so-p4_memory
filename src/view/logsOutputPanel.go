package view

import (
	"so-p4_memory/src/object"
	"so-p4_memory/src/view/lang"

	"github.com/gotk3/gotk3/gtk"
)

type LogTreeView int

func (l LogTreeView) String() string {
	switch l {
	case LOG_PROCESSES_TREE_VIEW:
		return lang.INPUT_LIST
	case READY_PROCESSES_TREE_VIEW:
		return lang.READY_LIST
	case RUNNING_PROCESSES_TREE_VIEW:
		return lang.RUNNING
	case BLOCKED_PROCESSES_TREE_VIEW:
		return lang.BLOCKED
	case FINISHED_PROCESSES_TREE_VIEW:
		return lang.FINISHED_LIST
	}
	panic("This should not happen!")
}

const (
	LOG_PROCESSES_TREE_VIEW LogTreeView = iota
	READY_PROCESSES_TREE_VIEW
	RUNNING_PROCESSES_TREE_VIEW
	BLOCKED_PROCESSES_TREE_VIEW
	FINISHED_PROCESSES_TREE_VIEW
)

var logTreeViews = [...]LogTreeView{
	LOG_PROCESSES_TREE_VIEW,
	READY_PROCESSES_TREE_VIEW,
	RUNNING_PROCESSES_TREE_VIEW,
	BLOCKED_PROCESSES_TREE_VIEW,
	FINISHED_PROCESSES_TREE_VIEW,
}

type LogsOutputPanel struct {
	Box            *gtk.Box
	headerBar      *gtk.HeaderBar
	scrolledWindow *gtk.ScrolledWindow

	logTreeViews         [len(logTreeViews)]*ProcessLogTreeView
	partitionLogTreeView *PartitionLogTreeView
}

func CreateProcessesOutputPanel() *LogsOutputPanel {
	panel := LogsOutputPanel{
		Box:            CreateBox(gtk.ORIENTATION_VERTICAL, ZeroMargin),
		headerBar:      CreateHeaderBar(),
		scrolledWindow: CreateScrolledWindow(),
	}
	panel.startLogTreeViews()

	panel.headerBar.SetHasSubtitle(false)
	panel.headerBar.SetTitle(lang.INPUT_LIST)
	panel.Box.PackStart(panel.headerBar, false, true, uint(ZeroMargin))
	panel.Box.PackEnd(panel.scrolledWindow, true, true, uint(ZeroMargin))

	panel.scrolledWindow.Add(panel.logTreeViews[LOG_PROCESSES_TREE_VIEW].TreeView)

	return &panel
}

func (c *LogsOutputPanel) ShowProcessesLogTreeView(log LogTreeView) {
	c.headerBar.SetTitle(log.String())
	// Remove current scrolled window children
	c.scrolledWindow.GetChildren().Foreach(func(item interface{}) {
		c.scrolledWindow.Remove(item.(gtk.IWidget))
	})
	c.scrolledWindow.Add(c.logTreeViews[log].TreeView)
	c.Box.ShowAll()
}

func (c *LogsOutputPanel) ShowPartitionsLogTreeView() {
	c.headerBar.SetTitle(lang.PARTITIONS)
	c.scrolledWindow.GetChildren().Foreach(func(item interface{}) {
		c.scrolledWindow.Remove(item.(gtk.IWidget))
	})
	c.scrolledWindow.Add(c.partitionLogTreeView.TreeView)
	c.Box.ShowAll()
}

func (c *LogsOutputPanel) UpdateProcessesLogTreeView(logs []*object.ProcessLog, logTreeView LogTreeView) {
	c.logTreeViews[logTreeView].Clear()
	for _, process := range logs {
		c.logTreeViews[logTreeView].AddRow(process)
	}
}

func (c *LogsOutputPanel) UpdateProcessesLogTreeViewByPartition(
	logs []*object.ProcessLog,
	logTreeView LogTreeView,
	partition *object.Partition,
) {
	c.logTreeViews[logTreeView].Clear()
	for _, process := range logs {
		if process.PartitionName == partition.Name {
			c.logTreeViews[logTreeView].AddRow(process)
		}
	}
}

func (c *LogsOutputPanel) UpdatePartitionsLogTreeView(logs []*object.Partition) {
	c.partitionLogTreeView.Clear()
	for _, partition := range logs {
		c.partitionLogTreeView.AddRow(partition)
	}
}

func (c *LogsOutputPanel) startLogTreeViews() {
	for _, logTreeView := range logTreeViews {
		c.logTreeViews[logTreeView] = CreateProcessLogTreeView()
	}

	c.partitionLogTreeView = CreatePartitionLogTreeView()
}
