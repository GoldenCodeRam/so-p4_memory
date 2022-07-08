package controller

import (
	"so-p4_memory/src/model"
	"so-p4_memory/src/object"
	"so-p4_memory/src/view/utils"
)

func (v *viewController) LogProcessDispatched(process *model.Process) {
}

func (v *viewController) LogProcessTimeout(process *model.Process) {
}

func (v *viewController) LogProcessBlocked(process *model.Process) {
}

func (v *viewController) LogProcessFinished(process *model.Process) {
}

func (v *viewController) LogProcessIOBlockedCompleted(process *model.Process) {
}

func (v *viewController) LogFinishedProcessing() {

}

func (v *viewController) ShowStateInputList() {

}

func (v *viewController) ShowStateReadyList() {

}

func (v *viewController) ShowStateRunningList() {

}

func (v *viewController) ShowStateBlockedList() {

}

func (v *viewController) ShowStateFinishedList() {

}

func (v *viewController) ShowTransitionReadyToRunningList() {

}

func (v *viewController) ShowTransitionRunningToReadyList() {

}

func (v *viewController) ShowTransitionRunningToBlockedList() {

}

func (v *viewController) ShowTransitionBlockedToReadyList() {

}

func (v *viewController) ShowTransitionRunningToFinishedList() {

}


func (v *viewController) AddProcess(process *object.Process) {
	err := v.Storage.AddProcessToReadyList(&model.Process{
		Process: process,
	})
	if err != nil {
		utils.ShowErrorDialog(err)
	} else {
		// TODO
	}
}

func (v *viewController) CreatePartition(partition *object.Partition) {
    // TODO: Also update the Combobox for the user to select a parititon for
    // the creation of a process.
}
