package so

func Execute(strategy int) {
	so := NewSystemOperation(strategy)

	pl := so.SystemCall(CREATE_PROCESS, nil)
	so.SystemCall(WRITE_PROCESS, pl)

	p2 := so.SystemCall(CREATE_PROCESS, nil)
	so.SystemCall(WRITE_PROCESS, p2)

	p3 := so.SystemCall(CREATE_PROCESS, nil)
	so.SystemCall(WRITE_PROCESS, p3)

	so.SystemCall(CLOSE_PROCESS, p2)
	so.SystemCall(CLOSE_PROCESS, p3)

	p4 := so.SystemCall(CREATE_PROCESS, nil)
	so.SystemCall(WRITE_PROCESS, p4)

	p5 := so.SystemCall(CREATE_PROCESS, nil)
	so.SystemCall(WRITE_PROCESS, p5)

	p6 := so.SystemCall(CREATE_PROCESS, nil)
	so.SystemCall(WRITE_PROCESS, p6)

}
