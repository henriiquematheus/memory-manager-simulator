package process

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// Process representa um processo no sistema
type Process struct {
	ID            string        // ID do processo
	SizeInMemory  int           // Tamanho do processo em memória
	AddressMemory AddressMemory // Endereço de memória do processo
}

// NewProcess cria uma nova instância de Process com um ID gerado aleatoriamente
func NewProcess() *Process {
	// Use the current time to initialize the seed for the random number generator
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	// Generate a random UUID for the process ID
	id := uuid.New().String()

	// List of available memory sizes
	numbers := []int{1, 2, 4, 6, 10, 20, 30, 50, 100}

	// Select a random size from the list
	sizeInMemory := numbers[r.Intn(len(numbers))]

	// Return a new instance of Process with the randomly generated ID and memory size
	return &Process{
		ID:           id,
		SizeInMemory: sizeInMemory,
	}
}

// GetID retorna o ID do processo
func (p *Process) GetID() string {
	return p.ID
}

// SetID define o ID do processo
func (p *Process) SetID(id string) {
	p.ID = id
}

// GetSizeInMemory retorna o tamanho do processo em memória
func (p *Process) GetSizeInMemory() int {
	return p.SizeInMemory
}

// SetSizeInMemory define o tamanho do processo em memória
func (p *Process) SetSizeInMemory(sizeInMemory int) {
	p.SizeInMemory = sizeInMemory
}

// GetAddressMemory retorna o endereço de memória do processo
func (p *Process) GetAddressMemory() AddressMemory {
	return p.AddressMemory
}

// SetAddressMemory define o endereço de memória do processo
func (p *Process) SetAddressMemory(am AddressMemory) {
	p.AddressMemory = am
}
