package view

import (
	"errors"
	"so-p4_memory/src/object"
	"so-p4_memory/src/view/lang"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type PartitionComboBox struct {
	comboBox  *gtk.ComboBox
	listStore *gtk.ListStore
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
		comboBox:  comboBox,
		listStore: listStore,
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

func (p *PartitionComboBox) GetSelectedPartition() (string, error) {
	tree, _ := p.comboBox.GetActiveIter()
	model, _ := p.comboBox.GetModel()
	value, _ := model.ToTreeModel().GetValue(tree, 0)
	partition, _ := value.GetString()

	if partition != "" {
		return partition, nil
	} else {
		return "", errors.New(lang.ERROR_PARTITION_NOT_SELECTED)
	}
}

func (p *PartitionComboBox) AddPartition(partition *object.Partition) {
	p.listStore.Set(
		p.listStore.Append(),
		[]int{0},
		[]interface{}{
			partition.Name,
		},
	)
	p.comboBox.SetSensitive(true)
	p.comboBox.SetActive(0)
}

func (p *PartitionComboBox) Reset() {
	p.listStore.Clear()
	p.comboBox.SetSensitive(false)
}
