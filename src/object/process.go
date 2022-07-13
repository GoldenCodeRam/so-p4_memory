package object

import (
	"fmt"
	"so-p4_memory/src/view/lang"
)

type ProcessState int
type IsBlocked bool

func (p ProcessState) String() string {
	switch p {
	case READY:
		return lang.READY
	case RUNNING:
		return lang.RUNNING
	case BLOCKED:
		return lang.BLOCKED
	case FINISHED:
		return lang.FINISHED
	case NOT_ENOUGH_SPACE:
		return lang.NOT_ENOUGH_SPACE
	}
	panic("This should not happen!")
}

func (i IsBlocked) String() string {
	if i {
		return lang.YES
	} else {
		return lang.NO
	}
}

const (
	READY ProcessState = iota
	RUNNING
	BLOCKED
	FINISHED
	NOT_ENOUGH_SPACE
)

type Process struct {
	Name          string
	Time          int
	TimeRemaining int
	IsBlocked     IsBlocked
	State         ProcessState
	Size          int

	Partition *Partition
}

func (p *Process) ToString() string {
	return fmt.Sprintf("%s %d %t\n", p.Name, p.Time, p.IsBlocked)
}
