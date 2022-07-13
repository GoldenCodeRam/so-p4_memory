package model

import (
	"errors"
	"log"
	"so-p4_memory/src/object"
	"so-p4_memory/src/view/lang"
)

type LogList int

const (
	READY_PROCESSES_LOG_LIST LogList = iota
	RUNNING_PROCESSES_LOG_LIST
	BLOCKED_PROCESSES_LOG_LIST
	FINISHED_PROCESSES_LOG_LIST
)

var logLists = [...]LogList{
	READY_PROCESSES_LOG_LIST,
	RUNNING_PROCESSES_LOG_LIST,
	BLOCKED_PROCESSES_LOG_LIST,
	FINISHED_PROCESSES_LOG_LIST,
}

type Storage struct {
	inputProcessesList []*Process
	validProcessesList []*Process
	processesLogsList  [len(logLists)][]*object.ProcessLog

	partitions         []*Partition
	nextPartitionPointer int
}

func (s *Storage) StartProcessorStorage() {
	log.Default().Println("Starting processor storage.")
	s.startLogLists()
	log.Default().Println("Log lists started.")
	s.enterInputProcessesToValidProcessesList()
	s.enterValidProcessesToReadyProcessesLogList()
	log.Default().Println("Finished processor storage.")
}

func (s *Storage) ResetProcessorStorage() {
	log.Default().Println("Resetting processor storage.")
	s.resetProcessesLogLists()
	s.partitions = make([]*Partition, 0)
	log.Default().Println("Finished resetting processor storage.")
}

func (s *Storage) AddProcessToInputList(process *Process) error {
	for _, p := range s.inputProcessesList {
		if p.Name == process.Name {
			return errors.New(lang.FormatString(lang.ERROR_PROCESS_NAME_X_ALREADY_ADDED, process.Name))
		}
	}

	s.inputProcessesList = append(s.inputProcessesList, process)
	return nil
}

func (s *Storage) ReturnProcessToValidProcesses(process *Process) {
	s.validProcessesList = append(s.validProcessesList, process)
}

func (s *Storage) GetNextReadyProcess() *Process {
	if len(s.validProcessesList) > 0 {
		nextProcess := s.validProcessesList[0]
		s.validProcessesList = s.validProcessesList[1:]
		return nextProcess
	} else {
		return nil
	}

}

func (s *Storage) GetLogList(logList LogList) []*object.ProcessLog {
	return s.processesLogsList[logList]
}

func (s *Storage) GetPartitionLogs() []*object.Partition {
    partitions := []*object.Partition{}
    for _, partition := range s.partitions {
        partitions = append(partitions, partition.Partition)
    }
    return partitions
}

func (s *Storage) GetNextAvailablePartition(process *Process) *Partition {
	jumps := len(s.partitions)
	currentPartitionIndex := s.nextPartitionPointer
	for i := 0; i < jumps; i++ {
		log.Default().Printf("Searching partition %s", s.partitions[currentPartitionIndex].Name)
		if s.partitions[currentPartitionIndex].Size >= process.Process.Size && s.partitions[currentPartitionIndex].Process == nil {
			s.nextPartitionPointer = (currentPartitionIndex + 1) % len(s.partitions)
			return s.partitions[currentPartitionIndex]
		}
		currentPartitionIndex = (currentPartitionIndex + 1) % len(s.partitions)
	}
	return nil
}

func (s *Storage) ExportInputProcessesListToLogList() []*object.ProcessLog {
	processes := []*object.ProcessLog{}
	for _, process := range s.inputProcessesList {
		processes = append(
			processes,
			object.CreateProcessLogFromProcess(process.Process),
		)
	}
	return processes
}

func (s *Storage) AddProcessToProcessLog(process *object.Process, list LogList) {
	s.processesLogsList[list] = append(
		s.processesLogsList[list],
		object.CreateProcessLogFromProcess(process),
	)
}

func (s *Storage) AddPartition(partition *Partition) error {
	for _, p := range s.partitions {
		if p.Name == partition.Name {
			return errors.New(lang.ERROR_COULDNT_ADD_PARTITION)
		}
	}

	s.partitions = append(s.partitions, partition)
	return nil
}

func (s *Storage) GetPartitionWithName(partitionName string) *Partition {
	for _, partition := range s.partitions {
		if partition.Name == partitionName {
			return partition
		}
	}
	panic("Partition not found!")
}

func (s *Storage) startLogLists() {
	for _, logList := range logLists {
		s.processesLogsList[logList] = []*object.ProcessLog{}
	}
}

func (s *Storage) enterInputProcessesToValidProcessesList() {
	for _, process := range s.inputProcessesList {
		if s.canProcessEnterPartition(process) {
			if process.State == object.READY {
				s.validProcessesList = append(s.validProcessesList, process)
			}
		} else {
			process.State = object.NOT_ENOUGH_SPACE
		}
	}
}

func (s *Storage) canProcessEnterPartition(process *Process) bool {
	for _, partition := range s.partitions {
		if partition.Size >= process.Size {
			return true
		}
	}
	return false
}

func (s *Storage) enterValidProcessesToReadyProcessesLogList() {
	for _, process := range s.validProcessesList {
		s.processesLogsList[READY_PROCESSES_LOG_LIST] = append(
			s.processesLogsList[READY_PROCESSES_LOG_LIST],
			object.CreateProcessLogFromProcess(process.Process),
		)
	}
}

func (s *Storage) resetProcessesLogLists() {
	for range logLists {
		s.processesLogsList = [len(logLists)][]*object.ProcessLog{}
	}
}
