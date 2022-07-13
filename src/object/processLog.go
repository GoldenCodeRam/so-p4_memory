package object

import (
	"so-p4_memory/src/view/lang"
	"strconv"
)

type ProcessLog struct {
	Name          string
	Time          int
	Size          int
	TimeRemaining int
	IsBlocked     string
	State         string

	PartitionName string
	PartitionSize string
}

func CreateProcessLogFromProcess(process *Process) *ProcessLog {
	processLog := ProcessLog{
		Name:          process.Name,
		Time:          process.Time,
		Size:          process.Size,
		TimeRemaining: process.TimeRemaining,
		IsBlocked:     process.IsBlocked.String(),
		State:         process.State.String(),
		PartitionName: lang.DOES_NOT_APPLY,
		PartitionSize: lang.DOES_NOT_APPLY,
	}

	if process.Partition != nil {
		processLog.PartitionName = process.Partition.Name
		processLog.PartitionSize = strconv.Itoa(process.Partition.Size)
	}

	return &processLog
}
