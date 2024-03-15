package process

import (
	"memory-manager-simulator/memory"
)

// SystemOperation representa o sistema operacional
type SystemOperation struct{}

var (
	globalMemoryManager *memory.MemoryManager
	globalCpuManager    *CpuManager
	globalScheduler     *Scheduler
)

// SystemCall executa uma chamada de sistema
func (so *SystemOperation) SystemCall(callType SystemCallType, p *Process) *Process {
	if callType == CREATE_PROCESS {
		if globalMemoryManager == nil {
			globalMemoryManager = memory.NewMemoryManager(memory.FIRST_FIT)
		}
		if globalCpuManager == nil {
			globalCpuManager = NewCpuManager()
		}
		return NewProcess()
	} else if callType == WRITE_PROCESS {
		globalMemoryManager.Write(p)
	} else if callType == CLOSE_PROCESS {
		globalMemoryManager.Delete(p)
	}
	return nil
}
