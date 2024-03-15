package memory

import (
	"fmt"

	"github.com/henriiquematheus/memory-manager-simulator/process"
)

// MemoryManager gerencia a memória física
type MemoryManager struct {
	physicalMemory [128]string // Memória física com 128 KB
	Strategy       int         // Adicionando a declaração da Strategy como um tipo inteiro
}

// Write escreve um processo na memória usando a estratégia atual
func (mm *MemoryManager) Write(p *process.Process) {
	switch mm.Strategy {
	case FIRST_FIT:
		mm.writeWithFirstFit(p)
	case BEST_FIT:
		mm.writeWithBestFit(p)
	case WORST_FIT:
		mm.writeWithWorstFit(p)
	}
}

// Delete remove um processo da memória
func (mm *MemoryManager) Delete(p *process.Process) {
	// Implemente a lógica para remover o processo da memória
}

func (mm *MemoryManager) writeWithFirstFit(p *process.Process) {
	fmt.Println("Escrevendo o processo na memória")
	actualSize := 0
	var bestPage *process.AddressMemory

	for i := 0; i < len(mm.physicalMemory); i++ {
		if mm.physicalMemory[i] == "" {
			actualSize++
		} else {
			if actualSize > 0 {
				start := i - actualSize
				end := i - 1
				address := process.NewAddressMemory(start, end)
				if p.SizeInMemory <= address.GetSize() {
					bestPage = address
				}
			}
			actualSize = 0
		}
	}

	if bestPage == nil || actualSize == 0 {
		fmt.Println("Não há espaço na memória")
	} else {
		fmt.Println("Processo inserido com sucesso")
		for i := bestPage.Start; i <= bestPage.End; i++ {
			mm.physicalMemory[i] = p.ID
		}
	}
	mm.printMemoryStatus()
}

// printMemoryStatus imprime o status da memória
func (mm *MemoryManager) printMemoryStatus() {
	for i := 0; i < len(mm.physicalMemory); i++ {
		fmt.Print(mm.physicalMemory[i], " | ")
	}
}

// writeWithBestFit escreve um processo na memória usando a estratégia Best Fit
func (mm *MemoryManager) writeWithBestFit(p *process.Process) {
	// Implemente a lógica de escrita usando a estratégia Best Fit
}

// writeWithWorstFit escreve um processo na memória usando a estratégia Worst Fit
func (mm *MemoryManager) writeWithWorstFit(p *process.Process) {
	// Implemente a lógica de escrita usando a estratégia Worst Fit
}
