package memory

import (
	"fmt"
	"memory-manager-simulator/process"
	"strconv"
	"strings"
)

func NewMemoryManager(strategy int) *MemoryManager {
	return &MemoryManager{
		Strategy: strategy,
	}
}

type MemoryManager struct {
	physicalMemory [256]string
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
	fmt.Println("Iniciando a remoção do processo na memória")
	found := false

	// Percorre cada posição na memória física
	for i := range mm.physicalMemory {

		if strings.HasPrefix(mm.physicalMemory[i], p.ID) {
			//Se a posição contem o ID do processo remove
			mm.physicalMemory[i] = ""
			found = true
		}
	}

	if found {
		fmt.Println("Processo removido com sucesso")
	} else {
		fmt.Println("Processo não encontrado na memória")
	}

	// Exibe o status da memória
	mm.printMemoryStatus()
	fmt.Println("Finalizando a remoção do processo na memória")
}

func (mm *MemoryManager) writeWithFirstFit(p *process.Process) {
	fmt.Println("Iniciando a escrita do processo na memória com a estratégia First Fit")
	fmt.Printf("Tamanho do processo (%s): %d\n", p.ID, p.SizeInMemory)

	start := -1
	var bestPage *process.AddressMemory

	for i := 0; i < len(mm.physicalMemory); i++ {
		if mm.physicalMemory[i] == "" {
			if start == -1 {
				start = i // Encontra o início do espaço vazio
			}
			// Verifica se o espaço é suficiente para o processo
			if i-start+1 >= p.SizeInMemory {
				bestPage = &process.AddressMemory{Start: start, End: i}
				break
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
			// Atribui ao índice atual da memória física o ID do processo com o fragmento sequencial
			fragmentNumber := i - bestPage.Start + 1
			mm.physicalMemory[i] = fmt.Sprintf("%s-%02d", p.ID, fragmentNumber)
		}

		// Atualiza o fragmento do processo
		p.IncrementFragment()

		mm.printMemoryStatus()
	}
	fmt.Println("Finalizando a escrita do processo na memória com a estratégia First Fit")
}

func (mm *MemoryManager) printMemoryStatus() {
	for i := 0; i < len(mm.physicalMemory); i++ {

		if mm.physicalMemory[i] == "" {
			fmt.Print("_ | ")
		} else {
			// Extrair o ID do processo e o número do fragmento
			idParts := strings.Split(mm.physicalMemory[i], "-")
			processID := idParts[0]
			fragmentNumber := idParts[1]

			fragmentInt, err := strconv.Atoi(fragmentNumber)
			if err != nil {
				fmt.Print(mm.physicalMemory[i], " | ")
			} else {

				fmt.Printf("%s-%02d | ", processID, fragmentInt)
			}
		}
	}
	fmt.Println()
}

func (mm *MemoryManager) writeWithBestFit(p *process.Process) {
	fmt.Println("Iniciando a escrita do processo na memória com a estratégia Best Fit")
	fmt.Printf("Tamanho do processo (%s): %d\n", p.ID, p.SizeInMemory)

	var bestPage *process.AddressMemory
	bestFitSize := len(mm.physicalMemory) + 1
	start := -1

	for i := 0; i < len(mm.physicalMemory); i++ {

		if mm.physicalMemory[i] == "" {
			if start == -1 {
				start = i // Início do espaço vazio
			}

			// Avança enquanto as posições seguintes também estiverem vazias
			j := i
			for ; j < len(mm.physicalMemory) && mm.physicalMemory[j] == ""; j++ {
			}

			size := j - start
			if size >= p.SizeInMemory && size < bestFitSize {
				// Atualiza a melhor partição encontrada
				bestPage = &process.AddressMemory{Start: start, End: j - 1}
				bestFitSize = size
			}

			// Avança o índice `i` para o final do espaço vazio encontrado
			i = j - 1
		} else {
			// Reinicia o início da busca por espaço vazio
			start = -1
		}
	}

	// Se uma melhor partição foi encontrada
	if bestPage != nil {
		fmt.Println("Processo inserido com sucesso")

		// Insere o processo na melhor partição encontrada
		for i := bestPage.Start; i < bestPage.Start+p.SizeInMemory; i++ {
			fragmentNumber := i - bestPage.Start + 1
			mm.physicalMemory[i] = fmt.Sprintf("%s-%02d", p.ID, fragmentNumber)
		}

		// Atualiza o fragmento do processo
		p.IncrementFragment()

		mm.printMemoryStatus()
	} else {
		fmt.Println("Não há espaço suficiente na memória para o processo")
	}

	fmt.Println("Finalizando a escrita do processo na memória com a estratégia Best Fit")
}

func (mm *MemoryManager) writeWithWorstFit(p *process.Process) {
	fmt.Println("Iniciando a escrita do processo na memória com a estratégia Worst Fit")
	fmt.Printf("Tamanho do processo (%s): %d\n", p.ID, p.SizeInMemory)

	var worstPage *process.AddressMemory
	largestFreeSize := 0 // Mantém o tamanho da maior partição encontrada
	start := -1

	for i := 0; i < len(mm.physicalMemory); i++ {
		// Verifica se a posição atual está vazia
		if mm.physicalMemory[i] == "" {
			if start == -1 {
				start = i
			}

			j := i
			for ; j < len(mm.physicalMemory) && mm.physicalMemory[j] == ""; j++ {
			}

			// Calcula o tamanho do espaço vazio
			size := j - start

			if size >= p.SizeInMemory && size > largestFreeSize {
				// Atualiza a maior partição encontrada
				worstPage = &process.AddressMemory{Start: start, End: j - 1}
				largestFreeSize = size
			}

			i = j - 1
		} else {
			// Reinicia o início da busca por espaço vazio
			start = -1
		}
	}

	if worstPage != nil {
		fmt.Println("Processo inserido com sucesso")

		// Insere o processo na maior partição encontrada
		for i := worstPage.Start; i < worstPage.Start+p.SizeInMemory; i++ {

			fragmentNumber := i - worstPage.Start + 1
			mm.physicalMemory[i] = fmt.Sprintf("%s-%02d", p.ID, fragmentNumber)
		}

		// Exibe o status da memória
		mm.printMemoryStatus()
	} else {

		fmt.Println("Não há espaço suficiente na memória para o processo")
	}

	fmt.Println("Finalizando a escrita do processo na memória com a estratégia Worst Fit")
}
