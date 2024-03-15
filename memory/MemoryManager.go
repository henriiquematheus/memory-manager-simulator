package memory

import (
	"fmt"
	"memory-manager-simulator/process"
)

func NewMemoryManager(strategy int) *MemoryManager {
	return &MemoryManager{
		Strategy: strategy,
	}
}

// MemoryManager gerencia a memória física
type MemoryManager struct {
	physicalMemory [128]string // Memória física com 128 KB
	Strategy       int         // Adicionando a declaração da Strategy como um tipo inteiro
}

// Write escreve um processo na memória usando a estratégia atual e retorna true se a escrita for bem-sucedida
func (mm *MemoryManager) Write(p *process.Process) bool {
	switch mm.Strategy {
	case FIRST_FIT:
		return mm.writeWithFirstFit(p)
	case BEST_FIT:
		return mm.writeWithBestFit(p)
	case WORST_FIT:
		return mm.writeWithWorstFit(p)
	default:
		return false
	}
}

// Delete remove um processo da memória
func (mm *MemoryManager) Delete(p *process.Process) {
	// Implemente a lógica para remover o processo da memória
}

func (mm *MemoryManager) writeWithFirstFit(p *process.Process) bool {
	fmt.Println("Escrevendo o processo na memória")
	actualSize := 0
	var bestPage *process.AddressMemory

	for i := 0; i < len(mm.physicalMemory); i++ {
		if mm.physicalMemory[i] == "" {
			actualSize++
			fmt.Printf("Espaço vazio encontrado: Tamanho atual: %d, Início: %d, Fim: %d\n", actualSize, i-actualSize, i)
		} else {
			if actualSize > 0 {
				start := i - actualSize
				end := i - 1
				address := &process.AddressMemory{Start: start, End: end}
				fmt.Printf("Processo com tamanho %d tentando ser alocado em espaço de tamanho %d\n", p.SizeInMemory, address.GetSize())
				if p.SizeInMemory <= address.GetSize() {
					bestPage = address
				}
			}
			actualSize = 0
		}
	}

	if bestPage == nil || actualSize == 0 {
		fmt.Println("Não há espaço na memória")
		return false
	}

	fmt.Println("Processo inserido com sucesso")
	for i := bestPage.Start; i <= bestPage.End; i++ {
		mm.physicalMemory[i] = p.ID
	}
	mm.printMemoryStatus()
	return true
}

// printMemoryStatus imprime o status da memória
func (mm *MemoryManager) printMemoryStatus() {
	for i := 0; i < len(mm.physicalMemory); i++ {
		fmt.Print(mm.physicalMemory[i], " | ")
	}
}

// writeWithBestFit escreve um processo na memória usando a estratégia Best Fit
func (mm *MemoryManager) writeWithBestFit(p *process.Process) bool {
	// Implemente a lógica de escrita usando a estratégia Best Fit
	return false
}

// writeWithWorstFit escreve um processo na memória usando a estratégia Worst Fit
func (mm *MemoryManager) writeWithWorstFit(p *process.Process) bool {
	// Implemente a lógica de escrita usando a estratégia Worst Fit
	return false
}
