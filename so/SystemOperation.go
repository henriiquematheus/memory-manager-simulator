package so

import (
	"memory-manager-simulator/memory"
	"memory-manager-simulator/process"
)

// SystemOperation representa o sistema operacional
type SystemOperation struct{}

var (
	globalMemoryManager *memory.MemoryManager
)

func NewSystemOperation() *SystemOperation {
	return &SystemOperation{}
}

// SystemCall executa uma chamada de sistema
func (so *SystemOperation) SystemCall(callType SystemCallType, p *process.Process) *process.Process {
	switch callType {
	case CREATE_PROCESS:
		if globalMemoryManager == nil {
			globalMemoryManager = memory.NewMemoryManager(memory.FIRST_FIT)
		}
		return process.NewProcess()
	case WRITE_PROCESS:
		globalMemoryManager.Write(p)
	case CLOSE_PROCESS:
		globalMemoryManager.Delete(p)
	}
	return nil
}
