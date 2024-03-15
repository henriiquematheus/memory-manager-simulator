package so

import (
	"memory-manager-simulator/process" // Importação do pacote memory
)

// SystemOperation representa o sistema operacional
type SystemOperation struct{}

var (
	globalMemoryManager *MemoryManager
	globalCpuManager    *CpuManager
	globalScheduler     *Scheduler
)

// SystemCall executa uma chamada de sistema
type SystemOperation struct{}

// SystemCall executa uma chamada de sistema
func (so *SystemOperation) SystemCall(callType SystemCallType, p *Process) *Process {
	if callType == CREATE_PROCESS {
		if globalMemoryManager == nil {
			globalMemoryManager = NewMemoryManager(FIRST_FIT)
		}
		if globalCpuManager == nil {
			globalCpuManager = NewCpuManager()
		}
		return process.NewProcess() // Usando NewProcess() do pacote process
	} else if callType == WRITE_PROCESS {
		globalMemoryManager.Write(p)
	} else if callType == CLOSE_PROCESS {
		globalMemoryManager.Delete(p)
	}
	return nil
}
