package model

import "so-p4_memory/src/object"

type ProcessorLogListeners interface {
	// Ready - Running
	LogProcessDispatched(process *Process)
	// Ready - SuspendedReady
	//LogProcessSuspendedReady(process *Process)
	// Running - Ready
	LogProcessTimeout(process *Process)
	// Running - Blocked
	LogProcessBlocked(process *Process)
	// Running - Finished
	LogProcessFinished(process *Process)
	// Blocked - Ready
	LogProcessIOBlockedCompleted(process *Process)

	// Finished all processes
	LogFinishedProcessing()
}

const PROCESSING_TIME = 5

type Processor struct {
	LogListeners ProcessorLogListeners
}

func (p *Processor) Process(process *Process, listeners ProcessorLogListeners) {
	switch process.State {
	case object.READY:
		p.makeProcessReadyTransition(process)
		break
	case object.RUNNING:
		p.makeProcessRunningTransition(process)
		break
	case object.BLOCKED:
		p.makeProcessBlockedTransition(process)
		break
	case object.FINISHED:
		panic("State FINISHED should never happen!")
	default:
		panic("Process with an undefined state!")
	}
}

func (p *Processor) makeProcessReadyTransition(process *Process) {
	process.State = object.RUNNING
	process.MakeProcess(PROCESSING_TIME)
	p.LogListeners.LogProcessDispatched(process)
}

func (p *Processor) makeProcessRunningTransition(process *Process) {
	if process.HasFinished() {
		process.State = object.FINISHED
		p.LogListeners.LogProcessFinished(process)
		process = nil
	} else if process.IsBlocked {
		process.State = object.BLOCKED
		p.LogListeners.LogProcessBlocked(process)
	} else {
		process.State = object.READY
		p.LogListeners.LogProcessTimeout(process)
		process = nil
	}
}

func (p *Processor) makeProcessBlockedTransition(process *Process) {
	process.State = object.READY
	p.LogListeners.LogProcessIOBlockedCompleted(process)
	process = nil
}
