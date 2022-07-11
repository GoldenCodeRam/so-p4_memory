package model

import (
	"math"
	"so-p4_memory/src/object"
)

type Process struct {
	*object.Process
}

func (p *Process) GetTimeRemaining() int {
	return p.Process.TimeRemaining
}

func (p *Process) HasFinished() bool {
	return p.Process.TimeRemaining <= 0
}

func (p *Process) MakeProcess(time int) {
	p.Process.TimeRemaining = int(math.Max(0.0, float64(p.Process.TimeRemaining-time)))
}

func (p *Process) CanEnterPartition() bool {
	return p.Process.Partition.Size >= p.Process.Size
}
