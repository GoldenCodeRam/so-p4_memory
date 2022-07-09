package view

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type PartitionComboBox struct {
	ComboBox  *gtk.ComboBox
	ListStore *gtk.ListStore
}

func CreatePartitionComboBox() *PartitionComboBox {
	comboBox, err := gtk.ComboBoxNewWithModelAndEntry(generateListStore())
	if err != nil {
		panic("Couldn't generate the combo box!")
	}

	partitionComboBox := PartitionComboBox{
		ComboBox: comboBox,
	}
	return &partitionComboBox
}

func generateListStore() *gtk.ListStore {
	listStore, err := gtk.ListStoreNew(
		glib.TYPE_INT,
		glib.TYPE_INT,
	)
	if err != nil {
		panic("Couldn't generate the list store element!")
	}
	return listStore
}
