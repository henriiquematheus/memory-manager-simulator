package so

func main() {
	so := NewSystemOperation()

	pl := so.SystemCall(CREATE_PROCESS, nil)
	so.SystemCall(WRITE_PROCESS, pl)

	p2 := so.SystemCall(CREATE_PROCESS, nil)
	so.SystemCall(WRITE_PROCESS, p2)

	p3 := so.SystemCall(CREATE_PROCESS, nil)
	so.SystemCall(WRITE_PROCESS, p3)
}
