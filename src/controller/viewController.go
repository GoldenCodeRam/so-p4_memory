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
	v.MainWindow.Window.SetDefaultSize(800, 600)
	v.MainWindow.Window.ShowAll()

	gtk.Main()
}

func (v *viewController) startProcessor() {
	log.Default().Println("Making processor tick...")
	for _, partition := range v.Storage.Partitions {
		for len(partition.Processes) > 0 || v.Processor.CurrentProcess != nil {
			log.Default().Printf("Starting with partition %d", partition.Number)

			if v.Processor.CurrentProcess == nil {
				process := model.Process{
					Process: partition.Processes[0],
				}
				v.Processor.CurrentProcess = &process
				partition.Processes = partition.Processes[1:]
			}

			v.Processor.Process(v)
		}
	}
}

func (v *viewController) updateAllStatesTreeViews(partition *object.Partition) {
	v.updateInputProcessesTreeView(partition)
	v.updateReadyProcessesTreeView(partition)
	v.updateRunningProcessesTreeView(partition)
	v.updateBlockedProcessesTreeView(partition)
	v.updateFinishedProcessesTreeView(partition)

	v.updateReadyToRunningProcessesTreeView(partition)
	v.updateRunningToReadyProcessesTreeView(partition)
	v.updateRunningToBlockedProcessesTreeView(partition)
	v.updateBlockedToReadyProcessesTreeView(partition)
	v.updateRunningToFinishedProcessesTreeView(partition)
}

func (v *viewController) updateInputProcessesTreeView(partition *object.Partition) {
	v.MainWindow.ProcessesOutputPanel.UpdateInputProcessesTreeView(
        partition,
		v.Storage.ExportInputProcessesListToLogList(),
	)
}

func (v *viewController) updateReadyProcessesTreeView(partition *object.Partition) {
	v.MainWindow.ProcessesOutputPanel.UpdateReadyProcessesTreeView(
        partition,
		v.Storage.ReadyProcessesLogList,
	)
}

func (v *viewController) updateRunningProcessesTreeView(partition *object.Partition) {
	v.MainWindow.ProcessesOutputPanel.UpdateRunningProcessesTreeView(
        partition,
		v.Storage.RunningProcessesLogList,
	)
}

func (v *viewController) updateBlockedProcessesTreeView(partition *object.Partition) {
	v.MainWindow.ProcessesOutputPanel.UpdateBlockedProcessesTreeView(
        partition,
		v.Storage.BlockedProcessesLogList,
	)
}

func (v *viewController) updateFinishedProcessesTreeView(partition *object.Partition) {
	v.MainWindow.ProcessesOutputPanel.UpdateFinishedProcessesTreeView(
        partition,
		v.Storage.FinishedProcessesLogList,
	)
}

func (v *viewController) updateReadyToRunningProcessesTreeView(partition *object.Partition) {
	v.MainWindow.ProcessesOutputPanel.UpdateReadyToRunningProcessesTreeView(
        partition,
		v.Storage.ReadyToRunningProcessesLogList,
	)
}

func (v *viewController) updateRunningToReadyProcessesTreeView(partition *object.Partition) {
	v.MainWindow.ProcessesOutputPanel.UpdateRunningToReadyProcessesTreeView(
        partition,
		v.Storage.RunningToReadyProcessesLogList,
	)
}

func (v *viewController) updateRunningToBlockedProcessesTreeView(partition *object.Partition) {
	v.MainWindow.ProcessesOutputPanel.UpdateRunningToBlockedProcessesTreeView(
        partition,
		v.Storage.RunningToBlockedProcessesLogList,
	)
}

func (v *viewController) updateBlockedToReadyProcessesTreeView(partition *object.Partition) {
	v.MainWindow.ProcessesOutputPanel.UpdateBlockedToReadyProcessesTreeView(
        partition,
		v.Storage.BlockedToReadyProcessesLogList,
	)
}

func (v *viewController) updateRunningToFinishedProcessesTreeView(partition *object.Partition) {
	v.MainWindow.ProcessesOutputPanel.UpdateRunningToFinishedProcessesTreeView(
        partition,
		v.Storage.RunningToFinishedProcessesLogList,
	)
}
