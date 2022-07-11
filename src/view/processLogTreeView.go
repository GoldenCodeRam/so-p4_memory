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
    PARTITION_NUMBER
    PARTITION_SIZE
)

type ProcessLogTreeView struct {
	TreeView  *gtk.TreeView
	listStore *gtk.ListStore
}

func CreateProcessLogTreeView() *ProcessLogTreeView {
	treeView, listStore := setupTreeView()
	return &ProcessLogTreeView{
		TreeView:  treeView,
		listStore: listStore,
	}
}

func (p *ProcessLogTreeView) AddRow(process *object.Process) {
	iter := p.listStore.Append()
	p.listStore.Set(
		iter,
		[]int{
            PROCESS_NAME,
            PROCESS_TIME,
            PROCESS_SIZE,
            PROCESS_BLOCKED,
            PROCESS_STATUS,
            PARTITION_NUMBER,
            PARTITION_SIZE,
        },
		[]interface{}{
            process.Name,
            process.Time,
            process.Size,
            process.IsBlocked.String(),
            process.State.String(),
            process.Partition.Number,
            process.Partition.Size,
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

func setupTreeView() (*gtk.TreeView, *gtk.ListStore) {
	treeView, _ := gtk.TreeViewNew()

	treeView.AppendColumn(createColumn("Nombre", PROCESS_NAME))
	treeView.AppendColumn(createColumn("Tiempo", PROCESS_TIME))
	treeView.AppendColumn(createColumn("Tamaño", PROCESS_SIZE))
	treeView.AppendColumn(createColumn("¿Se bloquea?", PROCESS_BLOCKED))
	treeView.AppendColumn(createColumn("Estado", PROCESS_STATUS))
	treeView.AppendColumn(createColumn("Número de partición", PARTITION_NUMBER))
	treeView.AppendColumn(createColumn("Tamaño de partición", PARTITION_SIZE))

	listStore, _ := gtk.ListStoreNew(
        glib.TYPE_STRING,
        glib.TYPE_INT,
        glib.TYPE_INT,
        glib.TYPE_STRING,
        glib.TYPE_STRING,
        glib.TYPE_INT,
        glib.TYPE_INT,
    )
	treeView.SetModel(listStore)

	return treeView, listStore
}

func createColumn(columnTitle string, id int) *gtk.TreeViewColumn {
	cellRenderer, _ := gtk.CellRendererTextNew()
	column, _ := gtk.TreeViewColumnNewWithAttribute(columnTitle, cellRenderer, "text", id)
	return column
}

