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

func (so *SystemOperation) SystemCall(callType SystemCallType, arg interface{}) interface{} {
	switch callType {
	case CREATE_PROCESS:
		size, ok := arg.(int)
		if !ok {
			return nil
		}
		return process.NewProcess(size)
	case WRITE_PROCESS:
		p, ok := arg.(*process.Process)
		if !ok {
			return nil
		}
		so.memoryManager.Write(p)
	case CLOSE_PROCESS:
		p, ok := arg.(*process.Process)
		if !ok {
			return nil
		}
		so.memoryManager.Delete(p)
	}
	return nil
}
