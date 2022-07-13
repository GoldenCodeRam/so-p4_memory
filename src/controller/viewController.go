package controller

import (
	"log"
	"so-p4_memory/src/model"
	"so-p4_memory/src/object"
	"so-p4_memory/src/view"
	"sync"

	"github.com/gotk3/gotk3/gtk"
)

type viewController struct {
	MainWindow *view.MainWindow
	Processor  *model.Processor
	Storage    *model.Storage
}

var viewControllerInstance *viewController
var lock = &sync.Mutex{}

func GetViewControllerInstance() *viewController {
	if viewControllerInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if viewControllerInstance == nil {
			viewControllerInstance = &viewController{}
			viewControllerInstance.Processor = &model.Processor{
				LogListeners: viewControllerInstance,
			}
			viewControllerInstance.Storage = &model.Storage{}
		}
	}
	return viewControllerInstance
}

func (v *viewController) StartApplication() {
	gtk.Init(nil)

	v.MainWindow = view.CreateMainWindow(v)
	v.MainWindow.StartMainWindow()

	gtk.Main()
}

func (v *viewController) startProcessor() {
	log.Default().Println("Making processor tick...")

	process := v.Storage.GetNextReadyProcess()

	for process != nil {
		v.Processor.CurrentPartition = v.Storage.GetNextAvailablePartition(process)
		v.Processor.CurrentPartition.Process = process
		process.Partition = v.Processor.CurrentPartition.Partition

		for v.Processor.CurrentPartition != nil {
			v.Processor.Process(v)
		}
		process = v.Storage.GetNextReadyProcess()
	}
}

func (v *viewController) updateAllProcessLogsTreeViews() {
	v.MainWindow.UpdateTreeView(v.Storage.ExportInputProcessesListToLogList(), view.LOG_PROCESSES_TREE_VIEW)
	v.MainWindow.UpdateTreeView(v.Storage.GetLogList(model.READY_PROCESSES_LOG_LIST), view.READY_PROCESSES_TREE_VIEW)
	v.MainWindow.UpdateTreeView(v.Storage.GetLogList(model.RUNNING_PROCESSES_LOG_LIST), view.RUNNING_PROCESSES_TREE_VIEW)
	v.MainWindow.UpdateTreeView(v.Storage.GetLogList(model.BLOCKED_PROCESSES_LOG_LIST), view.BLOCKED_PROCESSES_TREE_VIEW)
	v.MainWindow.UpdateTreeView(v.Storage.GetLogList(model.FINISHED_PROCESSES_LOG_LIST), view.FINISHED_PROCESSES_TREE_VIEW)
}

func (v *viewController) updateAllProcessLogsByPartitionTreeViews(partition *object.Partition) {
	v.MainWindow.UpdateProcessesLogTreeViewByPartition(
		v.Storage.ExportInputProcessesListToLogList(),
		view.LOG_PROCESSES_TREE_VIEW,
		partition,
	)
	v.MainWindow.UpdateProcessesLogTreeViewByPartition(
		v.Storage.GetLogList(model.READY_PROCESSES_LOG_LIST),
		view.READY_PROCESSES_TREE_VIEW,
		partition,
	)
	v.MainWindow.UpdateProcessesLogTreeViewByPartition(
		v.Storage.GetLogList(model.RUNNING_PROCESSES_LOG_LIST),
		view.RUNNING_PROCESSES_TREE_VIEW,
		partition,
	)
	v.MainWindow.UpdateProcessesLogTreeViewByPartition(
		v.Storage.GetLogList(model.BLOCKED_PROCESSES_LOG_LIST),
		view.BLOCKED_PROCESSES_TREE_VIEW,
		partition,
	)
	v.MainWindow.UpdateProcessesLogTreeViewByPartition(
		v.Storage.GetLogList(model.FINISHED_PROCESSES_LOG_LIST),
		view.FINISHED_PROCESSES_TREE_VIEW,
		partition,
	)
}
