package controller

import (
	"log"
	"so-p4_memory/src/model"
	"so-p4_memory/src/object"
	"so-p4_memory/src/view"
	"so-p4_memory/src/view/lang"
	"so-p4_memory/src/view/utils"
)

func (v *viewController) LogProcessDispatched(process *model.Process) {
	v.Storage.AddProcessToProcessLog(process.Process, model.RUNNING_PROCESSES_LOG_LIST)
	v.MainWindow.UpdateTreeView(
		v.Storage.GetLogList(model.RUNNING_PROCESSES_LOG_LIST),
		view.RUNNING_PROCESSES_TREE_VIEW,
	)
}

func (v *viewController) LogProcessTimeout(process *model.Process) {
	v.Storage.AddProcessToProcessLog(process.Process, model.READY_PROCESSES_LOG_LIST)
    v.Storage.ReturnProcessToValidProcesses(process)
	v.MainWindow.UpdateTreeView(
		v.Storage.GetLogList(model.READY_PROCESSES_LOG_LIST),
		view.READY_PROCESSES_TREE_VIEW,
	)
}

func (v *viewController) LogProcessBlocked(process *model.Process) {
	v.Storage.AddProcessToProcessLog(process.Process, model.BLOCKED_PROCESSES_LOG_LIST)
	v.MainWindow.UpdateTreeView(
		v.Storage.GetLogList(model.BLOCKED_PROCESSES_LOG_LIST),
		view.BLOCKED_PROCESSES_TREE_VIEW,
	)
}

func (v *viewController) LogProcessFinished(process *model.Process) {
	v.Storage.AddProcessToProcessLog(process.Process, model.FINISHED_PROCESSES_LOG_LIST)
	v.MainWindow.UpdateTreeView(
		v.Storage.GetLogList(model.FINISHED_PROCESSES_LOG_LIST),
		view.FINISHED_PROCESSES_TREE_VIEW,
	)
}

func (v *viewController) LogProcessIOBlockedCompleted(process *model.Process) {
	v.Storage.AddProcessToProcessLog(process.Process, model.READY_PROCESSES_LOG_LIST)
    v.Storage.ReturnProcessToValidProcesses(process)
	v.MainWindow.UpdateTreeView(
		v.Storage.GetLogList(model.READY_PROCESSES_LOG_LIST),
		view.READY_PROCESSES_TREE_VIEW,
	)
}

func (v *viewController) ShowStateInputList() {
	v.MainWindow.ShowProcessesLogTreeView(view.LOG_PROCESSES_TREE_VIEW)
}

func (v *viewController) ShowStateReadyList() {
	v.MainWindow.ShowProcessesLogTreeView(view.READY_PROCESSES_TREE_VIEW)
}

func (v *viewController) ShowStateRunningList() {
	v.MainWindow.ShowProcessesLogTreeView(view.RUNNING_PROCESSES_TREE_VIEW)
}

func (v *viewController) ShowStateBlockedList() {
	v.MainWindow.ShowProcessesLogTreeView(view.BLOCKED_PROCESSES_TREE_VIEW)
}

func (v *viewController) ShowStateFinishedList() {
	v.MainWindow.ShowProcessesLogTreeView(view.FINISHED_PROCESSES_TREE_VIEW)
}

func (v *viewController) ShowPartitionsList() {
    v.MainWindow.ShowPartitionsLogTreeView()
}

func (v *viewController) AddProcess(process *object.Process) {
	err := v.Storage.AddProcessToInputList(&model.Process{
		Process: process,
	})
	if err != nil {
		utils.ShowErrorDialog(err)
	} else {
		log.Default().Printf("Created process with name %s.", process.Name)
		v.updateAllProcessLogsTreeViews()
	}
}

func (v *viewController) CreatePartition(partition *object.Partition) {
	err := v.Storage.AddPartition(&model.Partition{
		Partition: partition,
	})

	if err != nil {
		log.Default().Printf(err.Error())
		utils.ShowErrorDialog(err)
	} else {
		log.Default().Printf("Created partition %s.", partition.Name)
		v.MainWindow.AddToPartitionList(partition, func() {
            v.updateAllProcessLogsByPartitionTreeViews(partition)
        })
        v.MainWindow.UpdatePartitionLogsTreeView(v.Storage.GetPartitionLogs())
	}
}

func (v *viewController) GetPartitionWithName(partitionName string) *object.Partition {
	return v.Storage.GetPartitionWithName(partitionName).Partition
}

func (v *viewController) StartProcessor() {
	v.Storage.StartProcessorStorage()
	v.updateAllProcessLogsTreeViews()

	v.startProcessor()
	v.updateAllProcessLogsTreeViews()
	utils.ShowInfoDialog(lang.PROCESSOR_FINISHED)
}

func (v *viewController) ShowWithoutPartitionFilters() {
	v.updateAllProcessLogsTreeViews()
}
