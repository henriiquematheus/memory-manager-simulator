// Conteúdo do arquivo systemoperation.go
package so

import (
	"fmt"
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
	case OPEN_PROCESS:
		fmt.Println("Abrindo processo:", p.ID)
	case READ_PROCESS:
		fmt.Println("Lendo processo:", p.ID)
	case CLOSE_PROCESS:
		fmt.Println("Fechando processo:", p.ID)
		globalMemoryManager.Delete(p)
	case CREATE_PROCESS:
		if globalMemoryManager == nil {
			globalMemoryManager = memory.NewMemoryManager(memory.FIRST_FIT)
		}
		fmt.Println("Processo criado:", p.ID)
		return process.NewProcess()
	case WRITE_PROCESS:
		if globalMemoryManager == nil {
			fmt.Println("Não há gerenciador de memória global disponível.")
			return nil
		}
		fmt.Println("Escrevendo processo na memória:", p.ID)
		success := globalMemoryManager.Write(p)
		if !success {
			fmt.Println("Falha ao escrever o processo na memória:", p.ID)
		}
	}
	return nil
}
