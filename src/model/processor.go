package model

import (
	"log"
	"so-p4_memory/src/object"
)

type ProcessorLogListeners interface {
	// Ready - Running
	LogProcessDispatched(process *Process)
	// Running - Ready
	LogProcessTimeout(process *Process)
	// Running - Blocked
	LogProcessBlocked(process *Process)
	// Running - Finished
	LogProcessFinished(process *Process)
	// Blocked - Ready
	LogProcessIOBlockedCompleted(process *Process)
}

const PROCESSING_TIME = 5

type Processor struct {
	LogListeners   ProcessorLogListeners
	CurrentProcess *Process
}

func (p *Processor) Process(listeners ProcessorLogListeners) {
    log.Default().Printf("Processing process %s:%s", p.CurrentProcess.Name, p.CurrentProcess.State)

	switch p.CurrentProcess.State {
	case object.READY:
		p.makeProcessReadyTransition()
		break
	case object.RUNNING:
		p.makeProcessRunningTransition()
		break
	case object.BLOCKED:
		p.makeProcessBlockedTransition()
		break
	case object.FINISHED:
		panic("State FINISHED should never happen!")
	default:
		panic("Process with an undefined state!")
	}
}

func (p *Processor) makeProcessReadyTransition() {
	p.CurrentProcess.State = object.RUNNING
	p.CurrentProcess.MakeProcess(PROCESSING_TIME)
	p.LogListeners.LogProcessDispatched(p.CurrentProcess)
}

func (p *Processor) makeProcessRunningTransition() {
	if p.CurrentProcess.HasFinished() {
		p.CurrentProcess.State = object.FINISHED
		p.LogListeners.LogProcessFinished(p.CurrentProcess)
		p.CurrentProcess = nil
	} else if p.CurrentProcess.IsBlocked {
		p.CurrentProcess.State = object.BLOCKED
		p.LogListeners.LogProcessBlocked(p.CurrentProcess)
	} else {
		p.CurrentProcess.State = object.READY
		p.LogListeners.LogProcessTimeout(p.CurrentProcess)
		p.CurrentProcess = nil
	}
}

func (p *Processor) makeProcessBlockedTransition() {
	p.CurrentProcess.State = object.READY
	p.LogListeners.LogProcessIOBlockedCompleted(p.CurrentProcess)
	p.CurrentProcess = nil
}

func (p *Processor) Restart() {
	p.CurrentProcess = nil
}
