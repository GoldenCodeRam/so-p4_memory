package view

import (
	"so-p4_memory/src/object"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type PartitionLogTreeView struct {
	TreeView  *gtk.TreeView
	listStore *gtk.ListStore
}

func CreatePartitionLogTreeView() *PartitionLogTreeView {
	treeView, listStore := setupPartitionLogTreeView()
	return &PartitionLogTreeView{
		TreeView:  treeView,
		listStore: listStore,
	}
}

func (p *PartitionLogTreeView) AddRow(partition *object.Partition) {
	iter := p.listStore.Append()

	p.listStore.Set(
		iter,
		[]int{
			0,
			1,
		},
		[]interface{}{
            partition.Name,
            partition.Size,
		},
	)
}

func (p *PartitionLogTreeView) Clear() {
	p.listStore.Clear()
}

func setupPartitionLogTreeView() (*gtk.TreeView, *gtk.ListStore) {
	treeView, _ := gtk.TreeViewNew()

	treeView.AppendColumn(createPartitionLogColumn("Nombre de partición", 0))
	treeView.AppendColumn(createPartitionLogColumn("Tamaño de partición", 1))

	listStore, _ := gtk.ListStoreNew(
		glib.TYPE_STRING,
		glib.TYPE_STRING,
	)
	treeView.SetModel(listStore)

	return treeView, listStore
}

func createPartitionLogColumn(columnTitle string, id int) *gtk.TreeViewColumn {
	cellRenderer, _ := gtk.CellRendererTextNew()
	column, _ := gtk.TreeViewColumnNewWithAttribute(columnTitle, cellRenderer, "text", id)
	return column
}
