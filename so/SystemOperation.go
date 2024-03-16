package so

import (
	"memory-manager-simulator/memory"
	"memory-manager-simulator/process"
)

type SystemOperation struct {
	memoryManager *memory.MemoryManager
}

func NewSystemOperation(strategy int) *SystemOperation {
	return &SystemOperation{
		memoryManager: memory.NewMemoryManager(strategy),
	}
}

func (so *SystemOperation) SystemCall(callType SystemCallType, p *process.Process) *process.Process {
	switch callType {
	case CREATE_PROCESS:
		if so.memoryManager == nil {
			so.memoryManager = memory.NewMemoryManager(memory.FIRST_FIT)
		}
		return process.NewProcess()
	case WRITE_PROCESS:
		so.memoryManager.Write(p)
	case CLOSE_PROCESS:
		so.memoryManager.Delete(p)
	}
	return nil
}
