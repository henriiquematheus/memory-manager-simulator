package process

// AddressMemory representa o endereço de memória de um processo
type AddressMemory struct {
	Start int // Início do endereço de memória
	End   int // Fim do endereço de memória
}

// GetStart retorna o valor do campo Start de AddressMemory
func (am *AddressMemory) GetStart() int {
	return am.Start
}

// SetStart define o valor do campo Start de AddressMemory
func (am *AddressMemory) SetStart(start int) {
	am.Start = start
}

// GetEnd retorna o valor do campo End de AddressMemory
func (am *AddressMemory) GetEnd() int {
	return am.End
}

// SetEnd define o valor do campo End de AddressMemory
func (am *AddressMemory) SetEnd(end int) {
	am.End = end
}

// GetSize retorna o tamanho do endereço de memória
func (am *AddressMemory) GetSize() int {
	return (am.End - am.Start) + 1
}
