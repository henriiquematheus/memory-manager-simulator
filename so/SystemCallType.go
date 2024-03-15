package so

// SystemCallType representa os tipos de chamadas de sistema
type SystemCallType int

// Constantes para tipos de chamadas de sistema
const (
	OPEN_PROCESS SystemCallType = iota
	READ_PROCESS
	CLOSE_PROCESS
	CREATE_PROCESS
	WRITE_PROCESS
)

// GetSystemCallTypeName retorna o nome do tipo de chamada de sistema com base na constante fornecida
func GetSystemCallTypeName(callType SystemCallType) string {
	switch callType {
	case OPEN_PROCESS:
		return "Open Process"
	case READ_PROCESS:
		return "Read Process"
	case CLOSE_PROCESS:
		return "Close Process"
	default:
		return "Unknown System Call Type"
	}
}
