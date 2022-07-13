package view

import (
	"so-p4_memory/src/object"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

const (
	PROCESS_NAME = iota
	PROCESS_TIME
	PROCESS_SIZE
	PROCESS_BLOCKED
	PROCESS_STATUS
    PROCESS_TIME_REMAINING
	PARTITION_NAME
	PARTITION_SIZE
)

type ProcessLogTreeView struct {
	TreeView  *gtk.TreeView
	listStore *gtk.ListStore
}

func CreateProcessLogTreeView() *ProcessLogTreeView {
	treeView, listStore := setupProcessLogTreeView()
	return &ProcessLogTreeView{
		TreeView:  treeView,
		listStore: listStore,
	}
}

func (p *ProcessLogTreeView) AddRow(processLog *object.ProcessLog) {
	iter := p.listStore.Append()

	p.listStore.Set(
		iter,
		[]int{
			PROCESS_NAME,
			PROCESS_TIME,
			PROCESS_SIZE,
			PROCESS_BLOCKED,
			PROCESS_STATUS,
            PROCESS_TIME_REMAINING,
			PARTITION_NAME,
			PARTITION_SIZE,
		},
		[]interface{}{
			processLog.Name,
			processLog.Time,
			processLog.Size,
			processLog.IsBlocked,
			processLog.State,
            processLog.TimeRemaining,
			processLog.PartitionName,
			processLog.PartitionSize,
		},
	)
}

func (p *ProcessLogTreeView) Clear() {
	p.listStore.Clear()
}

func (p *ProcessLogTreeView) RemoveRow(process *object.Process) {
	p.listStore.ForEach(func(model *gtk.TreeModel, path *gtk.TreePath, iter *gtk.TreeIter) bool {
		value, _ := model.GetValue(iter, PROCESS_NAME)
		valueString, _ := value.GetString()

		if valueString == process.Name {
			p.listStore.Remove(iter)
		}
		return false
	})
}

func setupProcessLogTreeView() (*gtk.TreeView, *gtk.ListStore) {
	treeView, _ := gtk.TreeViewNew()

	treeView.AppendColumn(createProcessLogColumn("Nombre", PROCESS_NAME))
	treeView.AppendColumn(createProcessLogColumn("Tiempo", PROCESS_TIME))
	treeView.AppendColumn(createProcessLogColumn("Tamaño", PROCESS_SIZE))
	treeView.AppendColumn(createProcessLogColumn("¿Se bloquea?", PROCESS_BLOCKED))
	treeView.AppendColumn(createProcessLogColumn("Estado", PROCESS_STATUS))
	treeView.AppendColumn(createProcessLogColumn("Tiempo restante", PROCESS_TIME_REMAINING))
	treeView.AppendColumn(createProcessLogColumn("Nombre de partición", PARTITION_NAME))
	treeView.AppendColumn(createProcessLogColumn("Tamaño de partición", PARTITION_SIZE))

	listStore, _ := gtk.ListStoreNew(
		glib.TYPE_STRING,
		glib.TYPE_INT,
		glib.TYPE_INT,
		glib.TYPE_STRING,
		glib.TYPE_STRING,
		glib.TYPE_INT,
		glib.TYPE_STRING,
		glib.TYPE_STRING,
	)
	treeView.SetModel(listStore)

	return treeView, listStore
}

func createProcessLogColumn(columnTitle string, id int) *gtk.TreeViewColumn {
	cellRenderer, _ := gtk.CellRendererTextNew()
	column, _ := gtk.TreeViewColumnNewWithAttribute(columnTitle, cellRenderer, "text", id)
	return column
}
