package main

import "github.com/end-r/firevm"

func main() {
	vm := firevm.Create("guardian.vm")
	vm.Stats()
}

type VM firevm.VM
