package so

func main() {
	so := so.NewSystemOperation()

	pl := so.SystemCall(so.CREATE_PROCESS, nil)
	so.SystemCall(so.WRITE_PROCESS, pl)

	p2 := so.SystemCall(so.CREATE_PROCESS, nil)
	so.SystemCall(so.WRITE_PROCESS, p2)

	p3 := so.SystemCall(so.CREATE_PROCESS, nil)
	so.SystemCall(so.WRITE_PROCESS, p3)
}
