package process_os

import (
	"memory-manager-simulator/memory"
	"memory-manager-simulator/so" // Importe o pacote so
)

// SystemOperation representa o sistema operacional
type SystemOperation struct{}

var (
	globalMemoryManager *memory.MemoryManager
	globalCpuManager    *CpuManager
	globalScheduler     *Scheduler
)

// SystemCall executa uma chamada de sistema
func (so *SystemOperation) SystemCall(callType so.SystemCallType, p *Process) *Process {
	if callType == so.CREATE_PROCESS {
		if globalMemoryManager == nil {
			globalMemoryManager = memory.NewMemoryManager(memory.FIRST_FIT)
		}
		if globalCpuManager == nil {
			globalCpuManager = NewCpuManager()
		}
		return NewProcess()
	} else if callType == so.WRITE_PROCESS {
		globalMemoryManager.Write(p)
	} else if callType == so.CLOSE_PROCESS {
		globalMemoryManager.Delete(p)
	}
	return nil
}
