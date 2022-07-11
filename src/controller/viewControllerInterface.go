package controller

import (
	"log"
	"so-p4_memory/src/model"
	"so-p4_memory/src/object"
	"so-p4_memory/src/view/lang"
	"so-p4_memory/src/view/utils"
)

func (v *viewController) LogProcessDispatched(process *model.Process) {
	v.Storage.AddProcessLogToRunningList(process)
	v.Storage.AddProcessLogToReadyToRunningList(process)
	v.updateRunningProcessesTreeView(nil)
}

func (v *viewController) LogProcessTimeout(process *model.Process) {
	v.Storage.AddProcessLogToReadyList(process)
	v.Storage.AddProcessToPartition(process)
	v.Storage.AddProcessLogToRunningToReadyList(process)
	v.updateReadyProcessesTreeView(nil)
}

func (v *viewController) LogProcessBlocked(process *model.Process) {
	v.Storage.AddProcessLogToBlockedList(process)
	v.Storage.AddProcessLogToRunningToBlockedList(process)
	v.updateBlockedProcessesTreeView(nil)
}

func (v *viewController) LogProcessFinished(process *model.Process) {
	v.Storage.AddProcessLogToFinishedList(process)
	v.Storage.AddProcessLogToRunningToFinishedList(process)
	v.updateFinishedProcessesTreeView(nil)
}

func (v *viewController) LogProcessIOBlockedCompleted(process *model.Process) {
	v.Storage.AddProcessLogToReadyList(process)
	v.Storage.AddProcessToPartition(process)
	v.Storage.AddProcessLogToBlockedToReadyList(process)
	v.updateReadyProcessesTreeView(nil)
}

func (v *viewController) ShowStateInputList() {
	v.MainWindow.ProcessesOutputPanel.ShowInputProcessesTreeView()
}

func (v *viewController) ShowStateReadyList() {
	v.MainWindow.ProcessesOutputPanel.ShowReadyProcessesTreeView()
}

func (v *viewController) ShowStateRunningList() {
	v.MainWindow.ProcessesOutputPanel.ShowRunningProcessesTreeView()
}

func (v *viewController) ShowStateBlockedList() {
	v.MainWindow.ProcessesOutputPanel.ShowBlockedProcessesTreeView()
}

func (v *viewController) ShowStateFinishedList() {
	v.MainWindow.ProcessesOutputPanel.ShowFinishedProcessesTreeView()
}

func (v *viewController) ShowTransitionReadyToRunningList() {
	v.MainWindow.ProcessesOutputPanel.ShowReadyToRunningProcessesTreeView()
}

func (v *viewController) ShowTransitionRunningToReadyList() {
	v.MainWindow.ProcessesOutputPanel.ShowRunningToReadyProcessesTreeView()
}

func (v *viewController) ShowTransitionRunningToBlockedList() {
	v.MainWindow.ProcessesOutputPanel.ShowRunningToBlockedProcessesTreeView()
}

func (v *viewController) ShowTransitionBlockedToReadyList() {
	v.MainWindow.ProcessesOutputPanel.ShowBlockedToReadyProcessesTreeView()
}

func (v *viewController) ShowTransitionRunningToFinishedList() {
	v.MainWindow.ProcessesOutputPanel.ShowRunningToFinishedProcessesTreeView()
}

func (v *viewController) AddProcess(process *object.Process) {
	err := v.Storage.AddProcessToInputList(&model.Process{
		Process: process,
	})
	if err != nil {
		utils.ShowErrorDialog(err)
	} else {
		log.Default().Printf("Created process with name %s.", process.Name)
		v.updateAllStatesTreeViews(nil)
	}
}

func (v *viewController) CreatePartition(partition *object.Partition) {
	err := v.Storage.AddPartition(partition)
	if err != nil {
		utils.ShowErrorDialog(err)
	} else {
		log.Default().Printf("Created partition with number %d.", partition.Number)
		v.MainWindow.CreateElementsNotebook.CreateProcessPanel.PartitionComboBox.AddPartition(partition)
		v.MainWindow.SelectListElementsNotebook.SelectOptionPartitionListPanel.AddPartitionSelectionButton(
			partition,
			func() {
				v.updateAllStatesTreeViews(partition)
			},
		)
	}
}

func (v *viewController) StartProcessor() {
	v.Storage.ResetProcessorStorage()
	v.Storage.StartProcessorStorage()
	v.updateAllStatesTreeViews(nil)

	v.startProcessor()
	v.updateAllStatesTreeViews(nil)
	utils.ShowInfoDialog(lang.PROCESSOR_FINISHED)
}

func (v *viewController) ShowWithoutPartitionFilters() {
	v.updateAllStatesTreeViews(nil)
}
