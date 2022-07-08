package model

import (
	"math"
	"so-p4_memory/src/object"
)

type Process struct {
	*object.Process
}

func (p *Process) MakeProcess(time int) {
	p.TimeRemaining = int(math.Max(0.0, float64(p.TimeRemaining-time)))
}
