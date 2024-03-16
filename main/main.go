package main

import (
	"memory-manager-simulator/memory"
	"memory-manager-simulator/so"
)

func main() {
	so.Execute(memory.FIRST_FIT)

	//so.Execute(memory.WORST_FIT)

	//so.Execute(memory.BEST_FIT)
}
