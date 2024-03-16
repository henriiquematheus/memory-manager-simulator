package process

type AddressMemory struct {
	Start int // Início do endereço de memória
	End   int // Fim do endereço de memória
}

func (am *AddressMemory) GetStart() int {
	return am.Start
}

func (am *AddressMemory) SetStart(start int) {
	am.Start = start
}

func (am *AddressMemory) GetEnd() int {
	return am.End
}

func (am *AddressMemory) SetEnd(end int) {
	am.End = end
}

// GetSize retorna o tamanho do endereço de memória
func (am *AddressMemory) GetSize() int {
	return (am.End - am.Start) + 1
}
