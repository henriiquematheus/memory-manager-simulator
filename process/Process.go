package process

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// Process representa um processo no sistema
type Process struct {
	ID            string
	SizeInMemory  int
	AddressMemory AddressMemory
}

// NewProcess cria uma nova instância de Process com um ID gerado aleatoriamente
func NewProcess() *Process {

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	id := uuid.New().String()

	numbers := []int{1, 2, 4, 6, 10, 20, 30, 50, 100}

	sizeInMemory := numbers[r.Intn(len(numbers))]

	return &Process{
		ID:           id,
		SizeInMemory: sizeInMemory,
	}
}

// GetID
func (p *Process) GetID() string {
	return p.ID
}

// SetID
func (p *Process) SetID(id string) {
	p.ID = id
}

// GetSizeInMemory
func (p *Process) GetSizeInMemory() int {
	return p.SizeInMemory
}

// SetSizeInMemory
func (p *Process) SetSizeInMemory(sizeInMemory int) {
	p.SizeInMemory = sizeInMemory
}

// GetAddressMemory
func (p *Process) GetAddressMemory() AddressMemory {
	return p.AddressMemory
}

// SetAddressMemory
func (p *Process) SetAddressMemory(am AddressMemory) {
	p.AddressMemory = am
}
