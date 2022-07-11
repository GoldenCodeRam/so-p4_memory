package model

import (
	"errors"
	"log"
	"so-p4_memory/src/object"
	"so-p4_memory/src/view/lang"
	"sort"
)

type Storage struct {
	InputProcessesList []*Process

	ReadyProcessesLogList    []*object.Process
	RunningProcessesLogList  []*object.Process
	BlockedProcessesLogList  []*object.Process
	FinishedProcessesLogList []*object.Process

	ReadyToRunningProcessesLogList    []*object.Process
	RunningToReadyProcessesLogList    []*object.Process
	RunningToBlockedProcessesLogList  []*object.Process
	BlockedToReadyProcessesLogList    []*object.Process
	RunningToFinishedProcessesLogList []*object.Process

	Partitions []*object.Partition
}

func (s *Storage) StartProcessorStorage() {
	log.Default().Println("Starting processor storage")
	s.sortPartitions()
	s.enterInputProcessListToPartitions()
	log.Default().Println("Finished processor storage")
}

func (s *Storage) ResetProcessorStorage() {
	log.Default().Println("Resetting processor storage")
	s.resetPartitionsProcesses()
	s.resetProcessesLogLists()
	log.Default().Println("Finished resetting processor storage")
}

func (s *Storage) AddProcessToInputList(process *Process) error {
	for _, p := range s.InputProcessesList {
		if p.Name == process.Name {
			return errors.New(lang.ERROR_COULDNT_ADD_PROCESS)
		}
	}

	s.InputProcessesList = append(s.InputProcessesList, process)
	return nil
}

func (s *Storage) AddProcessToPartition(process *Process) {
	for _, p := range s.Partitions {
		if p.Number == process.Partition.Number {
			p.Processes = append(p.Processes, process.Process)
		}
	}
}

func (s *Storage) ExportInputProcessesListToLogList() []*object.Process {
	processes := []*object.Process{}
	for _, process := range s.InputProcessesList {
		processes = append(processes, process.Process)
	}
	return processes
}

func (s *Storage) AddProcessLogToReadyList(process *Process) {
	s.ReadyProcessesLogList = append(s.ReadyProcessesLogList, process.Process)
}

func (s *Storage) AddProcessLogToRunningList(process *Process) {
	s.RunningProcessesLogList = append(s.RunningProcessesLogList, process.Process)
}

func (s *Storage) AddProcessLogToBlockedList(process *Process) {
	s.BlockedProcessesLogList = append(s.BlockedProcessesLogList, process.Process)
}

func (s *Storage) AddProcessLogToFinishedList(process *Process) {
	s.FinishedProcessesLogList = append(s.FinishedProcessesLogList, process.Process)
}

func (s *Storage) AddProcessLogToReadyToRunningList(process *Process) {
	s.ReadyToRunningProcessesLogList = append(s.ReadyToRunningProcessesLogList, process.Process)
}

func (s *Storage) AddProcessLogToRunningToReadyList(process *Process) {
	s.RunningToReadyProcessesLogList = append(s.RunningToReadyProcessesLogList, process.Process)
}

func (s *Storage) AddProcessLogToRunningToBlockedList(process *Process) {
	s.RunningToBlockedProcessesLogList = append(s.RunningToBlockedProcessesLogList, process.Process)
}

func (s *Storage) AddProcessLogToBlockedToReadyList(process *Process) {
	s.BlockedToReadyProcessesLogList = append(s.BlockedToReadyProcessesLogList, process.Process)
}

func (s *Storage) AddProcessLogToRunningToFinishedList(process *Process) {
	s.RunningToFinishedProcessesLogList = append(s.RunningToFinishedProcessesLogList, process.Process)
}

func (s *Storage) AddPartition(partition *object.Partition) error {
	for _, p := range s.Partitions {
		if p.Number == partition.Number {
			return errors.New(lang.ERROR_COULDNT_ADD_PARTITION)
		}
	}

	s.Partitions = append(s.Partitions, partition)
	return nil
}

func (s *Storage) enterInputProcessListToPartitions() {
	for _, partition := range s.Partitions {
		for _, process := range s.InputProcessesList {
			if partition.Number == process.Partition.Number {
				err := partition.AddProcess(process.Process)
				if err == nil {
					s.ReadyProcessesLogList = append(s.ReadyProcessesLogList, process.Process)
				}
			}
		}
	}
}

func (s *Storage) sortPartitions() {
	log.Default().Println("Sorting partitions")
	sort.SliceStable(s.Partitions, func(i, j int) bool {
		return s.Partitions[i].Number < s.Partitions[j].Number
	})
	log.Default().Println("Partitions sorted")
}

func (s *Storage) resetPartitionsProcesses() {
	for _, partition := range s.Partitions {
		partition.Processes = make([]*object.Process, 0)
	}
}

func (s *Storage) resetProcessesLogLists() {
	s.ReadyProcessesLogList = make([]*object.Process, 0)
	s.RunningProcessesLogList = make([]*object.Process, 0)
	s.BlockedProcessesLogList = make([]*object.Process, 0)
	s.FinishedProcessesLogList = make([]*object.Process, 0)

	s.RunningToReadyProcessesLogList = make([]*object.Process, 0)
	s.ReadyToRunningProcessesLogList = make([]*object.Process, 0)
	s.RunningToBlockedProcessesLogList = make([]*object.Process, 0)
	s.BlockedToReadyProcessesLogList = make([]*object.Process, 0)
	s.RunningToFinishedProcessesLogList = make([]*object.Process, 0)
}
