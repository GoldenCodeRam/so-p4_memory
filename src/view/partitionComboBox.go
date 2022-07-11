package view

import (
	"errors"
	"fmt"
	"so-p4_memory/src/object"
	"so-p4_memory/src/view/lang"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type PartitionComboBox struct {
	ComboBox  *gtk.ComboBox
	ListStore *gtk.ListStore
}

func CreatePartitionComboBox() *PartitionComboBox {
	listStore := generateListStore()
	comboBox, err := gtk.ComboBoxNewWithModelAndEntry(listStore)
	if err != nil {
		panic("Couldn't generate the combo box!")
	}

	comboBox.SetSensitive(false)
	comboBox.SetEntryTextColumn(0)
	partitionComboBox := PartitionComboBox{
		ComboBox:  comboBox,
		ListStore: listStore,
	}
	return &partitionComboBox
}

func generateListStore() *gtk.ListStore {
	listStore, err := gtk.ListStoreNew(
		glib.TYPE_STRING,
		glib.TYPE_POINTER,
	)
	if err != nil {
		panic("Couldn't generate the list store element!")
	}
	return listStore
}

func (p *PartitionComboBox) GetSelectedPartition() (*object.Partition, error) {
	tree, _ := p.ComboBox.GetActiveIter()
	model, _ := p.ComboBox.GetModel()
	value, _ := model.ToTreeModel().GetValue(tree, 1)
	partition := value.GetPointer()

	if partition != nil {
		return (*object.Partition)(partition), nil
	} else {
		return nil, errors.New(lang.ERROR_PARTITION_NOT_SELECTED)
	}
}

func (p *PartitionComboBox) AddPartition(partition *object.Partition) {
	p.ListStore.Set(
		p.ListStore.Append(),
		[]int{0, 1},
		[]interface{}{
			fmt.Sprintf("Partición %d", partition.Number),
			partition,
		},
	)
	p.ComboBox.SetSensitive(true)
	p.ComboBox.SetActive(0)
}

func (p *PartitionComboBox) Reset() {
	p.ListStore.Clear()
	p.ComboBox.SetSensitive(false)
}
