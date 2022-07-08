package object

import (
	"fmt"
)

type ProcessState int

const (
	READY ProcessState = iota
	RUNNING
	BLOCKED
	FINISHED
)

type Process struct {
	Name          string
	Time          int
	TimeRemaining int
	IsBlocked     bool
	State         ProcessState

	Partition PartitionNumber
}

func (p *Process) GetTimeRemaining() int {
	return p.TimeRemaining
}

func (p *Process) HasFinished() bool {
	return p.TimeRemaining == 0
}

func (p *Process) ToString() string {
	return fmt.Sprintf("%s %d %t\n", p.Name, p.Time, p.IsBlocked)
}
