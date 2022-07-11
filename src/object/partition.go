package object

import (
	"errors"
	"so-p4_memory/src/view/lang"
)

type PartitionNumber int

type Partition struct {
	Number  PartitionNumber
	Size    int
	Process *Process

    Processes []*Process
}

func (p *Partition) AddProcess(process *Process) error {
    if p.Size >= process.Size {
        p.Processes = append(p.Processes, process)
        return nil
    } else {
        return errors.New(lang.ERROR_COULDNT_ADD_PROCESS)
    }
}
