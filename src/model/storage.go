package model

import (
	"errors"
	"so-p4_memory/src/view/lang"
)

type Storage struct {
	EntryProcessesLogList    []*Process
	ReadyProcessesLogList    []*Process
	RunningProcessesLogList  []*Process
	BlockedProcessesLogList  []*Process
	FinishedProcessesLogList []*Process
}

func (s *Storage) AddProcessToReadyList(process *Process) error {
	for _, p := range s.ReadyProcessesLogList {
		if p.Name == process.Name {
			return errors.New(lang.ERROR_COULD_NOT_ADD_PROCESS)
		}
	}

	s.ReadyProcessesLogList = append(s.ReadyProcessesLogList, process)
	return nil
}

func (s *Storage) GetNextEntryProcess() *Process {
	if len(s.EntryProcessesLogList) > 0 {
		result := s.EntryProcessesLogList[0]
		s.EntryProcessesLogList = s.EntryProcessesLogList[1:]
		return result
	} else {
		return nil
	}
}

func (s *Storage) Reset() {
	s.EntryProcessesLogList = make([]*Process, 0)
	s.ReadyProcessesLogList = make([]*Process, 0)
	s.RunningProcessesLogList = make([]*Process, 0)
	s.BlockedProcessesLogList = make([]*Process, 0)
	s.FinishedProcessesLogList = make([]*Process, 0)
}
