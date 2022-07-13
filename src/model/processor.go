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
	LogListeners     ProcessorLogListeners
	CurrentPartition *Partition
}

func (p *Processor) Process(listeners ProcessorLogListeners) {
	log.Default().Printf("Processing process %s:%s", p.CurrentPartition.Process.Name, p.CurrentPartition.Process.State)
	switch p.CurrentPartition.Process.State {
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
	p.CurrentPartition.Process.State = object.RUNNING
	p.CurrentPartition.Process.MakeProcess(PROCESSING_TIME)
	p.LogListeners.LogProcessDispatched(p.CurrentPartition.Process)
}

func (p *Processor) makeProcessRunningTransition() {
	if p.CurrentPartition.Process.HasFinished() {
		p.CurrentPartition.Process.State = object.FINISHED
		p.LogListeners.LogProcessFinished(p.CurrentPartition.Process)
		p.clearCurrentPartition()
	} else if p.CurrentPartition.Process.IsBlocked {
		p.CurrentPartition.Process.State = object.BLOCKED
		p.LogListeners.LogProcessBlocked(p.CurrentPartition.Process)
	} else {
		p.CurrentPartition.Process.State = object.READY
		p.LogListeners.LogProcessTimeout(p.CurrentPartition.Process)
		p.clearCurrentPartition()
	}
}

func (p *Processor) makeProcessBlockedTransition() {
	p.CurrentPartition.Process.State = object.READY
	p.LogListeners.LogProcessIOBlockedCompleted(p.CurrentPartition.Process)
	p.clearCurrentPartition()
}

func (p *Processor) Restart() {
	p.clearCurrentPartition()
}

func (p *Processor) clearCurrentPartition() {
	p.CurrentPartition.Process.Partition = nil
	p.CurrentPartition.Process = nil
	p.CurrentPartition = nil
}
