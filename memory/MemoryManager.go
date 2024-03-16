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

type MemoryManager struct {
	physicalMemory [128]string
	Strategy       int
}

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

func (mm *MemoryManager) Delete(p *process.Process) {

}

func (mm *MemoryManager) writeWithFirstFit(p *process.Process) {
	fmt.Println("Iniciando a escrita do processo na memória com a estratégia First Fit")
	fmt.Printf("Tamanho do processo: %d\n", p.SizeInMemory)

	var bestPage *process.AddressMemory
	start := -1

	for i := 0; i < len(mm.physicalMemory); i++ {
		if mm.physicalMemory[i] == "" {
			if start == -1 {
				start = i // Início do espaço vazio
				fmt.Printf("Início do espaço vazio encontrado na posição: %d\n", start)
			}
			// Verificar se o tamanho do espaço vazio é suficiente
			size := i - start + 1

			if p.SizeInMemory <= size && (bestPage == nil || size < bestPage.GetSize()) {
				bestPage = &process.AddressMemory{Start: start, End: i}
				fmt.Printf("Melhor página atualizada: Início = %d, Fim = %d\n", bestPage.Start, bestPage.End)
			}
		} else {
			start = -1

		}
	}

	if bestPage == nil {
		fmt.Println("Não há espaço na memória para o processo")
	} else {
		fmt.Println("Processo inserido com sucesso")
		for i := bestPage.Start; i <= bestPage.End; i++ {
			mm.physicalMemory[i] = p.ID
		}
	}
	mm.printMemoryStatus()
	fmt.Println("Finalizando a escrita do processo na memória com a estratégia First Fit")
}

// printMemoryStatus imprime o status da memória
func (mm *MemoryManager) printMemoryStatus() {
	for i := 0; i < len(mm.physicalMemory); i++ {
		fmt.Print(mm.physicalMemory[i], " | ")
	}
}

func (mm *MemoryManager) writeWithBestFit(p *process.Process) {
	fmt.Println("Iniciando a escrita do processo na memória com a estratégia Best Fit")
	fmt.Printf("Tamanho do processo: %d\n", p.SizeInMemory)

	var bestPage *process.AddressMemory
	bestFitSize := len(mm.physicalMemory) + 1
	start := -1

	for i := 0; i < len(mm.physicalMemory); i++ {
		if mm.physicalMemory[i] == "" {
			if start == -1 {
				start = i // Início do espaço vazio
				fmt.Printf("Início do espaço vazio encontrado na posição: %d\n", start)
			}
			//verificar se é o melhor ajuste até agora
			j := i
			for ; j < len(mm.physicalMemory) && mm.physicalMemory[j] == ""; j++ {
			}
			size := j - start
			fmt.Printf("Tamanho do espaço vazio: %d\n", size)
			if size >= p.SizeInMemory && (size < bestFitSize || bestPage == nil) {
				bestPage = &process.AddressMemory{Start: start, End: j - 1}
				bestFitSize = size
				fmt.Printf("Melhor página atualizada: Início = %d, Fim = %d\n", bestPage.Start, bestPage.End)
			}
			i = j - 1
		} else {
			start = -1

		}
	}

	if bestPage == nil {
		fmt.Println("Não há espaço na memória para o processo")
	} else {
		fmt.Println("Processo inserido com sucesso")
		for i := bestPage.Start; i <= bestPage.End; i++ {
			mm.physicalMemory[i] = p.ID
		}
	}
	mm.printMemoryStatus()
	fmt.Println("Finalizando a escrita do processo na memória com a estratégia Best Fit")
}

func (mm *MemoryManager) writeWithWorstFit(p *process.Process) {
	fmt.Println("Iniciando a escrita do processo na memória com a estratégia Worst Fit")
	fmt.Printf("Tamanho do processo: %d\n", p.SizeInMemory)
	var worstPage *process.AddressMemory
	start := -1

	for i := 0; i < len(mm.physicalMemory); i++ {
		if mm.physicalMemory[i] == "" {
			if start == -1 {
				start = i // Início do espaço vazio
				fmt.Printf("Início do espaço vazio encontrado na posição: %d\n", start)
			}
			// Verificar se o tamanho do espaço vazio é suficiente
			size := i - start + 1

			if p.SizeInMemory <= size && (worstPage == nil || size > worstPage.GetSize()) {
				worstPage = &process.AddressMemory{Start: start, End: i}

			}
		} else {
			start = -1 // Reiniciamos o início do espaço vazio

		}
	}

	if worstPage == nil {
		fmt.Println("Não há espaço na memória para o processo")
	} else {
		fmt.Println("Processo inserido com sucesso")
		for i := worstPage.Start; i < worstPage.Start+p.SizeInMemory; i++ {
			mm.physicalMemory[i] = p.ID
		}
		mm.printMemoryStatus()
		fmt.Println("Finalizando a escrita do processo na memória com a estratégia Worst Fit")
	}
}
